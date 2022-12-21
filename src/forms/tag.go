package forms

type TagForm struct {
	ID   uint   `form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
}
