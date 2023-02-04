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
	if err != nil {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: "Register failed"},
			UserId:   -1,
			Token:    "",
		})
		return
	}
	user := users[0]
	token := user.Password + user.UserName
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: types.Response{StatusCode: 0, StatusMsg: "Register succeed"},
		UserId:   int64(user.ID),
		Token:    token,
	})

}

func Login(ctx context.Context, c *app.RequestContext) {
	var loginVar types.UserParam
	loginVar.UserName = c.Query("username")
	loginVar.PassWord = c.Query("password")
	token := loginVar.PassWord + loginVar.UserName

	if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
		return
	}
	user, err := service.LoginUser(ctx, loginVar)
	if err != nil { // user not exist
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: "Login failed, user not exist"},
			UserId:   -1,
			Token:    "",
		})
		return
	} else if loginVar.PassWord != user.Password {
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: types.Response{StatusCode: 1, StatusMsg: "Login failed, password not correct"},
			UserId:   -1,
			Token:    "",
		})
		return
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: types.Response{StatusCode: 0, StatusMsg: "Login succeed"},
			UserId:   int64(user.ID),
			Token:    token,
		})
	}

}
