package main

import (
	"context"
	"fmt"
	pb "loms/proto/loms_server"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect:", err)
		return
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	orderID := int64(6609119387103961710)
	req := &pb.OrderListRequest{OrderID: orderID}

	resp, err := client.OrderInfo(context.Background(), req)
	if err != nil {
		fmt.Println("could not get order info:", err)
		return
	}

	fmt.Printf("Response: Status: %s, User: %d\n", resp.Status, resp.User)
	for _, item := range resp.Items {
		fmt.Printf("Item SKU: %d, Count: %d\n", item.Sku, item.Count)
	}
}
