package db

import (
	"fmt"
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	UserID  uint `json:"user_id"`
	VideoID uint `json:"video_id"`
}

func (v *Favorite) TableName() string {
	return "favorite"
}
func IsFavorite(userId, videoId uint) bool {
	res := make([]*Favorite, 0)
	err := DB.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).Find(&res).Error
	if err != nil {
		return false
	}
	fmt.Printf("%#v\n", res)
	if len(res) > 0 {
		return true
	}
	return false
}
