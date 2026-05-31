package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		// 游戏管理
		{Path: "/chatMonitor/game/create", Description: "创建游戏", ApiGroup: "聊天监控", Method: "POST"},
		{Path: "/chatMonitor/game/update", Description: "更新游戏", ApiGroup: "聊天监控", Method: "PUT"},
		{Path: "/chatMonitor/game/delete", Description: "删除游戏", ApiGroup: "聊天监控", Method: "DELETE"},
		{Path: "/chatMonitor/game/list", Description: "游戏列表", ApiGroup: "聊天监控", Method: "GET"},
		// 聊天记录
		{Path: "/chatMonitor/chat/history", Description: "聊天历史", ApiGroup: "聊天监控", Method: "GET"},
		// 敏感词管理
		{Path: "/chatMonitor/sensitive/create", Description: "添加敏感词", ApiGroup: "聊天监控", Method: "POST"},
		{Path: "/chatMonitor/sensitive/import", Description: "批量导入敏感词", ApiGroup: "聊天监控", Method: "POST"},
		{Path: "/chatMonitor/sensitive/update", Description: "更新敏感词", ApiGroup: "聊天监控", Method: "PUT"},
		{Path: "/chatMonitor/sensitive/delete", Description: "删除敏感词", ApiGroup: "聊天监控", Method: "DELETE"},
		{Path: "/chatMonitor/sensitive/list", Description: "敏感词列表", ApiGroup: "聊天监控", Method: "GET"},
		// 白名单
		{Path: "/chatMonitor/whitelist/create", Description: "添加白名单", ApiGroup: "聊天监控", Method: "POST"},
		{Path: "/chatMonitor/whitelist/list", Description: "白名单列表", ApiGroup: "聊天监控", Method: "GET"},
		// 封禁管理
		{Path: "/chatMonitor/ban/create", Description: "创建封禁", ApiGroup: "聊天监控", Method: "POST"},
		{Path: "/chatMonitor/ban/revoke", Description: "解除封禁", ApiGroup: "聊天监控", Method: "PUT"},
		{Path: "/chatMonitor/ban/list", Description: "封禁列表", ApiGroup: "聊天监控", Method: "GET"},
		// 数据报表
		{Path: "/chatMonitor/stats/overview", Description: "统计概览", ApiGroup: "聊天监控", Method: "GET"},
		{Path: "/chatMonitor/stats/trend", Description: "消息趋势", ApiGroup: "聊天监控", Method: "GET"},
		{Path: "/chatMonitor/stats/violators", Description: "违规用户排行", ApiGroup: "聊天监控", Method: "GET"},
	}
	utils.RegisterApis(entities...)
}
