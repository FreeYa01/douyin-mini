package request

type FollowRes struct {
	ToUserID   int64  `json:"to_user_id"  form:"to_user_id"`
	ActionType int64  `json:"action_type" form:"action_type"`
}
