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

type CommentActionParam struct {
	VideoId     uint   `query:"video_id" vd:"$>0; msg:'Illegal format'"`
	ActionType  uint   `query:"action_type" vd:"($ == 1 || $ == 2); msg:'Illegal value or format'"`
	CommentText string `query:"comment_text"`
	CommentId   uint   `query:"comment_id"`
}

func CommentAction(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if exist == false {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: "please login"})
		return
	}

	var commentParam CommentActionParam
	if err := c.BindAndValidate(&commentParam); err != nil {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error()})
		return
	}

	if (commentParam.ActionType == 1 && commentParam.CommentText == "") ||
		(commentParam.ActionType == 2 && commentParam.CommentId == 0) {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: "param is error"})
		return
	}

	if err := service.CommentAction(ctx, commentParam.VideoId, commentParam.ActionType, commentParam.CommentId, commentParam.CommentText, userToken); err != nil {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, types.Response{StatusCode: 0,
		StatusMsg: "success"})
}

type CommentListResponse struct {
	types.Response
	CommentList []*service.VideoComment `json:"comment_list"`
}

func CommentList(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if userToken == nil || exist == false {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    types.Response{StatusCode: 1, StatusMsg: "please login"},
			CommentList: nil,
		})
		return
	}
	videoId, err := strconv.ParseUint(c.Query("video_id"), 10, 64)
	if err != nil || videoId <= 0 {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    types.Response{StatusCode: 1, StatusMsg: "video_id params format is error or value <= 0"},
			CommentList: nil,
		})
		return
	}
	videoCommentList, err := service.CommentList(ctx, uint(videoId), userToken)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    types.Response{StatusCode: 1, StatusMsg: err.Error()},
			CommentList: videoCommentList,
		})
		return
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    types.Response{StatusCode: 0, StatusMsg: "success"},
		CommentList: videoCommentList,
	})
}
