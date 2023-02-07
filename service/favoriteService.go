package service

import (
	"Simple-Douyin-Backend/service/db"
	"context"
	"errors"
)

func FavoriteAction(ctx context.Context, videoId uint, actionType uint, userToken interface{}) error {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return errors.New("user doesn't exist in db")
	}
	user := users[0]
	switch {
	case actionType == 1:
		if db.IsFavorite(user.ID, videoId) == true {
			return errors.New("all ready like this video")
		}
		if err := db.CreateFavorite(ctx, db.Favorite{VideoID: videoId, UserID: user.ID}); err != nil {
			return errors.New("favorite this video raise error in db")
		}
	case actionType == 2:
		if db.IsFavorite(user.ID, videoId) == false {
			return errors.New("all ready unlike this video")
		}
		if err := db.DeleteFavorite(ctx, db.Favorite{VideoID: videoId, UserID: user.ID}); err != nil {
			return errors.New("unlike this video raise error in db")
		}
	default:
		return errors.New("not support this action_type")
	}
	return nil
}
