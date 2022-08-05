package service

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/model/response"
	"douyin-mini/store"
	"douyin-mini/util/code"
	"encoding/json"
)

func Publish(video *model.Video) error{
	// 写入redis
	 time := video.CreatedAt.UnixMilli() / 1000
	 if err := InsertToRedis(time,*video); err != nil {
		 return err
	 }
	// 将视频信息写入mysql
	if ok := store.InsertVideo(video); !ok{
		return code.ErrInsert
	}
	return nil
}

func PublishList(userID int64)(videoList []response.VideoInfo,err error){
 // 	视频作者信息
	var author response.UserInfo
//	 从数据库中查询用户发布的所有视频,按发布时间降序排序
	vidModelInfo,err := store.GetVideoByAuthID(userID)
	if err != nil {
		global.Lg.Error("查询失败")
		return nil,err
	}
	// 初始化
	videoList = make([]response.VideoInfo,len(vidModelInfo))

	//	 生成视频id列表
	videoIDList := make([]int64,len(vidModelInfo))
	for i,v := range vidModelInfo{
		videoIDList[i] = v.VideoID
	}

	//  获取用户点赞该视频的情况
	favoriteList := store.IsFavoriteVideoByID(userID,videoIDList)
	favoriteMap := make(map[int64]bool,len(favoriteList))
	for _,v := range favoriteList{
		favoriteMap[v.VideoID] = v.IsFavorite
	}

	// 从数据库中获取用户信息
	userInfo,err := store.GetUserByID(userID)
	if err != nil {
		return nil,err
	}

	if vidModelInfo[0].AuthID == userID{
		author  = response.UserInfo{
			UserID: userInfo.UserID,
			UserName: userInfo.UserName,
			FollowCount: userInfo.FollowerCount,
			FollowerCount: userInfo.FollowerCount,
			ISFollow: global.FALSE,
		}
		for i,v:= range vidModelInfo {
			pubVideo := response.VideoInfo{
				VideoID: v.VideoID,
				Author:    author,
				PlayUrl: v.PlayUrl,
				CoverUrl: v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				CommentCount: v.CommentCount,
				IsFavorite:favoriteMap[v.VideoID],
				Title: v.VideoTitle,
			}
			videoList[i] = pubVideo
		}
		return
	}

//	 查询用户关注视频作者的情况
	followerTheVideoAuthor,err  := store.CurrentUserFollowTheVideoAuthor(userID,vidModelInfo[0].AuthID)
	if err != nil {
		return nil,err
	}
	// 作者信息
	author = response.UserInfo{
		UserID: userInfo.UserID,
		UserName: userInfo.UserName,
		FollowCount: userInfo.FollowerCount,
		FollowerCount: userInfo.FollowerCount,
		ISFollow: followerTheVideoAuthor.IsFollow,
	}


	for i,v:= range vidModelInfo {
		pubVideo := response.VideoInfo{
			VideoID: v.VideoID,
			Author:   author,
			PlayUrl: v.PlayUrl,
			CoverUrl: v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount: v.CommentCount,
			IsFavorite:favoriteMap[v.VideoID],
			Title: v.VideoTitle,
		}
		videoList[i] = pubVideo
	}
	return
}

func InsertToRedis(time int64,video model.Video) error {
	key := "video:next_time"
	val,err := json.Marshal(video)
	if err != nil {
		global.Lg.Error("序列化失败")
		return err
	}
	err = store.InsertToRedis(key,val,time)
	if err != nil{
		global.Lg.Error("redis插入数据失败")
		return err
	}
	return nil
}
