package controller

import (
	"douyin-mini/global"
	"douyin-mini/model/response"
	"douyin-mini/service"
	"douyin-mini/util"
	"douyin-mini/util/code"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Favorite(c *gin.Context)  {
	videoID,err := strconv.ParseInt(c.Query("video_id"),10,64)
	actionType,err := strconv.ParseInt(c.Query("action_type"),10,64)
	userID := c.GetInt64("userID")
	//  校验token
	err = util.VerifyToken(userID)
	if err != nil {
		code.SendResponse(c,err)
		return
	}

	if err != nil {
		global.Lg.Error("参数错误")
		code.SendResponse(c,err)
		return
	}
	// 点赞
	err = service.Favorite(userID,videoID,actionType)

	if err != nil {
		global.Lg.Error("操作失败")
		code.SendResponse(c,err)
		return
	}

	c.JSON(http.StatusOK,response.Response{
		StatusCode: code.OK.Code,
		StatusMsg: code.OK.Msg,
	})


}

// FavoriteList 点赞列表
func FavoriteList(c *gin.Context)  {
	userID := c.GetInt64("userID")
	//  校验token
	err := util.VerifyToken(userID)
	if err != nil {
		code.SendResponse(c,err)
		return
	}
		favList,err := service.FavoriteList(userID)
		if err != nil {
			code.SendResponse(c,code.ErrQuery)
			return
		}
		c.JSON(http.StatusOK,response.VideoList{
			Response:  response.Response{StatusCode: code.OK.Code,StatusMsg: code.OK.Msg},
			VideoList: favList,
		})
}