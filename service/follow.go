package service

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/model/request"
	"douyin-mini/model/response"
	"douyin-mini/store"
	"encoding/json"
	"strconv"
)

func Follow(userID int64,follow request.FollowRes)(err error) {
	// 关注
	if follow.ActionType == 1 {
		err = AddFollow(userID,follow.ToUserID)
	}else if follow.ActionType == 2{
		// 取消关注
		err = CancelFollow(userID,follow.ToUserID)
	}
	return  err
}

// AddFollow 关注
// 1.用户id，对方id。去关注表中设置。用户表需要设置用用户的关注总数以及被关注作者的粉丝数
// AddFollow 关注
// 1.用户id，对方id。去关注表中设置。用户表需要设置用用户的关注总数以及被关注作者的粉丝数
func AddFollow(userID,followID int64)( err error){
	//	 从redis中获取用户关注的状态
	ok,err := CurrentUserFollowTheVideoAuthorFromRedis(userID,followID)
	if err == nil  && ok{
		return nil
	}
	//	 redis中不存在该条数据,去数据库中查询
	_,err = store.CurrentUserFollowTheVideoAuthor(userID,followID)
	if err == nil{
		return
	}
	//	 创建对象
	followModel := model.Follow{
		ID: userID,
		FollowID: followID ,
		IsFollow: global.TRUE,
	}
	//  更新redis
	if err = InsertFollowToRedis(followModel);err != nil{
		return err
	}
	//  更新数据库
	if err = store.InsertFollow(followModel);err != nil {
		return
	}
	return
}

// CancelFollow 取消关注
// 1.判断用户是否在关注表中。2.删除该条关注记录
func CancelFollow(userID,followID int64) (err error) {
	//	用户关注该用户的情况
	_,err = store.CurrentUserFollowTheVideoAuthor(userID,followID)
	if err != nil {
		return
	}
//	 删除该条记录
	if err = store.DelFollowRecordByUIDAndFID(userID,followID);err != nil {
		return
	}
	return
}

// FollowList 关注列表
// 1.根据用户id获取用户关注的所有所有作者信息.
func FollowList(userID int64)(userList []response.Author,err error){
	followList,err := store.CurrentUserFollowAllVideoAuthor(userID)
	if err != nil{
		return
	}
//	生成用户id列表,查找用户信息
	userIDList := make([]int64,len(followList))
	for i,v := range followList{
		userIDList[i] = v.FollowID
	}
	userInfoList,err := store.GetUserListByID(userIDList)
	if err != nil {
		return
	}
//	生成author对象
	userList = make([]response.Author,len(userInfoList))
	for i,v := range userInfoList{
		author := response.Author{
			UserID: v.UserID,
			Name: v.UserName,
			FollowCount: v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow: global.TRUE,
		}
		userList[i] = author
	}
	return
}


func FollowerList(userID int64)(userList []response.Author,err error){
	followList,err := store.CurrentUserAllFollower(userID)
	// 查询失败,说明该用户没有粉丝
	if err != nil{
		return nil,err
	}
	//	生成用户id列表,查找用户信息
	userIDList := make([]int64,len(followList))
	for i,v := range followList{
		userIDList[i] = v.ID
	}
	userInfoList,err := store.GetUserListByID(userIDList)
	if err != nil {
		return
	}
	//	生成author对象
	userList = make([]response.Author,len(userInfoList))
	for i,v := range userInfoList{
		author := response.Author{
			UserID: v.UserID,
			Name: v.UserName,
			FollowCount: v.FollowCount,
			FollowerCount: v.FollowerCount,
			IsFollow: global.TRUE,
		}
		userList[i] = author
	}
	return
}

// CurrentUserFollowTheVideoAuthorFromRedis 从redis中查询当前用户是否关注视频作者
func CurrentUserFollowTheVideoAuthorFromRedis(userID,followID int64)(bool,error){
	// 类型转换
	uid := strconv.FormatInt(userID,10)
	fid := strconv.FormatInt(followID,10)
	key := "follow:uid:"+uid+":fid:"+fid
	strData,err := store.CurrentUserFollowTheVideoAuthorFromRedis(key)
	if err != nil {
		global.Lg.Info("redis中不存在该条数据,需要更新redis")
		return false,err
	}
	byteData := []byte(strData)
	var follow model.Follow
	// 反序列化,进行数据解析
	err = json.Unmarshal(byteData,follow)
	if err != nil {
		global.Lg.Error("反序列化失败")
		return false,err
	}
	return follow.IsFollow,err
}

func InsertFollowToRedis(follow model.Follow) error{
	followData,err := json.Marshal(&follow)
	if err != nil {
		global.Lg.Error("反序列化失败")
		return err
	}
	uid := strconv.FormatInt(follow.ID,10)
	fid := strconv.FormatInt(follow.FollowID,10)
	key := "follow:uid:"+uid+":fid:"+fid
	err = store.InsertFollowToRedis(key,followData)
	if err != nil {
		global.Lg.Error("redis插入数据失败")
		return err
	}
	return nil
}

