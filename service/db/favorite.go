package db

import (
	"context"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	VideoID uint `json:"video_id"`
}

func (f *Favorite) TableName() string {
	return "favorite"
}

func IsFavorite(userId, videoId uint) bool {
	res := make([]*Favorite, 0)
	err := DB.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).Find(&res).Error
	if err != nil {
		return false
	}
	if len(res) > 0 {
		return true
	}
	return false
}

func CreateFavorite(ctx context.Context, f Favorite) error {
	return DB.WithContext(ctx).Create(&f).Error
}

func DeleteFavorite(ctx context.Context, f Favorite) error {
	return DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", f.UserID, f.VideoID).Delete(&Favorite{}).Error
}

func VideoListByFavorite(ctx context.Context, userId uint) ([]*Video, error) {
	videos := make([]*Video, 0)
	err := DB.Where("id IN (?)", DB.Table("favorite").Select("video_id").
		Where("user_id = ?", userId)).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
