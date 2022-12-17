package model

import (
	"time"
)

type User struct {
	BaseModel
	Nickname string     `gorm:"type:varchar(20) comment '用户昵称';default:'用户昵称'"`
	Mobile   string     `gorm:"type:varchar(11);unique;not null;index:idx_mobile"`
	Password string     `gorm:"type:varchar(100);not null"`
	Avatar   string     `gorm:"type:varchar(200);"`
	Birthday *time.Time `gorm:"type:datetime"`
}
