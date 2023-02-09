package service

import (
	"Simple-Douyin-Backend/service/db"
	"context"
	"errors"
)

// RegisterUser service func
func RegisterUser(ctx context.Context, userNew *db.User) (*db.User, error) {
	// check database is user exist by username
	users, err := db.QueryUser(userNew.UserName)
	if len(users) != 0 || err != nil {
		return nil, errors.New("user exist")
	}
	if err := db.CreateUser(ctx, userNew); err != nil {
		return nil, errors.New("user create error")
	}
	return userNew, nil
}

func GetUserInfo(ctx context.Context, userId uint, userToken interface{}) (Author, error) {

	if userId == 0 {
		var userInfo Author
		users, err := db.QueryUser(userToken.(*db.User).UserName)
		if err != nil {
			return Author{}, errors.New("meUser doesn't exist in db")
		}
		me := users[0]
		userInfo = Author{
			Id:            me.ID,
			Name:          me.UserName,
			FollowCount:   me.FollowCount,
			FollowerCount: me.FollowerCount,
			IsFollow:      db.IsFollow(me.ID, me.ID),
		}
		return userInfo, nil
	}

	user, err := db.QueryUserByID(userId)
	if err != nil || user == nil {
		return Author{}, errors.New("fail in finding user Info")
	}
	var userInfo Author
	users, err := db.QueryUser(userToken.(*db.User).UserName)
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
