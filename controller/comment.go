package controller

import (
	"Simple-Douyin-Backend/mw"
	"Simple-Douyin-Backend/service"
	"Simple-Douyin-Backend/types"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
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
