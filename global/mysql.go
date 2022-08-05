package global

import (
	"douyin-mini/setting"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var SqlDB *gorm.DB

// InitMql 初始化数据库连接
func InitMql() {
	var err error
//  获取配置文件中对应mq对象
	mq := setting.Set.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",mq.UserName,mq.PassWord,mq.Host,mq.Port,mq.DBName)
	SqlDB,err = gorm.Open(mysql.Open(dsn),&gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,	// 指定表名为单数
		},
	})
	if err != nil {
		Lg.Info("mysql连接失败",zap.Error(err))
	}
}
// gorm2.0版本取消了关闭连接的机制,会自动关闭

