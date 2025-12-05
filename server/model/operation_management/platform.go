package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimPlatformModel struct {
	common.DimPlatformModel
}

func NewDimPlatformModel() *DimPlatformModel {
	model := &DimPlatformModel{}
	model.DimPlatformModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
