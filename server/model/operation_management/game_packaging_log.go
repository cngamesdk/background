package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/model/sql/log"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type OdsGamePackagingLogModel struct {
	log.OdsGamePackagingLogModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
	GameName     string `json:"game_name" gorm:"game_name"`
	AgentName    string `json:"agent_name" gorm:"agent_name"`
	SiteName     string `json:"site_name" gorm:"site_name"`
	StatusName   string `json:"status_name" gorm:"status_name"`
}

func NewOdsGamePackagingLogModel() *OdsGamePackagingLogModel {
	model := &OdsGamePackagingLogModel{}
	model.OdsGamePackagingLogModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

func (receiver *OdsGamePackagingLogModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *OdsGamePackagingLogModel) findHook(tx *gorm.DB) (err error) {
	if name, ok := sql.StatusMap[receiver.Status]; ok {
		receiver.StatusName = name
	}
	return
}
