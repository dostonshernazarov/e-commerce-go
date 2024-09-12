package notificationusecase

import (
	"context"
	"ekzamen_5/notification-service/internal/entity"
)

type notification interface {
	CreateNotification(ctx context.Context, req *entity.CreateNotification) error
	GetNotification(ctx context.Context, req *entity.GetNotificationReq) (*entity.GetNotificationResp, error)
	AddNotification(ctx context.Context, req *entity.AddNotificationReq) error
	SenEmail(ctx context.Context, to, senderName, content string) error
}

type NotificationUseCase struct {
	notification notification
}

func NewNotificationUseCase(ntf notification) *NotificationUseCase {
	return &NotificationUseCase{ntf}
}

func (n *NotificationUseCase) CreateNotification(ctx context.Context, req *entity.CreateNotification) error {
	return n.notification.CreateNotification(ctx, req)
}
func (n *NotificationUseCase) GetNotification(ctx context.Context, req *entity.GetNotificationReq) (*entity.GetNotificationResp, error) {
	return n.notification.GetNotification(ctx, req)
}

func (n *NotificationUseCase) AddNotification(ctx context.Context, req *entity.AddNotificationReq) error {
	return n.notification.AddNotification(ctx, req)
}

func (n *NotificationUseCase) SenEmail(ctx context.Context, to, senderName, content string) error {
	return n.notification.SenEmail(ctx, to, senderName, content)
}
