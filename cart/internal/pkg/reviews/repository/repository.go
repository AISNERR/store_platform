package repository

import (
	"context"
	"errors"
	"route256/cart/internal/pkg/reviews/model"
	"sync"
)

type Repository struct {
	mu      *sync.Mutex
	storage Storage
}

type Storage = map[model.UserID]map[model.SKU]model.ItemDetails

func NewCartRepository(capacity int) *Repository {
	return &Repository{
		storage: make(Storage, capacity),
		mu:      &sync.Mutex{},
	}
}

func (r *Repository) CreateCart(_ context.Context, cart_good model.Cart) (*model.Cart, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if cart_good.SKU < 1 {
		return nil, errors.New("sku must be defined")
	}

	if r.storage[cart_good.UserID] == nil {
		r.storage[cart_good.UserID] = make(map[model.SKU]model.ItemDetails)
	}

	val := r.storage[cart_good.UserID][cart_good.SKU]
	val.Count += cart_good.Count
	val.Name = cart_good.Name
	val.Price = float64(cart_good.Price)
	val.SkuId = cart_good.SKU
	r.storage[cart_good.UserID][cart_good.SKU] = val

	return &cart_good, nil

}

func (r *Repository) GetCart(ctx context.Context, user_id model.UserID) (map[model.SKU]model.ItemDetails, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	carts := r.storage[user_id]
	return carts, nil
}

func (r *Repository) DelCart(ctx context.Context, user model.UserID, sku model.SKU) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.storage[user][sku]; !exists {
		return model.ErrSkuNotFoundInCart
	}
	delete(r.storage[user], sku)
	return nil
}

func (r *Repository) DelWholeCart(ctx context.Context, user model.UserID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.storage[user]; !exists {
		return model.ErrNoCartFound
	}
	delete(r.storage, user)
	return nil
}
