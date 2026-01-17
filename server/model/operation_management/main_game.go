package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// DimMainGameModel 主游戏维度
type DimMainGameModel struct {
	common.DimMainGameModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
	RootGameName string `json:"root_game_name" gorm:"root_game_name"`
}

func NewDimMainGameModel() *DimMainGameModel {
	model := &DimMainGameModel{}
	model.DimMainGameModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
