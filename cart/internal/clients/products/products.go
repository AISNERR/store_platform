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
)

var ErrShouldRetry = errors.New("should retry request")

type Next interface {
	GetProduct(_ context.Context, sku int64) (model.ProductInfo, error)
}

type RequestBody struct {
	Token string `json:"token"`
	Sku   int64  `json:"sku" validate:"required,min=3,max=50"`
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

func (p *ProductsClient) GetProduct(_ context.Context, sku int64) (model.ProductInfo, error) {

	requestBody := RequestBody{
		Token: p.token,
		Sku:   sku,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return model.ProductInfo{}, fmt.Errorf("cannot marshal requestBody %w", err) //errors.Is errors.As
	}

	resp, err := http.Post(p.url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return model.ProductInfo{}, fmt.Errorf("http.Post err: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 420 || resp.StatusCode == 429 {
		return model.ProductInfo{}, fmt.Errorf("status code is %d %w", resp.StatusCode, ErrShouldRetry)
	} else if resp.StatusCode != http.StatusOK {
		return model.ProductInfo{}, fmt.Errorf("status code is %d", resp.StatusCode)
	}

	var prod model.ProductInfo

	response_product, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.ProductInfo{}, fmt.Errorf("io.ReadAll(resp.Body) err: %w", err)
	}
	err = json.Unmarshal(response_product, &prod)
	if err != nil {
		return model.ProductInfo{}, fmt.Errorf("unmarshal resp body err: %w", err)
	}
	return prod, nil
}
