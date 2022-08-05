package controller

import (
	"douyin-mini/global"
	"douyin-mini/model"
	response2 "douyin-mini/model/response"
	"douyin-mini/service"
	"douyin-mini/setting"
	"douyin-mini/util"
	"douyin-mini/util/code"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
)

// Publish 投稿视频
func Publish(c *gin.Context)  {
//	 获取参数
	title := c.PostForm("title")
	file,err := c.FormFile("data")
//	 获取作者id
	authID := c.GetInt64("userID")
//  校验token
	err = util.VerifyToken(authID)
	if err != nil {
		code.SendResponse(c,err)
		return
	}
	if err != nil{
		global.Lg.Error("获取文件失败")
		code.SendResponse(c,code.ErrParamsIncorrect)
		return
	}
	//判断视频标题长度
	if len(title) >= global.VIDEO_TITLE_MAX_LENGTH {
		global.Lg.Error("标题过长")
		code.SendResponse(c,code.ErrParamsIncorrect)
		return
	}
	//	 生成文件的ID
	id := util.GenID()
	// ID转换为字符串
	name := strconv.FormatInt(id,10)
	// 从上下文获取视频格式
	videoName := name + c.GetString("FileType")
	coverName := name + global.COVER_JPG_FORMAT
	//	获取文件的存储、封面地址
	videoPath := filepath.Join(setting.Set.Sys.SaveFilePath,videoName)
	coverPath := filepath.Join(setting.Set.Sys.CoverSavePath,coverName)
	// 将 \ 替换为 /
	videoPath = filepath.ToSlash(videoPath)
	coverPath = filepath.ToSlash(coverPath)
//	 保存视频到本地
	if err = c.SaveUploadedFile(file,videoPath);err != nil{
		global.Lg.Error("视频上传失败")
		code.SendResponse(c,code.ErrUpload)
		return
	}

//  生成视频封面
	if err = util.GetSnapshot(videoPath,coverPath, global.VIDEO_FRAME_NUM);err != nil{
		fmt.Println(err)
		global.Lg.Error("截取封面失败")
		code.SendResponse(c,code.ErrUpload)
		return
	}
	video := &model.Video{
		VideoID: id,
		AuthID: authID,
		VideoTitle: title,
		PlayUrl: global.ADDR+videoPath,
		CoverUrl: global.ADDR+coverPath,
	}
   if err = service.Publish(video);err != nil{
	   global.Lg.Error("插入数据失败")
	   code.SendResponse(c,err)
	   return
   }
//	 4.返回响应
	c.JSON(http.StatusOK, response2.Response{
		StatusCode: code.OK.Code,
		StatusMsg: code.OK.Msg,
	})
}

// PublishList 发布列表
func PublishList(ctx *gin.Context)  {
	userIDToken:= ctx.GetInt64("userID")
	//userID,_ :=  strconv.ParseInt(ctx.Query("user_id"),10,64)
	// 校验token
	err := util.VerifyToken(userIDToken)
	if err != nil {
		code.SendResponse(ctx,err)
		return
	}
	data,err := service.PublishList(userIDToken)

	if err != nil {
		global.Lg.Error("获取发布的所有视频失败")
		code.SendResponse(ctx,err)
		return
	}
	ctx.JSON(http.StatusOK, response2.VideoList{
		Response:  response2.Response{StatusCode: code.OK.Code, StatusMsg: code.OK.Msg},
		VideoList: data,
	})
}
