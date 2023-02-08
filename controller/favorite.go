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

type FavoriteActionParam struct {
	VideoId    uint `query:"video_id" vd:"$>0; msg:'Illegal format'"`
	ActionType uint `query:"action_type" vd:"($ == 1 or $ == 2); msg:'Illegal value or format'"`
}

func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	var favoriteParam FavoriteActionParam

	if err := c.BindAndValidate(&favoriteParam); err != nil {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error()})
		return
	}

	if exist == false {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: "please login"})
		return
	}

	if err := service.FavoriteAction(ctx, uint(favoriteParam.VideoId), favoriteParam.ActionType, userToken); err != nil {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, types.Response{StatusCode: 0,
		StatusMsg: "success"})
}

func FavoriteList(ctx context.Context, c *app.RequestContext) {
	userToken, exist := c.Get(mw.IdentityKey)
	if exist == false {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: "please login"})
		return
	}

	userId, err := strconv.ParseUint(c.Query("user_id"), 10, 64)
	if err != nil || userId <= 0 {
		c.JSON(http.StatusOK, types.Response{StatusCode: 1,
			StatusMsg: err.Error() + "or user_id <= 0"})
		return
	}

	viewVideoList, err := service.FavoriteList(ctx, uint(userId), userToken)
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
