package entity

import "time"

type (
	Message struct {
		CreateAt   time.Time `json:"create_at" bson:"create_at"`
		SenderName string    `json:"sender_name" bson:"sender_name"`
		Status     string    `json:"status" bson:"status"`
	}
	CreateNotification struct {
		UserId string `json:"user_id" bson:"user_id"`
	}
	Notification struct {
		UserId   string    `json:"user_id" bson:"user_id"`
		Offset   int64     `json:"offset" bson:"offset"`
		Messages []Message `json:"messages" bson:"messages"`
	}
	EmptyMessage       struct{}
	GetNotificationReq struct {
		UserId string `json:"user_id" bson:"user_id"`
		Offset int64  `json:"offset" bson:"offset"`
	}
	GetNotificationResp struct {
		Messages []Message `json:"notifications" bson:"notifications"`
	}
	CreateMessageReq struct {
		Status     string `json:"status" bson:"status"`
		SenderName string `json:"sender_name" bson:"sender_name"`
	}
	AddNotificationReq struct {
		UserId        string            `json:"user_id" bson:"user_id"`
		CreateMessage *CreateMessageReq `json:"notification" bson:"notification"`
	}
	EmailNotificationReq struct {
		SenderName string    `json:"sender_name" bson:"sender_name"`
		SenderAt   time.Time `json:"sender_at" bson:"sender_at"`
		Tittle     string    `json:"Tittle" bson:"Tittle"`
		Content    string    `json:"content" bson:"content"`
		Recipient  []string  `json:"recipient" bson:"recipient"`
	}
)
