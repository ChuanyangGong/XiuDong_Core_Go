package response

import (
	"time"
)

type BaseRsp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type UserInfoRspData struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Nickname  string    `json:"nickname"`
	Mobile    string    `json:"mobile"`
	Avatar    string    `json:"avatar"`
	Birthday  time.Time `json:"birthday"`
}

type UserListRspData struct {
	Total int               `json:"total"`
	Users []UserInfoRspData `json:"users"`
}

type GetUserListRsp struct {
	BaseRsp
	Data UserListRspData `json:"data"`
}
