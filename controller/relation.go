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

func RelationAction(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if userToken == nil || exist == false {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1, StatusMsg: "please login"})
		return
	}
	toUserId, err := strconv.ParseUint(c.Query("to_user_id"), 10, 64)

	if err != nil || toUserId < 1 {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1, StatusMsg: "to_user_id params format is wrong"})
		return
	}
	actionType, err := strconv.ParseUint(c.Query("action_type"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1, StatusMsg: "action_type params format is wrong"})
		return
	}
	if actionType != 1 && actionType != 2 {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1, StatusMsg: "action is wrong"})
		return
	}

	if err := service.RelationAction(ctx, uint(toUserId), uint(actionType), userToken); err != nil {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, types.Response{StatusCode: 0,
		StatusMsg: "success"})
}

type FollowListResponse struct {
	types.Response
	UserList []*service.ViewUser `json:"user_list"`
}

func FollowList(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if exist == false {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: "please login"})
		return
	}
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil || userId < 0 {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error() + "or user_id <= 0"})
		return
	}
	followList, err := service.FollowList(uint(userId), userToken)
	if err != nil {
		c.JSON(http.StatusOK, FollowListResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: err.Error()},
			UserList: followList,
		})
		return
	}
	c.JSON(http.StatusOK, FollowListResponse{
		Response: types.Response{StatusCode: 0, StatusMsg: "success"},
		UserList: followList,
	})
}

type FollowerListResponse struct {
	types.Response
	UserList []*service.ViewUser `json:"user_list"`
}

func FollowerList(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if exist == false {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: "please login"})
		return
	}
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil || userId <= 0 {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error() + "or user_id <= 0"})
		return
	}
	followerList, err := service.FollowerList(uint(userId), userToken)
	if err != nil {
		c.JSON(http.StatusOK, FollowerListResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: err.Error()},
			UserList: followerList,
		})
		return
	}
	c.JSON(http.StatusOK, FollowerListResponse{
		Response: types.Response{StatusCode: 0, StatusMsg: "success"},
		UserList: followerList,
	})
}

type FriendListResponse struct {
	types.Response
	UserList []*service.ViewUser `json:"user_list"`
}

func FriendList(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if exist == false {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: "please login"})
		return
	}
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil || userId <= 0 {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error() + "or user_id <= 0"})
		return
	}

	friendList, err := service.FriendList(uint(userId), userToken)
	if err != nil {
		c.JSON(http.StatusOK, FriendListResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: err.Error()},
			UserList: friendList,
		})
		return
	}
	c.JSON(http.StatusOK, FriendListResponse{
		Response: types.Response{StatusCode: 0, StatusMsg: "success"},
		UserList: friendList,
	})
}
