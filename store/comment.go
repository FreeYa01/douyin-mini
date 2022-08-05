package store

import (
	"douyin-mini/global"
	"douyin-mini/model"
	"douyin-mini/util/code"
	"errors"
	"gorm.io/gorm"
)


func GetCommentByCommentID(commentID int64)(comment model.Comment,err error) {
	rowsAffected := global.SqlDB.First(&comment,"comment_id = ?",commentID).RowsAffected;if rowsAffected == global.OPERATION_FAILED{
		return comment,code.ErrQuery
	}

	errors.Is(err,gorm.ErrRecordNotFound)
	return
}


func GetCommentByVID(videoID int64)(comment []model.Comment,err error) {
	rowsAffected := global.SqlDB.Find(&comment,"video_id = ?",videoID).RowsAffected;if rowsAffected == global.OPERATION_FAILED{
		return comment,code.ErrQuery
	}
	return
}
func InsertComment(comm model.Comment) error {
	rowsAffected := global.SqlDB.Create(&comm).RowsAffected;if rowsAffected == global.OPERATION_FAILED{
		return code.ErrInsert
	}
	return nil
}

func DelCommentByID(commentID int64) error {
	rowsAffected := global.SqlDB.Where("comment_id = ?",commentID).Delete(&model.Comment{}).RowsAffected;if rowsAffected == global.OPERATION_FAILED{
		return code.ErrDEl
	}
	return nil
}
