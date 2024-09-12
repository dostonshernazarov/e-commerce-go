package notificationserver

import (
	"context"
	"ekzamen_5/notification-service/internal/entity"
	notificationusecase "ekzamen_5/notification-service/internal/usecase/notification"
	notificationpb "github.com/D1Y0RBEKORIFJONOV/ekzamen-5protos/gen/go/notification"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"log"
	"time"
)

type notificationServer struct {
	notificationpb.UnimplementedNotificationServiceServer
	notification *notificationusecase.NotificationUseCase
	redisClient  *redis.Client
}

func RegisterNotificationServer(grpcServer *grpc.Server, notification *notificationusecase.NotificationUseCase) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("cannot connect to redis: %v", err)
	}
	notificationpb.RegisterNotificationServiceServer(grpcServer, &notificationServer{
		notification: notification,
		redisClient:  redisClient,
	})
}

func (n *notificationServer) CreateNotification(ctx context.Context,
	req *notificationpb.CreateNotificationReq) (*notificationpb.EmptyMessage, error) {
	err := n.notification.CreateNotification(ctx, &entity.CreateNotification{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &notificationpb.EmptyMessage{}, nil
}

func (n *notificationServer) GetNotification(ctx context.Context, req *notificationpb.GetNotificationReq) (*notificationpb.GetNotificationRes, error) {
	messages, err := n.notification.GetNotification(ctx, &entity.GetNotificationReq{
		UserId: req.UserId,
		Offset: req.Offset,
	})
	if err != nil {
		return nil, err
	}

	var res notificationpb.GetNotificationRes
	for _, msg := range messages.Messages {
		res.Messages = append(res.Messages, &notificationpb.Message{
			CreatedAt:  msg.CreateAt.Format(time.RFC3339),
			SenderName: msg.SenderName,
			Status:     msg.Status,
		})
	}
	return &res, nil
}

func (n *notificationServer) AddNotification(ctx context.Context, req *notificationpb.AddNotificationReq) (*notificationpb.EmptyMessage, error) {
	err := n.notification.AddNotification(ctx, &entity.AddNotificationReq{
		UserId: req.UserId,
		CreateMessage: &entity.CreateMessageReq{
			SenderName: req.Messages.SenderName,
			Status:     req.Messages.Status,
		},
	})
	if err != nil {
		return nil, err
	}

	err = n.redisClient.Publish(ctx, "notifications", "check").Err()
	if err != nil {
		return nil, err
	}

	return &notificationpb.EmptyMessage{}, nil
}

func (n *notificationServer) SendEmailNotification(ctx context.Context, req *notificationpb.SendEmailNotificationReq) (*notificationpb.SendEmailNotificationRes, error) {
	err := n.notification.SenEmail(ctx, req.Email, req.SenderName, req.Notification)
	if err != nil {
		return &notificationpb.SendEmailNotificationRes{
			Successfully: true,
		}, err
	}
	return &notificationpb.SendEmailNotificationRes{
		Successfully: true,
	}, nil
}
