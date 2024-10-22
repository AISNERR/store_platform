package model

type UserID = int64

type ItemDetails struct {
	UserId int64  `json:"user"`
	SkuId  int32  `json:"sku_id"`
	Count  int64  `json:"count"`
	Status string `json:"status"`
}
