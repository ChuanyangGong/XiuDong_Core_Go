package forms

import "time"

type PerformanceForm struct {
	ID          uint      `form:"id" json:"id"`
	Cover       string    `form:"cover" json:"cover"`
	Title       string    `form:"title" json:"title" binding:"required"`
	StartTime   time.Time `form:"startTime" json:"startTime" binding:"required" time_format:"2006-01-02"`
	PlacementId uint      `form:"placementId" json:"placementId" binding:"required"`
	Tags        []uint    `form:"tags" json:"tags"`
}

type TicketFileForm struct {
	Name          string     `form:"name" json:"name" binding:"required"`
	Price         int        `form:"price" json:"price" binding:"required"`
	DiscountPrice *int       `form:"discountPrice" json:"discountPrice"`
	StartSellAt   *time.Time `form:"startSellAt" json:"startSellAt" time_format:"2006-01-02" binding:"required"`
	StopSellAt    *time.Time `form:"stopSellAt" json:"stopSellAt" time_format:"2006-01-02" binding:"required"`
	TicketNumber  int        `form:"ticketNumber" json:"ticketNumber" binding:"required"`
	Comment       string     `form:"comment" json:"comment"`
}

type TicketFiles struct {
	PerformanceId   uint             `form:"performanceId" json:"performanceId" binding:"required"`
	TicketFileForms []TicketFileForm `form:"tickets" json:"tickets"`
}
