package notificationusecase

import (
	"context"
	"ekzamen_5/notification-service/internal/entity"
)

type notificationRepository interface {
	SaveNotification(ctx context.Context, notification *entity.Notification) error
	GetNotification(ctx context.Context, req *entity.GetNotificationReq) ([]entity.Message, error)
	AddNotification(ctx context.Context, userID string, message entity.Message) error
	GetOffsetNotification(ctx context.Context, userId string) (int64, error)
	UpdateOffsetNotification(ctx context.Context, userID string, offset int64) error
}

type NotificationRepository struct {
	notificationRepo notificationRepository
}

func NewNotificationRepository(notificationRepo notificationRepository) *NotificationRepository {
	return &NotificationRepository{notificationRepo: notificationRepo}
}

func (n *NotificationRepository) SaveNotification(ctx context.Context, notification *entity.Notification) error {
	return n.notificationRepo.SaveNotification(ctx, notification)
}
func (n *NotificationRepository) GetNotification(ctx context.Context, req *entity.GetNotificationReq) ([]entity.Message, error) {
	return n.notificationRepo.GetNotification(ctx, req)
}
func (n *NotificationRepository) AddNotification(ctx context.Context, userID string, message entity.Message) error {
	return n.notificationRepo.AddNotification(ctx, userID, message)
}

func (n *NotificationRepository) GetOffsetNotification(ctx context.Context, userID string) (int64, error) {
	return n.notificationRepo.GetOffsetNotification(ctx, userID)
}

func (n *NotificationRepository) UpdateOffsetNotification(ctx context.Context, userID string, offset int64) error {
	return n.notificationRepo.UpdateOffsetNotification(ctx, userID, offset)
}
