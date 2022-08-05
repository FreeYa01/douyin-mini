package store

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/util/code"
	"github.com/go-redis/redis"
)


// CurrentUserFollowVideosAuthor 查询当前用户关注视频列表作者的情况
func CurrentUserFollowVideosAuthor(userID int64,authorIDList []int64)(followList []model.Follow){
	if rowsAffected :=  global.SqlDB.Find(&followList,"id = ? and follow_id IN ?",userID,authorIDList).RowsAffected;rowsAffected == global.OPERATION_FAILED{
		return nil
	}
	return followList
}
// CommentUserFollowCurrentVideoAuthor 评论区用户关注当前视频的情况
func CommentUserFollowCurrentVideoAuthor(userID []int64,authorID int64)(followList []model.Follow)  {
	if rowsAffected :=  global.SqlDB.Find(&followList,"follow_id = ? and id  IN ? ",authorID,userID).RowsAffected;rowsAffected == global.OPERATION_FAILED{
		return nil
	}
	return followList
}

// CurrentUserFollowTheVideoAuthor 当前用户关注视频作者情况
func CurrentUserFollowTheVideoAuthor(userID int64,authorID int64)(followList *model.Follow,err error){
	if rowsAffected :=  global.SqlDB.Find(&followList,"id = ? and follow_id = ? ",userID,authorID).RowsAffected;rowsAffected == global.OPERATION_FAILED{
		return nil,code.ErrQuery
	}
	return followList,nil
}

// CurrentUserFollowAllVideoAuthor 当前用户关注的所有作者
func CurrentUserFollowAllVideoAuthor(userID int64)(followList []model.Follow,err error)  {
	if rowsAffected :=  global.SqlDB.Find(&followList,"id = ? ",userID).RowsAffected;rowsAffected == global.OPERATION_FAILED{
		return nil,code.ErrQuery
	}
	return followList,nil
}

// CurrentUserAllFollower 当前用户的所有粉丝
func CurrentUserAllFollower(FollowID int64)(followList []model.Follow,err error)  {
	if rowsAffected :=  global.SqlDB.Distinct("id").Find(&followList,"follow_id = ? ",FollowID).RowsAffected;rowsAffected == global.OPERATION_FAILED{
		return nil,code.ErrQuery
	}
	return followList,nil
}
func InsertFollow(follow model.Follow) error {
	if rowsAffected := global.SqlDB.Create(&follow).RowsAffected; rowsAffected == global.OPERATION_FAILED{
		return code.ErrInsert
	}
	return nil
}

func DelFollowRecordByUIDAndFID(userID,FollowID int64)  error{
	if rowsAffected := global.SqlDB.Where("id = ? and follow_id = ?",userID,FollowID).Delete(&model.Follow{}).RowsAffected;rowsAffected == global.OPERATION_FAILED{
		return code.ErrDEl
	}
	return nil
}

// InsertToRedis 存入视频
func InsertToRedis(key string,val []byte,time int64) error {
	err := global.RedisDB.ZAdd(key,redis.Z{
		Score: float64(time),
		Member: val,
	}).Err()
	if err != nil {
		return err
	}
	return nil
}





