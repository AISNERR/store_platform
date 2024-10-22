package service

import (
	"context"
	"fmt"
	"route256/cart/internal/pkg/reviews/model"
)

func (s *CartService) Checkout(ctx context.Context, userID model.UserID) (int64, error) {
	cart, err := s.GetCart(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("s.GetCart user %d, err: %w", err)
	}
	orderID, err := s.lomsClient.OrderCreate(ctx, userID, cart)
	if err != nil {
		return 0, fmt.Errorf("s.lomsClient.OrderCreate err: %w", err)
	}
	return orderID, nil
}
