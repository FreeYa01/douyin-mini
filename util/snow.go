package util

import (
	"douyin-mini/global"
	sf "github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"time"
)
var node *sf.Node
func Init(startTime string, machineID int64){
	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		global.Lg.Error("雪花算法解析转换日期形式失败",zap.Error(err))
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}
