package products

import (
	"context"
	"route256/cart/internal/pkg/reviews/model"
	"testing"
	"time"
"go.uber.org/goleak"

)

func TestGetProductForItemsAsync_NoLeak(t *testing.T) {
	defer goleak.VerifyNone(t)  
	ctx := context.Background()
	items := []model.ItemDetails{
		{SkuId: 2958025},
	}
	client := New("testtoken", "http://route256.pavl.uk:8080/get_product")
	_, err := GetProductForItemsAsync(ctx, items, client.GetProduct)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	time.Sleep(500 * time.Millisecond)
}
