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

type CommentListResponse struct {
	types.Response
	VideoCommentList []*service.VideoComment `json:"video_comment"`
}

func CommentList(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if userToken == nil || exist == false {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:         types.Response{StatusCode: 1, StatusMsg: "please login"},
			VideoCommentList: nil,
		})
		return
	}
	videoId, err := strconv.ParseUint(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:         types.Response{StatusCode: 1, StatusMsg: "video_id params format is error"},
			VideoCommentList: nil,
		})
		return
	}
	videoCommentList, err := service.CommentListService(ctx, uint(videoId), userToken)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:         types.Response{StatusCode: 1, StatusMsg: err.Error()},
			VideoCommentList: videoCommentList,
		})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:         types.Response{StatusCode: 0, StatusMsg: "success"},
		VideoCommentList: videoCommentList,
	})
}
