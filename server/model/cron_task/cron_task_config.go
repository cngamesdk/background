package cron_task

import (
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/model/sql/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimCronTaskConfigModel struct {
	cron_task.DimCronTaskConfigModel
	StatusName        string `json:"status_name" gorm:"-"`
	TaskTypeName      string `json:"task_type_name" gorm:"-"`
	ExecutionModeName string `json:"execution_mode_name" gorm:"-"`
}

func NewDimCronTaskConfigModel() *DimCronTaskConfigModel {
	model := &DimCronTaskConfigModel{}
	model.DimCronTaskConfigModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

func (receiver *DimCronTaskConfigModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *DimCronTaskConfigModel) findHook(tx *gorm.DB) (err error) {
	if name, ok := sql.StatusMap[receiver.Status]; ok {
		receiver.StatusName = name
	}
	receiver.TaskTypeName = cron_task.GetTaskTypeName(receiver.TaskType)
	receiver.ExecutionModeName = cron_task.GetExecutionModeName(receiver.ExecutionMode)
	return
}
