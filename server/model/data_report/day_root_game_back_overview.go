package data_report

import (
	"github.com/cngamesdk/go-core/model/sql/report"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DwsDayRootGameBackOverviewLogModel struct {
	report.DwsDayRootGameBackOverviewLogModel
}

func NewDwsDayRootGameBackOverviewLogModel() *DwsDayRootGameBackOverviewLogModel {
	model := &DwsDayRootGameBackOverviewLogModel{}
	model.DwsDayRootGameBackOverviewLogModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
