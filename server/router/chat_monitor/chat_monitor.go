package chat_monitor

import (
	"context"
	"sync"

	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router/chat_monitor/initialize"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	ChatMonitorRouter
}

type ChatMonitorRouter struct{}

var once sync.Once

func (r *ChatMonitorRouter) InitChatMonitorRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	// 注册菜单和API（只执行一次）
	once.Do(func() {
		ctx := context.Background()
		initialize.Menu(ctx)
		initialize.Api(ctx)
	})

	chatMonitorApi := v1.ApiGroupApp.ChatMonitorApiGroup

	// 需要操作记录的路由
	recordRouter := Router.Group("chatMonitor").Use(middleware.OperationRecord())
	{
		recordRouter.POST("game/create", chatMonitorApi.GameApi.CreateGame)
		recordRouter.PUT("game/update", chatMonitorApi.GameApi.UpdateGame)
		recordRouter.DELETE("game/delete", chatMonitorApi.GameApi.DeleteGame)

		recordRouter.POST("sensitive/create", chatMonitorApi.SensitiveApi.CreateSensitiveWord)
		recordRouter.POST("sensitive/import", chatMonitorApi.SensitiveApi.ImportSensitiveWords)
		recordRouter.PUT("sensitive/update", chatMonitorApi.SensitiveApi.UpdateSensitiveWord)
		recordRouter.DELETE("sensitive/delete", chatMonitorApi.SensitiveApi.DeleteSensitiveWord)

		recordRouter.POST("whitelist/create", chatMonitorApi.WhitelistApi.CreateWhitelist)

		recordRouter.POST("ban/create", chatMonitorApi.BanApi.CreateBan)
		recordRouter.PUT("ban/revoke", chatMonitorApi.BanApi.RevokeBan)
	}

	// 只需要鉴权的路由
	privateRouter := Router.Group("chatMonitor")
	{
		privateRouter.GET("game/list", chatMonitorApi.GameApi.GetGameList)

		privateRouter.GET("chat/history", chatMonitorApi.ChatApi.GetChatHistory)

		privateRouter.GET("sensitive/list", chatMonitorApi.SensitiveApi.GetSensitiveWordList)
		privateRouter.GET("whitelist/list", chatMonitorApi.WhitelistApi.GetWhitelistList)

		privateRouter.GET("ban/list", chatMonitorApi.BanApi.GetBanList)

		privateRouter.GET("stats/overview", chatMonitorApi.StatsApi.GetStatsOverview)
		privateRouter.GET("stats/trend", chatMonitorApi.StatsApi.GetStatsTrend)
		privateRouter.GET("stats/violators", chatMonitorApi.StatsApi.GetViolators)
	}
}
