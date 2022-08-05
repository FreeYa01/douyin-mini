package response

// UserResponse 用户登录、注册响应结构体
type UserResponse struct {
	UserID 		int64   	`json:"user_id"`
	Token  		string 	    `json:"token"`
	Response

}
// UserInfoResponse 用户信息响应结构体
type UserInfoResponse struct {
	Response
	User UserInfo `json:"user"`
}

type UserInfo struct {
	ISFollow        bool  	`json:"is_follow" form:"is_follow"`
	UserID       	int64 	`json:"id" form:"id"`
	FollowCount  	int64  	`json:"follow_count" form:"follow_count"`
	FollowerCount   int64  	`json:"follower_count" form:"follower_count"`
	UserName       string 	`json:"name" form:"name"`
}
