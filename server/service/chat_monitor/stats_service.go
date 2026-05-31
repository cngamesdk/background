package chat_monitor

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat_monitor"
)

type StatsService struct{}

type StatsOverview struct {
	TodayMessages  int64 `json:"todayMessages"`
	TodaySensitive int64 `json:"todaySensitive"`
	TodayActive    int64 `json:"todayActive"`
	TodayBans      int64 `json:"todayBans"`
	TotalGames     int64 `json:"totalGames"`
}

type TrendItem struct {
	Date           string `json:"date"`
	TotalMessages  int64  `json:"totalMessages"`
	SensitiveCount int64  `json:"sensitiveCount"`
	ActiveUsers    int64  `json:"activeUsers"`
}

type ViolatorItem struct {
	SenderUID  string `json:"senderUid"`
	SenderName string `json:"senderName"`
	AppID      string `json:"appId"`
	Count      int64  `json:"count"`
}

func (s *StatsService) GetOverview(appID string) (*StatsOverview, error) {
	data := &StatsOverview{}
	today := "CURDATE()"

	msgDB := global.GVA_DB.Model(&chat_monitor.ChatMessage{}).Where("DATE(sent_at) = " + today)
	if appID != "" {
		msgDB = msgDB.Where("app_id = ?", appID)
	}
	msgDB.Count(&data.TodayMessages)

	senDB := global.GVA_DB.Model(&chat_monitor.ChatMessage{}).Where("DATE(sent_at) = " + today + " AND is_sensitive = 1")
	if appID != "" {
		senDB = senDB.Where("app_id = ?", appID)
	}
	senDB.Count(&data.TodaySensitive)

	activeDB := global.GVA_DB.Model(&chat_monitor.ChatMessage{}).Where("DATE(sent_at) = " + today)
	if appID != "" {
		activeDB = activeDB.Where("app_id = ?", appID)
	}
	activeDB.Distinct("sender_uid").Count(&data.TodayActive)

	banDB := global.GVA_DB.Model(&chat_monitor.BanRecord{}).Where("DATE(created_at) = " + today)
	if appID != "" {
		banDB = banDB.Where("app_id = ?", appID)
	}
	banDB.Count(&data.TodayBans)

	global.GVA_DB.Model(&chat_monitor.Game{}).Where("status = 1").Count(&data.TotalGames)
	return data, nil
}

func (s *StatsService) GetTrend(appID string, days int) ([]TrendItem, error) {
	var results []TrendItem
	if days <= 0 {
		days = 7
	}

	db := global.GVA_DB.Model(&chat_monitor.DailyStats{}).Where("stat_date >= DATE_SUB(CURDATE(), INTERVAL ? DAY)", days)
	if appID != "" {
		db = db.Where("app_id = ?", appID)
	}

	err := db.Select("stat_date as date, SUM(total_messages) as total_messages, SUM(sensitive_count) as sensitive_count, SUM(active_users) as active_users").
		Group("stat_date").Order("stat_date ASC").Find(&results).Error
	return results, err
}

func (s *StatsService) GetViolators(appID string, limit int) ([]ViolatorItem, error) {
	var results []ViolatorItem
	if limit <= 0 {
		limit = 20
	}

	db := global.GVA_DB.Model(&chat_monitor.ChatMessage{}).Where("is_sensitive = 1")
	if appID != "" {
		db = db.Where("app_id = ?", appID)
	}

	err := db.Select("sender_uid, sender_name, app_id, COUNT(*) as count").
		Group("sender_uid, sender_name, app_id").Order("count DESC").Limit(limit).Find(&results).Error
	return results, err
}
