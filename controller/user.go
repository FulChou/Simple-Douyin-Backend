package controller

import (
	"Simple-Douyin-Backend/service"
	"Simple-Douyin-Backend/types"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

type UserLoginResponse struct {
	types.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	var registerVar types.UserParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		return
	}
	users, err := service.RegisterUser(ctx, registerVar)
	user := users[0]
	token := user.Password + user.UserName

	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: "注册失败"},
			UserId:   -1,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: types.Response{StatusCode: 0, StatusMsg: "注册成功"},
			UserId:   int64(user.ID),
			Token:    token,
		})
	}
}

func Login(ctx context.Context, c *app.RequestContext) {

}
