package store

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/model/request"
	"douyin-mini/util/code"
	"errors"
	"gorm.io/gorm"
)
// GetUserByName 根据用户名查询用户
func GetUserByName(ur *request.UserRes) (user *model.User,err error) {
	err = global.SqlDB.Where("user_name = ?",ur.UserName).First(&user).Error
	// 如果用户不存在,返回错误
	if errors.Is(err,gorm.ErrRecordNotFound){
		return user,err
	}
	return user,nil
}

// InsertUser 插入数据,单条
func InsertUser(user *model.User) bool{
	rowsAffected := global.SqlDB.Create(user).RowsAffected
	return rowsAffected != global.OPERATION_FAILED
}

// GetUserByID 根据用户id查询用户
func GetUserByID(ID int64)(user model.User,err error) {
  if rowsAffected := global.SqlDB.Where("user_id = ?",ID).Find(&user).RowsAffected; rowsAffected == global.OPERATION_FAILED {
	  return user,code.ErrUserNotExist
  }
  return user,nil
}

func GetUserListByID(IDList []int64)(user []model.User,err error) {
	if rowsAffected := global.SqlDB.Where("user_id IN ? ",IDList).Find(&user).RowsAffected; rowsAffected == global.OPERATION_FAILED {
		return user,code.ErrUserNotExist
	}
	return user,nil
}

func GetUserInfoFromRedis(authorID string)(user string,err error){
	user,err = global.RedisDB.Get(authorID).Result()
	if err != nil {
		return
	}
	return
}

func UpdateUserToRedis(key string,val []byte) error  {
	err := global.RedisDB.Set(key,val,global.REDIS_USERINFO_OVERDUE).Err()
	return err
}

func CurrentUserFollowTheVideoAuthorFromRedis(key string)(string,error) {
	res,err := global.RedisDB.HGet(key,"followObj").Result()
	if err != nil {
		return res,code.ErrQuery
	}
	return res,err
}

func InsertFollowToRedis(key string,val []byte) error{
	if err := global.RedisDB.HSet(key,"followObj",val).Err();  err != nil{
		return err
	}
	return nil
}