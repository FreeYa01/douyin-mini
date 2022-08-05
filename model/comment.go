package model

import (
	"time"
)

// Comment 评论列表

type Comment struct {
	CommentID   int64           `gorm:"primaryKey;NOT NULL;comment:comment_id"`
	UserID 		int64			`gorm:"index:idx_com_uid;priority:1;comment:用户id"`
	VideoID 	int64			`gorm:"index:idx_com_vid;priority:2;comment:视频id"`
	UserName 	string  		`gorm:"type:varchar(255);index:idx_com_name;priority:3;comment:用户名称"`
	Content  	string  		`gorm:"comment:评论内容"`
	CreatedAt   time.Time   	`gorm:"comment:创建时间"`
	UpdatedAt   time.Time   	`gorm:"comment:更新时间"`
	DeletedAt  *time.Time   	`gorm:"index;comment:删除时间"`
}
