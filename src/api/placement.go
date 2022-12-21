package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"XDCore/src/global/response"
	"XDCore/src/model"
	"XDCore/src/service"
)

func PlacementModelToInfo(place *model.Placement) *response.PlacementInfoRspData {
	return &response.PlacementInfoRspData{
		Name:    place.Name,
		Address: place.Address,
		City:    place.City,
		BaseInfoRsp: response.BaseInfoRsp{
			UpdatedAt: place.UpdatedAt,
			CreatedAt: place.CreatedAt,
			ID:        int32(place.ID),
		},
	}
}

// 获取场地列表
func GetPlacementList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	placementListRes, err := service.GetPlacementListService(&service.PageInfo{
		Page:     uint32(page),
		PageSize: uint32(pageSize),
	})
	if err != nil {
		zap.S().Errorf("GetPlacementList Error: %v", err)
		ctx.String(http.StatusInternalServerError, "服务器内部出错")
		return
	}

	// 构造结果
	placementListRspData := response.PlacementListRspData{
		Total:      int(placementListRes.Total),
		Placements: make([]response.PlacementInfoRspData, 0),
	}
	for _, v := range *placementListRes.Placements {
		placementListRspData.Placements = append(placementListRspData.Placements, *PlacementModelToInfo(&v))
	}
	rsp := &response.BaseRsp{
		Success: true,
		Data:    placementListRspData,
	}
	ctx.JSON(http.StatusOK, rsp)
}
