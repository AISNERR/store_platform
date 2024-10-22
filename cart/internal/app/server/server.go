package server

import (
	"context"
	"route256/cart/internal/pkg/reviews/model"
)

type CartService interface {
	AddCart(ctx context.Context, cart model.Cart) (*model.Cart, error)
	GetCart(ctx context.Context, sku model.UserID) (*model.CartInfo, error)
	DelCart(ctx context.Context, user_id model.UserID, sku_id model.SKU) error
	DelWholeCart(ctx context.Context, user_id model.UserID) error
	Checkout(ctx context.Context, userID model.UserID) (int64, error)
}

type Validator interface {
	Struct(v any) error
}

type Server struct {
	cartService CartService
	validator   Validator
}

func New(cartService CartService, validator Validator) *Server {
	return &Server{
		cartService: cartService,
		validator:   validator,
	}
}
