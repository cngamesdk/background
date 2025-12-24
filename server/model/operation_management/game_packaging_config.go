package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimGamePackagingConfigModel struct {
	common.DimGamePackagingConfigModel
}

func NewDimGamePackagingConfigModel() *DimGamePackagingConfigModel {
	model := &DimGamePackagingConfigModel{}
	model.DimGamePackagingConfigModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
