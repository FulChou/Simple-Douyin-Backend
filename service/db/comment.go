package db

import (
	"context"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	VideoID uint   `json:"video_id"`
	Content string `json:"content"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func FindComment(ID uint) (*Comment, error) {
	var comment *Comment
	if err := DB.Where("id = ?", ID).Find(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func CreateComment(ctx context.Context, c Comment) error {
	return DB.WithContext(ctx).Create(&c).Error
}

func DeleteComment(ctx context.Context, commentId uint) error {
	return DB.WithContext(ctx).Where("id = ?", commentId).Delete(&Comment{}).Error
}
