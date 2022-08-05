package store

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/util/code"
)

func InsertVideo(video *model.Video) bool{
	rowsAffected := global.SqlDB.Create(&video).RowsAffected
	return  rowsAffected != global.OPERATION_FAILED
}
// GetVideoByAuthID 根据作者id,获取该作者发布的所有视频
func GetVideoByAuthID(authID int64)(videoInfo [] model.Video,err error){
	rowsAffected := global.SqlDB.Select([]string{"video_id","auth_id","video_title","play_url","cover_url","favorite_count","comment_count"}).Find(&videoInfo,"auth_id = ?",authID).Order("created_at desc").RowsAffected
	if rowsAffected == global.OPERATION_FAILED {
		return nil,code.ErrQuery
	}
	return
}
// GetVideosByVID 根据视频id获取视频信息
func GetVideosByVID(videoID []int64)(videoInfo []model.Video,err error){
	rowsAffected := global.SqlDB.Select([]string{"video_id","auth_id","video_title","play_url","cover_url","favorite_count","comment_count"}).Find(&videoInfo,"video_id IN ?",videoID).RowsAffected
	if rowsAffected == global.OPERATION_FAILED {
		return nil,code.ErrQuery
	}
	return
}

func GetVideoByVID(videoID int64)(videoInfo model.Video,err error){
	rowsAffected := global.SqlDB.Select([]string{"video_id","auth_id","video_title","play_url","cover_url","favorite_count","comment_count"}).Find(&videoInfo,"video_id = ?",videoID).RowsAffected
	if rowsAffected == global.OPERATION_FAILED {
		return videoInfo,code.ErrQuery
	}
	return
}



