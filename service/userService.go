package service

import (
	"Simple-Douyin-Backend/service/db"
	"context"
	"errors"
)

// RegisterUser service func
func RegisterUser(ctx context.Context, userNew *db.User) (*db.User, error) {
	// token := registerVar.UserName + registerVar.PassWord
	// check database is user exist by username
	users, err := db.QueryUser(ctx, userNew.UserName)
	if len(users) != 0 || err != nil {
		return nil, errors.New("user exist")
	}
	if err := db.CreateUser(ctx, userNew); err != nil {
		return nil, errors.New("user create error")
	}
	return userNew, nil
}

func LoginUser(ctx context.Context, user db.User) (*db.User, error) {
	users, err := db.QueryUser(ctx, user.UserName)
	if len(users) == 0 || err != nil {
		return nil, errors.New("user not exist")
	}
	return users[0], nil
}
