package db

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	FollowCount   uint64 `json:"follow_count" `
	FollowerCount uint64 `json:"follower_count"`
}

func (u *User) TableName() string {
	return UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, users *User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// QueryUser query list of user info by name
func QueryUser(userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("user_name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func QueryUserByID(ID uint) (*User, error) {
	var res *User
	if err := DB.Where("id = ?", ID).First(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func CheckUser(account, password string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.Where("user_name = ?", account).Where("password = ?", password).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateUserFollows(userId uint, toUserId uint, count int) error {
	// update My_follow_count
	user, err := QueryUserByID(userId)
	if err != nil {
		return errors.New("user doesn't exist in db")
	}
	if count == -1 && user.FollowCount == 0 {
		return errors.New("follow_count already zero")
	}

	if err := DB.Model(&User{}).Where("id = ?", userId).Update("follow_count", int(user.FollowCount)+count).Error; err != nil {
		return errors.New("update user follow_count failed")
	}

	// update follower_count
	toUser, err := QueryUserByID(toUserId)
	if err != nil {
		return errors.New("toUser doesn't exist in db")
	}
	if count == -1 && user.FollowerCount == 0 {
		return errors.New("follower_count already zero")
	}
	if err := DB.Model(&User{}).Where("id = ?", toUserId).Update("follower_count", int(toUser.FollowerCount)+count).Error; err != nil {
		return errors.New("update user follower_count failed")
	}
	return nil
}
