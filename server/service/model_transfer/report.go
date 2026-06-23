package model_transfer

import (
	"context"
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer/api"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ReportService struct{}

// DailyReport 日报表查询
func (s *ReportService) DailyReport(ctx context.Context, req *api.DailyReportReq) ([]api.DailyReportResp, int64, error) {
	var list []api.DailyReportResp
	var total int64

	db := global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiDailyReport{})

	// 筛选条件
	if req.TokenName != "" {
		db = db.Where("token_name LIKE ?", "%"+req.TokenName+"%")
	}
	if req.StartDate != "" {
		db = db.Where("stat_date >= ?", req.StartDate)
	}
	if req.EndDate != "" {
		db = db.Where("stat_date <= ?", req.EndDate)
	}
	if req.Provider != "" {
		db = db.Where("provider = ?", req.Provider)
	}
	if req.Model != "" {
		db = db.Where("model = ?", req.Model)
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("获取日报总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页查询
	if err := db.Scopes(req.Paginate()).Order("stat_date, token_id DESC").Find(&list).Error; err != nil {
		global.GVA_LOG.Error("查询日报列表失败", zap.Error(err))
		return nil, 0, err
	}

	// 计算成功率
	for i := range list {
		if list[i].RequestCount > 0 {
			list[i].SuccessRate = float64(list[i].SuccessCount) * 100.0 / float64(list[i].RequestCount)
		}
	}

	return list, total, nil
}

// TokenUsage Token使用详情
func (s *ReportService) TokenUsage(ctx context.Context, req *api.TokenUsageReq) ([]api.TokenUsageResp, int64, error) {
	var list []api.TokenUsageResp
	var total int64

	db := global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiUsageLog{}).
		Where("token_id = ?", req.TokenID)

	// 时间范围
	if req.StartDate != "" {
		db = db.Where("DATE(request_time) >= ?", req.StartDate)
	}
	if req.EndDate != "" {
		db = db.Where("DATE(request_time) <= ?", req.EndDate)
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("获取使用记录总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页查询
	if err := db.Scopes(func(d *gorm.DB) *gorm.DB {
		offset := (req.Page - 1) * req.PageSize
		return d.Offset(offset).Limit(req.PageSize)
	}).Order("request_time DESC").Find(&list).Error; err != nil {
		global.GVA_LOG.Error("查询使用记录失败", zap.Error(err))
		return nil, 0, err
	}

	return list, total, nil
}

// SummaryReport 汇总报表
func (s *ReportService) SummaryReport(ctx context.Context, req *api.SummaryReportReq) (*api.SummaryReportResp, error) {
	resp := &api.SummaryReportResp{}

	// 总体统计
	var totalStats struct {
		TotalRequests int64
		SuccessCount  int64
		ErrorCount    int64
		TotalTokens   int64
		AvgDurationMs float64
	}

	if err := global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiUsageLog{}).
		Select(`
			COUNT(*) as total_requests,
			SUM(CASE WHEN status_code = 200 THEN 1 ELSE 0 END) as success_count,
			SUM(CASE WHEN status_code != 200 THEN 1 ELSE 0 END) as error_count,
			COALESCE(SUM(total_tokens), 0) as total_tokens,
			COALESCE(AVG(duration_ms), 0) as avg_duration_ms
		`).
		Where("DATE(request_time) BETWEEN ? AND ?", req.StartDate, req.EndDate).
		Scan(&totalStats).Error; err != nil {
		global.GVA_LOG.Error("查询汇总统计失败", zap.Error(err))
		return nil, err
	}

	resp.TotalRequests = totalStats.TotalRequests
	resp.SuccessCount = totalStats.SuccessCount
	resp.ErrorCount = totalStats.ErrorCount
	resp.TotalTokens = totalStats.TotalTokens
	resp.AvgDurationMs = totalStats.AvgDurationMs
	if resp.TotalRequests > 0 {
		resp.SuccessRate = mathutil.Percent(float64(resp.SuccessCount), float64(resp.TotalRequests), 2)
	}

	// 按提供商统计
	var providerStats []api.ProviderStat
	if err := global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiUsageLog{}).
		Select(`
			provider,
			COUNT(*) as request_count,
			COALESCE(SUM(total_tokens), 0) as total_tokens,
			SUM(CASE WHEN status_code = 200 THEN 1 ELSE 0 END) * 100.0 / COUNT(*) as success_rate
		`).
		Where("DATE(request_time) BETWEEN ? AND ?", req.StartDate, req.EndDate).
		Group("provider").
		Scan(&providerStats).Error; err != nil {
		global.GVA_LOG.Error("查询提供商统计失败", zap.Error(err))
		return nil, err
	}

	resp.ProviderStats = providerStats

	return resp, nil
}
