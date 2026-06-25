package task

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

func DailyReport(db *gorm.DB) error {
	ctx := context.Background()

	// 计算前一天的日期
	yesterday := time.Now().AddDate(0, 0, -1)
	statDate := yesterday.Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02") + " 23:59:59"

	global.GVA_LOG.Info("开始生成日报", zap.String("statDate", statDate))

	// 从使用日志表聚合数据
	type DailyStats struct {
		StatDate      string
		TokenID       int64
		TokenName     string
		Provider      string
		Model         string
		RequestCount  int
		SuccessCount  int
		ErrorCount    int
		TotalTokens   int64
		AvgDurationMs float32
	}

	var stats []DailyStats

	// 聚合查询
	sql := `
		SELECT
			date(l.request_time) AS stat_date,
			l.token_id,
			t.name as token_name,
			l.provider,
			l.model,
			COUNT(*) as request_count,
			SUM(CASE WHEN l.status_code = 200 THEN 1 ELSE 0 END) as success_count,
			SUM(CASE WHEN l.status_code != 200 THEN 1 ELSE 0 END) as error_count,
			COALESCE(SUM(l.total_tokens), 0) as total_tokens,
			COALESCE(AVG(l.duration_ms), 0) as avg_duration_ms
		FROM ods_ai_usage_log l
		LEFT JOIN dim_ai_token t ON l.token_id = t.id
		WHERE l.request_time BETWEEN ? AND ?
		GROUP BY stat_date, l.token_id, t.name, l.provider, l.model
	`

	if err := db.WithContext(ctx).Raw(sql, statDate, endDate).Scan(&stats).Error; err != nil {
		global.GVA_LOG.Error("查询日报数据失败", zap.Error(err))
		return err
	}

	global.GVA_LOG.Info("查询到日报数据", zap.Int("count", len(stats)))

	// 批量插入或更新日报数据
	for _, stat := range stats {
		report := model_transfer.AiDailyReport{
			StatDate:      stat.StatDate,
			TokenID:       stat.TokenID,
			TokenName:     stat.TokenName,
			Provider:      stat.Provider,
			Model:         stat.Model,
			RequestCount:  stat.RequestCount,
			SuccessCount:  stat.SuccessCount,
			ErrorCount:    stat.ErrorCount,
			TotalTokens:   stat.TotalTokens,
			AvgDurationMs: cast.ToInt64(stat.AvgDurationMs),
		}

		// 使用 ON DUPLICATE KEY UPDATE 或 Upsert
		// 先检查是否存在
		var existing model_transfer.AiDailyReport
		err := db.WithContext(ctx).
			Where("DATE(stat_date) = ? AND token_id = ? AND provider = ? AND model = ?",
				stat.StatDate, stat.TokenID, stat.Provider, stat.Model).
			First(&existing).Error

		if err == nil {
			// 记录已存在，更新
			report.ID = existing.ID
			report.CreatedAt = existing.CreatedAt
			if err := db.WithContext(ctx).Save(&report).Error; err != nil {
				global.GVA_LOG.Error("更新日报数据失败",
					zap.Int64("tokenId", stat.TokenID),
					zap.String("provider", stat.Provider),
					zap.String("model", stat.Model),
					zap.Error(err))
				return err
			}
			global.GVA_LOG.Debug("更新日报数据成功",
				zap.Int64("tokenId", stat.TokenID),
				zap.String("provider", stat.Provider),
				zap.String("model", stat.Model))
		} else {
			// 记录不存在，创建新记录
			if err := db.WithContext(ctx).Create(&report).Error; err != nil {
				global.GVA_LOG.Error("创建日报数据失败",
					zap.Int64("tokenId", stat.TokenID),
					zap.String("provider", stat.Provider),
					zap.String("model", stat.Model),
					zap.Error(err))
				return err
			}
			global.GVA_LOG.Debug("创建日报数据成功",
				zap.Int64("tokenId", stat.TokenID),
				zap.String("provider", stat.Provider),
				zap.String("model", stat.Model))
		}
	}

	global.GVA_LOG.Info("日报生成完成",
		zap.String("statDate", statDate),
		zap.Int("count", len(stats)))
	return nil
}
