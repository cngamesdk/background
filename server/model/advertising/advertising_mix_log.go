package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type OdsAdvertisingMixLogModel struct {
	advertising.OdsAdvertisingMixLogModel
}

func NewOdsAdvertisingMixLogModel() *OdsAdvertisingMixLogModel {
	model := &OdsAdvertisingMixLogModel{}
	model.OdsAdvertisingMixLogModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
