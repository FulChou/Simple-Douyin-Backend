// Code generated by hertz generator.

package main

import (
	"Simple-Douyin-Backend/biz/handler"
	"Simple-Douyin-Backend/controller"
	"Simple-Douyin-Backend/mw"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	//r.GET("/ping", handler.Ping)

	// your code ...
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	auth := r.Group("", mw.JwtMiddleware.MiddlewareFunc())
	auth.GET("/refresh_token", mw.JwtMiddleware.RefreshHandler)
	auth.GET("/ping", handler.Ping)

	auth.POST("/douyin/publish/action/", controller.Publish)

	apiRouter := r.Group("/douyin")
	apiRouter.POST("/user/login/", mw.JwtMiddleware.LoginHandler)
	// basic apis
	//apiRouter.GET("/feed/", controller.Feed)
	//apiRouter.GET("/user/", controller.UserInfo)

	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.GET("/publish/list/", mw.JwtMiddleware.MiddlewareFunc(), controller.PublishList)
	//
	//// extra apis - I
	//apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	//apiRouter.GET("/favorite/list/", controller.FavoriteList)
	//apiRouter.POST("/comment/action/", controller.CommentAction)
	//apiRouter.GET("/comment/list/", controller.CommentList)
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", controller.RelationAction)
	//apiRouter.GET("/relation/follow/list/", controller.FollowList)
	//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	//apiRouter.GET("/relation/friend/list/", controller.FriendList)
	//apiRouter.GET("/message/chat/", controller.MessageChat)
	//apiRouter.POST("/message/action/", controller.MessageAction)
}
