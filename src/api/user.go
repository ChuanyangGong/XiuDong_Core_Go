package api

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.uber.org/zap"

	"XDCore/src/forms"
	"XDCore/src/global"
	"XDCore/src/global/request"
	"XDCore/src/global/response"
	"XDCore/src/middlewares"
	"XDCore/src/model"
	"XDCore/src/service"
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

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*model.JwtClaims)
	zap.S().Infof("用户 %d 调用获取用户列表 page: %d, pageSize: %d",
		currentUser.ID, page, pageSize)

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

	// 获取用户
	user, err := service.GetUserByMobile(pwdLoginForm.Mobile)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "用户名或密码错误",
		})
		return
	}
	ok, _ := service.CheckPassword(pwdLoginForm.Password, user.Password)
	if ok {
		// 生成 token
		j := middlewares.NewJWT()
		claims := model.JwtClaims{
			ID:       uint(user.ID),
			Nickname: user.Nickname,
			Mobile:   user.Mobile,
			StandardClaims: jwt.StandardClaims{
				NotBefore: time.Now().Unix(),
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				Issuer:    "XIUDONG_Core",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			zap.S().Errorf("token 生成失败：%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "token 生成失败",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": ok,
			"data": map[string]string{
				"token": token,
			},
			"msg": "登录成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": ok,
			"msg":     "用户名或密码错误",
		})
	}
}
