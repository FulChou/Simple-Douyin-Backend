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
