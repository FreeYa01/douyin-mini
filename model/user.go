package model

import "time"

// User 用户表
type User struct {
	UserID 			int64 		`gorm:"primaryKey;autoIncrement:false;comment:视频id"`
	UserName 		string      `gorm:"type:varchar(255);index:idx_usr_name;comment:用户名"`
	UserPwd  		string      `gorm:"type:varchar(255);comment:用户密码"`
	FollowCount  	int64		`gorm:"default:0;comment:关注总数"`
	FollowerCount 	int64		`gorm:"default:0;comment:粉丝总数"`
	FavoriteCount   int64        `gorm:"default:0;comment:获赞总数"`
	CreatedAt       time.Time   `gorm:"comment:创建时间"`
	UpdatedAt       time.Time   `gorm:"comment:更新时间"`
	DeletedAt      *time.Time   `gorm:"index;comment:删除时间"`
}
