package global

import (
	"douyin-mini/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)
var Lg *zap.Logger
func InitLog(){
	writeSyncer := getLogWritter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder,writeSyncer,zapcore.InfoLevel)
	Lg = zap.New(core,zap.AddCaller()) // zap.AddCaller显示行号
}

// 设置日志编码
func  getEncoder() zapcore.Encoder  {
	// 获得zapcore.EncoderConfig结构体
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000") // 设置时间格式
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder // 显示文件完整路径
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按照级别显示不同颜色
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 设置日志输出位置
func getLogWritter() zapcore.WriteSyncer  {
	file, _ := os.OpenFile(setting.Set.Log.Path+setting.Set.Log.Filename,os.O_CREATE|os.O_APPEND,0755)
	//zw := zapcore.AddSync(&lumberjack.Logger{
	//	Filename:setting.Set.Log.Path+setting.Set.Log.Filename,
	//	MaxSize: setting.Set.Log.MaxSize,
	//	MaxAge: setting.Set.Log.MaxAge,
	//	MaxBackups: setting.Set.Log.MaxBackups,
	//	Compress: false,
	//})
	return zapcore.AddSync(file)
}


