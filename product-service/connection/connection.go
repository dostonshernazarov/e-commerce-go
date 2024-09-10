package connection

import (
	"log"
	"net"
	"os"
	"product-service/internal/database/mongodb"
	"product-service/internal/service"
	"product-service/protos"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func ConnGrpc() {
	lis, err := net.Listen("tcp", os.Getenv("product_server"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	mongoDB, err := mongodb.NewMongoRepo(os.Getenv("mongo_url"))
	if err != nil {
		log.Fatal(err)
	}

	productService := service.NewProductService(mongoDB)
	s := grpc.NewServer()
	protos.RegisterProductServiceServer(s, productService)

	log.Println("Server is listening on port: ", os.Getenv("product_server"))
	if err = s.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
