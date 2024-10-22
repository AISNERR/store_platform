package server

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"route256/cart/internal/app/server"
	"route256/cart/internal/clients/products"
	"route256/cart/internal/pkg/reviews/model"
	"route256/cart/internal/pkg/reviews/repository"
	"route256/cart/internal/pkg/reviews/service"
	"time"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ItemS struct {
	suite.Suite
	productsClient           *products.ProductsClient
	productClientWithRetries *products.WithRetries
	cartRepository           *repository.Repository
	cartService              *service.CartService
	server                   *server.Server
	validator                *validator.Validate
}

func (s *ItemS) SetupSuite() {
	s.productsClient = products.New("testtoken", "http://route256.pavl.uk:8080/get_product")
	s.productClientWithRetries = products.NewProductWithRetries(3, time.Second, s.productsClient)
	s.cartRepository = repository.NewCartRepository(100)
	s.cartService = service.NewService(s.productClientWithRetries, s.cartRepository)
	s.validator = validator.New()
	s.server = server.New(s.cartService, s.validator)
}

func (s *ItemS) TestDeleteGood_Success() {

	userID := int64(123)
	sku := int64(234)
	cart := model.Cart{
		SKU:    234,
		Name:   "T-shirt",
		Count:  2,
		Price:  235,
		UserID: userID,
	}
	s.cartRepository.CreateCart(context.Background(), cart)
	
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/user/%d/cart/%d", userID, sku), nil)
	rr := httptest.NewRecorder()

	s.server.DeleteGood(rr, req)

	
	require.Equal(s.T(), http.StatusOK, rr.Code)
}
