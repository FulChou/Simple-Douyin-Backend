package service

import (
	"Simple-Douyin-Backend/service/db"
	"context"
	"errors"
	"fmt"
)

func VideoPublish(ctx context.Context, title, videoPath string, userToken interface{}) error {
	users, err := db.QueryUser(ctx, userToken.(*db.User).UserName)
	if err != nil {
		return errors.New("user doesn't exist in db")
	}
	user := users[0]

	// create video raw in db
	if err := db.CreateVideo(ctx, db.Video{PlayUrl: videoPath, UserId: user.ID, Title: title}); err != nil {
		return errors.New("create video record fail in database")
	}
	return nil
}

type Author struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint64 `json:"follow_count"`
	FollowerCount uint64 `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type ViewVideo struct {
	Id            uint   `json:"id"`
	Title         string `json:"title"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Author        Author `json:"author"`
}

func PublishListService(ctx context.Context, userId uint) ([]ViewVideo, error) {
	videoList, err := db.VideoListBy(ctx, userId)
	if err != nil || len(videoList) == 0 {
		return nil, errors.New("fail in finding video list")
	}
	viewVideoList := make([]ViewVideo, 0)

	for _, video := range videoList {
		user, err := db.QueryUserByID(ctx, video.UserId)
		if err != nil {
			user = &db.User{}
		}
		fmt.Printf("user: %#v\n", user)
		viewVideoList = append(viewVideoList, ViewVideo{
			Id: video.ID, Title: video.Title, PlayUrl: video.PlayUrl,
			CoverUrl: video.CoverUrl, FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount, IsFavorite: db.IsFavorite(user.ID, video.ID),
			Author: Author{Id: user.ID, Name: user.UserName, FollowCount: user.FollowCount,
				FollowerCount: user.FollowerCount, IsFollow: false,
			}})
	}
	return viewVideoList, nil
}
