package data_report

import (
	"github.com/cngamesdk/go-core/model/sql/report"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DwsDayRootGameBackPayActiveLogModel struct {
	report.DwsDayRootGameBackPayActiveLogModel
}

func NewDwsDayRootGameBackPayActiveLogModel() *DwsDayRootGameBackPayActiveLogModel {
	model := &DwsDayRootGameBackPayActiveLogModel{}
	model.DwsDayRootGameBackPayActiveLogModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
