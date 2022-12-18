package api

import (
	"XDCore/src/forms"
	"XDCore/src/global"
	"XDCore/src/global/request"
	"XDCore/src/global/response"
	"XDCore/src/model"
	"XDCore/src/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

func UserModelToInfo(user *model.User) *response.UserInfoRspData {
	return &response.UserInfoRspData{
		Nickname: user.Nickname,
		Mobile:   user.Mobile,
		Birthday: *user.Birthday,
		Avatar:   user.Avatar,
	}
}

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// 获取用户列表
func GetUserList(ctx *gin.Context) {
	// 构造请求
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	req := &request.GetUserListReq{
		Page:     uint32(page),
		PageSize: uint32(pageSize),
	}

	usrListRes, err := service.GetUserListService(&service.PageInfo{
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		zap.S().Errorf("GetUserListService Error: %v", err)
		ctx.String(http.StatusInternalServerError, "服务器内部出错")
		return
	}

	// 构造结果
	rsp := &response.GetUserListRsp{
		BaseRsp: response.BaseRsp{
			Success: true,
		},
		Data: response.UserListRspData{
			Total: int(usrListRes.Total),
			Users: make([]response.UserInfoRspData, 0),
		},
	}
	for _, v := range *usrListRes.Users {
		rsp.Data.Users = append(rsp.Data.Users, *UserModelToInfo(&v))
	}
	ctx.JSON(http.StatusOK, rsp)
}

// 登录
func PasswordLogin(ctx *gin.Context) {
	pwdLoginForm := forms.PasswordLoginForm{}
	if err := ctx.ShouldBind(&pwdLoginForm); err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": removeTopStruct(errs.Translate(global.Trans)),
			})
		}
		return
	}
}
