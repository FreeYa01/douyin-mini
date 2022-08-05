package middleware

import (
	"bytes"
	"douyin-mini/global"
	"douyin-mini/util/code"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"path"
	"strconv"
	"strings"
)

// 获取传入字节的二进制
func bytesToHexString(src []byte) string  {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte,0)
	for _,v := range  src {
		sub := v & 0xFF
		hv  := hex.EncodeToString(append(temp,sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0),10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

// GetFileType 获取文件类型,
// 用文件前面的几个字节来进行判断，fsrc:文件字节流【仅用前面几个字节】
func GetFileType(fsrc []byte) string  {
	var fileType string
	fileCode := bytesToHexString(fsrc)
	global.FILE_TYPE_MAP.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode,strings.ToLower(k)) ||
			strings.HasPrefix(k,strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})
	return fileType
}

// FileCheck 检查文件类型
func FileCheck () gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data,err := ctx.FormFile("data")
		// 检查上传文件大小
		if data.Size >= global.FILE_MAX_SIZE {
			code.SendResponse(ctx,code.ErrUnExceededSize)
			ctx.Abort()
			return
		}
		if err != nil {
			code.SendResponse(ctx,err)
			ctx.Abort()
			return
		}
	//	 获取文件后缀
		fileSuffix := path.Ext(data.Filename)
		if _,ok := global.FILE_TYPE_LIST[fileSuffix]; !ok {
			code.SendResponse(ctx,code.ErrTypeIncorrect)
			ctx.Abort()
			return
		}
		f,err := data.Open()
		buffer := make([]byte,30)
		_,err = f.Read(buffer)
		fileType := GetFileType(buffer)

		if fileType == "" {
			code.SendResponse(ctx,code.ErrTypeIncorrect)
			ctx.Abort()
			return
		}
		// 保存文件类型
		ctx.Set("FileType",fileType)
		// 向下执行
		ctx.Next()
	}
}