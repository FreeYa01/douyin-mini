package response
// Author 作者
type Author struct {
	UserID       	int64   `json:"id"`
	FollowCount 	int64   `json:"follow_count"`
	FollowerCount 	int64   `json:"follower_count"`
	Name	 		string  `json:"name"`
	IsFollow    	 bool   `json:"is_follow"`
}
// VideoInfo 视频信息
type VideoInfo struct {
	VideoID 		int64	    `json:"id"`
	Author			UserInfo 	`json:"author"`
	Title   		string		`json:"title"`
	PlayUrl     	string		`json:"play_url"`
	CoverUrl		string		`json:"cover_url"`
	FavoriteCount	int64		`json:"favorite_count"`
	CommentCount	int64		`json:"comment_count"`
	IsFavorite      bool		`json:"isFavorite"`
}




