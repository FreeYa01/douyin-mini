package code

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Errno  定义错误码
type Errno struct {
	Code int
	Msg string
	Err error
}

// Response 响应
type Response struct {
	Code    int         `json:"status_code"`
	Message string      `json:"status_msg"`
}

// Err 定义错误
type Err struct {
	Code int       // 错误码
	Msg  string   //展示给用户看的信息
	Errord error //保留内部错误信息
}

// 实现错误接口
func (err Errno) Error() string{
	return  err.Msg
}

func (err *Err) Error() string{
	return fmt.Sprintf("Err-code:%d,message:%s,error:%s",err.Code,err.Msg,err.Errord)
}

// DecodeErr 解码错误,获取code和message
func DecodeErr(err error)(int,string) {
	if err == nil{
		return OK.Code,OK.Msg
	}
	switch typed := err.(type) {
	 case *Err:
		 if typed.Code == ErrBind.Code{
			 typed.Msg = typed.Msg +" 具体是 " + typed.Errord.Error()
		 }
		 return typed.Code,typed.Msg
	case *Errno:
		return typed.Code,typed.Msg
	}
	return InternalServerError.Code,err.Error()
}

// SendResponse 统一封装返回结果集
func SendResponse(c *gin.Context, err error) {
	code, message := DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}


