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
	"strconv"
)


// Comment 发布评论
func Comment(c *gin.Context)  {
	var comm request.CommentRes
	err := c.ShouldBindQuery(&comm)
	if err != nil {
		global.Lg.Error("参数绑定失败")
		code.SendResponse(c,code.ErrParamsIncorrect)
		return
	}
	userID := c.GetInt64("userID")
		if err != nil {
			global.Lg.Error("类型转换失败")
			code.SendResponse(c,code.ErrTypeIncorrect)
			return
		}
	//  校验token
	err = util.VerifyToken(userID)
	if err != nil {
		code.SendResponse(c,err)
		return
	}
	// commentInfo 要响应的结果
	var commentInfo response.CommentInfo
	// 发布评论
	if comm.ActionType == 1 {
		commentInfo,err = service.AddComment(userID,comm)
	}else if comm.ActionType == 2{
	// 取消评论
		commentInfo,err = service.DelComment(comm.CommentID)
	}
	if err != nil{
		global.Lg.Error("评论操作失败")
		code.SendResponse(c,err)
		return
	}

	c.JSON(http.StatusOK,response.CommentReq{
			Response:response.Response{StatusCode: code.OK.Code,StatusMsg: code.OK.Msg},
			CommentInfo:commentInfo,
	})

}

func CommentList(c *gin.Context) {
	//  校验token
	err := util.VerifyToken(c.GetInt64("userID"))
	if err != nil {
		code.SendResponse(c,err)
		return
	}
	videoID,err := strconv.ParseInt(c.Query("video_id"),10,64)
	if err != nil {
		global.Lg.Error("类型转换失败")
		code.SendResponse(c,code.ErrTypeIncorrect)
		return
	}
	commInfoList,err := service.CommentList(videoID)
	if err != nil{
		global.Lg.Error("获取评论列表失败")
		code.SendResponse(c,err)
		return
	}
	c.JSON(http.StatusOK,response.CommentList{
		Response:response.Response{StatusCode: code.OK.Code,StatusMsg: code.OK.Msg},
		CommentList: commInfoList,
	})

}
