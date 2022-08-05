package model

import "time"

// Favorite 点赞表
type Favorite struct {
	UserID 		int64			`gorm:"primaryKey;autoIncrement:false;index:idx_fav_uid;priority:1;comment:用户id"`
	VideoID 	int64	    	`gorm:"primaryKey;autoIncrement:false;index:idx_fav_vid;priority:3;comment:视频id"`
	IsFavorite   bool           `gorm:"comment:是否点赞"`
	CreatedAt   time.Time   	`gorm:"comment:创建时间"`
	UpdatedAt   time.Time   	`gorm:"comment:更新时间"`
	DeletedAt  *time.Time   	`gorm:"index;comment:删除时间"`
}
