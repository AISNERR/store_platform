package repository

import (
	"database/sql"
	"errors"
	"loms/internal/pkg/model"
	"sync"
)

type Repository struct {
	mu *sync.Mutex
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
		mu: &sync.Mutex{},
	}
}

func (r *Repository) GetOrder(orderId int64) ([]model.ItemDetails, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var items []model.ItemDetails

	// Здесь выполняется запрос к базе данных для получения деталей заказа
	query := "SELECT item_id, item_name, quantity, status FROM orders WHERE order_id = ?"
	rows, err := r.db.Query(query, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.ItemDetails
		if err := rows.Scan(&item.Count, &item.SkuId, &item.UserId, &item.Status); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if len(items) == 0 {
		return nil, errors.New("order not found")
	}

	return items, nil
}

func (r *Repository) PaymentOrder(orderId int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Здесь выполняется обновление статуса заказа на "оплачен"
	query := "UPDATE orders SET status = 'payed' WHERE order_id = ?"
	result, err := r.db.Exec(query, orderId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}

func (r *Repository) CancellationOrder(orderId int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Здесь выполняется обновление статуса заказа на "отменен"
	query := "UPDATE orders SET status = 'cancelled' WHERE order_id = ?"
	result, err := r.db.Exec(query, orderId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("order not found")
	}

	return nil
}
