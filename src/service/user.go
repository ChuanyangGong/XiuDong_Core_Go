package service

import (
	"XDCore/src/global"
	"XDCore/src/model"

	"gorm.io/gorm"
)

// 函数传参
type PageInfo struct {
	Page     uint32
	PageSize uint32
}

type UserListRes struct {
	Total uint32
	Users *[]model.User
}

// 获取用户列表
func GetUserListService(pageInfo *PageInfo) (*UserListRes, error) {
	// 获取总用户个数
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &UserListRes{}
	rsp.Total = uint32(result.RowsAffected)

	// 获取分页数据
	global.DB.Scopes(Paginate(int(pageInfo.Page), int(pageInfo.PageSize))).Find(&users)
	rsp.Users = &users

	return rsp, nil
}

// 构造分页器
func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
