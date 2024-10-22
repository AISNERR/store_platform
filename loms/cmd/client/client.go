package main

import (
	"context"
	"log"
	"time"

	pb "loms/proto/loms_server"

	"google.golang.org/grpc"
)

func mai() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)

	req := &pb.Request{
		User: 45554,
		Items: []*pb.Request_Item{
			{Sku: 101, Count: 2},
			{Sku: 102, Count: 1},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.OrderCreate(ctx, req)
	if err != nil {
		log.Fatalf("could not create order: %v", err)
	}

	log.Printf("Order created with ID: %d", res.OrderID)
}
