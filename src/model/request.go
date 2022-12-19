package model

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	ID       uint
	Nickname string
	Mobile   string
	jwt.StandardClaims
}
