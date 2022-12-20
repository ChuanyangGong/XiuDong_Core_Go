package model

import "gorm.io/gorm"

type Order struct {
	BaseModel
	TicketFileId int `gorm:"primaryKey"`
	UserId       int `gorm:"primaryKey"`
	Number       int `gorm:"type:int comment '购票数量'"`
	Status       int `gorm:"type:int comment '购票状态：0 待支付, 1 已支付, 2 已取消'"`
}

func (Order) BeforeCreate(db *gorm.DB) error {
	err := db.SetupJoinTable(&TicketFile{}, "Users", &Order{})
	if err != nil {
		return err
	}
	return nil
}
