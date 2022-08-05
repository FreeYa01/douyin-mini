package util

import (
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
)
// Encrypt 对用户密码加密,采用bcrypt加密方式
func Encrypt(pwd string) (string,error){
	encryptBytes,err := bcrypt.GenerateFromPassword([]byte(pwd),bcrypt.DefaultCost)
	if err != nil{
		return "",err
	}
	return string(encryptBytes),err
}
// Verify 验证用户密码是否正确,将当前用户输入的密码进行加密后 和原始加密后的密码编码进行比较
func Verify(originalPwd,currentPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(originalPwd),[]byte(currentPwd))
	return err == nil
}
