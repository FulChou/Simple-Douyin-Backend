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

	video, err := db.GetVideoByID(videoId)
	if err != nil || video == nil {
		return errors.New("video doesn't exist in db")
	}

	switch {
	case actionType == 1:
		if db.IsFavorite(user.ID, videoId) == true {
			return errors.New("all ready like this video")
		}
		if err := db.CreateFavorite(ctx, db.Favorite{VideoID: videoId, UserID: user.ID}); err != nil {
			return errors.New("favorite this video raise error in db")
		}
		if err := db.UpdateFavoriteCountBy(videoId, 1); err != nil {
			return err
		}
	case actionType == 2:
		if db.IsFavorite(user.ID, videoId) == false {
			return errors.New("all ready unlike this video")
		}
		if err := db.DeleteFavorite(ctx, db.Favorite{VideoID: videoId, UserID: user.ID}); err != nil {
			return errors.New("unlike this video raise error in db")
		}
		if err := db.UpdateFavoriteCountBy(videoId, -1); err != nil {
			return err
		}
	default:
		return errors.New("not support this action_type")
	}
	return nil
}

func FavoriteList(ctx context.Context, userId uint, userToken interface{}) ([]*ViewVideo, error) {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return nil, errors.New("user doesn't exist in db")
	}
	me := users[0]

	videoList, err := db.VideoListByFavorite(ctx, userId)
	if err != nil || len(videoList) == 0 {
		return nil, errors.New("fail in finding video list")
	}
	viewVideoList, err := Conver2ViewVideo(videoList, me)
	if err != nil {
		return nil, errors.New("error in createViewVideo")
	}
	return viewVideoList, nil

}
