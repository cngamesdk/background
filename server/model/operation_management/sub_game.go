package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// DimGameModel 子游戏维度
type DimGameModel struct {
	common.DimGameModel
	GameTypeStr  string `json:"game_type_str" gorm:"-"`
	OsStr        string `json:"os_str" gorm:"-"`
	PlatformName string `json:"platform_name" gorm:"platform_name"`
	MainGameName string `json:"main_game_name" gorm:"main_game_name"`
}

func NewDimGameModel() *DimGameModel {
	model := &DimGameModel{}
	model.DimGameModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

func (receiver *DimGameModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *DimGameModel) findHook(tx *gorm.DB) (err error) {
	receiver.GameTypeStr = common.GetGameTypeName(receiver.GameType)
	receiver.OsStr = common.GetGameOsName(receiver.Os)
	return
}
