package model

type Placement struct {
	BaseModel
	Name         string `gorm:"type:varchar(50)"`
	City         string `gorm:"type:varchar(20)"`
	Address      string `gorm:"type:varchar(160)"`
	Performances []Performance
}
