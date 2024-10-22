package loms

import (
	"context"
	"log"
	pb "route256/cart/proto/loms_server"
	"time"
	"google.golang.org/grpc"
)

type Item struct {
	Sku_id int64
	Count  int64
}

type LomsClient struct {
	lomsClinet pb.OrderServiceClient
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
		lomsClinet: pb.NewOrderServiceClient(conn),
	}
}

func (p *LomsClient) OrderCreate(user_id int64, items []Item) {
	var requestItems []*pb.Request_Item

	for _, item := range items {
		requestItem := &pb.Request_Item{
			Sku:   uint32(item.Sku_id),
			Count: uint64(item.Count),
		}
		requestItems = append(requestItems, requestItem)
	}
	req := &pb.Request{
		User:  user_id,
		Items: requestItems,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := p.lomsClinet.OrderCreate(ctx, req)
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}

	log.Printf("Order created with ID: %d", res.OrderID)
}
