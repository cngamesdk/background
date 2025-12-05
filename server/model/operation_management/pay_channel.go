package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimPayChannelModel struct {
	common.DimPayChannelModel
}

func NewDimPayChannelModel() *DimPayChannelModel {
	model := &DimPayChannelModel{}
	model.DimPayChannelModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
