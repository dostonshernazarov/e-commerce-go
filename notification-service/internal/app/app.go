package app

import (
	grpcapp "ekzamen_5/notification-service/internal/app/grpc"
	"ekzamen_5/notification-service/internal/config"
	"ekzamen_5/notification-service/internal/infastructure/repository/mongodb"
	notifactionservice "ekzamen_5/notification-service/internal/services/notifaction"
	notificationusecase "ekzamen_5/notification-service/internal/usecase/notification"
	"log/slog"
)

type App struct {
	GrpcApp *grpcapp.App
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	db, err := mongodb.NewMongoDB(cfg, logger)
	if err != nil {
		panic(err)
	}
	dbUseCase := notificationusecase.NewNotificationRepository(db)

	service := notifactionservice.NewNotification(logger, dbUseCase)
	serviceUseCase := notificationusecase.NewNotificationUseCase(service)
	server := grpcapp.NewApp(cfg.RPCPort, logger, serviceUseCase)
	return &App{
		GrpcApp: server,
	}
}
