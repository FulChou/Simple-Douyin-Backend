package db

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return UserTableName
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}
