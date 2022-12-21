package service

import (
	"XDCore/src/global"
	"XDCore/src/model"
	"fmt"

	"go.uber.org/zap"
)

// type PerformanceListRes struct {
// 	Total      uint32
// 	Performances *[]model.Performance
// }

// 获取演出列表
// func GetPerformanceList(pageInfo *PageInfo) (*PerformanceListRes, error) {
// 	// 获取演出总数
// 	var performances []model.Performance
// 	result := global.DB.Find(&performances)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	rsp := &PerformanceListRes{}
// 	rsp.Total = uint32(result.RowsAffected)

// 	// 获取分页数据
// 	global.DB.Scopes(Paginate(int(pageInfo.Page), int(pageInfo.PageSize))).Find(&performances)
// 	rsp.Performances = &performances

// 	return rsp, nil
// }

// 新增演出
func CreatePerformance(perfm model.Performance) (*model.Performance, error) {
	result := global.DB.Create(&perfm)
	if result.Error != nil {
		zap.S().Errorf("create performance err:%v", result.Error)
		return nil, result.Error
	}
	return &perfm, nil
}

// 更新演出
func UpdatePerformance(perfm model.Performance) (*model.Performance, error) {
	result := global.DB.Updates(&perfm)
	if result.Error != nil {
		zap.S().Errorf("update performance err:%v", result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("不存在该演出")
	}
	return &perfm, nil
}

// 删除演出
func DeletePerformanceById(id uint) error {
	result := global.DB.Delete(&model.Performance{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("不存在该演出")
	}
	return nil
}
