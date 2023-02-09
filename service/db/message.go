package db

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	ToUserID uint   `json:"to_user_id"`
	Content  string `json:"content"`
}

func (m *Message) TableName() string {
	return "message"
}

func CreateMessage(m Message) error {
	return DB.Create(&m).Error
}
