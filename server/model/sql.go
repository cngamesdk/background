package model

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/material"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"gorm.io/gorm"
)

// 链接平台
func JoinPlatform(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as platform on %s.platform_id = platform.id", (&operation_management.DimPlatformModel{}).TableName(), alias))
}

// 链接子游戏
func JoinGame(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as game on %s.game_id = game.id", (&operation_management.DimGameModel{}).TableName(), alias))
}

// 链接主游戏
func JoinMainGame(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as main_game on %s.main_game_id = main_game.id", (&operation_management.DimMainGameModel{}).TableName(), alias))
}

// 链接根游戏
func JoinRootGame(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as root_game on %s.root_game_id = root_game.id", (&operation_management.DimRootGameModel{}).TableName(), alias))
}

// 链接主体
func JoinCompany(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as company on %s.company_id = company.id", (&operation_management.DimCompanyModel{}).TableName(), alias))
}

// 链接通用配置
func JoinProductCommonConfig(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as common_config on %s.product_config_id = common_config.id", (&operation_management.DimProductCommonConfigurationModel{}).TableName(), alias))
}

// 链接联运渠道
func JoinCoPublishing(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as channel on %s.channel_id = channel.id", (&operation_management.DimPublishingChannelConfigModel{}).TableName(), alias))
}

// 链接渠道
func JoinAgent(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as agent on %s.agent_id = agent.id", (&advertising.DimAgentModel{}).TableName(), alias))
}

// 链接广告位
func JoinSite(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as site on %s.site_id = site.id", (&advertising.DimSiteModel{}).TableName(), alias))
}

// 链接媒体
func JoinMedia(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as ad on %s.advertising_media_id = ad.id", (&advertising.DimAdvertisingMediaModel{}).TableName(), alias))
}

// 链接媒体通过媒体码
func JoinMediaByCode(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as ad on %s.code = ad.code", (&advertising.DimAdvertisingMediaModel{}).TableName(), alias))
}

// 链接渠道组
func JoinChannelGroup(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as channel_group on %s.channel_group_id = channel_group.id", (&advertising.DimChannelGroupModel{}).TableName(), alias))
}

// 链接素材题材
func JoinMaterialTheme(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as theme on %s.theme_id = theme.id", (&material.DimMaterialThemeModel{}).TableName(), alias))
}
