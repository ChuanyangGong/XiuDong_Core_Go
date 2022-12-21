package forms

type PlacementForm struct {
	ID      uint   `form:"id" json:"id"`
	Name    string `form:"name" json:"name" binding:"required"`
	City    string `form:"city" json:"city"`
	Address string `form:"address" json:"address"`
}
