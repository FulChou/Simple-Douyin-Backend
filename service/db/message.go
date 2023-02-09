package db

import (
	"context"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	ToUserID uint   `json:"to_user_id"`
	Content  string `json:"content"`
}

func (c *Message) TableName() string {
	return "message"
}

func CreateMessage(ctx context.Context, c Message) error {
	return DB.WithContext(ctx).Create(&c).Error
}

func MessagesByUserID(myId uint, toUserId uint) ([]*Message, error) {
	messages := make([]*Message, 0)
	if err := DB.Where("user_id = ? AND to_user_id = ?", myId, toUserId).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
