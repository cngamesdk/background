package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer"
)

// DailyReportReq 日报查询请求
type DailyReportReq struct {
	request.PageInfo
	TokenName string `json:"tokenName" form:"tokenName"` // Token名称（模糊查询）
	StartDate string `json:"startDate" form:"startDate"` // 开始日期
	EndDate   string `json:"endDate" form:"endDate"`     // 结束日期
	Provider  string `json:"provider" form:"provider"`   // 提供商
	Model     string `json:"model" form:"model"`         // 模型
}

// DailyReportResp 日报响应
type DailyReportResp struct {
	model_transfer.AiDailyReport
	SuccessRate    float64 `json:"successRate" gorm:"-"`    // 成功率
}

// TokenUsageReq Token使用详情请求
type TokenUsageReq struct {
	request.PageInfo
	TokenID   int64  `json:"tokenId" form:"tokenId" binding:"required"`
	StartDate string `json:"startDate" form:"startDate"`
	EndDate   string `json:"endDate" form:"endDate"`
}

// TokenUsageResp Token使用详情响应
type TokenUsageResp struct {
	TraceID        string `json:"traceId"`
	Provider       string `json:"provider"`
	Model          string `json:"model"`
	TotalTokens    int    `json:"totalTokens"`
	RequestTime    string `json:"requestTime"`
	DurationMs     int    `json:"durationMs"`
	StatusCode     int    `json:"statusCode"`
	ErrorMsg       string `json:"errorMsg"`
}

// SummaryReportReq 汇总报表请求
type SummaryReportReq struct {
	StartDate string `json:"startDate" form:"startDate" binding:"required"`
	EndDate   string `json:"endDate" form:"endDate" binding:"required"`
}

// SummaryReportResp 汇总报表响应
type SummaryReportResp struct {
	TotalRequests  int64   `json:"totalRequests"`  // 总请求数
	SuccessCount   int64   `json:"successCount"`   // 成功次数
	ErrorCount     int64   `json:"errorCount"`     // 失败次数
	TotalTokens    int64   `json:"totalTokens"`    // 总Token数
	AvgDurationMs  float64 `json:"avgDurationMs"`  // 平均耗时
	SuccessRate    float64 `json:"successRate"`    // 成功率
	ProviderStats  []ProviderStat `json:"providerStats"`  // 提供商统计
}

// ProviderStat 提供商统计
type ProviderStat struct {
	Provider      string  `json:"provider"`
	RequestCount  int64   `json:"requestCount"`
	TotalTokens   int64   `json:"totalTokens"`
	SuccessRate   float64 `json:"successRate"`
}
