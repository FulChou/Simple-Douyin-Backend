package service

import (
	"Simple-Douyin-Backend/service/db"
	"context"
	"errors"
)

func CommentAction(ctx context.Context, videoId, actionType, commentId uint, commentText string, userToken interface{}) error {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return errors.New("user doesn't exist in db")
	}
	user := users[0]
	switch {
	case actionType == 1:
		if err := db.CreateComment(ctx, db.Comment{VideoID: videoId, UserID: user.ID, Content: commentText}); err != nil {
			return errors.New("create this comment raise error in db")
		}
	case actionType == 2:
		if res, err := db.FindComment(commentId); res == nil || err != nil {
			return errors.New("not find this comment ")
		}
		if err := db.DeleteComment(ctx, commentId); err != nil {
			return errors.New("delete this comment raise error in db")
		}
	default:
		return errors.New("not support this action_type")
	}
	return nil
}
