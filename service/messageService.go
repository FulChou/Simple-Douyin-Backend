package service

import (
	"Simple-Douyin-Backend/service/db"
	"context"
	"errors"
)

func MessageAction(ctx context.Context, toUserId uint, actionType uint, messageText string, userToken interface{}) error {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return errors.New("user doesn't exist in db")
	}
	user := users[0]
	_, err = db.QueryUserByID(toUserId)
	if err != nil {
		return errors.New("toUser doesn't exist in db")
	}
	if actionType == 1 {
		if err := db.CreateMessage(ctx, db.Message{UserID: user.ID, ToUserID: toUserId, Content: messageText}); err != nil {
			return errors.New("create this comment raise error in db")
		}
	} else {
		return errors.New("not support this action_type")
	}
	return nil
}

type ViewMessage struct {
	Id         uint   `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
}

func MessageList(toUserId uint, userToken interface{}) ([]*ViewMessage, error) {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return nil, errors.New("user doesn't exist in db")
	}
	me := users[0]
	_, err = db.QueryUserByID(toUserId)
	if err != nil {
		return nil, errors.New("toUser doesn't exist in db")
	}
	initMessageList, err := db.MessagesByUserID(me.ID, toUserId)
	if err != nil {
		return nil, errors.New("fail in finding chat messages")
	}
	if len(initMessageList) == 0 {
		return nil, errors.New("there is no history message")
	}
	messageList := make([]*ViewMessage, 0)
	for _, message := range initMessageList {
		messageList = append(messageList, &ViewMessage{
			Id:         message.ID,
			Content:    message.Content,
			CreateTime: message.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return messageList, nil
}
