package service

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/model/response"
	"douyin-mini/store"
	"douyin-mini/util/code"
	"encoding/json"
	"go.uber.org/zap"
)

func Feed (userID,latestTime int64) (videoList []response.VideoInfo,nextTime int64,err error) {
	// 从redis中获取视频流数据
	videoList,nextTime,err = FeedVideosFromRedis(userID,latestTime)
	//  去数据库中查询
	if err != nil {
		videoList,nextTime,err = FeedVideosFromMysql(userID,latestTime)
	}
	return
}

// FeedVideosFromRedis 从redis中获取数据
func FeedVideosFromRedis(userID,latestTime int64)(videoList []response.VideoInfo,nextTime int64,err error) {
	//	从redis中批量获取数据
	videoModelRedis, err := FeedFromRedis(latestTime)
	if err != nil {
		return
	}

	//	获取视频作者信息
	userList := make([]*response.UserInfo, len(videoModelRedis))
	for i, v := range videoModelRedis {
		userList[i], err = GetUserInfoFromRedis(v.AuthID)
		if err != nil {
			// 从mysql中获取
			userList[i],err = AuthorInfoFromMysql(userID,v.AuthID)
			if err == nil {
				return
			}
		//	 更新redis
			_ = UpdateUserToRedis(userList[i])
		}
	}
	//    赋值要返回的视频信息列表
	for i,v := range videoModelRedis{
		vidInfo := response.VideoInfo{
			VideoID: v.VideoID,
			Author: *userList[i],
			Title: v.VideoTitle,
			PlayUrl: v.PlayUrl,
			CoverUrl: v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount: v.CommentCount,
		}
		// 从redis中获取用户点赞该视频的情况
		err  = GetFavoriteFromRedis(vidInfo.Author.UserID,v.VideoID)
		if err != nil {
			vidInfo.IsFavorite = global.FALSE
		}
		vidInfo.IsFavorite = global.TRUE
		videoList = append(videoList,vidInfo)
	}
	return videoList,videoModelRedis[len(videoModelRedis)-1].CreatedAt.UnixMilli() /1000,nil
}

// AuthorInfoFromMysql 从数据库中获取作者信息
func AuthorInfoFromMysql(userID,authorID int64)(author *response.UserInfo,err error) {
	// 获取用户信息
	ur,err := GetUserInfoFromMysql(authorID)
	follow,err := store.CurrentUserFollowTheVideoAuthor(userID,authorID)
	// 错误判断
	if err != nil {
		global.Lg.Error("获取信息失败",zap.Error(err))
		return
	}
	author = &response.UserInfo{
		UserID: ur.UserID,
		UserName: ur.UserName,
		FollowCount: ur.FollowCount,
		FollowerCount: ur.FollowerCount,
		ISFollow:follow.IsFollow,
	}
	return author,err
}

// FeedVideosFromMysql 从数据库中获取数据
func FeedVideosFromMysql(userID,latestTime int64) (videoList []response.VideoInfo,nextTime int64,err error) {
		//		分批取出定量视频数据
		vidModelList,err := store.GetVideoListByTime(latestTime)
		if err != nil || len(vidModelList) == 0 {
			global.Lg.Error("获取视频列表失败")
			return nil,0,code.ErrQuery
		}

		//      筛选出符合条件是视频
		videoSiftTimeList := make([]model.Video,0)
		//      视频、作者id列表
		authorIDList  := make([]int64,0)
		videoIDList  := make([]int64,0)
		for _,v := range vidModelList{
			if v.CreatedAt.UnixMilli() < latestTime{
				videoSiftTimeList = append(videoSiftTimeList,v)
				authorIDList = append(authorIDList,v.AuthID)
				videoIDList = append(videoIDList,v.VideoID)
			}
		}

	//    从数据库中找出视频作者的信息
		userList,err := store.GetUserListByID(authorIDList)
		if err != nil{
			global.Lg.Error("获取用户信息失败")
			return nil,0,code.ErrQuery
		}
	//  用户未登录,点赞,关注均为默认设置false

	//  从数据库中查询用户点赞该视频的情况
		favoriteList := store.IsFavoriteVideoByID(userID,videoIDList)
		favoriteMap := make(map[int64]bool,len(favoriteList))
		for _,v := range favoriteList{
			favoriteMap[v.VideoID] = v.IsFavorite
		}
	//  从数据库中查询用户关注视频作者的情况
		followList := store.CurrentUserFollowVideosAuthor(userID,authorIDList)
		followMap := make(map[int64]bool,len(followList))
		for _,v := range followList{
			followMap[v.FollowID] = v.IsFollow
		}

	//	  创建作者信息
		authorList := make([]response.UserInfo,len(userList))
		authorMap := make(map[int64]response.UserInfo,len(authorList))
		for i,v  := range userList{
			auth := response.UserInfo{
				UserID: v.UserID,
				FollowCount: v.FollowCount,
				FollowerCount: v.FollowerCount,
				UserName: v.UserName,
				ISFollow:followMap[v.UserID],
			}
			authorList[i] = auth
		}
	//   更新authorMap的值,根据作者id,获取作者信息
		for _,v := range authorList {
			authorMap[v.UserID] = v
		}
	//    赋值要返回的视频信息列表
		for _,v := range videoSiftTimeList{
			vidInfo := response.VideoInfo{
				VideoID: v.VideoID,
				Author:authorMap[v.AuthID],
				Title: v.VideoTitle,
				PlayUrl: v.PlayUrl,
				CoverUrl: v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				CommentCount: v.CommentCount,
				IsFavorite: favoriteMap[v.VideoID],
			}
			videoList = append(videoList,vidInfo)
		}
	//  nextTime:本次视频流中发布最早视频的毫秒数。
		nextTime = vidModelList[len(vidModelList)-1].CreatedAt.UnixMilli()
		return videoList,nextTime,err
	}


// FeedFromRedis 从redis中获取视频流的所有数据
func FeedFromRedis(nextTime int64)(videoList []model.Video,err error) {
	videoModel,err := store.FeedFromRedis(nextTime)
//	反序列化
   videoList = make([]model.Video,len(videoModel))
	for i,v := range videoModel{
		err  = json.Unmarshal([]byte(v),videoList[i])
		if err != nil  {
			global.Lg.Error("序列化失败")
			return
		}
	}
	return videoList,err
}