package response

type PlacementInfoRspData struct {
	BaseInfoRsp
	Name    string `json:"name"`
	City    string `json:"city"`
	Address string `json:"address"`
}

type PlacementListRspData struct {
	Total      int                    `json:"total"`
	Placements []PlacementInfoRspData `json:"placements"`
}
