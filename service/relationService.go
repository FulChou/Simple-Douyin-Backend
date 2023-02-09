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
	_, err = db.QueryUserByID(toUserId)
	// fmt.Println(toUser)
	if err != nil {
		return errors.New("toUser doesn't exist in db")
	}
	switch {
	case actionType == 1:
		if db.IsFollow(user.ID, toUserId) == true {
			return errors.New("already followed this user")
		}
		if err := db.CreateRelation(ctx, db.Follow{UserID: user.ID, FollowUserID: toUserId}); err != nil {
			return errors.New("follow this user raise error in db")
		}
		// Update follow_count && follower_count in Table user
		err := db.UpdateUserFollows(user.ID, toUserId, 1)
		if err != nil {
			return err
		}

	case actionType == 2:
		if db.IsFollow(user.ID, toUserId) == false {
			return errors.New("already unfollowed this user")
		}
		if err := db.DeleteRelation(ctx, db.Follow{UserID: user.ID, FollowUserID: toUserId}); err != nil {
			return errors.New("unfollow this user raise error in db")
		}
		// Update follow_count && follower_count in Table user
		err := db.UpdateUserFollows(user.ID, toUserId, -1)
		if err != nil {
			return err
		}

	default:
		return errors.New("not support this action_type")
	}
	return nil
}

type ViewUser struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint64 `json:"follow_count"`
	FollowerCount uint64 `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func FollowList(userId uint, userToken interface{}) ([]*ViewUser, error) {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return nil, errors.New("user doesn't exist in db")
	}
	// get initial follows
	if userId == 0 {
		userId = users[0].ID
	}
	initFollows, err := db.GetFollowList(userId)
	if err != nil {
		return nil, errors.New("fail in finding follows")
	}
	if len(initFollows) == 0 {
		return nil, errors.New("user has no follow")
	}

	// Add user information to follow
	me := users[0]
	followList := make([]*ViewUser, 0)
	for _, follow := range initFollows {
		user, err := db.QueryUserByID(follow.FollowUserID)
		if err != nil {
			user = &db.User{}
		}
		followList = append(followList, &ViewUser{
			ID:            user.ID,
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      db.IsFollow(me.ID, user.ID),
		})
	}
	return followList, nil
}

func FollowerList(userId uint, userToken interface{}) ([]*ViewUser, error) {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return nil, errors.New("user doesn't exist in db")
	}
	if userId == 0 {
		userId = users[0].ID
	}
	// get initial follows
	initFollowers, err := db.GetFollowerList(userId)
	if err != nil {
		return nil, errors.New("fail in finding Followers")
	}
	if len(initFollowers) == 0 {
		return nil, errors.New("user has no follower")
	}

	// Add user information to follower
	me := users[0]
	followerList := make([]*ViewUser, 0)
	for _, follow := range initFollowers {
		user, err := db.QueryUserByID(follow.UserID)
		if err != nil {
			user = &db.User{}
		}
		followerList = append(followerList, &ViewUser{
			ID:            user.ID,
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      db.IsFollow(me.ID, user.ID),
		})
	}
	return followerList, nil
}

func FriendList(userId uint, userToken interface{}) ([]*ViewUser, error) {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return nil, errors.New("user doesn't exist in db")
	}
	if userId == 0 {
		userId = users[0].ID
	}
	// get initial friends
	initFriends, err := db.GetFriendList(userId)
	if err != nil {
		return nil, errors.New("fail in finding Friends")
	}
	if len(initFriends) == 0 {
		return nil, errors.New("user has no friend")
	}

	// Add user information to friends
	me := users[0]
	friendList := make([]*ViewUser, 0)
	for _, follow := range initFriends {
		user, err := db.QueryUserByID(follow.UserID)
		if err != nil {
			user = &db.User{}
		}
		friendList = append(friendList, &ViewUser{
			ID:            user.ID,
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      db.IsFollow(me.ID, user.ID),
		})
	}
	return friendList, nil
}
