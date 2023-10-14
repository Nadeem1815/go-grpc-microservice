authproto:
	protoc --go_out=. --go-grpc_out=. ./pkg/auth/pb/auth.proto

orderproto:
	protoc --go_out=. --go-grpc_out=. ./pkg/order/pb/order.proto

productproto:
	protoc --go_out=. --go-grpc_out=. ./pkg/product/pb/product.proto

server:
	go run cmd/main.go