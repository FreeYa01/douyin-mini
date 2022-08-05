package code
var (
	OK 							= &Errno{Code: 0,   Msg: "OK"}

	//系统错误, 前缀为 100
	InternalServerError 		= &Errno{Code: 10001,   Msg: "内部服务器错误"}
	ErrBind            		    = &Errno{Code: 10002,   Msg: "请求参数错误"}
	ErrTokenSign                = &Errno{Code: 10003,   Msg: "签名 jwt 时发生错误"}
	ErrTokenAnalysis			= &Errno{Code: 10004,   Msg: "解析token时发生错误"}
	ErrTokenOverDue				= &Errno{Code: 10005,   Msg: "token已过期"}
	ErrEncrypt                  = &Errno{Code: 10006,   Msg: "加密用户密码时发生错误"}
	ErrUpload                   = &Errno{Code: 10007,   Msg: "上传失败"}



	// 数据库错误, 前缀为 201
	ErrDatabase                 = &Errno{Code: 20100,   Msg: "数据库错误"}
	ErrFill                     = &Errno{Code: 20101,   Msg: "从数据库填充 struct 时发生错误"}
	ErrInsert                    = &Errno{Code: 20102,   Msg: "插入失败"}
	ErrQuery					= &Errno{Code: 20103, Msg:"查询失败"}
	ErrUpdate					= &Errno{Code: 20104, Msg:"更新失败"}
	ErrDEl						= &Errno{Code: 20105, Msg:"删除失败"}


	// 认证错误, 前缀是 202
	ErrValidation               = &Errno{Code: 20201,   Msg: "验证失败"}
	ErrTokenInvalid             = &Errno{Code: 20202,   Msg: "jwt 是无效的"}

	// 用户错误, 前缀为 203
	ErrUserNotExist         	= &Errno{Code: 20301,   Msg: "用户不存在"}
	ErrUserExist			    = &Errno{Code: 20302,   Msg: "用户已存在"}
	ErrReigster      			= &Errno{Code: 20303,   Msg: "注册失败错误"}
	ErrPasswordIncorrect        = &Errno{Code: 20304,   Msg: "密码错误"}

//	 参数错误,前缀为204
	ErrParamsIncorrect	   		= &Errno{Code: 20401,   Msg: "参数错误"}

//	 操作错误,前缀为205
	ErrUnExceededSize           = &Errno{Code:205001,   Msg:"超出最大限制"}
	ErrTypeIncorrect            = &Errno{Code:205002,   Msg: "类型不匹配"}

)