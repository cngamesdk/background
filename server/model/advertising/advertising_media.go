package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimAdvertisingMediaModel struct {
	advertising.DimAdvertisingMediaModel
}

func NewDimAdvertisingMediaModel() *DimAdvertisingMediaModel {
	model := &DimAdvertisingMediaModel{}
	model.DimAdvertisingMediaModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}