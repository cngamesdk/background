package model_transfer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ReportApi struct{}

// DailyReport 日报表
// @Tags ModelTransfer
// @Summary 日报表查询
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.DailyReportReq true "日报表查询"
// @Success 200 {object} response.Response{data=response.PageResult} "获取成功"
// @Router /model-transfer/report/daily [post]
func (a *ReportApi) DailyReport(c *gin.Context) {
	var req api.DailyReportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := reportService.DailyReport(c.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("查询日报失败", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, c)
}

// TokenUsage Token使用详情
// @Tags ModelTransfer
// @Summary Token使用详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.TokenUsageReq true "Token使用详情"
// @Success 200 {object} response.Response{data=response.PageResult} "获取成功"
// @Router /model-transfer/report/token-usage [post]
func (a *ReportApi) TokenUsage(c *gin.Context) {
	var req api.TokenUsageReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := reportService.TokenUsage(c.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("查询使用详情失败", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, c)
}

// SummaryReport 汇总报表
// @Tags ModelTransfer
// @Summary 汇总报表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.SummaryReportReq true "汇总报表查询"
// @Success 200 {object} response.Response{data=api.SummaryReportResp} "获取成功"
// @Router /model-transfer/report/summary [post]
func (a *ReportApi) SummaryReport(c *gin.Context) {
	var req api.SummaryReportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	summary, err := reportService.SummaryReport(c.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("查询汇总报表失败", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(summary, c)
}
