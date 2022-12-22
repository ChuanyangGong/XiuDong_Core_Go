package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"XDCore/src/forms"
	"XDCore/src/global"
	"XDCore/src/global/response"
	"XDCore/src/model"
	"XDCore/src/service"
)

func TicketInfoToModel(place *forms.TicketFileForm) *model.TicketFile {
	return &model.TicketFile{
		Name:          place.Name,
		Price:         place.Price,
		DiscountPrice: place.DiscountPrice,
		StartSellAt:   place.StartSellAt,
		StopSellAt:    place.StopSellAt,
		TicketNumber:  place.TicketNumber,
		Comment:       place.Comment,
	}
}

// 新增和更新场地
func CreateUpdateTicket(ctx *gin.Context) {
	tickets := forms.TicketFiles{}
	if err := ctx.ShouldBind(&tickets); err != nil {
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

	ticketModels := []model.TicketFile{}
	for _, ticket := range tickets.TicketFileForms {
		ticketModels = append(ticketModels, *TicketInfoToModel(&ticket))
	}
	ticketModels, err := service.UpdateTicketFile(ticketModels, tickets.PerformanceId)
	if err != nil {
		zap.S().Errorf("更新票档信息失败（%v）：%v", tickets, err)
		ctx.JSON(http.StatusOK, response.BaseRsp{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response.BaseRsp{
		Success: true,
	})
	return
}
