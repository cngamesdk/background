package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/publishing"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimPublishingChannelConfigModel struct {
	publishing.DimPublishingChannelConfigModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
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
	PlatformName string `json:"platform_name" gorm:"platform_name"`
	GameName     string `json:"game_name" gorm:"game_name"`
	ChannelName  string `json:"channel_name" gorm:"channel_name"`
	AgentName    string `json:"agent_name" gorm:"agent_name"`
	SiteName     string `json:"site_name" gorm:"site_name"`
}

func NewDimPublishingChannelGameConfigModel() *DimPublishingChannelGameConfigModel {
	model := &DimPublishingChannelGameConfigModel{}
	model.DimPublishingChannelGameConfigModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
