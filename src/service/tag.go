package service

import (
	"XDCore/src/global"
	"XDCore/src/model"
	"fmt"

	"go.uber.org/zap"
)

type TagListRes struct {
	Total uint32
	Tags  *[]model.Tag
}

// 获取标签列表
func GetTagList(pageInfo *PageInfo) (*TagListRes, error) {
	// 获取标签总数
	var tags []model.Tag
	result := global.DB.Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &TagListRes{}
	rsp.Total = uint32(result.RowsAffected)

	// 获取分页数据
	global.DB.Scopes(Paginate(int(pageInfo.Page), int(pageInfo.PageSize))).Find(&tags)
	rsp.Tags = &tags

	return rsp, nil
}

// 新增标签
func CreateTag(place model.Tag) (*model.Tag, error) {
	result := global.DB.Create(&place)
	if result.Error != nil {
		zap.S().Errorf("create tag err:%v", result.Error)
		return nil, result.Error
	}
	return &place, nil
}

// 更新标签
func UpdateTag(place model.Tag) (*model.Tag, error) {
	result := global.DB.Updates(&place)
	if result.Error != nil {
		zap.S().Errorf("update tag err:%v", result.Error)
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("不存在该标签")
	}
	return &place, nil
}

// 删除标签
func DeleteTagById(id uint) error {
	result := global.DB.Delete(&model.Tag{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("不存在该标签")
	}
	return nil
}
