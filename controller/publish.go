package controller

import (
	"Simple-Douyin-Backend/mw"
	"Simple-Douyin-Backend/service"
	"Simple-Douyin-Backend/service/minio"
	"Simple-Douyin-Backend/types"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
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
	finalName := fmt.Sprintf("%s_%s", title, filename)
	savePath := filepath.Join("./static/", finalName)

	// save to local
	if err := c.SaveUploadedFile(data, savePath); err != nil {
		c.JSON(http.StatusOK, types.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// save to minio
	bucketName := "dousheng"
	if err := minio.FileUploader(ctx, bucketName, finalName, savePath); err != nil {
		fmt.Println("save to minio failed")
		c.JSON(http.StatusOK, types.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// get URL from minio
	url, err := minio.GetFileUrl(bucketName, finalName, 0)
	if err != nil {
		log.Printf("get url failed")
	} else {
		log.Printf("User uploaded a file", url)
	}

	if err := service.VideoPublish(ctx, title, url.String(), userToken); err != nil {
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
	NextTime      int                  `json:"next_time"`
	ViewVideoList []*service.ViewVideo `json:"video_list"`
}

func PublishList(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if userToken == nil || exist == false {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:      types.Response{StatusCode: 1, StatusMsg: "please login"},
			ViewVideoList: nil,
		})
		return
	}
	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:      types.Response{StatusCode: 1, StatusMsg: "user_id params format is error"},
			ViewVideoList: nil,
		})
		return
	}
	viewVideoList, err := service.PublishList(ctx, uint(userId), userToken)
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
	userToken, _ := c.Get(mw.IdentityKey)
	viewVideoList, err := service.VideoListByTimeStr(latestTime, userToken)

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
