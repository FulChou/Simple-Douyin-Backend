package db

import (
	"context"
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	UserID       uint `json:"user_id"`
	FollowUserID uint `json:"follow_user_id"`
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

func CreateRelation(ctx context.Context, f Follow) error {
	return DB.WithContext(ctx).Create(&f).Error
}

func DeleteRelation(ctx context.Context, f Follow) error {
	return DB.WithContext(ctx).Where("user_id = ? AND follow_user_id = ?", f.UserID, f.FollowUserID).Delete(&f).Error
}

func GetFollowList(userId uint) ([]*Follow, error) {
	followList := make([]*Follow, 0)
	err := DB.Where("user_id = ?", userId).Find(&followList).Error
	if err != nil {
		return nil, err
	}
	return followList, nil
}

func GetFollowerList(followerId uint) ([]*Follow, error) {
	followerList := make([]*Follow, 0)
	err := DB.Where("follow_user_id = ?", followerId).Find(&followerList).Error
	if err != nil {
		return nil, err
	}
	return followerList, nil
}
