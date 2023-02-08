package service

import (
	"Simple-Douyin-Backend/service/db"
	"context"
	"errors"
)

func RelationAction(ctx context.Context, toUserId uint, actionType uint, userToken interface{}) error {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return errors.New("user doesn't exist in db")
	}
	user := users[0]
	switch {
	case actionType == 1:
		if db.IsFollow(user.ID, toUserId) == true {
			return errors.New("already followed this user")
		}
		if err := db.CreateRelation(ctx, db.Follow{UserID: user.ID, FollowUserID: toUserId}); err != nil {
			return errors.New("follow this user raise error in db")
		}
	case actionType == 2:
		if db.IsFollow(user.ID, toUserId) == false {
			return errors.New("already unfollowed this user")
		}
		if err := db.DeleteRelation(ctx, db.Follow{UserID: user.ID, FollowUserID: toUserId}); err != nil {
			return errors.New("unfollow this user raise error in db")
		}
	default:
		return errors.New("not support this action_type")
	}
	return nil
}
