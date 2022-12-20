package model

type Tag struct {
	BaseModel
	Name         string         `gorm:"type:varchar(10);unique;not null"`
	Performances []*Performance `gorm:"many2many:performance_tag;"`
}
