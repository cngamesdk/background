package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  0,
			Path:      "chatMonitor",
			Name:      "chatMonitor",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      30,
			Meta:      model.Meta{Title: "聊天监控", Icon: "chat-dot-round"},
		},
		{
			Path:      "chatMonitorRealtime",
			Name:      "chatMonitorRealtime",
			Hidden:    false,
			Component: "view/chatMonitor/monitor/index.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "实时监控", Icon: "monitor"},
		},
		{
			Path:      "chatMonitorGame",
			Name:      "chatMonitorGame",
			Hidden:    false,
			Component: "view/chatMonitor/game/index.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "游戏管理", Icon: "menu"},
		},
		{
			Path:      "chatMonitorHistory",
			Name:      "chatMonitorHistory",
			Hidden:    false,
			Component: "view/chatMonitor/history/index.vue",
			Sort:      3,
			Meta:      model.Meta{Title: "历史查询", Icon: "search"},
		},
		{
			Path:      "chatMonitorSensitive",
			Name:      "chatMonitorSensitive",
			Hidden:    false,
			Component: "view/chatMonitor/sensitive/index.vue",
			Sort:      4,
			Meta:      model.Meta{Title: "词库管理", Icon: "warning"},
		},
		{
			Path:      "chatMonitorBan",
			Name:      "chatMonitorBan",
			Hidden:    false,
			Component: "view/chatMonitor/ban/index.vue",
			Sort:      5,
			Meta:      model.Meta{Title: "封禁管理", Icon: "lock"},
		},
		{
			Path:      "chatMonitorStats",
			Name:      "chatMonitorStats",
			Hidden:    false,
			Component: "view/chatMonitor/statistics/index.vue",
			Sort:      6,
			Meta:      model.Meta{Title: "数据报表", Icon: "data-line"},
		},
	}
	utils.RegisterMenus(entities...)
}
