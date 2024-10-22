package loms

import (
	"context"
	"fmt"
	pb "route256/cart/proto/loms_server"

	"route256/cart/internal/pkg/reviews/model"

	grpc "google.golang.org/grpc"
)

type LOMSClient struct {
	pbCli pb.OrderServiceClient
}

func New(cc grpc.ClientConnInterface) *LOMSClient {
	cli := pb.NewOrderServiceClient(cc)
	return &LOMSClient{
		pbCli: cli,
	}
}

func (l *LOMSClient) OrderCreate(ctx context.Context, user model.UserID, cart *model.CartInfo) (int64, error) {
	req := toDTO(user, cart)
	resp, err := l.pbCli.OrderCreate(ctx, req)
	if err != nil {
		return 0, fmt.Errorf("grpcCall err OrderCreate: %w", err)
	}
	return resp.OrderID, nil
}

func toDTO(user model.UserID, cart *model.CartInfo) *pb.Request {
	items := make([]*pb.Request_Item, 0, len(cart.Items))

	for _, v := range cart.Items {
		item := pb.Request_Item{
			Sku:   uint32(v.SkuId),
			Count: uint64(v.Count),
		}
		items = append(items, &item)
	}

	return &pb.Request{
		User:  int64(user),
		Items: items,
	}
}
