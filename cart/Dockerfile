FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd/server/server.go
FROM gcr.io/distroless/base
COPY --from=builder /app/server /server
CMD ["/server"]
EXPOSE 8080