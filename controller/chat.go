package controller

import (
	"Simple-Douyin-Backend/mw"
	"Simple-Douyin-Backend/service"
	"Simple-Douyin-Backend/types"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

type MessageActionParam struct {
	ToUserId   uint   `query:"to_user_id" vd:"$>0; msg:'Illegal format'"`
	ActionType uint   `query:"action_type" vd:"$ == 1; msg:'Illegal value or format'"`
	Content    string `query:"content"`
}

func MessageAction(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if exist == false {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: "please login"})
		return
	}

	var messageParam MessageActionParam
	if err := c.BindAndValidate(&messageParam); err != nil {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error()})
		return
	}

	if messageParam.Content == "" {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: "chat message is empty"})
		return
	}

	if err := service.MessageAction(ctx, messageParam.ToUserId, messageParam.ActionType, messageParam.Content, userToken); err != nil {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, types.Response{StatusCode: 0,
		StatusMsg: "success"})
}

type MessageChatResponse struct {
	types.Response
	MessageList []*service.ViewMessage `json:"view_message"`
}

func MessageChat(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if userToken == nil || exist == false {
		c.JSON(http.StatusOK, MessageChatResponse{
			Response:    types.Response{StatusCode: 1, StatusMsg: "please login"},
			MessageList: nil,
		})
		return
	}
	toUserId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 64)
	if err != nil || toUserId <= 0 {
		c.JSON(http.StatusOK, MessageChatResponse{
			Response:    types.Response{StatusCode: 1, StatusMsg: "video_id params format is error or value <= 0"},
			MessageList: nil,
		})
		return
	}
	messageList, err := service.MessageList(uint(toUserId), userToken)
	if err != nil {
		c.JSON(http.StatusOK, MessageChatResponse{
			Response:    types.Response{StatusCode: 1, StatusMsg: err.Error()},
			MessageList: messageList,
		})
		return
	}
	c.JSON(http.StatusOK, MessageChatResponse{
		Response:    types.Response{StatusCode: 0, StatusMsg: "success"},
		MessageList: messageList,
	})
}
