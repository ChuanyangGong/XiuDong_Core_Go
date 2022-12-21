package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"XDCore/src/forms"
	"XDCore/src/global"
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

func PlacementInfoToModel(place *forms.PlacementForm) *model.Placement {
	return &model.Placement{
		BaseModel: model.BaseModel{
			ID: place.ID,
		},
		City:    place.City,
		Name:    place.Name,
		Address: place.Address,
	}
}

// 获取场地列表
func GetPlacementList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	placementListRes, err := service.GetPlacementList(&service.PageInfo{
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

// 新增和更新场地
func CreateUpdatePlacement(ctx *gin.Context) {
	placementForm := forms.PlacementForm{}
	if err := ctx.ShouldBind(&placementForm); err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			ctx.JSON(http.StatusOK, response.BaseRsp{
				Success: false,
				Msg:     err.Error(),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": removeTopStruct(errs.Translate(global.Trans)),
			})
		}
		return
	}

	var rspData *model.Placement = nil
	var err error
	if placementForm.ID == 0 {
		rspData, err = service.CreatePlacement(*PlacementInfoToModel(&placementForm))
	} else {
		rspData, err = service.UpdatePlacement(*PlacementInfoToModel(&placementForm))
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	rsp := &response.BaseRsp{
		Success: true,
		Data:    PlacementModelToInfo(rspData),
	}
	ctx.JSON(http.StatusOK, rsp)
}

// 删除场地
func DeletePlacement(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		zap.S().Errorf("解析场地id出错：%v\n", err)
		ctx.JSON(http.StatusBadRequest, response.BaseRsp{
			Success: false,
			Msg:     "请输入正确的场地id",
		})
		return
	}

	err = service.DeletePlacementById(uint(id))
	if err != nil {
		zap.S().Errorf("删除场地 %d 出错：%v\n", id, err)
		ctx.JSON(http.StatusOK, response.BaseRsp{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.BaseRsp{
		Success: true,
	})
}
