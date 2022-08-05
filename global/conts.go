package global

import (
	"sync"
	"time"
)

var(
	VIDEO_TITLE_MAX_LENGTH 		= 160 				// 视频最大长度
	FALSE                       = false
	TRUE                        = true
	ADDR                        = "8.142.189.85:8888/"
	TIME_FORMAT                 = "2006-01-02 15:04:05"
	COVER_JPG_FORMAT     		= ".jpg"			// 封面格式
	VIDEO_FRAME_NUM				=  1				// 截取帧数
	OPERATION_FAILED			= int64(0)			// 操作失败 = 0
	Token_OVERDUE        		= int64(10) 		// token过期
	REDIS_USERINFO_OVERDUE		= time.Hour*720     // redis过期时间
	FILE_MAX_SIZE	       	    = int64(10 << 20) 	// 上传视频大小限制为20MB
	FILE_TYPE_MAP         		sync.Map		    // 视频格式白名单
	FILE_TYPE_LIST       		 = map[string]bool{".mp4": true, ".avi": true, ".wmv": true, ".mpeg": true,
		".mov": true, ".flv": true, ".rmvb": true, ".3gb": true, ".vob": true, ".m4v": true} // 文件类型
)
