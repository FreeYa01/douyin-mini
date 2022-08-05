package model

import "time"

type Follow struct {
	ID          int64     `gorm:"primary_key;autoIncrement:false;NOT NULL;comment:用户的标识"`
	FollowID    int64     `gorm:"primary_key;autoIncrement:false;NOT NULL;comment:所关注的用户"`
	IsFollow    bool      `gorm:"NOT NULL"`
	CreatedAt   time.Time `gorm:"comment:创建时间"`
	UpdatedAt   time.Time `gorm:"comment:更新时间"`

}