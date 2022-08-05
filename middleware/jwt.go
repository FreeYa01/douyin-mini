package middleware

import (
	"douyin-mini/global"
	"douyin-mini/util"
	"douyin-mini/util/code"
	"github.com/gin-gonic/gin"
	"time"
)
// JwtAuthMiddleWare  jwt认证中间件
func JwtAuthMiddleWare() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token := ctx.PostForm("token")
		if token == ""{
			token = ctx.Query("token")
		}
		if token == "" {
		 code.SendResponse(ctx,code.ErrTokenInvalid)
			// 阻止继续执行
			ctx.Abort()
			return
		}
		// 解析token
		claims,err := util.ParseToken(token)
		if err != nil {
			code.SendResponse(ctx,err)
			ctx.Abort()
		}
		// 过期
		if claims.ExpiresAt - time.Now().Unix() <= global.Token_OVERDUE{
			code.SendResponse(ctx,code.ErrTokenOverDue)
			ctx.Abort()
		}
		// 存入上下文中,key为userID
		ctx.Set("userID", claims.UserID)
		ctx.Next()
	}
}