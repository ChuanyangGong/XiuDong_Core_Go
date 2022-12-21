package service

import (
	"strings"

	"github.com/anaskhan96/go-password-encoder"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"XDCore/src/global"
	"XDCore/src/model"
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
func GetUserList(pageInfo *PageInfo) (*UserListRes, error) {
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

// 通过 mobile 查询用户
func GetUserByMobile(mobile string) (*model.User, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		zap.S().Errorf("mobile(%s) 查询用户出错：%v", mobile, result.Error)
		return nil, result.Error
	}

	return &user, nil
}

// 通过 id 查询用户
func GetUserById(id int) (*model.User, error) {
	var user model.User
	result := global.DB.First(&user, id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		zap.S().Errorf("id(%d) 查询用户出错：%v", id, result.Error)
		return nil, result.Error
	}

	return &user, nil
}

// 创建用户
func CreateUser(user *model.User) (*model.User, error) {
	result := global.DB.Where(&model.User{Mobile: user.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	result = global.DB.Create(user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	return user, nil
}

// 检查密码是否正确
func CheckPassword(pwd string, encryptedPwd string) (bool, error) {
	passwdInfo := strings.Split(encryptedPwd, "$")
	check := password.Verify(pwd, passwdInfo[2], passwdInfo[3], global.PwdOption)
	return check, nil
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
