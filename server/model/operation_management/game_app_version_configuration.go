package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimGameAppVersionConfiguration struct {
	common.DimGameAppVersionConfiguration
}

func NewDimGameAppVersionConfiguration() *DimGameAppVersionConfiguration {
	model := &DimGameAppVersionConfiguration{}
	model.DimGameAppVersionConfiguration.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
