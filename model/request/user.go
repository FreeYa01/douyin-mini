package request

// UserRes 用户登录、注册信息
type UserRes struct {
	UserName  string  `json:"username"  form:"username" binding:"required"`
	UserPwd   string  `json:"password"  form:"password" binding:"required"`
}

type UserPersonInfo struct {
	UserID   int64   `json:"user_id" form:"user_id"`
	Token    string  `json:"token" form:"token"`
}

