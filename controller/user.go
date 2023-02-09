package controller

import (
	"Simple-Douyin-Backend/mw"
	"Simple-Douyin-Backend/service"
	"Simple-Douyin-Backend/service/db"
	"Simple-Douyin-Backend/types"
	"Simple-Douyin-Backend/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

type UserParam struct {
	UserName string `form:"username" json:"username" query:"username" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
	PassWord string `form:"password" json:"password" query:"password" vd:"(len($) > 0 && len($) < 30); msg:'Illegal format'"`
}

type UserLoginResponse struct {
	types.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	types.Response
	User service.Author `json:"user"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar UserParam
	if err := c.BindAndValidate(&registerVar); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: err.Error()},
			UserId:   -1,
		})
		return
	}
	user := &db.User{
		UserName: registerVar.UserName,
		Password: utils.MD5(registerVar.PassWord),
	}

	user, err := service.RegisterUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: err.Error()},
			UserId:   -1,
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: types.Response{StatusCode: 0, StatusMsg: "Register succeed"},
		UserId:   int64(user.ID),
	})

}

func UserInfo(ctx context.Context, c *app.RequestContext) {
	userToken, _ := c.Get(mw.IdentityKey)
	if userToken == nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: "please login"},
			User:     service.Author{},
		})
		return
	}

	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: "user_id params format is error"},
			User:     service.Author{},
		})
		return
	}
	userInfo, err := service.GetUserInfo(ctx, uint(userId), userToken)

	if err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: err.Error()},
			User:     userInfo,
		})
		return
	}
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: types.Response{StatusCode: 0, StatusMsg: "success"},
		User:     userInfo,
	})
}
