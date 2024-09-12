gen-order:
	@protoc \
	--proto_path=../e-commerce-proto \
	--go_out=./protos \
	--go-grpc_out=./protos \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	../e-commerce-proto/order/order.proto

gen-user:
	@protoc \
	--proto_path=../e-commerce-proto \
	--go_out=./protos \
	--go-grpc_out=./protos \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	../e-commerce-proto/user/user.proto

gen-product:
	@protoc \
	--proto_path=../e-commerce-proto \
	--go_out=./protos \
	--go-grpc_out=./protos \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	../e-commerce-proto/product/product.proto


CURRENT_DIR := $(shell pwd)
DB_URL := "postgres://postgres:azamat@localhost:5432/mydatabase?sslmode=disable"

migrate-up:
	migrate -path migrations/ -database $(DB_URL) up


migrate-down:
	migrate -path migrations/ -database  $(DB_URL) down

migrate-force:
	migrate -path migrations/ -database  $(DB_URL) force 1

run:
	@go run cmd/main.go