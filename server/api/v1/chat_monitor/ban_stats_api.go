package chat_monitor

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat_monitor"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	chatMonitorService "github.com/flipped-aurora/gin-vue-admin/server/service/chat_monitor"
	"github.com/gin-gonic/gin"
)

type BanApi struct{}
type StatsApi struct{}

var banService = chatMonitorService.ServiceGroup{}.BanService
var statsService = chatMonitorService.ServiceGroup{}.StatsService

// CreateBan 创建封禁
// @Tags     ChatMonitor
// @Summary  创建封禁/禁言
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body chat_monitor.BanCreateReq true "封禁信息"
// @Success  200 {object} response.Response{msg=string} "创建成功"
// @Router   /chatMonitor/ban/create [post]
func (a *BanApi) CreateBan(c *gin.Context) {
	var req chat_monitor.BanCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := c.GetUint("userID")
	if err := banService.CreateBan(req, userID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("封禁成功", c)
}

// RevokeBan 解除封禁
// @Tags     ChatMonitor
// @Summary  解除封禁
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {object} response.Response{msg=string} "解封成功"
// @Router   /chatMonitor/ban/revoke [put]
func (a *BanApi) RevokeBan(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := banService.RevokeBan(req.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("解封成功", c)
}

// GetBanList 获取封禁列表
// @Tags     ChatMonitor
// @Summary  获取封禁记录列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query chat_monitor.BanRecordSearch true "查询参数"
// @Success  200 {object} response.Response{data=response.PageResult} "获取成功"
// @Router   /chatMonitor/ban/list [get]
func (a *BanApi) GetBanList(c *gin.Context) {
	var search chat_monitor.BanRecordSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := banService.GetBanList(search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     search.Page,
		PageSize: search.PageSize,
	}, "获取成功", c)
}

// GetStatsOverview 统计概览
// @Tags     ChatMonitor
// @Summary  获取统计概览
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    appId query string false "游戏ID"
// @Success  200 {object} response.Response{data=chatMonitorService.StatsOverview} "获取成功"
// @Router   /chatMonitor/stats/overview [get]
func (a *StatsApi) GetStatsOverview(c *gin.Context) {
	appID := c.Query("appId")
	data, err := statsService.GetOverview(appID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// GetStatsTrend 趋势图
// @Tags     ChatMonitor
// @Summary  获取消息趋势数据
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query chat_monitor.StatsTrendReq true "查询参数"
// @Success  200 {object} response.Response{data=[]chatMonitorService.TrendItem} "获取成功"
// @Router   /chatMonitor/stats/trend [get]
func (a *StatsApi) GetStatsTrend(c *gin.Context) {
	var req chat_monitor.StatsTrendReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := statsService.GetTrend(req.AppID, req.Days)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}

// GetViolators 违规用户排行
// @Tags     ChatMonitor
// @Summary  获取违规用户排行
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query chat_monitor.ViolatorsReq true "查询参数"
// @Success  200 {object} response.Response{data=[]chatMonitorService.ViolatorItem} "获取成功"
// @Router   /chatMonitor/stats/violators [get]
func (a *StatsApi) GetViolators(c *gin.Context) {
	var req chat_monitor.ViolatorsReq
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := statsService.GetViolators(req.AppID, req.Limit)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(data, "获取成功", c)
}
