package service

import (
	"Simple-Douyin-Backend/service/db"
	"context"
	"errors"
	"strconv"
	"time"
)

// VideoPublish save video information to MySQL
func VideoPublish(ctx context.Context, title, videoPath string, userToken interface{}) error {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return errors.New("user doesn't exist in db")
	}
	user := users[0]
	coverUrl := "https://c-ssl.duitang.com/uploads/item/201606/29/20160629123842_wZeyR.jpeg"
	if err := db.CreateVideo(ctx, db.Video{PlayUrl: videoPath, CoverUrl: coverUrl, UserId: user.ID, Title: title}); err != nil {
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

func PublishList(ctx context.Context, userId uint, userToken interface{}) ([]*ViewVideo, error) {
	users, err := db.QueryUser(userToken.(*db.User).UserName)
	if err != nil {
		return nil, errors.New("user doesn't exist in db")
	}
	me := users[0]

	videoList, err := db.VideoListByUserID(ctx, userId)
	if err != nil || len(videoList) == 0 {
		return nil, errors.New("fail in finding video list")
	}
	viewVideoList, err := Conver2ViewVideo(videoList, me)
	if err != nil {
		return nil, errors.New("error in createViewVideo")
	}
	return viewVideoList, nil

}

func Conver2ViewVideo(videoList []*db.Video, me *db.User) ([]*ViewVideo, error) {
	viewVideoList := make([]*ViewVideo, 0)
	for _, video := range videoList {
		user, err := db.QueryUserByID(video.UserId)
		if err != nil {
			user = &db.User{}
		}
		viewVideoList = append(viewVideoList, &ViewVideo{
			Id: video.ID, Title: video.Title, PlayUrl: video.PlayUrl,
			CoverUrl: video.CoverUrl, FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount, IsFavorite: db.IsFavorite(me.ID, video.ID),
			Author: Author{Id: user.ID,
				Name:          user.UserName,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      db.IsFollow(me.ID, user.ID),
			}})
	}
	return viewVideoList, nil
}

func VideoListByTimeStr(timeStr string, userToken interface{}) ([]*ViewVideo, error) {
	me := new(db.User)

	if userToken != nil {
		users, err := db.QueryUser(userToken.(*db.User).UserName)
		if err != nil {
			return nil, errors.New("user doesn't exist in db")
		}
		me = users[0]
	}

	var lastTime time.Time
	if timeStr == "" {
		lastTime = time.Now()
	} else {
		timeStr = timeStr[:len(timeStr)-3]
		tUnix, err := strconv.Atoi(timeStr)
		if err != nil {
			return nil, errors.New("params latest_time format error")
		}
		lastTime = time.Unix(int64(tUnix), 0)
	}

	videoList, err := db.VideoListByTime(lastTime)

	if err != nil {
		return nil, errors.New("video not fond, please check last time params")
	}
	viewVideoList, err := Conver2ViewVideo(videoList, me)
	if err != nil {
		return nil, errors.New("erro in createViewVideo")
	}
	return viewVideoList, nil
}
