package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimGamePackagingConfigModel struct {
	common.DimGamePackagingConfigModel
	PlatformName  string `json:"platform_name" gorm:"platform_name"`
	GameName      string `json:"game_name" gorm:"game_name"`
	MediaName     string `json:"media_name" gorm:"media_name"`
	StatusName    string `json:"status_name" gorm:"-"`
	UseStatusName string `json:"use_status_name" gorm:"-"`
}

func NewDimGamePackagingConfigModel() *DimGamePackagingConfigModel {
	model := &DimGamePackagingConfigModel{}
	model.DimGamePackagingConfigModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

func (receiver *DimGamePackagingConfigModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *DimGamePackagingConfigModel) findHook(tx *gorm.DB) (err error) {
	if name, ok := sql.StatusMap[receiver.Status]; ok {
		receiver.StatusName = name
	}
	if name, ok := sql.StatusMap[receiver.UseStatus]; ok {
		receiver.UseStatusName = name
	}
	return
}
