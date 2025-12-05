package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimAdvertisingMediaAccountModel struct {
	advertising.DimAdvertisingMediaAccountModel
}

func NewDimAdvertisingMediaAccountModel() *DimAdvertisingMediaAccountModel {
	model := &DimAdvertisingMediaAccountModel{}
	model.DimAdvertisingMediaAccountModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}