package forms

import "time"

type PerformanceForm struct {
	ID          uint      `form:"id" json:"id"`
	Cover       string    `form:"cover" json:"cover"`
	Title       string    `form:"title" json:"title" binding:"required"`
	StartTime   time.Time `form:"startTime" json:"startTime" binding:"required" time_format:"2006-01-02"`
	PlacementId uint      `form:"placementId" json:"placementId" binding:"required"`
}
