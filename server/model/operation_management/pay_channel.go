package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimPayChannelModel struct {
	common.DimPayChannelModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
	CompanyName  string `json:"company_name" gorm:"company_name"`
	PayTypeName  string `json:"pay_type_name" gorm:"-"`
	StatusName   string `json:"status_name" gorm:"-"`
}

func NewDimPayChannelModel() *DimPayChannelModel {
	model := &DimPayChannelModel{}
	model.DimPayChannelModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

func (receiver *DimPayChannelModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *DimPayChannelModel) findHook(tx *gorm.DB) (err error) {
	receiver.PayTypeName = common.GetPayTypeName(receiver.PayType)
	receiver.StatusName = common.GetPayStatusName(receiver.Status)
	return
}
