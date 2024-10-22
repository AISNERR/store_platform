// package main

// import (
// 	"log"
// 	"net"

// 	"loms/internal/app/server"
// 	"loms/internal/pkg/repository"
// 	pb "loms/proto/loms_server"
// 	"net/http"

// 	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
// 	httpSwagger "github.com/swaggo/http-swagger"
// 	"google.golang.org/grpc"
// )

// func main() {
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	repo := repository.NewOrderRepository(100)
// 	server := server.Server{Repo: repo}

// 	s := grpc.NewServer()
// 	mux := runtime.NewServeMux()
// 	pb.RegisterOrderServiceServer(s, &server)

// 	log.Println("gRPC server is running on port 50051")
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}

// 	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

// 	log.Println("HTTP server is running on port 8080")
// 	if err := http.ListenAndServe(":8080", mux); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

package main

import (
	"database/sql"
	"log"
	"net"

	"loms/internal/app/server"
	"loms/internal/pkg/repository"
	pb "loms/proto/loms_server"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"  
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
)

func main() {

	connStr := "user=postgres dbname=postgres sslmode=disable password=12345"  
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
		 
	}  
	defer db.Close()

	repo := repository.NewRepository(db)
	grpcServer := server.Server{Repo: repo}


	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &grpcServer)
	mux := runtime.NewServeMux()

	 
	// err = pb.RegisterOrderServiceHandlerServer(context.Background(), mux, &grpcServer)
	// if err != nil {
	// 	log.Fatalf("failed to register gRPC gateway: %v", err)
	// }

	 
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

 
	go func() {
		log.Println("gRPC server is running on port 50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()
	log.Println("HTTP server is running on port 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
