package db

import (
	"context"
	"gorm.io/gorm"
	"time"
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

type Comment struct {
	gorm.Model
	UserId  uint   `json:"user_id"`
	Content string `json:"content"`
}

func (v *Comment) TableName() string {
	return "comment"
}

func CreateVideo(ctx context.Context, v Video) error {
	return DB.WithContext(ctx).Create(&v).Error
}

func VideoListByUserID(ctx context.Context, userId uint) ([]*Video, error) {
	videos := make([]*Video, 0)
	if err := DB.Where("user_id = ?", userId).
		Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func VideoListByTime(lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, 0)
	if err := DB.Where("created_at < ?", lastTime).Order("created_at DESC").
		Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil

}

func CommentsByVideoID(videoId uint) ([]*Comment, error) {
	comments := make([]*Comment, 0)
	if err := DB.Where("video_id = ?", videoId).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
