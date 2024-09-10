package user

// import (
// 	"context"
// 	"fmt"
// 	"order-service/internal/config"
// 	userpb "order-service/protos/user"

// 	configloader "github.com/Oyatillohgayratov/config-loader"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

func UserClinet(req int) (int32, error) {
	// 	cfg := config.Config{}
	// 	err := configloader.LoadYAMLConfig("config.yaml", &cfg)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	port := fmt.Sprintf("%s:%s", cfg.UserServer.Http.Host, cfg.UserServer.Http.Port)

	// 	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 	if err != nil {
	// 		return 0, err
	// 	}

	// 	client := userpb.NewUsrServiceClient(conn)

	// 	res, err := client.GetUser(context.Background(), &userpb.UserRequest{Id: int32(req)})
	// 	if err != nil {
	// 		return 0, err
	// 	}

	// return res.Id, nil
	return 0, nil

}
