package repository

import (
	"context"
	"loms/internal/pkg/model"
	"math/rand"
	"time"
)

func GenerateRandomInt64() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63()
}

func (r *Repository) CreateOrder(
	ctx context.Context,
	cart_good []model.ItemDetails) (int64, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	orderId := GenerateRandomInt64()

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	for _, item := range cart_good {
		query := "INSERT INTO orders (order_id, item_id, item_name, quantity, status) VALUES (?, ?, ?, ?, ?)"
		_, err := tx.Exec(query, orderId, item.SkuId, item.UserId, item.Count, "awaiting payment")
		if err != nil {
			return 0, err
		}
	}

	return orderId, nil
}
