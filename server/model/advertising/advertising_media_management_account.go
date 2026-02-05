package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimAdvertisingMediaManagementAccountModel struct {
	advertising.DimAdvertisingMediaManagementAccountModel
}

func NewDimAdvertisingMediaManagementAccountModel() *DimAdvertisingMediaManagementAccountModel {
	model := &DimAdvertisingMediaManagementAccountModel{}
	model.DimAdvertisingMediaManagementAccountModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
