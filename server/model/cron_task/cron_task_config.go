package cron_task

import (
	"github.com/cngamesdk/go-core/model/sql/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimCronTaskConfigModel struct {
	cron_task.DimCronTaskConfigModel
}

func NewDimCronTaskConfigModel() *DimCronTaskConfigModel {
	model := &DimCronTaskConfigModel{}
	model.DimCronTaskConfigModel.Db = func() *gorm.DB {
		return global.GVA_DB

	}
	return model
}

