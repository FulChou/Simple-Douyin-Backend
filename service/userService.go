package service

import (
	"Simple-Douyin-Backend/service/db"
	"Simple-Douyin-Backend/types"
	"context"
	"errors"
)

// RegisterUser service func
func RegisterUser(ctx context.Context, registerVar types.UserParam) ([]*db.User, error) {
	// token := registerVar.UserName + registerVar.PassWord
	// check database is user exist by username
	users, err := db.QueryUser(ctx, registerVar.UserName)
	if len(users) != 0 || err != nil {
		return nil, errors.New("user exist")
	}

	user := []*db.User{{
		UserName: registerVar.UserName,
		Password: registerVar.PassWord,
	}}
	if err := db.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
