package global

import (
	"douyin-mini/setting"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
)
var RedisDB *redis.Client
func InitRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr: setting.Set.Redis.Addr,
		Password: "",
		DB: setting.Set.Redis.DB,
	})
//	测试连通性
	_,err := RedisDB.Ping().Result()
	if err != nil {
		Lg.Error("redis连接失败",zap.Error(err))

	}
}
