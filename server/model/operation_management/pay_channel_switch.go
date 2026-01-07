package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimPayChannelSwitchModel struct {
	common.DimPayChannelSwitchModel
}

func NewDimPayChannelSwitchModel() *DimPayChannelSwitchModel {
	model := &DimPayChannelSwitchModel{}
	model.DimPayChannelSwitchModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
