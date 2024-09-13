package user

import (
	"context"
	"fmt"
	"order-service/internal/config"
	userpb "order-service/protos/user"

	configloader "github.com/Oyatillohgayratov/config-loader"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet(req string) (string, error) {
	cfg := config.Config{}
	err := configloader.LoadYAMLConfig("config.yaml", &cfg)
	if err != nil {
		return "0", err
	}
	port := fmt.Sprintf("%s:%s", cfg.UserServer.Http.Host, cfg.UserServer.Http.Port)

	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "0", err
	}

	client := userpb.NewUserServiceClient(conn)

	res, err := client.GetUserById(context.Background(), &userpb.GetUserByIdRequest{Id: req})
	if err != nil {
		return "0", err
	}
	
	return res.User.Id, nil
}
