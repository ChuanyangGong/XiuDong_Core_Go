package response

import (
	"time"
)

type UserInfoRspData struct {
	BaseInfoRsp
	Nickname string    `json:"nickname"`
	Mobile   string    `json:"mobile"`
	Avatar   string    `json:"avatar"`
	Birthday time.Time `json:"birthday"`
}

type UserListRspData struct {
	Total int               `json:"total"`
	Users []UserInfoRspData `json:"users"`
}
