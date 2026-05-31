package chat_monitor

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat_monitor"
	"go.uber.org/zap"
)

type CallbackService struct{}

// CallbackPayload 回调通知载荷
type CallbackPayload struct {
	Action   string `json:"action"`
	AppID    string `json:"app_id"`
	BanType  int    `json:"ban_type"`
	Target   string `json:"target"`
	Reason   string `json:"reason"`
	Duration int    `json:"duration"`
	ExpireAt string `json:"expire_at,omitempty"`
}

// NotifyGame 通知游戏研发封禁/解封
func (s *CallbackService) NotifyGame(record chat_monitor.BanRecord, action string) {
	var game chat_monitor.Game
	if err := global.GVA_DB.Where("app_id = ?", record.AppID).First(&game).Error; err != nil {
		global.GVA_LOG.Warn("回调通知: 游戏不存在", zap.String("app_id", record.AppID))
		return
	}
	if game.CallbackURL == "" {
		global.GVA_LOG.Warn("回调通知: 游戏未配置回调地址", zap.String("app_id", record.AppID))
		return
	}

	banType := 0
	if record.BanType != nil {
		banType = *record.BanType
	}
	duration := 0
	if record.Duration != nil {
		duration = *record.Duration
	}

	payload := CallbackPayload{
		Action:   action,
		AppID:    record.AppID,
		BanType:  banType,
		Target:   record.Target,
		Reason:   record.Reason,
		Duration: duration,
		ExpireAt: record.ExpireAt,
	}

	data, _ := json.Marshal(payload)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(game.CallbackURL, "application/json", bytes.NewReader(data))
	if err != nil {
		global.GVA_LOG.Error("回调通知失败", zap.String("app_id", record.AppID), zap.Error(err))
		global.GVA_DB.Model(&record).Update("callback_status", 2)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		global.GVA_DB.Model(&record).Update("callback_status", 1)
		global.GVA_LOG.Info("回调通知成功", zap.String("app_id", record.AppID), zap.String("action", action))
	} else {
		global.GVA_DB.Model(&record).Update("callback_status", 2)
		global.GVA_LOG.Warn("回调通知返回非200", zap.String("app_id", record.AppID), zap.Int("status", resp.StatusCode))
	}
}
