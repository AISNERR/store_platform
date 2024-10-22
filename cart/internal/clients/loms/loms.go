package loms

import (
	"context"
	"fmt"
	pb "route256/cart/v1"

	"route256/cart/internal/pkg/reviews/model"

	grpc "google.golang.org/grpc"
)

type LOMSClient struct {
	pbCli pb.LOMSServiceClient
}

func New(cc grpc.ClientConnInterface) *LOMSClient {
	cli := pb.NewLOMSServiceClient(cc)
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
	return resp.OrderId, nil
}

func toDTO(user model.UserID, cart *model.CartInfo) *pb.OrderCreateRequest {
	items := make([]*pb.Item, 0, len(cart.Items))

	for _, v := range cart.Items {
		item := pb.Item{
			Sku:   uint32(v.SkuId),
			Count: uint32(v.Count),
		}
		items = append(items, &item)
	}

	return &pb.OrderCreateRequest{
		UserId:  int64(user),
		Items: items,
	}
}
