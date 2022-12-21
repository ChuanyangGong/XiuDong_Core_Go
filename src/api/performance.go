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

func PerformanceModelToInfo(perfm *model.Performance, place *model.Placement) *response.PerformanceInfoRspData {
	rsp := &response.PerformanceInfoRspData{
		Cover:     perfm.Cover,
		Title:     perfm.Title,
		StartTime: *perfm.StartTime,
		BaseInfoRsp: response.BaseInfoRsp{
			UpdatedAt: perfm.UpdatedAt,
			CreatedAt: perfm.CreatedAt,
			ID:        int32(perfm.ID),
		},
	}
	if place != nil {
		rsp.Placement = PlacementModelToInfo(place)
	}
	return rsp
}

func PerformanceInfoToModel(perfm *forms.PerformanceForm) *model.Performance {
	return &model.Performance{
		BaseModel: model.BaseModel{
			ID: perfm.ID,
		},
		Cover:       perfm.Cover,
		Title:       perfm.Title,
		StartTime:   &perfm.StartTime,
		PlacementId: perfm.PlacementId,
	}
}

// 获取演出列表
// func GetPerformanceList(ctx *gin.Context) {
// 	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
// 	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
// 	performanceListRes, err := service.GetPerformanceList(&service.PageInfo{
// 		Page:     uint32(page),
// 		PageSize: uint32(pageSize),
// 	})
// 	if err != nil {
// 		zap.S().Errorf("GetPerformanceList Error: %v", err)
// 		ctx.String(http.StatusInternalServerError, "服务器内部出错")
// 		return
// 	}

// 	// 构造结果
// 	performanceListRspData := response.PerformanceListRspData{
// 		Total:      int(performanceListRes.Total),
// 		Performances: make([]response.PerformanceInfoRspData, 0),
// 	}
// 	for _, v := range *performanceListRes.Performances {
// 		performanceListRspData.Performances = append(performanceListRspData.Performances, *PerformanceModelToInfo(&v))
// 	}
// 	rsp := &response.BaseRsp{
// 		Success: true,
// 		Data:    performanceListRspData,
// 	}
// 	ctx.JSON(http.StatusOK, rsp)
// }

// 新增和更新演出
func CreateUpdatePerformance(ctx *gin.Context) {
	perfmForm := forms.PerformanceForm{}
	if err := ctx.ShouldBind(&perfmForm); err != nil {
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

	var rspData *model.Performance = nil
	var err error
	if perfmForm.ID == 0 {
		rspData, err = service.CreatePerformance(*PerformanceInfoToModel(&perfmForm))
	} else {
		rspData, err = service.UpdatePerformance(*PerformanceInfoToModel(&perfmForm))
	}

	if err != nil {
		ctx.JSON(http.StatusOK, response.BaseRsp{
			Success: false,
			Msg:     err.Error(),
		})
	} else {
		rsp := &response.BaseRsp{
			Success: true,
			Data:    PerformanceModelToInfo(rspData, nil),
		}
		ctx.JSON(http.StatusOK, rsp)
	}
}

// 删除演出
func DeletePerformance(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		zap.S().Errorf("解析演出id出错：%v\n", err)
		ctx.JSON(http.StatusBadRequest, response.BaseRsp{
			Success: false,
			Msg:     "请输入正确的演出id",
		})
		return
	}

	err = service.DeletePerformanceById(uint(id))
	if err != nil {
		zap.S().Errorf("删除演出 %d 出错：%v\n", id, err)
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
