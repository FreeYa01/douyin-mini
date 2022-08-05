package request
type CommentRes struct {
	VideoID     int64      `json:"video_id"     form:"video_id"`
	ActionType  int64      `json:"action_type"  form:"action_type"`
	CommentID   int64      `json:"comment_id"   form:"comment_id"`
	CommentText string     `json:"comment_text" form:"comment_text"`
}
