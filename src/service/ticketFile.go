package service

import (
	"XDCore/src/global"
	"XDCore/src/model"

	"gorm.io/gorm"
)

func UpdateTicketFile(tickets []model.TicketFile, perfmId uint) ([]model.TicketFile, error) {
	for idx := range tickets {
		tickets[idx].PerformanceId = perfmId
	}

	// 使用事务更新票档信息
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 删除已有的 ticket
		rst := tx.Where("performance_id = ?", perfmId).Delete(&model.TicketFile{})
		if rst.Error != nil {
			return rst.Error
		}

		// 插入新增的
		rst = tx.Create(tickets)
		if rst.Error != nil {
			return rst.Error
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tickets, nil
}
