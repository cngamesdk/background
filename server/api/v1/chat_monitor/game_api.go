package chat_monitor

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat_monitor"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	chatMonitorService "github.com/flipped-aurora/gin-vue-admin/server/service/chat_monitor"
	"github.com/gin-gonic/gin"
)

type GameApi struct{}

var gameService = chatMonitorService.ServiceGroup{}.GameService

// CreateGame 创建游戏
// @Tags     ChatMonitor
// @Summary  创建游戏接入配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body chat_monitor.Game true "游戏信息"
// @Success  200 {object} response.Response{msg=string} "创建成功"
// @Router   /chatMonitor/game/create [post]
func (a *GameApi) CreateGame(c *gin.Context) {
	var game chat_monitor.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gameService.CreateGame(game); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// UpdateGame 更新游戏
// @Tags     ChatMonitor
// @Summary  更新游戏配置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body chat_monitor.Game true "游戏信息"
// @Success  200 {object} response.Response{msg=string} "更新成功"
// @Router   /chatMonitor/game/update [put]
func (a *GameApi) UpdateGame(c *gin.Context) {
	var game chat_monitor.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gameService.UpdateGame(game); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteGame 删除游戏
// @Tags     ChatMonitor
// @Summary  删除游戏
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body request.GetById true "游戏ID"
// @Success  200 {object} response.Response{msg=string} "删除成功"
// @Router   /chatMonitor/game/delete [delete]
func (a *GameApi) DeleteGame(c *gin.Context) {
	var req struct {
		ID uint `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := gameService.DeleteGame(req.ID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetGameList 获取游戏列表
// @Tags     ChatMonitor
// @Summary  获取游戏列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data query chat_monitor.GameSearch true "查询参数"
// @Success  200 {object} response.Response{data=response.PageResult} "获取成功"
// @Router   /chatMonitor/game/list [get]
func (a *GameApi) GetGameList(c *gin.Context) {
	var search chat_monitor.GameSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := gameService.GetGameList(search)
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

// placeholder
var _ = global.GVA_DB
