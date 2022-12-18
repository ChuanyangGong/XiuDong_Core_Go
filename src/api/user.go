package api

import (
	"XDCore/src/global/request"
	"XDCore/src/global/response"
	"XDCore/src/model"
	"XDCore/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
