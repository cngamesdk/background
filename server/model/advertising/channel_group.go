package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimChannelGroupModel struct {
	advertising.DimChannelGroupModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
	MediaName    string `json:"media_name" gorm:"media_name"`
}

func NewDimChannelGroupModel() *DimChannelGroupModel {
	model := &DimChannelGroupModel{}
	model.DimChannelGroupModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
