package controller

import (
	"douyin-mini/global"
	"douyin-mini/model/request"
	"douyin-mini/model/response"
	"douyin-mini/service"
	"douyin-mini/util"
	"douyin-mini/util/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Follow(c *gin.Context)  {
	var follow request.FollowRes
	err := c.ShouldBindQuery(&follow)
	if err != nil {
		global.Lg.Error("参数绑定失败")
		code.SendResponse(c,code.ErrParamsIncorrect)
		return
	}

	userID := c.GetInt64("userID")
	//  校验token
	err = util.VerifyToken(userID)
	if err != nil {
		code.SendResponse(c,err)
		return
	}
	if err != nil {
		global.Lg.Error("类型转换失败")
		code.SendResponse(c,code.ErrTypeIncorrect)
		return
	}
	err = service.Follow(userID,follow)
	if err != nil{
		global.Lg.Error("关注操作失败")
		code.SendResponse(c,err)
		return
	}

	c.JSON(http.StatusOK,response.Response{
		StatusCode: code.OK.Code,
		StatusMsg: code.OK.Msg,
	})

}

// FollowList 关注列表
func FollowList(c *gin.Context)  {
	userID := c.GetInt64("userID")
	//  校验token
	err := util.VerifyToken(userID)
	if err != nil {
		code.SendResponse(c,err)
		return
	}
    userList,err := service.FollowList(userID)
	if err != nil {
		global.Lg.Error("关注操作失败")
		code.SendResponse(c,err)
		return
	}
	c.JSON(http.StatusOK,response.FollowList{
		Response: response.Response{StatusCode: code.OK.Code,StatusMsg: code.OK.Msg},
		Author:   userList,
	})
}



// FollowerList 粉丝列表
func FollowerList(c *gin.Context)  {
	userID := c.GetInt64("userID")
	userList,err := service.FollowerList(userID)
	if err != nil {
		global.Lg.Error("获取粉丝列表失败")
		code.SendResponse(c,err)
		return
	}
	c.JSON(http.StatusOK,response.FollowList{
		Response: response.Response{StatusCode: code.OK.Code,StatusMsg: code.OK.Msg},
		Author:   userList,
	})
}