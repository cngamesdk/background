package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine"
	"go.uber.org/zap"
)

// SeedActivityData 初始化活动引擎种子数据（活动类型、模板、奖励道具示例）
func SeedActivityData(ctx context.Context) {
	seedActivityTypes(ctx)
	seedTemplates(ctx)
	seedRewardItems(ctx)
}

func seedActivityTypes(ctx context.Context) {
	types := []map[string]interface{}{
		{"type_code": "recharge", "type_name": "累计充值", "sort_order": 1, "status": "normal"},
		{"type_code": "login", "type_name": "登录奖励", "sort_order": 2, "status": "normal"},
		{"type_code": "signin", "type_name": "签到", "sort_order": 3, "status": "normal"},
		{"type_code": "share", "type_name": "分享奖励", "sort_order": 4, "status": "normal"},
		{"type_code": "custom", "type_name": "自定义活动", "sort_order": 5, "status": "normal"},
	}
	for _, t := range types {
		var count int64
		global.GVA_DB.WithContext(ctx).Table("dim_activity_type").Where("type_code = ?", t["type_code"]).Count(&count)
		if count == 0 {
			global.GVA_DB.WithContext(ctx).Table("dim_activity_type").Create(t)
		}
	}
	global.GVA_LOG.Info("activity types seeded")
}

func seedTemplates(ctx context.Context) {
	templates := []activity_engine.OdsActivityTemplate{
		{
			TemplateName:      "累计充值648奖励",
			ActivityType:      "recharge",
			Description:       "累计充值达到指定金额，领取阶梯奖励",
			TriggerConfig:     `{"event_type":"recharge","conditions":[]}`,
			CalculationConfig: `{"mode":"accumulate","field":"amount","reset_cycle":"never","dedup_key":""}`,
			RewardConfig:      `{"strategy":"tiered","tiers":[{"threshold":68,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":1000}]},{"threshold":128,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":3000}]},{"threshold":328,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":50}]},{"threshold":648,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":200},{"item_code":"skin_box","item_name":"皮肤礼盒","quantity":1}]}]}`,
			ConstraintConfig:  `{"user_segments":["all"],"daily_claim_max":0,"total_claim_max":1,"cooldown_sec":0,"time_windows":[]}`,
			Status:            "normal",
		},
		{
			TemplateName:      "7日签到",
			ActivityType:      "signin",
			Description:       "连续7天签到，每日领取不同奖励",
			TriggerConfig:     `{"event_type":"signin","conditions":[]}`,
			CalculationConfig: `{"mode":"accumulate","field":"count","reset_cycle":"never","dedup_key":""}`,
			RewardConfig:      `{"strategy":"tiered","tiers":[{"threshold":1,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":100}]},{"threshold":2,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":200}]},{"threshold":3,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":300}]},{"threshold":4,"probability":1,"items":[{"item_code":"exp_card","item_name":"经验卡","quantity":1}]},{"threshold":5,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":500}]},{"threshold":6,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":20}]},{"threshold":7,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":50},{"item_code":"skin_box","item_name":"皮肤礼盒","quantity":1}]}]}`,
			ConstraintConfig:  `{"user_segments":["all"],"daily_claim_max":1,"total_claim_max":7,"cooldown_sec":0,"time_windows":[]}`,
			Status:            "normal",
		},
		{
			TemplateName:      "每日登录奖励",
			ActivityType:      "login",
			Description:       "每天登录即可领取固定奖励，每日重置",
			TriggerConfig:     `{"event_type":"login","conditions":[]}`,
			CalculationConfig: `{"mode":"daily_reset","field":"count","reset_cycle":"daily","dedup_key":""}`,
			RewardConfig:      `{"strategy":"fixed","tiers":[{"threshold":1,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":200},{"item_code":"stamina","item_name":"体力","quantity":30}]}]}`,
			ConstraintConfig:  `{"user_segments":["all"],"daily_claim_max":1,"total_claim_max":0,"cooldown_sec":0,"time_windows":[]}`,
			Status:            "normal",
		},
		{
			TemplateName:      "分享有礼",
			ActivityType:      "share",
			Description:       "分享游戏到社交平台，每日可领取一次奖励",
			TriggerConfig:     `{"event_type":"share","conditions":[]}`,
			CalculationConfig: `{"mode":"daily_reset","field":"count","reset_cycle":"daily","dedup_key":""}`,
			RewardConfig:      `{"strategy":"fixed","tiers":[{"threshold":1,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":10}]}]}`,
			ConstraintConfig:  `{"user_segments":["all"],"daily_claim_max":1,"total_claim_max":0,"cooldown_sec":0,"time_windows":[]}`,
			Status:            "normal",
		},
		{
			TemplateName:      "首充礼包",
			ActivityType:      "recharge",
			Description:       "首次充值任意金额即可领取豪华礼包，终身仅一次",
			TriggerConfig:     `{"event_type":"recharge","conditions":[{"field":"amount","operator":"gte","value":1}]}`,
			CalculationConfig: `{"mode":"accumulate","field":"count","reset_cycle":"never","dedup_key":""}`,
			RewardConfig:      `{"strategy":"fixed","tiers":[{"threshold":1,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":500},{"item_code":"gold","item_name":"金币","quantity":10000},{"item_code":"vip_card","item_name":"VIP体验卡(3天)","quantity":1}]}]}`,
			ConstraintConfig:  `{"user_segments":["all"],"daily_claim_max":0,"total_claim_max":1,"cooldown_sec":0,"time_windows":[]}`,
			Status:            "normal",
		},
	}

	for i := range templates {
		var count int64
		global.GVA_DB.WithContext(ctx).Model(&activity_engine.OdsActivityTemplate{}).
			Where("template_name = ?", templates[i].TemplateName).Count(&count)
		if count == 0 {
			global.GVA_DB.WithContext(ctx).Create(&templates[i])
		}
	}
	global.GVA_LOG.Info("activity templates seeded", zap.Int("count", len(templates)))
}

func seedRewardItems(ctx context.Context) {
	items := []activity_engine.OdsActivityRewardItem{
		{PlatformID: 0, ItemCode: "gold", ItemName: "金币", ItemType: "currency", Status: "normal"},
		{PlatformID: 0, ItemCode: "diamond", ItemName: "钻石", ItemType: "currency", Status: "normal"},
		{PlatformID: 0, ItemCode: "stamina", ItemName: "体力", ItemType: "currency", Status: "normal"},
		{PlatformID: 0, ItemCode: "exp_card", ItemName: "经验卡", ItemType: "item", Status: "normal"},
		{PlatformID: 0, ItemCode: "skin_box", ItemName: "皮肤礼盒", ItemType: "item", Status: "normal"},
		{PlatformID: 0, ItemCode: "vip_card", ItemName: "VIP体验卡(3天)", ItemType: "item", Status: "normal"},
		{PlatformID: 0, ItemCode: "coupon_10", ItemName: "10元代金券", ItemType: "coupon", Status: "normal"},
		{PlatformID: 0, ItemCode: "coupon_50", ItemName: "50元代金券", ItemType: "coupon", Status: "normal"},
		{PlatformID: 0, ItemCode: "avatar_frame", ItemName: "限定头像框", ItemType: "item", Status: "normal"},
		{PlatformID: 0, ItemCode: "title_badge", ItemName: "专属称号", ItemType: "item", Status: "normal"},
	}

	for i := range items {
		var count int64
		global.GVA_DB.WithContext(ctx).Model(&activity_engine.OdsActivityRewardItem{}).
			Where("item_code = ? AND platform_id = ?", items[i].ItemCode, items[i].PlatformID).Count(&count)
		if count == 0 {
			global.GVA_DB.WithContext(ctx).Create(&items[i])
		}
	}
	global.GVA_LOG.Info("reward items seeded", zap.Int("count", len(items)))
}
