package models

import "time"

type Comment struct {
	CommentID   int32     `json:"comment_id"`
	PostID      int32     `json:"post_id"`
	UserID      int32     `json:"user_id"`
	Message     string    `json:"message"`
	CommentLike int32     `json:"comment_like"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateComment struct {
	CommentID int32   `json:"comment_id"`
	PostID    int32   `json:"post_id"`
	UserID    int32   `json:"user_id"`
	Message   *string `json:"message,omitempty"`  // Optional update
}

type DeleteComment struct {
	CommentID int32 `json:"comment_id"`
}

type GetComment struct {
	CommentID   int32     `json:"comment_id"`
	PostID      int32     `json:"post_id"`
	UserID      int32     `json:"user_id"`
	Message     string    `json:"message"`
	CommentLike int32     `json:"comment_like"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
