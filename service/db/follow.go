package db

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	UserID   uint `json:"user_id"`
	FollowID uint `json:"follow_user_id"`
}

func (v *Follow) TableName() string {
	return "follow"
}

func IsFollow(myID uint, userID uint) bool {
	res := make([]*Follow, 0)
	err := DB.Model(&Follow{}).Where("user_id = ? AND follow_user_id = ?", myID, userID).Find(&res).Error
	if err != nil {
		return false
	}
	if len(res) > 0 {
		return true
	}
	return false
}
