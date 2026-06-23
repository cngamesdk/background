package model_transfer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
	"time"
)

// AiDailyReport 日报表
type AiDailyReport struct {
	ID             int64     `gorm:"primarykey;autoIncrement" json:"id"`
	StatDate       string    `gorm:"type:date;not null;index:idx_stat_date" json:"statDate"`
	TokenID        int64     `gorm:"type:bigint;not null;index:idx_token_id" json:"tokenId"`
	TokenName      string    `gorm:"type:varchar(255);not null" json:"tokenName"`
	Provider       string    `gorm:"type:varchar(32);not null" json:"provider"`
	Model          string    `gorm:"type:varchar(128);not null" json:"model"`
	RequestCount   int       `gorm:"type:int;not null;default:0" json:"requestCount"`
	SuccessCount   int       `gorm:"type:int;not null;default:0" json:"successCount"`
	ErrorCount     int       `gorm:"type:int;not null;default:0" json:"errorCount"`
	TotalTokens    int64     `gorm:"type:bigint;not null;default:0" json:"totalTokens"`
	RequestTokens  int64     `gorm:"type:bigint;not null;default:0" json:"requestTokens"`
	ResponseTokens int64     `gorm:"type:bigint;not null;default:0" json:"responseTokens"`
	AvgDurationMs  int64       `gorm:"type:int;not null;default:0" json:"avgDurationMs"`
	CreatedAt      time.Time `gorm:"type:datetime;not null;autoCreateTime" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"type:datetime;not null;autoUpdateTime" json:"updatedAt"`
}

func (AiDailyReport) TableName() string {
	return "dws_ai_daily_report"
}

// NewAiDailyReportModel 创建模型实例
func NewAiDailyReportModel() *AiDailyReport {
	return &AiDailyReport{}
}

// Db 获取数据库连接
func (m *AiDailyReport) Db() *gorm.DB {
	return global.GVA_DB
}
