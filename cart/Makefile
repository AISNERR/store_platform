BINARY_NAME=server
SOURCE_DIR=cmd/server
SOURCE_FILE=$(SOURCE_DIR)/server.go
TEST_DIR=route256/cart/internal/pkg/reviews/service
PROTO_DIR=proto
PROTO_FILE=$(PROTO_DIR)/loms.proto


all: build

generate:
	protoc --go_out=$(PROTO_DIR) --go-grpc_out=$(PROTO_DIR) --grpc-gateway_out . --openapiv2_out .  $(PROTO_FILE)
	 


build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) $(SOURCE_FILE)

run: build ./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME) coverage.out coverage.html


coverage: 
	go test $(TEST_DIR) -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

coverage-race:
	go test $(TEST_DIR) -coverprofile=coverage.out -race
	go tool cover -html=coverage.out -o coverage.html


lint:
	gocyclo .
	gocyclo $(SOURCE_DIR)
	gocyclo -avg .

cognition:
	gocognit .
	gocognit $(SOURCE_DIR)
	gocognit -avg .


