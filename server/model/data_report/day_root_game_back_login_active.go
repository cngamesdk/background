package data_report

import (
	"github.com/cngamesdk/go-core/model/sql/report"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DwsDayRootGameBackLoginActiveLogModel struct {
	report.DwsDayRootGameBackLoginActiveLogModel
}

func NewDwsDayRootGameBackLoginActiveLogModel() *DwsDayRootGameBackLoginActiveLogModel {
	model := &DwsDayRootGameBackLoginActiveLogModel{}
	model.DwsDayRootGameBackLoginActiveLogModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
