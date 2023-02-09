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

//
//func GetFriendList(userId uint) ([]*Follow, error) {
//	friendList := make([]*Follow, 0)
//	err := DB.Where("user_id = ? AND follow_user_id IN (?)", userId, DB.Table("follow").
//		Where("follow_user_id = ? AND deleted_at IS NUll", userId)).Find(&friendList).Error
//	if err != nil {
//		return nil, err
//	}
//	return friendList, nil
//}

func GetFriendList(userId uint) ([]*Follow, error) {
	friendList := make([]*Follow, 0)
	err := DB.Debug().Where("user_id IN (?) AND follow_user_id = ?", DB.Table("follow").Select("follow_user_id").
		Where("user_id = ? AND deleted_at IS NUll", userId), userId).Find(&friendList).Error
	if err != nil {
		return nil, err
	}
	return friendList, nil
}

// a ... { 1, 2, 3}
//  { 1, 2, 3}   a
//
