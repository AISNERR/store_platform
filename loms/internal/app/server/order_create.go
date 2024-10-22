package server

import (
	"context"
	"log"
	"loms/internal/pkg/model"
	"loms/internal/pkg/repository"
	pb "loms/proto/loms_server"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	Repo *repository.Repository
}

func (s *Server) OrderCreate(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Received Request: User ID: %d, Items: %v", req.User, req.Items)

	var itemDetailsList []model.ItemDetails

	for _, item := range req.Items {
		cartGood := model.ItemDetails{
			UserId: int64(req.User),
			SkuId:  int32(item.Sku),
			Count:  int64(item.Count),
			Status: "New",
		}
		itemDetailsList = append(itemDetailsList, cartGood)
	}

	orderId, err := s.Repo.CreateOrder(ctx, itemDetailsList)
	log.Printf("orderId: %d", orderId)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		return nil, err
	}

	return &pb.Response{OrderID: orderId}, nil
}

func (s *Server) OrderInfo(ctx context.Context, req *pb.OrderListRequest) (*pb.OrderListResponse, error) {
	orderID := req.OrderID

	items, err := s.Repo.GetOrder(orderID)
	if err != nil {
		return nil, err
	}
	response := &pb.OrderListResponse{
		Status: items[0].Status,
		User:   items[0].UserId,
	}

	for _, item := range items {
		response.Items = append(response.Items, &pb.OrderListResponse_Item{
			Sku:   uint32(item.SkuId),
			Count: uint64(item.Count),
		})
	}

	return response, nil
}

func (s *Server) OrderPay(ctx context.Context, req *pb.OrderPayRequest) (*pb.OrderPayResponse, error) {
    err := s.Repo.PaymentOrder(req.OrderID)
    if err != nil {
        return &pb.OrderPayResponse{
        }, err
    }
    return &pb.OrderPayResponse{
    }, nil
}

func (s *Server) OrderCancel(ctx context.Context, req *pb.OrderCancelRequest) (*pb.OrderCancelResponse, error) {
    err := s.Repo.CancellationOrder(req.OrderID)
    if err != nil {
        return &pb.OrderCancelResponse{
        }, err
    }
    return &pb.OrderCancelResponse{
    }, nil
}

