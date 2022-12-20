package model

import "time"

type Performance struct {
	BaseModel
	Cover       string     `gorm:"type:varchar(160) comment '演出封面地址'"`
	Title       string     `gorm:"type:varchar(50) comment '演出名称'"`
	StartTime   *time.Time `gorm:"type:datetime comment '演出开始时间'"`
	Tags        []*Tag     `gorm:"many2many:performance_tag;"`
	TicketFiles []TicketFile
	PlacementId uint
}

type TicketFile struct {
	BaseModel
	Name          string     `gorm:"type:varchar(20) comment '票档名称'"`
	Price         int        `gorm:"type:int comment '票价'"`
	DiscountPrice *int       `gorm:"type:int comment '折扣票价'"`
	StartSellAt   *time.Time `gorm:"type:datetime comment '开售时间'"`
	StopSellAt    *time.Time `gorm:"type:datetime comment '停售时间'"`
	TicketNumber  int        `gorm:"type:int comment '总票数'"`
	PerformanceId uint
}
