package model

import (
	"time"
)

type Video struct {
	VideoID 	 	int64		`gorm:"primaryKey;autoIncrement:false;comment:视频id"`
	AuthID          int64       `gorm:"index:idx_video_uid,comment:作者id"`
	VideoTitle 	 	string  	`gorm:"type:varchar(255);comment:视频标题 " `
	PlayUrl  	 	string  	`gorm:"type:varchar(255);comment:视频地址 " `
    CoverUrl  	 	string  	`gorm:"type:varchar(255);comment:封面地址 " `
	FavoriteCount 	int64     	`gorm:"default:0;comment:点赞总数 "`
	CommentCount 	int64    	`gorm:"default:0;comment:评论总数 "`
	CreatedAt       time.Time   `gorm:"comment:创建时间"`
	UpdatedAt       time.Time   `gorm:"comment:更新时间"`
	DeletedAt       *time.Time  `gorm:"index;comment:删除时间"`
}
