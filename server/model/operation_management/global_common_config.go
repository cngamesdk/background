package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimGlobalCommonConfigModel struct {
	common.DimGlobalCommonConfigModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
}

func NewDimGlobalCommonConfigModel() *DimGlobalCommonConfigModel {
	model := &DimGlobalCommonConfigModel{}
	model.DimGlobalCommonConfigModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
