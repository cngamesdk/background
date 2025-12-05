package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/publishing"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimPublishingChannelConfigModel struct {
	publishing.DimPublishingChannelConfigModel
}

func NewDimPublishingChannelConfigModel() *DimPublishingChannelConfigModel {
	model := &DimPublishingChannelConfigModel{}
	model.DimPublishingChannelConfigModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

type DimPublishingChannelGameConfigModel struct {
	publishing.DimPublishingChannelGameConfigModel
}

func NewDimPublishingChannelGameConfigModel() *DimPublishingChannelGameConfigModel {
	model := &DimPublishingChannelGameConfigModel{}
	model.DimPublishingChannelGameConfigModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}