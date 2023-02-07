package controller

import (
	"Simple-Douyin-Backend/mw"
	"Simple-Douyin-Backend/service"
	"Simple-Douyin-Backend/types"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"path/filepath"
	"strconv"
)

func Publish(ctx context.Context, c *app.RequestContext) {
	title := c.PostForm("title")
	userToken, _ := c.Get(mw.IdentityKey)
	if userToken == nil {
		// can not find user has been login
		c.JSON(http.StatusOK, types.Response{
			StatusCode: 1,
			StatusMsg:  "user need login",
		})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, types.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", 0, filename)
	savePath := filepath.Join("./public/", title+"_"+finalName)
	if err := c.SaveUploadedFile(data, savePath); err != nil {
		c.JSON(http.StatusOK, types.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if err := service.VideoPublish(ctx, title, savePath, userToken); err != nil {
		c.JSON(http.StatusOK, types.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, types.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})

}

type VideoListResponse struct {
	types.Response
	ViewVideoList []*service.ViewVideo `json:"video_list"`
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	userToken, _ := c.Get(mw.IdentityKey)
	if userToken == nil {
		// user need login
		return
	}
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil {
		// user id is error, need be number
		return
	}
	viewVideoList, err := service.PublishListService(ctx, uint(userId))
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:      types.Response{StatusCode: 1, StatusMsg: err.Error()},
			ViewVideoList: viewVideoList,
		})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response:      types.Response{StatusCode: 0, StatusMsg: "success"},
		ViewVideoList: viewVideoList,
	})
}

func Feed(ctx context.Context, c *app.RequestContext) {
	var latestTime string
	latestTime = c.Query("latest_time")
	//token, _ = c.Get(mw.IdentityKey)
	viewVideoList, err := service.VideoListByTimeStr(latestTime)
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:      types.Response{StatusCode: 1, StatusMsg: err.Error()},
			ViewVideoList: viewVideoList,
		})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response:      types.Response{StatusCode: 0, StatusMsg: "success"},
		ViewVideoList: viewVideoList,
	})

	//userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	//if err != nil {
	//	// user id is error, need be number
	//	return
	//}
	//viewVideoList, err := service.PublishListService(ctx, uint(userId))
	//if err != nil {
	//	c.JSON(http.StatusOK, VideoListResponse{
	//		Response:      types.Response{StatusCode: 1, StatusMsg: err.Error()},
	//		ViewVideoList: viewVideoList,
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, VideoListResponse{
	//	Response:      types.Response{StatusCode: 0, StatusMsg: "success"},
	//	ViewVideoList: viewVideoList,
	//})
}
