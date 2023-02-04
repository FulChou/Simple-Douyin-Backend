package controller

import (
	"Simple-Douyin-Backend/service/db"
	"Simple-Douyin-Backend/types"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"path/filepath"
)

func Publish(ctx context.Context, c *app.RequestContext) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	fmt.Println("token", token)
	//if _, exist := usersLoginInfo[token]; !exist {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, types.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	//user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", 0, filename)
	saveFile := filepath.Join("./public/", title+"_"+finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, types.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// create video raw in db
	if err := db.CreateVideo(ctx, db.Video{PlayUrl: saveFile, UserId: 110}); err != nil {
		c.JSON(http.StatusOK, types.Response{
			StatusCode: 1,
			StatusMsg:  "create video record fail in database",
		})
		return
	}

	c.JSON(http.StatusOK, types.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})

}
