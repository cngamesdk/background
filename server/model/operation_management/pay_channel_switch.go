package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimPayChannelSwitchModel struct {
	common.DimPayChannelSwitchModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
	PayTypeName  string `json:"pay_type_name" gorm:"-"`
	StatusName   string `json:"status_name" gorm:"-"`
}

func NewDimPayChannelSwitchModel() *DimPayChannelSwitchModel {
	model := &DimPayChannelSwitchModel{}
	model.DimPayChannelSwitchModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

func (receiver *DimPayChannelSwitchModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *DimPayChannelSwitchModel) findHook(tx *gorm.DB) (err error) {
	receiver.PayTypeName = common.GetPayTypeName(receiver.PayType)
	if name, ok := sql.StatusMap[receiver.Status]; ok {
		receiver.StatusName = name
	}
	return
}
