package loms

import (
	"context"
	"log"
	pb  "route256/cart/v1"
	"time"
	"google.golang.org/grpc"
)

type Item struct {
	Sku_id int64
	Count  int64
}

type LomsClient struct {
	lomsClinet pb.LOMSServiceClient
}

type ConnRPC struct{
	ConnRPC *grpc.ClientConn
}

func NewConnRPC(conn *grpc.ClientConn) *ConnRPC {
	return &ConnRPC{
		ConnRPC: conn,
	}
}


func NewOrderServiceClient(conn *grpc.ClientConn) *LomsClient {
	return &LomsClient{
		lomsClinet: pb.NewLOMSServiceClient(conn),
	}
}

func (p *LomsClient) OrderCreate(user_id int64, items []Item) {
	var requestItems []*pb.Item

	for _, item := range items {
		requestItem := &pb.Item{
			Sku:   uint32(item.Sku_id),
			Count: uint32(item.Count),
		}
		requestItems = append(requestItems, requestItem)
	}
	req := &pb.OrderCreateRequest{
		UserId:  user_id,
		Items: requestItems,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := p.lomsClinet.OrderCreate(ctx, req)
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}

	log.Printf("Order created with ID: %d", res.OrderId)
}
