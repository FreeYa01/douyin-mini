package response

type FeedList struct {
	Response
	NextTime   int64      `json:"next_time"`
	VideoList []VideoInfo `json:"video_list"`
}
