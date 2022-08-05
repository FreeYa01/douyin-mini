package controller

import (
	"douyin-mini/global"
	"douyin-mini/model/request"
	response2 "douyin-mini/model/response"
	"douyin-mini/service"
	"douyin-mini/util"
	"douyin-mini/util/code"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)
// Login 登录,返回token和userID
func Login(c *gin.Context)  {
	var usr request.UserRes
	// 获取用户名和密码
	if err := c.ShouldBindQuery(&usr); err != nil{
		global.Lg.Error("参数出错",zap.Error(code.ErrParamsIncorrect))
		code.SendResponse(c,err)
		return
	}
	// 查询用户是否注册,返回用户id和token
	user,err := service.Login(&usr)
	if err != nil{
		global.Lg.Error("获取id和token失败",zap.Error(err))
		code.SendResponse(c,err)
		return
	}
	//	生成token
	token,err := util.GenerationToken(user.UserID)
	if err != nil {
		code.SendResponse(c,err)
	}
	c.JSON(http.StatusOK, response2.UserResponse{
		Response: response2.Response{StatusCode: code.OK.Code, StatusMsg: code.OK.Msg},
		UserID:   user.UserID,
		Token:    token,
	})
}
// Register 注册成功返回token和userID
func Register(c *gin.Context)  {
	var usr request.UserRes
	// 获取用户名和密码
	err := c.ShouldBindQuery(&usr); if err != nil{
		global.Lg.Error("参数出错",zap.Error(code.ErrParamsIncorrect))
		return
	}
	// 注册用户信息到数据库
	user,err := service.Register(&usr)
	if err != nil{
		global.Lg.Error("用户注册失败",zap.Error(err))
		code.SendResponse(c,err)
		return
	}
	// 注册成功,生成token
	token,err := util.GenerationToken(user.UserID)
	if err != nil{
		global.Lg.Info("生成token失败")
		code.SendResponse(c,err)
	}
	c.JSON(http.StatusOK, response2.UserResponse{
		Response: response2.Response{StatusCode: code.OK.Code, StatusMsg: code.OK.Msg},
		UserID:   user.UserID,
		Token:    token,
	})

}
// UserInfo 验证用户是否授权,获取用户信息
func UserInfo(c *gin.Context)  {
	userID := c.GetInt64("userID")
	// 校验token
	err := util.VerifyToken(userID)
	if err != nil {
		code.SendResponse(c,err)
		return
	}
	if err != nil{
		global.Lg.Error("类型转换失败",zap.Error(code.ErrParamsIncorrect))
		code.SendResponse(c,code.ErrParamsIncorrect)
		return
	}

	// 查询用户信息
	user,err := service.GetUserInfo(userID)
	if err != nil{
		global.Lg.Error("获取用户信息失败",zap.Error(err))
		code.SendResponse(c,err)
		return
	}

	c.JSON(http.StatusOK, response2.UserInfoResponse{
		Response: response2.Response{StatusCode: code.OK.Code, StatusMsg: code.OK.Msg},
		User:     *user,
	})
}