package grpcapp

import (
	notificationserver "ekzamen_5/notification-service/internal/grpc/notification"
	notificationusecase "ekzamen_5/notification-service/internal/usecase/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type App struct {
	GRPCServer *grpc.Server
	logger     *slog.Logger
	Port       string
}

func NewApp(port string, logger *slog.Logger, notification *notificationusecase.NotificationUseCase) *App {
	grpcServer := grpc.NewServer()
	notificationserver.RegisterNotificationServer(grpcServer, notification)
	reflection.Register(grpcServer)
	return &App{
		GRPCServer: grpcServer,
		logger:     logger,
		Port:       port,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.App.Run"
	log := a.logger.With(
		slog.String("method", op),
		slog.String("port", a.Port))

	l, err := net.Listen("tcp", a.Port)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("starting gRPC server on port", slog.String("port", a.Port))
	err = a.GRPCServer.Serve(l)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}
func (a *App) Stop() {
	log := a.logger.With("port", a.Port)
	log.Info("stopping server")
	a.GRPCServer.GracefulStop()
}
