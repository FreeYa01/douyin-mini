package util

import (
	"douyin-mini/store"
	"douyin-mini/util/code"
)

func VerifyToken(tokenID int64) error {
	if _,err := store.GetUserByID(tokenID);err != nil{
		return code.ErrTokenInvalid
	}
	return nil
}
