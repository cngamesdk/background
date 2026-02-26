package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimAdvertisingMediaAuthModel struct {
	advertising.DimAdvertisingMediaAuthModel
}

func NewDimAdvertisingMediaAuthModel() *DimAdvertisingMediaAuthModel {
	model := &DimAdvertisingMediaAuthModel{}
	model.DimAdvertisingMediaAuthModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
