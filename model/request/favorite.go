package request

type FavoriteActionRes struct {
	VideoID    string 		`json:"video_id"`
	ActionType string        `json:"action_type"`
}
