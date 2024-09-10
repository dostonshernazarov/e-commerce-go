package product

// import (
// 	"context"
// 	"fmt"
// 	"order-service/internal/config"
// 	productpb "order-service/protos/product"

// 	configloader "github.com/Oyatillohgayratov/config-loader"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func ProductClinet(req int32) (int32, error) {
// 	cfg := config.Config{}
// 	err := configloader.LoadYAMLConfig("config.yaml", &cfg)
// 	if err != nil {
// 		return 0, err
// 	}
// 	port := fmt.Sprintf("%s:%s", cfg.ProductServer.Http.Host, cfg.ProductServer.Http.Port)

// 	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		return 0, err
// 	}

// 	client := productpb.NewProductServiceClient(conn)

// 	res, err := client.GetProducts(context.Background(), &productpb.GetProductRequest{Id: req})
// 	if err != nil {
// 		return 0, err
// 	}

// 	return res.Product.Id, nil
// }