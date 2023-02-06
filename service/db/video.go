package db

import (
	"context"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserId        uint   `json:"user_id"`
	Title         string `json:"title"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
}

func (v *Video) TableName() string {
	return VideoTableName
}

func CreateVideo(ctx context.Context, v Video) error {
	return DB.WithContext(ctx).Create(&v).Error
}
