BINARY_NAME=server
SOURCE_DIR=cmd/server
SOURCE_FILE=$(SOURCE_DIR)/server.go
PROTO_DIR=proto
PROTO_FILE=$(PROTO_DIR)/loms.proto


all: build

generate:
	protoc --go_out=$(PROTO_DIR) --go-grpc_out=$(PROTO_DIR) --grpc-gateway_out . --openapiv2_out .  $(PROTO_FILE)

build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) $(SOURCE_FILE)


run: build ./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)  