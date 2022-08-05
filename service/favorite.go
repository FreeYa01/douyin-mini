package service

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/model/response"
	"douyin-mini/store"
	"fmt"
	"strconv"
)

// Favorite 点赞功能
func Favorite(userID,videoID,actionType int64) (err error){
	// 点赞
	if actionType == 1 {
		err = AddFavorite(userID,videoID)
	}else if actionType == 2{
		// 取消点赞
		err = CancelFavorite(userID,videoID)
	}
	return err
}

// AddFavorite 点赞
func AddFavorite(userID,videoID int64) error {
	// 从缓存中查询用户是否点赞，限制用户点赞
	err := GetFavoriteFromRedis(userID,videoID)
	if err == nil {
		return  nil
	}
	//  没有,插入到Favorite表中
	fav := model.Favorite{
		UserID:     userID,
		VideoID:    videoID,
		IsFavorite: global.TRUE,
	}
	// 插入到数据库中
	err = store.InsertFavorite(fav)
	if err != nil {
		return err
	}
	// 更新到redis
	err = UpdateFavoriteFromRedis(userID,videoID)
	return  err
}

// CancelFavorite 取消点赞
func CancelFavorite(userID,videoID int64) error {
// 先删除缓存
	err := CancelFavoriteFormRedis(userID,videoID)
	if err != nil {
		return err
	}
//  删除数据库
	if err = store.DelFavorite(userID,videoID);err != nil{
		return err
	}
	return nil
}

// FavoriteList 点赞列表
func FavoriteList(userID int64)(videoList []response.VideoInfo,err error){
//	 从redis中获取用户点赞的所有视频id
//	 从redis中获取用户点赞的所有视频信息
//	 从redis中获取用户的所有信息

//	 如果不在在数据库中查询,更新redis

//  从点赞表中获取用户点赞记录
	favoriteList,err := store.GetFavoriteByUID(userID)
	if err != nil {
		return nil,err
	}
//	获取用户点赞的视频ID列表
	favoriteIDList := make([]int64,len(favoriteList))
	for i,v := range favoriteList {
		favoriteIDList[i] = v.VideoID
	}
//	 根据点赞的视频ID列表获取视频信息
	videoInfo,err := store.GetVideosByVID(favoriteIDList)
	if err != nil {
		return nil,err
	}
//	根据视频信息获取作者id列表
	authorIDList := make([]int64,len(videoInfo))
	for i,v := range videoInfo {
		authorIDList[i] = v.AuthID
	}
//	根据视频作者信息,查询用户关注视频作者的情况
	followList := store.CurrentUserFollowVideosAuthor(userID,authorIDList)
	if err != nil {
		return nil,err
	}
	followMap := make(map[int64]bool,len(followList))
	for _,v := range followList{
		followMap[v.FollowID] = v.IsFollow
	}
//	 获取作者信息
	authorList,err := store.GetUserListByID(authorIDList)
	authorMap := make(map[int64]model.User,len(authorList))
	for _,v := range authorList{
		authorMap[v.UserID] = v
	}
	if err != nil {
		return nil,err
	}
	for i,v := range videoInfo{
		vInfo := response.VideoInfo{
			VideoID: v.VideoID,
			Author: response.UserInfo{
				UserID:v.AuthID,
				UserName: authorMap[v.AuthID].UserName,
				FollowCount: authorMap[v.AuthID].FollowCount,
				FollowerCount: authorMap[v.AuthID].FollowerCount,
				ISFollow: followMap[v.AuthID],
			},
			PlayUrl:v.PlayUrl,
			CoverUrl: v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount: v.CommentCount,
			IsFavorite: favoriteList[i].IsFavorite,
			Title: v.VideoTitle,
		}
		videoList = append(videoList,vInfo)
	}
	return videoList,nil
}
// GetFavoriteFromRedis 从redis里查询点赞状态
func GetFavoriteFromRedis(userID,videoID int64) (err error) {
	// 类型转换
	uid := strconv.FormatInt(userID,10)
	vid := strconv.FormatInt(videoID,10)
	key := fmt.Sprintf("favorite:uid:%s:vid:%s",uid,vid)
	err = store.GetFavoriteFromRedis(key,videoID)
	if err != nil {
		global.Lg.Info("redis中没有该条记录，需要存入redis中")
	}
	return err
}

// UpdateFavoriteFromRedis 更新用户点赞的状态
func UpdateFavoriteFromRedis(userID,videoID int64)(err error) {
	// 类型转换
	uid := strconv.FormatInt(userID,10)
	vid := strconv.FormatInt(videoID,10)
	key := fmt.Sprintf("favorite:uid:%s:vid:%s",uid,vid)
	err = store.UpdateFavoriteFormRedis(key,videoID)
	if err != nil{
		global.Lg.Error("更新redis失败")
	}
	return err
}
// CancelFavoriteFormRedis 取消用户点赞的状态
func CancelFavoriteFormRedis(userID,videoID int64)(err error) {
	// 类型转换
	uid := strconv.FormatInt(userID,10)
	vid := strconv.FormatInt(videoID,10)
	key := fmt.Sprintf("favorite:uid:%s:vid:%s",uid,vid)
	err = store.CancelFavoriteFormRedis(key,videoID)
	if err != nil{
		global.Lg.Error("删除redis缓存失败")
	}
	return err
}
