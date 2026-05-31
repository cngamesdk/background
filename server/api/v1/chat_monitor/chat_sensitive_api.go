package chat_monitor

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat_monitor"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	chatMonitorService "github.com/flipped-aurora/gin-vue-admin/server/service/chat_monitor"
	"github.com/gin-gonic/gin"
)

type ChatApi struct{}
type SensitiveApi struct{}
type WhitelistApi struct{}

var chatService = chatMonitorService.ServiceGroup{}.ChatService
var sensitiveService = chatMonitorService.ServiceGroup{}.SensitiveService
var whitelistService = chatMonitorService.ServiceGroup{}.WhitelistService

// GetChatHistory 获取聊天记录
// @Tags     ChatMonitor
// @Summary  获取聊天历史记录
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query chat_monitor.ChatMessageSearch true "查询参数"
// @Success  200 {object} response.Response{data=response.PageResult} "获取成功"
// @Router   /chatMonitor/chat/history [get]
func (a *ChatApi) GetChatHistory(c *gin.Context) {
	var search chat_monitor.ChatMessageSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := chatService.GetChatHistory(search)
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

// CreateSensitiveWord 创建敏感词
// @Tags     ChatMonitor
// @Summary  添加敏感词
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body chat_monitor.SensitiveWord true "敏感词信息"
// @Success  200 {object} response.Response{msg=string} "创建成功"
// @Router   /chatMonitor/sensitive/create [post]
func (a *SensitiveApi) CreateSensitiveWord(c *gin.Context) {
	var word chat_monitor.SensitiveWord
	if err := c.ShouldBindJSON(&word); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sensitiveService.CreateSensitiveWord(word); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// ImportSensitiveWords 批量导入敏感词
// @Tags     ChatMonitor
// @Summary  批量导入敏感词
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body chat_monitor.SensitiveImportReq true "导入数据"
// @Success  200 {object} response.Response{data=map[string]int64} "导入成功"
// @Router   /chatMonitor/sensitive/import [post]
func (a *SensitiveApi) ImportSensitiveWords(c *gin.Context) {
	var req chat_monitor.SensitiveImportReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	count, err := sensitiveService.ImportSensitiveWords(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{"imported": count}, "导入成功", c)
}

// UpdateSensitiveWord 更新敏感词
// @Tags     ChatMonitor
// @Summary  更新敏感词
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body chat_monitor.SensitiveWord true "敏感词信息"
// @Success  200 {object} response.Response{msg=string} "更新成功"
// @Router   /chatMonitor/sensitive/update [put]
func (a *SensitiveApi) UpdateSensitiveWord(c *gin.Context) {
	var word chat_monitor.SensitiveWord
	if err := c.ShouldBindJSON(&word); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sensitiveService.UpdateSensitiveWord(word); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteSensitiveWord 删除敏感词
// @Tags     ChatMonitor
// @Summary  删除敏感词
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200 {object} response.Response{msg=string} "删除成功"
// @Router   /chatMonitor/sensitive/delete [delete]
func (a *SensitiveApi) DeleteSensitiveWord(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sensitiveService.DeleteSensitiveWord(req.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetSensitiveWordList 获取敏感词列表
// @Tags     ChatMonitor
// @Summary  获取敏感词列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query chat_monitor.SensitiveWordSearch true "查询参数"
// @Success  200 {object} response.Response{data=response.PageResult} "获取成功"
// @Router   /chatMonitor/sensitive/list [get]
func (a *SensitiveApi) GetSensitiveWordList(c *gin.Context) {
	var search chat_monitor.SensitiveWordSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sensitiveService.GetSensitiveWordList(search)
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

// CreateWhitelist 添加白名单
// @Tags     ChatMonitor
// @Summary  添加白名单词
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body chat_monitor.Whitelist true "白名单信息"
// @Success  200 {object} response.Response{msg=string} "创建成功"
// @Router   /chatMonitor/whitelist/create [post]
func (a *WhitelistApi) CreateWhitelist(c *gin.Context) {
	var word chat_monitor.Whitelist
	if err := c.ShouldBindJSON(&word); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := whitelistService.CreateWhitelist(word); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// GetWhitelistList 获取白名单列表
// @Tags     ChatMonitor
// @Summary  获取白名单列表
// @Security ApiKeyAuth
// @Produce  application/json
// @Param    data query chat_monitor.WhitelistSearch true "查询参数"
// @Success  200 {object} response.Response{data=response.PageResult} "获取成功"
// @Router   /chatMonitor/whitelist/list [get]
func (a *WhitelistApi) GetWhitelistList(c *gin.Context) {
	var search chat_monitor.WhitelistSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := whitelistService.GetWhitelistList(search)
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
