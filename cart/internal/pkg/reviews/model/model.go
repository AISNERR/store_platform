package model

import "errors"

type UserID = int64
type SKU = int64
type Name = string
type Price = int64

type Cart struct {
	SKU    SKU    `json:"sku" validate:"required,min=1"`
	UserID UserID `json:"user_id" validate:"required,min=1"`
	Count  int64  `json:"count" validate:"required,min=1"`
	Name   Name
	Price  Price
}

type ProductInfo struct {
	Name  string
	Price int64
}

type ItemDetails struct {
	SkuId int64   `json:"sku_id"`
	Count int64   `json:"count"`
	Price float64 `json:"price"`
	Name  string  `json:"name"`
}

type RequestDelete struct {
	SKU    int64
	UserID int64
}

var ErrSkuNotFoundInCart = errors.New("no sku found in cart")
var ErrNoCartFound = errors.New("no cart found")
var ErrNoProductSku = errors.New("invalid sku")

type CartInfo struct {
	Items      []ItemDetails `json:"items"`
	TotalPrice float64       `json:"total_price"`
}
