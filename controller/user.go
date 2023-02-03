package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(ctx context.Context, c *app.RequestContext) {
	//fmt.Println(c.Query("username"))
	var registerVar UserParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")
	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		return
	}
	token := registerVar.UserName + registerVar.PassWord
	//// check database
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
	//	})
	//} else { // add new
	//	atomic.AddInt64(&userIdSequence, 1)
	//	newUser := User{
	//		Id:   userIdSequence,
	//		Name: username,
	//	}
	//	usersLoginInfo[token] = newUser
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0, StatusMsg: "注册成功"},
		UserId:   1,
		Token:    token,
	})
	//}

	hlog.Info("test", registerVar, c.Param("username"))

}

func Login(ctx context.Context, c *app.RequestContext) {

}
