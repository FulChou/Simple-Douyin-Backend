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
	// fmt.Printf("%#+v", user[0])
	if err := db.CreateUser(ctx, user); err != nil {
		// fmt.Println("fail")
		return nil, err
	}
	return user, nil
}

func LoginUser(ctx context.Context, loginVar types.UserParam) (*db.User, error) {
	users, err := db.QueryUser(ctx, loginVar.UserName)
	if len(users) == 0 || err != nil {
		return nil, errors.New("user not exist")
	}

	return users[0], nil
}
