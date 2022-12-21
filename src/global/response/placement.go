package response

import "time"

type StandardRsp struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type BaseInfoRsp struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PlacementListRspData struct {
	Total      int                    `json:"total"`
	Placements []PlacementInfoRspData `json:"placements"`
}

type PlacementInfoRspData struct {
	BaseInfoRsp
	Name    string `json:"name"`
	City    string `json:"city"`
	Address string `json:"address"`
}
