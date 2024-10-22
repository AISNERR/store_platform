package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	server2 "route256/cart/internal/app/server"
	"route256/cart/internal/clients/loms"
	"route256/cart/internal/clients/products"
	"route256/cart/internal/http/middleware"
	"route256/cart/internal/pkg/reviews/repository"
	"route256/cart/internal/pkg/reviews/service"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	connRPC := loms.NewConnRPC(conn)

	productsClient := products.New("testtoken", "http://route256.pavl.uk:8080/get_product")
	productClientWithRetries := products.NewProductWithRetries(3, time.Second, productsClient)
	cartRepository := repository.NewCartRepository(100)
	lomsClient := loms.New(conn)

	cartService := service.NewService(productClientWithRetries, cartRepository, lomsClient, connRPC)
	validator := validator.New()
	server := server2.New(cartService, validator)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /user/{user_id}/cart/{sku_id}", server.CreateGood)
	mux.HandleFunc("GET /user/{user_id}/cart", server.GetGood)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku_id}", server.DeleteGood)
	mux.HandleFunc("DELETE /user/{user_id}/cart", server.DeleteCart)
	mux.HandleFunc("POST /checkout", server.CheckoutCart)

	logMux := middleware.NewLogMux(mux)

	log.Println("server starting")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe(":8080", logMux); err != nil {
			log.Fatalf("server failed: %v", err)
		}
	}()
	sig := <-sigChan
	log.Printf("received signal: %s, shutting down...", sig)
	time.Sleep(1 * time.Second)
	log.Println("All tasks stop. Graceful shutdown succeed")
}
