package db

import (
	"context"
	"errors"
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

func GetVideoByID(id uint) (*Video, error) {
	var video *Video
	if err := DB.Where("id = ?", id).First(&video).Error; err != nil {
		return nil, err
	}
	return video, nil
}

func UpdateCommentCount(videoId uint, count int) error {
	video, err := GetVideoByID(videoId)
	if err != nil {
		return err
	}
	if count == -1 && video.CommentCount == 0 {
		return errors.New("comment_count already zero")
	}
	if err := DB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", int(video.CommentCount)+count).Error; err != nil {
		return errors.New("update comment_count failed")
	}
	return nil
}
