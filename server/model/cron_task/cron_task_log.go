package cron_task

import (
	"github.com/cngamesdk/go-core/model/sql/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type OdsCronTaskLogModel struct {
	cron_task.OdsCronTaskLogModel
}

func NewOdsCronTaskLogModel() *OdsCronTaskLogModel {
	model := &OdsCronTaskLogModel{}
	model.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
