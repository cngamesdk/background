package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine"
	"go.uber.org/zap"
)

// SeedActivityExamples 初始化不同类型的活动案例数据
func SeedActivityExamples(ctx context.Context) {
	activities := []activity_engine.OdsActivityConfig{
		{
			PlatformID:   0,
			GameID:       0,
			ActivityName: "【案例】累计充值648元送豪礼",
			ActivityType: "recharge",
			Description:  "活动期间累计充值达到指定金额，可领取对应档位奖励",
			StartTime:    "2026-06-01 00:00:00",
			EndTime:      "2026-06-30 23:59:59",
			TriggerConfig: `{"event_type":"recharge","conditions":[]}`,
			CalculationConfig: `{"mode":"accumulate","field":"amount","reset_cycle":"never","dedup_key":""}`,
			RewardConfig: `{"strategy":"tiered","tiers":[{"threshold":68,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":1000}]},{"threshold":128,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":3000}]},{"threshold":328,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":50},{"item_code":"exp_card","item_name":"经验卡","quantity":2}]},{"threshold":648,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":200},{"item_code":"skin_box","item_name":"皮肤礼盒","quantity":1}]}]}`,
			ConstraintConfig: `{"user_segments":["all"],"daily_claim_max":0,"total_claim_max":1,"cooldown_sec":0,"time_windows":[]}`,
			Priority:       10,
			Status:         "normal",
			GrayscaleRatio: 100,
		},
		{
			PlatformID:   0,
			GameID:       0,
			ActivityName: "【案例】7日签到送好礼",
			ActivityType: "signin",
			Description:  "连续7天签到，每天领取不同奖励，第7天领取终极大奖",
			StartTime:    "2026-06-01 00:00:00",
			EndTime:      "2026-07-31 23:59:59",
			TriggerConfig: `{"event_type":"signin","conditions":[]}`,
			CalculationConfig: `{"mode":"accumulate","field":"count","reset_cycle":"never","dedup_key":""}`,
			RewardConfig: `{"strategy":"tiered","tiers":[{"threshold":1,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":100}]},{"threshold":2,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":200}]},{"threshold":3,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":300}]},{"threshold":4,"probability":1,"items":[{"item_code":"exp_card","item_name":"经验卡","quantity":1}]},{"threshold":5,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":500}]},{"threshold":6,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":20}]},{"threshold":7,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":50},{"item_code":"skin_box","item_name":"皮肤礼盒","quantity":1}]}]}`,
			ConstraintConfig: `{"user_segments":["all"],"daily_claim_max":1,"total_claim_max":7,"cooldown_sec":0,"time_windows":[]}`,
			Priority:       8,
			Status:         "normal",
			GrayscaleRatio: 100,
		},
		{
			PlatformID:   0,
			GameID:       0,
			ActivityName: "【案例】每日登录领金币",
			ActivityType: "login",
			Description:  "每天登录游戏即可领取金币和体力奖励，每日重置",
			StartTime:    "2026-06-01 00:00:00",
			EndTime:      "2026-12-31 23:59:59",
			TriggerConfig: `{"event_type":"login","conditions":[]}`,
			CalculationConfig: `{"mode":"daily_reset","field":"count","reset_cycle":"daily","dedup_key":""}`,
			RewardConfig: `{"strategy":"fixed","tiers":[{"threshold":1,"probability":1,"items":[{"item_code":"gold","item_name":"金币","quantity":200},{"item_code":"stamina","item_name":"体力","quantity":30}]}]}`,
			ConstraintConfig: `{"user_segments":["all"],"daily_claim_max":1,"total_claim_max":0,"cooldown_sec":0,"time_windows":[]}`,
			Priority:       5,
			Status:         "normal",
			GrayscaleRatio: 100,
		},
		{
			PlatformID:   0,
			GameID:       0,
			ActivityName: "【案例】分享有礼",
			ActivityType: "share",
			Description:  "每日分享游戏到微信/QQ，即可领取钻石奖励",
			StartTime:    "2026-06-01 00:00:00",
			EndTime:      "2026-09-30 23:59:59",
			TriggerConfig: `{"event_type":"share","conditions":[]}`,
			CalculationConfig: `{"mode":"daily_reset","field":"count","reset_cycle":"daily","dedup_key":""}`,
			RewardConfig: `{"strategy":"fixed","tiers":[{"threshold":1,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":10}]}]}`,
			ConstraintConfig: `{"user_segments":["all"],"daily_claim_max":1,"total_claim_max":0,"cooldown_sec":0,"time_windows":[]}`,
			Priority:       3,
			Status:         "normal",
			GrayscaleRatio: 100,
		},
		{
			PlatformID:   0,
			GameID:       0,
			ActivityName: "【案例】首充礼包",
			ActivityType: "recharge",
			Description:  "首次充值任意金额即可领取超值礼包，终身仅限一次",
			StartTime:    "2026-01-01 00:00:00",
			EndTime:      "2027-12-31 23:59:59",
			TriggerConfig: `{"event_type":"recharge","conditions":[{"field":"amount","operator":"gte","value":1}]}`,
			CalculationConfig: `{"mode":"accumulate","field":"count","reset_cycle":"never","dedup_key":""}`,
			RewardConfig: `{"strategy":"fixed","tiers":[{"threshold":1,"probability":1,"items":[{"item_code":"diamond","item_name":"钻石","quantity":500},{"item_code":"gold","item_name":"金币","quantity":10000},{"item_code":"vip_card","item_name":"VIP体验卡(3天)","quantity":1}]}]}`,
			ConstraintConfig: `{"user_segments":["all"],"daily_claim_max":0,"total_claim_max":1,"cooldown_sec":0,"time_windows":[]}`,
			Priority:       20,
			Status:         "normal",
			GrayscaleRatio: 100,
		},
		{
			PlatformID:   0,
			GameID:       0,
			ActivityName: "【案例】周末充值翻倍(概率掉落)",
			ActivityType: "recharge",
			Description:  "周末充值有概率获得额外钻石奖励，充值越多概率越高",
			StartTime:    "2026-06-01 00:00:00",
			EndTime:      "2026-08-31 23:59:59",
			TriggerConfig: `{"event_type":"recharge","conditions":[{"field":"amount","operator":"gte","value":30}]}`,
			CalculationConfig: `{"mode":"accumulate","field":"amount","reset_cycle":"never","dedup_key":""}`,
			RewardConfig: `{"strategy":"probability","tiers":[{"threshold":30,"probability":0.3,"items":[{"item_code":"diamond","item_name":"钻石","quantity":30}]},{"threshold":98,"probability":0.5,"items":[{"item_code":"diamond","item_name":"钻石","quantity":100}]},{"threshold":298,"probability":0.8,"items":[{"item_code":"diamond","item_name":"钻石","quantity":300},{"item_code":"avatar_frame","item_name":"限定头像框","quantity":1}]}]}`,
			ConstraintConfig: `{"user_segments":["all"],"daily_claim_max":0,"total_claim_max":3,"cooldown_sec":0,"time_windows":[{"start_hour":0,"end_hour":24}]}`,
			Priority:       7,
			Status:         "normal",
			GrayscaleRatio: 50,
		},
	}

	for i := range activities {
		var count int64
		global.GVA_DB.WithContext(ctx).Model(&activity_engine.OdsActivityConfig{}).
			Where("activity_name = ?", activities[i].ActivityName).Count(&count)
		if count == 0 {
			global.GVA_DB.WithContext(ctx).Create(&activities[i])
		}
	}
	global.GVA_LOG.Info("activity examples seeded", zap.Int("count", len(activities)))
}
