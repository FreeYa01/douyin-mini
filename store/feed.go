package store

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/util/code"
	"github.com/go-redis/redis"
	"strconv"
)

func GetVideoListByTime(latestTime int64)(videoList []model.Video,err error)  {
	err = global.SqlDB.Limit(60).Find(&videoList).Error
	if err != nil{
		return nil,err
	}
	return
}

func FeedFromRedis(nextTime int64)(videoList []string,err error)  {
	// 根据分数获取视频流列表
	videoList,err = global.RedisDB.ZRangeByScore("video:next_time",redis.ZRangeBy{
		Min: "0" ,
		Max: strconv.FormatFloat(float64(nextTime-2),'f', 3, 64),
		Offset: 0,
		Count: 30,

	}).Result()
	if err != nil ||len(videoList) == 0 {
		return videoList,code.ErrQuery
	}
	return
}
