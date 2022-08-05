package response

type FollowList struct {
	Response
	Author   []Author      `json:"user_list"`
}
