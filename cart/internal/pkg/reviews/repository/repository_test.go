package repository

import (
	"context"
	"route256/cart/internal/pkg/reviews/model"
	"sync"
	"testing"
	"github.com/go-playground/assert/v2"
)

func TestRepository_CreateCart(t *testing.T) {
	repo := &Repository{
		mu:      &sync.Mutex{},
		storage: make(Storage),
	}

	cart := model.Cart{
		UserID: 1,
		SKU:    100,
		Count:  2,
		Price:  10.0,
		Name:   "Item 100",
	}

	createdCart, err := repo.CreateCart(context.Background(), cart)
	assert.Equal(t, err, nil)
	assert.Equal(t, createdCart.Count, cart.Count)

	cartDetails, _ := repo.GetCart(context.Background(), 1)
	if cartDetails[100].Count != 2 {
		t.Fatalf("expected count to be 2, got %d", cartDetails[100].Count)
	}
}

func TestRepository_GetCart(t *testing.T) {
	repo := &Repository{
		mu:      &sync.Mutex{},
		storage: make(Storage),
	}

	_, _ = repo.CreateCart(context.Background(), model.Cart{
		UserID: 1,
		SKU:    100,
		Count:  2,
		Price:  10.0,
		Name:   "Item 100"})

	cart, err := repo.GetCart(context.Background(), 1)
	assert.Equal(t, len(cart), 1)
	assert.Equal(t, err, nil)
}

func TestRepository_DelCart(t *testing.T) {
	repo := &Repository{
		mu:      &sync.Mutex{},
		storage: make(Storage),
	}

	_, _ = repo.CreateCart(context.Background(), model.Cart{
		UserID: 1,
		SKU:    100,
		Count:  2,
		Price:  10.0,
		Name:   "Item 100"})

	err := repo.DelCart(context.Background(), 1, 100)
	assert.Equal(t, err, nil)

	err = repo.DelCart(context.Background(), 1, 100)
	assert.Equal(t, err, model.ErrSkuNotFoundInCart)
}

func BenchmarkCreateCart(b *testing.B) {
    repo := NewCartRepository(100)  
    cartItem := model.Cart{
        UserID: 1,
        SKU:    100,
        Count:  1,
        Name:   "Test",
        Price:  10.0,
    }

    b.ResetTimer()  

    for i := 0; i < b.N; i++ {
        _, err := repo.CreateCart(context.Background(), cartItem)
        if err != nil {
            b.Fatalf("failed: %v", err)
        }
        cartItem.SKU++
    }
}

func BenchmarkDelCart(b *testing.B) {
    repo := NewCartRepository(100)  
    cartItem := model.Cart{
        UserID: 1,
        SKU:    100,
        Count:  1,
        Name:   "Test Item",
        Price:  10.0,
    }
    _, _ = repo.CreateCart(context.Background(), cartItem)

    b.ResetTimer()  

    for i := 0; i < b.N; i++ {
        err := repo.DelCart(context.Background(), cartItem.UserID, cartItem.SKU)
        if err != nil {
            b.Fatalf("failed to delete: %v", err)
        }
        _, _ = repo.CreateCart(context.Background(), cartItem)
    }
}
