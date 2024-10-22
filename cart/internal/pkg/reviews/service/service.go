package service

// slog
import (
	"context"
	"errors"
	"fmt"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/pkg/reviews/model"
	"sort"
)

type CartRepository interface {
	CreateCart(_ context.Context, cart model.Cart) (*model.Cart, error)
	GetCart(_ context.Context, user_id model.UserID) (map[model.SKU]model.ItemDetails, error)
	DelCart(ctx context.Context, user_id model.UserID, sku model.SKU) error
	DelWholeCart(ctx context.Context, user model.UserID) error
}

type ProductClient interface {
	GetProduct(ctx context.Context, sku int64) (model.ProductInfo, error)
}

type LOMSClient interface {
	OrderCreate(ctx context.Context, userID model.UserID, cart *model.CartInfo) (int64, error)
}

type CartService struct {
	repository    CartRepository
	productClient ProductClient
	lomsClient    LOMSClient
	ConnRPC       *loms.ConnRPC
}

func NewService(
	pc ProductClient,
	repository CartRepository,
	lomsClient LOMSClient,
	connRPC *loms.ConnRPC) *CartService {
	return &CartService{
		repository:    repository,
		productClient: pc,
		lomsClient:    lomsClient,
		ConnRPC:       connRPC,
	}
}

func (s *CartService) AddCart(ctx context.Context, cart model.Cart) (*model.Cart, error) {
	if cart.SKU < 1 || cart.UserID <= 0 || cart.Count < 1 {
		return nil, errors.New("fail validation")
	}

	productInfo, err := s.productClient.GetProduct(ctx, cart.SKU)
	if err != nil {
		return nil, model.ErrNoProductSku
	}
	if productInfo.Name == "" {
		return nil, errors.New("product info bad name")
	}
	cart.Name = productInfo.Name
	cart.Price = productInfo.Price

	return s.repository.CreateCart(ctx, cart)
}

func buildCart(responseGetCart map[int64]model.ItemDetails) *model.CartInfo {
	lisItems := make([]model.ItemDetails, 0, len(responseGetCart))
	totalPrice := 0.0
	for _, itemCart := range responseGetCart {
		lisItems = append(lisItems, itemCart)
		skuPrice := itemCart.Price * float64(itemCart.Count)
		totalPrice += skuPrice
	}
	sort.Slice(lisItems, func(i, j int) bool {
		return lisItems[i].Price > lisItems[j].Price
	})
	cartInfo := model.CartInfo{
		Items:      lisItems,
		TotalPrice: totalPrice,
	}
	return &cartInfo
}

func (s *CartService) GetCart(ctx context.Context, user_id model.UserID) (*model.CartInfo, error) {
	if user_id <= 0 {
		return nil, errors.New("fail validation")
	}
	responseGetCart, err := s.repository.GetCart(ctx, user_id)
	if err != nil {
		return nil, fmt.Errorf("s.repository.GetCart err: %w", err)
	}
	cartInfo := buildCart(responseGetCart)
	return cartInfo, nil
}

func (s *CartService) DelCart(ctx context.Context, user_id model.UserID, sku model.SKU) error {
	return s.repository.DelCart(ctx, user_id, sku)
}

func (s *CartService) DelWholeCart(ctx context.Context, user_id model.UserID) error {
	return s.repository.DelWholeCart(ctx, user_id)
}
