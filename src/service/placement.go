package service

import (
	"XDCore/src/global"
	"XDCore/src/model"
)

type PlacementListRes struct {
	Total      uint32
	Placements *[]model.Placement
}

func GetPlacementListService(pageInfo *PageInfo) (*PlacementListRes, error) {
	// 获取场地总数
	var placements []model.Placement
	result := global.DB.Find(&placements)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &PlacementListRes{}
	rsp.Total = uint32(result.RowsAffected)

	// 获取分页数据
	global.DB.Scopes(Paginate(int(pageInfo.Page), int(pageInfo.PageSize))).Find(&placements)
	rsp.Placements = &placements

	return rsp, nil
}
