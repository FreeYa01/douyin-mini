package setting

import (
	"github.com/fsnotify/fsnotify"
	viper "github.com/spf13/viper"
	"go.uber.org/zap"
)
// mapstructure将map转为struct

type Setting struct{
	*Sys     `mapstructure:"sys"`
	*Mysql   `mapstructure:"mysql"`
	*Redis   `mapstructure:"redis"`
	*Log     `mapstructure:"log"`
	*Auth    `mapstructure:"auth"`
}
type Sys struct {
	Port 		int     `mapstructure:"port"`
	Host 		string  `mapstructure:"host"`
	StartTime   string  `mapstructure:"start-time"`
	MachineID   int64   `mapstructure:"machine-id"`
	SaveFilePath string  `mapstructure:"file-save-path"`
	CoverSavePath string  `mapstructure:"cover-save-path"`
}

type Mysql struct {
	Host 		string	`mapstructure:"host"`
	DBName 		string	`mapstructure:"db-name"`
	UserName 	string  `mapstructure:"username"`
	PassWord 	string  `mapstructure:"password"`
	Port 		 int    `mapstructure:"port"`
	MaxIdleConns int    `mapstructure:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns"`
}

type Redis struct {
	Addr		string	 `mapstructure:"addr"`
	DB 			int      `mapstructure:"db"`
	Port 		int      `mapstructure:"port"`
}

type Log struct {
	MaxSize     int      `mapstructure:"max_size"`
	MaxBackups  int		 `mapstructure:" max_backups"`
	MaxAge      int		 `mapstructure:"max_age"`
	Level 		string	 `mapstructure:"level"`
	Filename    string	 `mapstructure:"filename"`
	Prex  		string	 `mapstructure:"prex"`
	Path  		string	 `mapstructure:"path"`
}

type Auth struct {
	JwtSecret  string
}

var Set  = new(Setting)
// InitSetting 初始化
func InitSetting(filePath string) {
	//	 根据路径查找配置文件
	viper.SetConfigFile(filePath)
//	 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		zap.L().Info("读取配置文件失败",zap.Error(err))
		return
	}
	// 反序列化,将二进制数据转换为结构体对象。
	if err := viper.Unmarshal(Set); err != nil{
		zap.L().Info("配置文件序列化失败",zap.Error(err))
		return
	}
	//	开启热启动,修改配置文件会自动进行热启动
	viper.WatchConfig()
//  监控配置文件的修改
	viper.OnConfigChange(func(in fsnotify.Event){
		zap.L().Debug("配置文件被修改")
	})
}
