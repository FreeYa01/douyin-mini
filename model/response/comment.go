package response

type CommentReq struct {
	Response
	CommentInfo  `json:"comment"`
}

type CommentList struct {
	Response
	CommentList []CommentInfo `json:"comment_list"`
}

// CommentInfo 评论信息
type  CommentInfo struct {
	ID   		int64   `json:"id"`
	Author              `json:"user"`
	Content 	string  `json:"content"`
	CreateDate 	string  `json:"Create_date"`
}
