package response

import "time"

type PerformanceInfoRspData struct {
	BaseInfoRsp
	Cover     string                `json:"cover"`
	Title     string                `json:"title"`
	StartTime time.Time             `json:"startTime"`
	Placement *PlacementInfoRspData `json:"placement,omitempty"`
}

type PerformanceListRspData struct {
	Total        int                      `json:"total"`
	Performances []PerformanceInfoRspData `json:"performances"`
}
