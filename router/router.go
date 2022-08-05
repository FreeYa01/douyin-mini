package router

import (
	"douyin-mini/controller"
	"douyin-mini/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)
func InitRouter() *gin.Engine{
	r := gin.Default()

	api := r.Group("/douyin")
//   用户
	userApi := api.Group("/user")
	{
	    userApi.POST("/register", controller.Register)
		userApi.POST("/login", controller.Login)
	    userApi.GET("/", controller.UserInfo).Use(middleware.JwtAuthMiddleWare())
	}


//	点赞
	favoriteApi := api.Group("/favorite").Use(middleware.JwtAuthMiddleWare())
	{
		favoriteApi.POST("/action", controller.Favorite)
		favoriteApi.GET("/list", controller.FavoriteList)
	}

//	评论
	commentApi := api.Group("/comment").Use(middleware.JwtAuthMiddleWare())
	{
		commentApi.POST("/action", controller.Comment)
		commentApi.GET("/list", controller.CommentList)
	}

//	关注
	relationApi := api.Group("/relation").Use(middleware.JwtAuthMiddleWare())
	{
		relationApi.POST("/action", controller.Follow)
		relationApi.GET("/follow/list", controller.FollowList)
		relationApi.GET("/follower/list", controller.FollowerList)
	}

	//	视频
	api.GET("/feed/", controller.Feed).Use(middleware.JwtAuthMiddleWare())
	api.GET("/publish/list", controller.PublishList).Use(middleware.JwtAuthMiddleWare())
	api.POST("/publish/action/", controller.Publish,middleware.FileCheck(),middleware.JwtAuthMiddleWare())

	return r
}

/***
	不能重复点赞
*/
