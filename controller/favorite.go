package controller

import (
	"Simple-Douyin-Backend/mw"
	"Simple-Douyin-Backend/service"
	"Simple-Douyin-Backend/types"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
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
