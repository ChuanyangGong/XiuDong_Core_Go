package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"XDCore/src/global"
	"XDCore/src/model"
)

func main() {
	dsn := "root:openGauss_123@tcp(60.205.227.224)/CY2212_DMCore?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	// 全局配置 logger
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	// 创建 User 表
	_ = db.AutoMigrate(&model.User{})

	// 导入 100 个用户
	salt, encodedPwd := password.Encode("root123456", global.PwdOption)
	pwd := fmt.Sprintf("$pdkdf2-sha512$%s$%s", salt, encodedPwd)
	users := make([]*model.User, 0, 128)
	for i := 1; i <= 100; i += 1 {
		now := time.Now()
		user := model.User{
			Nickname: fmt.Sprintf("用户%d", i),
			Mobile:   fmt.Sprintf("13258262%d", 300+i),
			Password: pwd,
			Birthday: &now,
		}
		users = append(users, &user)
	}
	db.Create(&users)
}
