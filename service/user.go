package service

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/model/request"
	"douyin-mini/model/response"
	"douyin-mini/store"
	"douyin-mini/util"
	"douyin-mini/util/code"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func Login(usr *request.UserRes)(*model.User,error){
	// 判断用户是否存在
	usrInfo,err := store.GetUserByName(usr)
	// 如果用户不存在,返回错误
	if errors.Is(err,gorm.ErrRecordNotFound) {
		return nil,err
	}
	// 验证用户密码是否正确
	if !util.Verify(usrInfo.UserPwd,usr.UserPwd){
		return nil,code.ErrPasswordIncorrect
	}
// 返回对象
	return usrInfo,err
}

func Register(usr *request.UserRes)(user *model.User,err error){
	// 查询用户信息,判断用户是否存在
	_,err = store.GetUserByName(usr)
	if !errors.Is(err,gorm.ErrRecordNotFound) {
		return nil,code.ErrUserExist
	}
	// 对用户密码进行加密
	pwd,err := util.Encrypt(usr.UserPwd)

	// 创建用户对象,生成用户id
	user = &model.User{
		UserID:   util.GenID(),
		UserPwd:  pwd,
		UserName: usr.UserName,
	}
	// 将用户信息存入数据库
	if ok := store.InsertUser(user);!ok{
		return nil,code.ErrReigster
	}
	// 返回对象
	return
}

func GetUserInfo(userID int64)( userInfo *response.UserInfo,err error){
//	 从redis中查询用户信息
	userInfo,err = GetUserInfoFromRedis(userID)
	//  redis中不存在用户,去查mysql
	var user model.User
	if err != nil {
		user,_ = GetUserInfoFromMysql(userID)
		userInfo = &response.UserInfo{
			UserID: user.UserID,
			UserName: user.UserName,
			FollowCount: user.FollowCount,
			FollowerCount: user.FollowCount,
			ISFollow:  global.FALSE,
		}
		//	更新redis缓存
		err = UpdateUserToRedis(userInfo)
		if err != nil {
			return nil,code.ErrUpdate
		}
		return userInfo,nil
	}
	return userInfo,nil
}

func GetUserInfoFromRedis(authorID int64) (user *response.UserInfo,err error) {
	// 类型转换
	id := strconv.FormatInt(authorID,10)
	userStr,err := store.GetUserInfoFromRedis(id)
	if err != nil {
		global.Lg.Info("用户还未添加到换存中")
		return
	}
	err = json.Unmarshal([]byte(userStr),&user)
	if err != nil {
		global.Lg.Error("类型转换失败")
		return
	}
	return
}

func UpdateUserToRedis(user *response.UserInfo) error {
	userID := strconv.FormatInt(user.UserID,10)
	userData,err := json.Marshal(&user)
	if err != nil {
		global.Lg.Error("序列化失败")
		return code.ErrTypeIncorrect
	}
	 err = store.UpdateUserToRedis("user:userID:"+userID,userData)
	 if err != nil {
	 	global.Lg.Error("更新redis失败")
		return err
	}
	return nil
}

func GetUserInfoFromMysql(userID int64) (user model.User,err error) {
	//	 查不到去数据库中查询
	user,err = store.GetUserByID(userID)
	return user,err
}

