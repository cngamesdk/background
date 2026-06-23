package model_transfer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
	"time"
)

// AiUsageLog 使用记录表
type AiUsageLog struct {
	ID             int64     `gorm:"primarykey;autoIncrement" json:"id"`
	TraceID        string    `gorm:"type:varchar(64);not null;index:idx_trace_id" json:"traceId"`
	TokenID        int64     `gorm:"type:bigint;not null;index:idx_token_id" json:"tokenId"`
	TokenName      string    `gorm:"type:varchar(255);not null" json:"tokenName"`
	Provider       string    `gorm:"type:varchar(32);not null;index:idx_provider_model" json:"provider"`
	Model          string    `gorm:"type:varchar(128);not null;index:idx_provider_model" json:"model"`
	RequestTokens  int       `gorm:"type:int;not null;default:0" json:"requestTokens"`
	ResponseTokens int       `gorm:"type:int;not null;default:0" json:"responseTokens"`
	TotalTokens    int       `gorm:"type:int;not null;default:0" json:"totalTokens"`
	RequestTime    time.Time `gorm:"type:datetime;not null;index:idx_request_time" json:"requestTime"`
	ResponseTime   time.Time `gorm:"type:datetime;not null" json:"responseTime"`
	DurationMs     int       `gorm:"type:int;not null" json:"durationMs"`
	StatusCode     int       `gorm:"type:int;not null" json:"statusCode"`
	ErrorMsg       string    `gorm:"type:text" json:"errorMsg"`
	IP             string    `gorm:"type:varchar(64)" json:"ip"`
	UserAgent      string    `gorm:"type:varchar(512)" json:"userAgent"`
	CreatedAt      time.Time `gorm:"type:datetime;not null;autoCreateTime" json:"createdAt"`
}

func (AiUsageLog) TableName() string {
	return "ods_ai_usage_log"
}

// NewAiUsageLogModel 创建模型实例
func NewAiUsageLogModel() *AiUsageLog {
	return &AiUsageLog{}
}

// Db 获取数据库连接
func (m *AiUsageLog) Db() *gorm.DB {
	return global.GVA_DB
}
