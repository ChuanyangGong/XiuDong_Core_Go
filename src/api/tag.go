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

func TagModelToInfo(tag *model.Tag) *response.TagInfoRspData {
	return &response.TagInfoRspData{
		Name: tag.Name,
		BaseInfoRsp: response.BaseInfoRsp{
			UpdatedAt: tag.UpdatedAt,
			CreatedAt: tag.CreatedAt,
			ID:        int32(tag.ID),
		},
	}
}

func TagInfoToModel(tag *forms.TagForm) *model.Tag {
	return &model.Tag{
		BaseModel: model.BaseModel{
			ID: tag.ID,
		},
		Name: tag.Name,
	}
}

// 获取标签列表
func GetTagList(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	tagListRes, err := service.GetTagList(&service.PageInfo{
		Page:     uint32(page),
		PageSize: uint32(pageSize),
	})
	if err != nil {
		zap.S().Errorf("GetTagList Error: %v", err)
		ctx.String(http.StatusInternalServerError, "服务器内部出错")
		return
	}

	// 构造结果
	tagListRspData := response.TagListRspData{
		Total: int(tagListRes.Total),
		Tags:  make([]response.TagInfoRspData, 0),
	}
	for _, v := range *tagListRes.Tags {
		tagListRspData.Tags = append(tagListRspData.Tags, *TagModelToInfo(&v))
	}
	rsp := &response.BaseRsp{
		Success: true,
		Data:    tagListRspData,
	}
	ctx.JSON(http.StatusOK, rsp)
}

// 新增和更新标签
func CreateUpdateTag(ctx *gin.Context) {
	tagForm := forms.TagForm{}
	if err := ctx.ShouldBind(&tagForm); err != nil {
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

	var rspData *model.Tag = nil
	var err error
	if tagForm.ID == 0 {
		rspData, err = service.CreateTag(*TagInfoToModel(&tagForm))
	} else {
		rspData, err = service.UpdateTag(*TagInfoToModel(&tagForm))
	}

	if err != nil {
		ctx.JSON(http.StatusOK, response.BaseRsp{
			Success: false,
			Msg:     err.Error(),
		})
	} else {
		rsp := &response.BaseRsp{
			Success: true,
			Data:    TagModelToInfo(rspData),
		}
		ctx.JSON(http.StatusOK, rsp)
	}
}

// 删除标签
func DeleteTag(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		zap.S().Errorf("解析标签id出错：%v\n", err)
		ctx.JSON(http.StatusBadRequest, response.BaseRsp{
			Success: false,
			Msg:     "请输入正确的标签id",
		})
		return
	}

	err = service.DeleteTagById(uint(id))
	if err != nil {
		zap.S().Errorf("删除标签 %d 出错：%v\n", id, err)
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
