package store

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/util/code"
	"errors"
	"gorm.io/gorm"
)

// IsFavoriteVideoByID  根据用户id和视频id列表,获取对于该视频的点赞情况
func IsFavoriteVideoByID(userID int64,videoIDList []int64) (favoriteList []model.Favorite){
	rowsAffected := global.SqlDB.Where("user_id = ? and video_id IN ?",userID,videoIDList).Find(&favoriteList).RowsAffected
	if rowsAffected == global.OPERATION_FAILED {
		return nil
	}
	return
}

func GetFavoriteByUID (userID int64)(favoriteList []model.Favorite,err error) {
	rowsAffected := global.SqlDB.Find(&favoriteList,"user_id = ?",userID).RowsAffected
	if rowsAffected == global.OPERATION_FAILED{
		return nil,code.ErrQuery
	}
	return
}

// GetFavoriteByUIDAndVID   根据用户id和视频id查询点赞信息
func GetFavoriteByUIDAndVID(userID,videoID int64)(favorite model.Favorite,err error){
	err = global.SqlDB.First(&favorite,"user_id = ? and video_id = ?",userID,videoID).Error
	if errors.Is(err,gorm.ErrRecordNotFound){
		return favorite,err
	}
	return
}

// InsertFavorite 插入
func InsertFavorite(fav model.Favorite) error{
	rowsAffected := global.SqlDB.Create(&fav).RowsAffected
	if rowsAffected == global.OPERATION_FAILED{
		return code.ErrInsert
	}
	return nil
}

// DelFavorite 删除点赞列表
func DelFavorite(userID,videoID int64) error {
	if rowsAffected := global.SqlDB.Where("user_id = ? and video_id = ?",userID,videoID).Delete(&model.Favorite{}).RowsAffected;rowsAffected == global.OPERATION_FAILED{
		return code.ErrDEl
	}
	return nil
}

// GetFavoriteFromRedis 从redis获取数据
func GetFavoriteFromRedis(key string,val int64)error {
	// 判断key中成员是否在redis中
	if ok,_ := global.RedisDB.SIsMember(key,val).Result(); !ok{
		return code.ErrQuery
	}
	return nil
}

// UpdateFavoriteFormRedis 更新redis数据
func UpdateFavoriteFormRedis(key string,val int64) error  {
	if err := global.RedisDB.SAdd(key,val).Err();err != nil{
		return err
	}
	return nil
}

// CancelFavoriteFormRedis 删除缓存
func CancelFavoriteFormRedis(key string,val int64) error  {
	if err := global.RedisDB.SRem(key,val).Err();err != nil{
		return err
	}
	return nil
}


