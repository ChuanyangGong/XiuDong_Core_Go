package middlewares

import (
	"XDCore/src/global"
	"XDCore/src/model"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("x-token")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "请先登录",
			})
			ctx.Abort()
			return
		}
		j := NewJWT()
		// 解析 token
		claims, err := j.ParseToken(token)
		if err != nil {
			msg := "未登录"
			if err == TokenExpired {
				msg = "授权已过期"
			}

			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": msg,
			})
			ctx.Abort()
			return
		}
		ctx.Set("claims", claims)
		ctx.Set("userId", claims.ID)
		ctx.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not activate yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Can't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.JwtConfig.SigningKey),
	}
}

func (j *JWT) CreateToken(claims model.JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*model.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.JwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
