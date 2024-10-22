package grpcadapter

import (
	"context"
	"loms/internal/model"
	grpc_loms "loms/pkg/api/loms/v1"
)

func (g *GrpcAdapter) OrderPay(ctx context.Context, in *grpc_loms.OrderPayRequest) (*grpc_loms.OrderPayResponse, error) {
	err := g.lomsManager.OrderPay(ctx, model.OrderID(in.OrderId))
	if err != nil {
		return nil, err
	}
	return &grpc_loms.OrderPayResponse{}, nil
}
