package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"route256/cart/internal/pkg/reviews/model"

	pb "route256/cart/proto/loms_server"

	"google.golang.org/grpc"
)

type CheckoutRequest struct {
	UserId model.UserID `json:"user_id"`
}

type CheckoutResponse struct {
	OrderId string `json:"order_id"`
}

type Item struct {
	Sku_id int64
	Count  int64
}

type CheckoutRPCRequest struct {
	UserId int64
	Items  []Item
}

type LomsClient struct {
	lomsClinet pb.OrderServiceClient
}

func NewOrderServiceClient(conn *grpc.ClientConn) *LomsClient {
	return &LomsClient{
		lomsClinet: pb.NewOrderServiceClient(conn),
	}
}

func (s *Server) CheckoutCart(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var createRequest CheckoutRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /checkout out failed: %s", errOut.Error())
			return
		}
		return
	}
	err = json.Unmarshal(body, &createRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /checkout out failed: %s", errOut.Error())
			return
		}
		return
	}

	orderID, err := s.cartService.Checkout(ctx, createRequest.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /checkout out failed: %s", errOut.Error())
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]int64{
        "user_id": orderID,
    }
	if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
