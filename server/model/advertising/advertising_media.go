package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimAdvertisingMediaModel struct {
	advertising.DimAdvertisingMediaModel
	PlatformName    string `json:"platform_name" gorm:"platform_name"`
	CommonMediaName string `json:"common_media_name" gorm:"-"`
}

func NewDimAdvertisingMediaModel() *DimAdvertisingMediaModel {
	model := &DimAdvertisingMediaModel{}
	model.DimAdvertisingMediaModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

func (receiver *DimAdvertisingMediaModel) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *DimAdvertisingMediaModel) findHook(tx *gorm.DB) (err error) {
	if name, ok := advertising.CommonMediasMap[receiver.BelongCommonMedia]; ok {
		receiver.CommonMediaName = name
	}
	return
}
