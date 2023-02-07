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

//func LoginUser(ctx context.Context, user db.User) (*db.User, error) {
//	users, err := db.QueryUser(ctx, user.UserName)
//	if len(users) == 0 || err != nil {
//		return nil, errors.New("user not exist")
//	}
//	return users[0], nil
//}

func GetUserInfo(ctx context.Context, userId uint, userToken interface{}) (Author, error) {
	user, err := db.QueryUserByID(userId)
	if err != nil || user == nil {
		return Author{}, errors.New("fail in finding user Info")
	}
	var userInfo Author
	users, err := db.QueryUser(ctx, userToken.(*db.User).UserName)
	if err != nil {
		return Author{}, errors.New("meUser doesn't exist in db")
	}
	myID := users[0].ID
	userInfo = Author{
		Id:            user.ID,
		Name:          user.UserName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      db.IsFollow(myID, userId),
	}
	return userInfo, nil
}
