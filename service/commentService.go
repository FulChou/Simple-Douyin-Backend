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

type VideoComment struct {
	Id         uint   `json:"id"`
	User       Author `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

func CommentList(ctx context.Context, videoId uint, userToken interface{}) ([]*VideoComment, error) {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return nil, errors.New("user doesn't exist in db")
	}

	// get initial comments
	initComments, err := db.CommentsByVideoID(videoId)
	if err != nil {
		return nil, errors.New("fail in finding video comments")
	}
	if len(initComments) == 0 {
		return nil, errors.New("there is no comment")
	}

	// Add user information to comments
	me := users[0]
	commentList := make([]*VideoComment, 0)
	for _, comment := range initComments {
		user, err := db.QueryUserByID(comment.UserID)
		if err != nil {
			user = &db.User{}
		}
		commentList = append(commentList, &VideoComment{
			Id: comment.ID,
			User: Author{
				Id:            user.ID,
				Name:          user.UserName,
				FollowCount:   user.FollowerCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      db.IsFollow(me.ID, user.ID),
			},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return commentList, nil
}
