package service

import (
	"XDCore/src/global"
	"XDCore/src/model"
	"fmt"

	"go.uber.org/zap"
)

type PlacementListRes struct {
	Total      uint32
	Placements *[]model.Placement
}

// 获取场地列表
func GetPlacementList(pageInfo *PageInfo) (*PlacementListRes, error) {
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

// 新增场地
func CreatePlacement(place model.Placement) (*model.Placement, error) {
	result := global.DB.Create(&place)
	if result.Error != nil {
		zap.S().Errorf("create placement err:%v", result.Error)
		return nil, result.Error
	}
	return &place, nil
}

// 更新场地
func UpdatePlacement(place model.Placement) (*model.Placement, error) {
	result := global.DB.Updates(&place)
	if result.Error != nil {
		zap.S().Errorf("update placement err:%v", result.Error)
		return nil, result.Error
	}
	return &place, nil
}

// 删除场地
func DeletePlacementById(id uint) error {
	result := global.DB.Delete(&model.Placement{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("不存在该场地")
	}
	return nil
}
