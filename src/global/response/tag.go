package response

type TagInfoRspData struct {
	BaseInfoRsp
	Name    string `json:"name"`
	City    string `json:"city"`
	Address string `json:"address"`
}

type TagListRspData struct {
	Total int              `json:"total"`
	Tags  []TagInfoRspData `json:"tags"`
}
