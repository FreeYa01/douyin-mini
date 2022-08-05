package model

import (
	"douyin-mini/global"
	"go.uber.org/zap"
)

func InitMysqlTable()  {
	if err := global.SqlDB.AutoMigrate(&User{},&Video{},&Favorite{},&Comment{},&Follow{});err != nil {
		global.Lg.Info("数据表创建失败",zap.Error(err))
		return
	}
}
