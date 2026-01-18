package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimProductCommonConfigurationModel struct {
	common.DimProductCommonConfigurationModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
}

func NewDimProductCommonConfigurationModel() *DimProductCommonConfigurationModel {
	model := &DimProductCommonConfigurationModel{}
	model.DimProductCommonConfigurationModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
