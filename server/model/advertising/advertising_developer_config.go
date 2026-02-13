package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimAdvertisingDeveloperConfigModel struct {
	advertising.DimAdvertisingDeveloperConfigModel
}

func NewDimAdvertisingDeveloperConfigModel() *DimAdvertisingDeveloperConfigModel {
	model := &DimAdvertisingDeveloperConfigModel{}
	model.DimAdvertisingDeveloperConfigModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	model.DimAdvertisingDeveloperConfigModel.AesKey = func() string {
		return global.GVA_CONFIG.Common.AesCryptKey
	}
	return model
}

func (receiver *DimAdvertisingDeveloperConfigModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.DimAdvertisingDeveloperConfigModel.AfterFind(tx)
}
