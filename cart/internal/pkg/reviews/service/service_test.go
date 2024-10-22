package service

// import (
// 	"context"
// 	"errors"
// 	"route256/cart/internal/pkg/reviews/model"
// 	"route256/cart/internal/pkg/reviews/service/mock"
// 	"testing"

// 	"github.com/go-playground/assert/v2"
// 	"github.com/gojuno/minimock"
// )

// func TestAddCart(t *testing.T) {
// 	ctrl := minimock.NewController(t)
// 	defer ctrl.Finish()

// 	cartMock := mock.NewCartRepositoryMock(t)
// 	productClient := mock.NewProductClientMock(t)

// 	service := NewService(productClient, cartMock)
// 	tests := []struct {
// 		name                string
// 		cart                model.Cart
// 		productInfo         model.ProductInfo
// 		createCartMockResp  model.Cart
// 		mockGetProductError error
// 		mockCreateCartError error
// 		expectedError       error
// 	}{
// 		{
// 			name: "success",
// 			cart: model.Cart{
// 				SKU:    1076963,
// 				UserID: 31337,
// 				Count:  273,
// 			},
// 			productInfo: model.ProductInfo{
// 				Name:  "TestProduct",
// 				Price: 10.0,
// 			},
// 			createCartMockResp: model.Cart{
// 				SKU:    1076963,
// 				UserID: 31337,
// 				Count:  273,
// 				Name:   "TestProduct",
// 				Price:  10.0,
// 			},
// 			mockGetProductError: nil,
// 			mockCreateCartError: nil,
// 			expectedError:       nil,
// 		},
// 		{
// 			name: "zero_user_id",
// 			cart: model.Cart{
// 				SKU:    1076963,
// 				UserID: 0,
// 				Count:  1,
// 			},
// 			productInfo: model.ProductInfo{
// 				Name:  "zero",
// 				Price: 10.0,
// 			},
// 			createCartMockResp: model.Cart{
// 				SKU:    1076963,
// 				UserID: 0,
// 				Count:  1,
// 				Name:   "zero",
// 				Price:  10.0,
// 			},
// 			mockGetProductError: nil,
// 			mockCreateCartError: nil,
// 			expectedError:       errors.New("fail validation"),
// 		},
// 		{
// 			name: "product_client_err",
// 			cart: model.Cart{
// 				SKU:    1076963,
// 				UserID: 345,
// 				Count:  1,
// 			},
// 			productInfo: model.ProductInfo{
// 				Name:  "t-shirt",
// 				Price: 10.0,
// 			},
// 			createCartMockResp: model.Cart{
// 				SKU:    1076963,
// 				UserID: 345,
// 				Count:  1,
// 				Name:   "t-shirt",
// 				Price:  10.0,
// 			},
// 			mockGetProductError: model.ErrNoProductSku,
// 			mockCreateCartError: nil,
// 			expectedError:       model.ErrNoProductSku,
// 		},
// 		{
// 			name: "bad_name_err",
// 			cart: model.Cart{
// 				SKU:    1076963,
// 				UserID: 345,
// 				Count:  1,
// 			},
// 			productInfo: model.ProductInfo{
// 				Name:  "",
// 				Price: 10.0,
// 			},
// 			createCartMockResp: model.Cart{
// 				SKU:    1076963,
// 				UserID: 345,
// 				Count:  1,
// 				Name:   "",
// 				Price:  10.0,
// 			},
// 			mockGetProductError: nil,
// 			mockCreateCartError: nil,
// 			expectedError:       errors.New("product info bad name"),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			productClient.GetProductMock.Return(tt.productInfo, tt.mockGetProductError)

// 			if tt.expectedError == nil && tt.mockCreateCartError == nil {
// 				cartMock.CreateCartMock.Return(&tt.createCartMockResp, nil)
// 			} else if tt.mockCreateCartError != nil {
// 				cartMock.CreateCartMock.Return(nil, tt.mockCreateCartError)
// 			}

// 			result, err := service.AddCart(context.Background(), tt.cart)
// 			if result == nil || err != nil {
// 				assert.Equal(t, err, tt.expectedError)
// 			} else {
// 				assert.Equal(t, string(result.Name), string(tt.productInfo.Name))
// 			}
// 		})
// 	}
// }

// func TestGetCart(t *testing.T) {
// 	ctrl := minimock.NewController(t)
// 	defer ctrl.Finish()

// 	cartMock := mock.NewCartRepositoryMock(t)
// 	productClient := mock.NewProductClientMock(t)

// 	service := NewService(productClient, cartMock)
// 	tests := []struct {
// 		name                string
// 		user_id             int64
// 		cartInfo            model.CartInfo
// 		mockGetCartResponse map[model.SKU]model.ItemDetails
// 		mockGetCartError    error
// 		expectedError       error
// 	}{
// 		{
// 			name:    "success",
// 			user_id: 3344,
// 			cartInfo: model.CartInfo{
// 				Items: []model.ItemDetails{model.ItemDetails{
// 					SkuId: 12345,
// 					Count: 1,
// 					Price: 7384,
// 					Name:  "2344",
// 				}},
// 				TotalPrice: 7384,
// 			},
// 			mockGetCartResponse: map[model.SKU]model.ItemDetails{
// 				12345: {
// 					SkuId: 12345,
// 					Count: 1,
// 					Price: 7384,
// 					Name:  "2344",
// 				},
// 			},
// 			mockGetCartError: nil,
// 			expectedError:    nil,
// 		},
// 		{
// 			name:    "fail_validation",
// 			user_id: 0,
// 			cartInfo: model.CartInfo{
// 				Items: []model.ItemDetails{model.ItemDetails{
// 					SkuId: 12345,
// 					Count: 1,
// 					Price: 7384,
// 					Name:  "2344",
// 				}},
// 				TotalPrice: 7384,
// 			},
// 			mockGetCartResponse: map[model.SKU]model.ItemDetails{
// 				12345: {
// 					SkuId: 12345,
// 					Count: 1,
// 					Price: 7384,
// 					Name:  "2344",
// 				},
// 			},
// 			mockGetCartError: nil,
// 			expectedError:    errors.New("fail validation"),
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			cartMock.GetCartMock.Return(tt.mockGetCartResponse, tt.mockGetCartError)

// 			result, err := service.GetCart(context.Background(), tt.user_id)

// 			if tt.expectedError != nil {
// 				assert.Equal(t, tt.expectedError.Error(), err.Error())
// 			} else {
// 				assert.Equal(t, result, tt.cartInfo)
// 			}
// 		})
// 	}

// }
