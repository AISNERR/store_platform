package products

import (
	"context"
	"errors"
	"sync"
	"testing"

	"route256/cart/internal/pkg/reviews/model"
)

type MockProductsClient struct {
	products map[model.SKU]model.ProductInfo
	mutex    sync.Mutex
}

func NewMockProductsClient() *MockProductsClient {
	return &MockProductsClient{
		products: make(map[model.SKU]model.ProductInfo),
	}
}

func (m *MockProductsClient) GetProduct(ctx context.Context, sku int64) (model.ProductInfo, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if product, exists := m.products[model.SKU(sku)]; exists {
		return product, nil
	}
	return model.ProductInfo{}, errors.New("product not found")
}

func TestGetProductForItemsAsync(t *testing.T) {
	mockClient := NewMockProductsClient()
	mockClient.products[1] = model.ProductInfo{Price: 1, Name: "Product1"}
	mockClient.products[2] = model.ProductInfo{Price: 2, Name: "Product2"}
	mockClient.products[3] = model.ProductInfo{Price: 1627, Name: "Product3"}
	mockClient.products[4] = model.ProductInfo{Price: 2378, Name: "Product4"}

	items1 := []model.ItemDetails{
		{SkuId: 2},
		{SkuId: 1},
	}
	items2 := []model.ItemDetails{
		{SkuId: 1},
		{SkuId: 2},
	}
	items3 := []model.ItemDetails{
		{SkuId: 3},
		{SkuId: 4},
	}

	apiCall := func(ctx context.Context, sku model.SKU) (model.ProductInfo, error) {
		return mockClient.GetProduct(ctx, int64(sku))
	}

	t.Run("Test successful product retrieval", func(t *testing.T) { 
		productInfo, err := GetProductForItemsAsync(context.Background(), items1, apiCall)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if len(productInfo) != 2 {
			t.Fatalf("expected 2 products, got %d", len(productInfo))
		}
	})

	t.Run("Test error handling when product not found", func(t *testing.T) {
		t.Parallel()  
		mockClient.products = make(map[model.SKU]model.ProductInfo)
		productInfo, err := GetProductForItemsAsync(context.Background(), items2, apiCall)
		if err == nil {
			t.Fatalf("expected error, got none")
		}

		if len(productInfo) != 0 {
			t.Fatalf("expected 0 products, got %d", len(productInfo))
		}
	})

	t.Run("Test error handling when product not found", func(t *testing.T) {
		t.Parallel()  
		mockClient.products = make(map[model.SKU]model.ProductInfo)
		productInfo, err := GetProductForItemsAsync(context.Background(), items3, apiCall)
		if err == nil {
			t.Fatalf("expected error, got none")
		}

		if len(productInfo) != 0 {
			t.Fatalf("expected 0 products, got %d", len(productInfo))
		}
	})
}
