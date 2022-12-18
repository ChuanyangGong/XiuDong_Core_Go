package global

import (
	"XDCore/src/config"
	"crypto/sha512"

	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/gorm"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	DB           *gorm.DB
	PwdOption    *password.Options
)

func init() {
	PwdOption = &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
}
