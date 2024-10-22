package products

import (
	"context"
	"errors"
	"route256/cart/internal/pkg/reviews/model"
	"time"
)

type ProductClient interface {
	GetProduct(ctx context.Context, sku int64) (model.ProductInfo, error)
}

type WithRetries struct {
	retryCount int64
	pause      time.Duration
	next       ProductClient
}

func NewProductWithRetries(
	retryCount int64,
	pause time.Duration,
	next ProductClient,
) *WithRetries {
	return &WithRetries{
		retryCount: retryCount,
		pause:      pause,
		next:       next,
	}
}

func (w *WithRetries) GetProduct(ctx context.Context, sku int64) (model.ProductInfo, error) {
	for i := 0; i < int(w.retryCount)-1; i++ {
		res, err := w.next.GetProduct(ctx, sku)
		if errors.Is(err, ErrShouldRetry) {
			time.Sleep(w.pause)
			continue
		}
		return res, err
	}
	return w.next.GetProduct(ctx, sku)
}
