package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer"
)

func bizModel() error {
	db := global.GVA_DB

	// AI模型中转表迁移
	err := db.AutoMigrate(
		&model_transfer.AiToken{},
		&model_transfer.AiUsageLog{},
		&model_transfer.AiDailyReport{},
	)
	if err != nil {
		return err
	}

	// 初始化AI模型中转菜单和API数据
	InitModelTransferData()

	return nil
}
