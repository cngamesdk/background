package model

import (
	"fmt"
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

// 链接通用配置
func JoinProductCommonConfig(tx *gorm.DB, alias string) {
	tx.Joins(fmt.Sprintf("join %s as common_config on %s.product_config_id = common_config.id", (&operation_management.DimProductCommonConfigurationModel{}).TableName(), alias))
}
