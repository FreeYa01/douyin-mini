package response
// VideoList 发布列表
type VideoList struct {
	Response
	VideoList  [] VideoInfo `json:"video_list"`
}
