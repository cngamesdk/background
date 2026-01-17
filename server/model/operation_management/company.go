package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimCompanyModel struct {
	common.DimCompanyModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
}

func NewDimCompanyModel() *DimCompanyModel {
	model := &DimCompanyModel{}
	model.DimCompanyModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
