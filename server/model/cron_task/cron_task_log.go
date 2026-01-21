package cron_task

import (
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/model/sql/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type OdsCronTaskLogModel struct {
	cron_task.OdsCronTaskLogModel
	StatusName string `json:"status_name" gorm:"-"`
}

func NewOdsCronTaskLogModel() *OdsCronTaskLogModel {
	model := &OdsCronTaskLogModel{}
	model.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

func (receiver *OdsCronTaskLogModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *OdsCronTaskLogModel) findHook(tx *gorm.DB) (err error) {
	if name, ok := sql.StatusMap[receiver.Status]; ok {
		receiver.StatusName = name
	}
	return
}
