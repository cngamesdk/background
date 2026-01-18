package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// DimRootGameModel 根游戏维度
type DimRootGameModel struct {
	common.DimRootGameModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
}

func NewDimRootGameModel() *DimRootGameModel {
	model := &DimRootGameModel{}
	model.DimRootGameModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
