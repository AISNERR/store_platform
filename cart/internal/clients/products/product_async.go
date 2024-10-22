package products

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"route256/cart/internal/pkg/reviews/model"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var ErrShouldRetry = errors.New("should retry request")

type RequestBody struct {
	Token string `json:"token"`
	Sku   int64  `json:"sku"`
}

type ProductsClient struct {
	url   string
	token string
}

func New(token, url string) *ProductsClient {
	return &ProductsClient{
		url:   url,
		token: token,
	}
}

func (p *ProductsClient) GetProduct(ctx context.Context, sku int64) (model.ProductInfo, error) {
	requestBody := RequestBody{
		Token: p.token,
		Sku:   sku,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return model.ProductInfo{}, fmt.Errorf("cannot marshal requestBody: %w", err)
	}

	resp, err := http.Post(p.url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return model.ProductInfo{}, fmt.Errorf("http.Post err: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 420 || resp.StatusCode == 429 {
		return model.ProductInfo{}, fmt.Errorf("status code is %d: %w", resp.StatusCode, ErrShouldRetry)
	} else if resp.StatusCode != http.StatusOK {
		return model.ProductInfo{}, fmt.Errorf("status code is %d", resp.StatusCode)
	}

	var prod model.ProductInfo
	responseProduct, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.ProductInfo{}, fmt.Errorf("io.ReadAll(resp.Body) err: %w", err)
	}
	err = json.Unmarshal(responseProduct, &prod)

	if err != nil {
		return model.ProductInfo{}, fmt.Errorf("unmarshal resp body err: %w", err)
	}
	return prod, nil
}

func GetProductForItemsAsync(ctx context.Context, items []model.ItemDetails, apiCall func(ctx context.Context, sku model.SKU) (model.ProductInfo, error)) (map[model.SKU]model.ProductInfo, error) {
	if len(items) < 1 {
		return make(map[model.SKU]model.ProductInfo), nil
	}
	var g errgroup.Group
	productInfoBySKU := make(map[model.SKU]model.ProductInfo)
	productInfoMutex := sync.Mutex{}
	rateLimiter := time.Tick(100 * time.Millisecond)

	for _, item := range items {
		sku := item.SkuId
		g.Go(func() error {
			<-rateLimiter

			productInfo, err := apiCall(ctx, sku)
			if err != nil {
				return err
			}

			productInfoMutex.Lock()
			productInfoBySKU[sku] = productInfo
			productInfoMutex.Unlock()

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return productInfoBySKU, err
	}
	return productInfoBySKU, nil
}
