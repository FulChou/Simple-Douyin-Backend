package service

import (
	"Simple-Douyin-Backend/service/db"
	"context"
	"errors"
)

func VideoPublish(ctx context.Context, title, videoPath string, userToken interface{}) error {
	users, err := db.QueryUser(ctx, userToken.(*db.User).UserName)
	if err != nil {
		return errors.New("user doesn't exist in db")
	}
	user := users[0]

	// create video raw in db
	if err := db.CreateVideo(ctx, db.Video{PlayUrl: videoPath, UserId: user.ID, Title: title}); err != nil {
		return errors.New("create video record fail in database")
	}
	return nil

}

type ViewVideo struct {
}

func PublishListService(user_id uint) []*ViewVideo {
	return nil
}
