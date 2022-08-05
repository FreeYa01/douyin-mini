package controller

import (
	"douyin-mini/global"
	"douyin-mini/model/response"
	"douyin-mini/service"
	"douyin-mini/util/code"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Feed 视频流接口,实现刷视频功能
func Feed(c *gin.Context)  {
	// 当前时间,转换为字符串
	//currentTime := strconv.FormatInt(time.Now().UnixMilli(),10)
	// 参数为空,采用当前时间作为值
	latestTimeString := c.Query("latest_time")
	latestTime,err := strconv.ParseInt(latestTimeString,10,64)
	if err != nil {
		fmt.Println(err)
		global.Lg.Error("类型转换失败")
		code.SendResponse(c,code.ErrTypeIncorrect)
		return
	}
	// 获取用户id
	userID := c.GetInt64("userID")

	// 获取视频列表
		videoList,nextTime,err := service.Feed(userID,latestTime)
		if err != nil{
			global.Lg.Error("获取视频列表失败")
			code.SendResponse(c,code.ErrQuery)
			return
		}
		c.JSON(http.StatusOK,response.FeedList{
			Response:  response.Response{StatusCode: code.OK.Code,StatusMsg: code.OK.Msg},
			NextTime: nextTime,
			VideoList: videoList,
		})

}
