package service

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/model/request"
	"douyin-mini/model/response"
	"douyin-mini/store"
	"douyin-mini/util"
	"time"
)

func AddComment(userID int64,comm request.CommentRes)(commInfo response.CommentInfo,err error){
	// 从库中找到用户信息
	user,err  := store.GetUserByID(userID)
	if err != nil{
		return
	}
	// 用户关注视频作者信息？
	video,err := store.GetVideoByVID(comm.VideoID)
	if err != nil {
		return
	}
	favorite,err := store.CurrentUserFollowTheVideoAuthor(userID,video.AuthID)
	var isFollow bool
	if err != nil {
		isFollow  = global.FALSE
	}else{
		isFollow = favorite.IsFollow
	}
	// 创建评论对象
	comment := model.Comment{
		CommentID: util.GenID(),
		UserID:    userID,
		VideoID:   comm.VideoID,
		Content:   comm.CommentText,
		CreatedAt: time.Now(),
	}
	// 存到数据库
	if err = store.InsertComment(comment); err != nil{
		return
	}
	//创建user对象
	author := response.Author{
		UserID: user.UserID,
		Name: user.UserName,
		FollowCount: user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:isFollow,
	}
	//创建commInfo对象,作为返回值
	commInfo = response.CommentInfo{
		ID: comment.CommentID,
		Author:author,
		Content: comment.Content,
		CreateDate:comment.CreatedAt.Format(global.TIME_FORMAT),
	}
	return
}

func DelComment(commentID int64)(commInfo response.CommentInfo,err error) {
	// 查询要删除的评论信息
	comm,err := store.GetCommentByCommentID(commentID)
	if err != nil {
		return
	}
	// 从库中找到用户信息
	user,err  := store.GetUserByID(comm.UserID)
	// 用户关注视频作者信息？
	video,err := store.GetVideoByVID(comm.VideoID)
	if err != nil {
		return
	}
	favorite,err := store.CurrentUserFollowTheVideoAuthor(comm.UserID,video.AuthID)
	var isFollow bool
	if err != nil {
		isFollow  = global.FALSE
	}else{
		isFollow = favorite.IsFollow
	}
	// 删除评论:为什么删除操作放在后面可以，前面就不行-------------------------------------------------------------------------------------------------------------------
	if err = store.DelCommentByID(commentID); err != nil {
		return
	}
	//创建user对象
	author := response.Author{
		UserID: user.UserID,
		Name: user.UserName,
		FollowCount: user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:isFollow,
	}
	//创建commInfo对象,作为返回值
	commInfo = response.CommentInfo{
		ID: comm.CommentID,
		Author:author,
		Content: comm.Content,
		CreateDate:comm.CreatedAt.Format(global.TIME_FORMAT),
	}
	return
}

func CommentList(videoID int64)(commInfoList []response.CommentInfo,err error) {
//	 根据视频id获取所有评论信息
	commList,err := store.GetCommentByVID(videoID)
	if err != nil{
		return
	}
//	根据视频id获取视频作者信息
   videoAuthor,err := store.GetVideoByVID(videoID)
   if err != nil {
	   return
   }
//	 生成用户id列表
	userIDList := make([]int64,len(commList))
	for i,v := range commList {
		userIDList[i] = v.UserID
	}
//	 查找用户信息
	userInfoList,err := store.GetUserListByID(userIDList)
	userInfoMap := make(map[int64]model.User,len(userInfoList))
	for _,v := range userInfoList{
		userInfoMap[v.UserID] = v
	}
//	评论区用户关注视频作者的相关信息
	followList := store.CommentUserFollowCurrentVideoAuthor(userIDList,videoAuthor.AuthID)
	followMap := make(map[int64]bool,len(followList))
	for _,v := range followList{
		followMap[v.ID] = v.IsFollow
	}
//  最后结果集
	commInfoList = make([]response.CommentInfo,len(commList))
	for i,v := range commList{
		commInfo := response.CommentInfo{
			ID: v.CommentID,
			Author:response.Author{
				UserID: userInfoMap[v.UserID].UserID,
				Name:  userInfoMap[v.UserID].UserName,
				FollowCount: userInfoMap[v.UserID].FollowCount,
				FollowerCount:  userInfoMap[v.UserID].FollowerCount,
				IsFollow: followMap[v.UserID],
			},
			Content: v.Content,
			CreateDate: v.CreatedAt.Format(global.TIME_FORMAT),
		}
		commInfoList[i] = commInfo
	}
	return
}
