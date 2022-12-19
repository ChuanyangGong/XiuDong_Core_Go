package global

import (
	"crypto/sha512"

	"github.com/anaskhan96/go-password-encoder"
	ut "github.com/go-playground/universal-translator"
	"gorm.io/gorm"

	"XDCore/src/config"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	DB           *gorm.DB
	PwdOption    *password.Options
	Trans        ut.Translator
)

func init() {
	PwdOption = &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
}
