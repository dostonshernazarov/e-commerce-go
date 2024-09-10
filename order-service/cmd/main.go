package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"order-service/internal/config"
	orderspb "order-service/protos/order"
	"order-service/service"
	"order-service/storage/connpostgres"
	"os"

	configloader "github.com/Oyatillohgayratov/config-loader"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var logger *slog.Logger

func main() {
	cfg := config.Config{}

	err := configloader.LoadYAMLConfig("config.yaml", &cfg)
	if err != nil {
		logger.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}
	DBstring := cfg.LoadConfig()

	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	db, err := connpostgres.New(DBstring)
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	OrderService := service.NewOrderService(db)
	port := fmt.Sprintf("%s:%s", cfg.OrderServer.Http.Host, cfg.OrderServer.Http.Port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor),
	)
	orderspb.RegisterOrdersServiceServer(s, OrderService)
	reflection.Register(s)

	log.Println("main: server running  port", port)

	err = s.Serve(lis)

	if err != nil {
		log.Fatal("failed while listening:", (err))
	}
}

func grpcInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	m, err := handler(ctx, req)
	if err != nil {
		logger.Error("RPC failed with error", slog.String("error", err.Error()))
	}
	return m, err
}

// {
// 	"id":"string",
// 	"user_id":"string",
// 	"products":[
// 		"product_id","product_id","product_id"
// 	]
// 	"locato":"string",
// 	"status":"status_number"
// }
