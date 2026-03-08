package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type OdsAdvertisingOceanengineAppLogModel struct {
	advertising.OdsAdvertisingOceanengineAppLogModel
}

func NewOdsAdvertisingOceanengineAppLogModel() *OdsAdvertisingOceanengineAppLogModel {
	model := &OdsAdvertisingOceanengineAppLogModel{}
	model.OdsAdvertisingOceanengineAppLogModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
