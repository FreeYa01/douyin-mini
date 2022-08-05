package main

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/router"
	"douyin-mini/setting"
	"douyin-mini/util"
	"fmt"
	"gorm.io/gorm/logger"
	"time"
)
func main(){
	// 这里的地址可以改变成活动的，不用写死，可以优化
	setting.InitSetting("./config/config.yaml")
	// 初始化日志库`
	global.InitLog()
	// 白名单
	global.InitType()
	// 雪花算法初始化
	util.Init(setting.Set.Sys.StartTime,setting.Set.Sys.MachineID)
	//	连接数据库
	global.InitMql()
	global.SqlDB.Logger.LogMode(logger.Error)
	// 初始化表格
	model.InitMysqlTable()
	global.InitRedis()
	//	初始化路由
	r := router.InitRouter()
	fmt.Println(time.Now().UnixMilli())
	//	启动路由
	r.Run(fmt.Sprintf(":%d",setting.Set.Sys.Port))
}
