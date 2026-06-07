package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/activity_engine"
	"github.com/flipped-aurora/gin-vue-admin/server/router/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/router/chat_monitor"
	"github.com/flipped-aurora/gin-vue-admin/server/router/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/router/data_report"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/live_chat"
	"github.com/flipped-aurora/gin-vue-admin/server/router/material"
	"github.com/flipped-aurora/gin-vue-admin/server/router/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system_management"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System              system.RouterGroup
	Example             example.RouterGroup
	OperationManagement operation_management.RouterGroup
	SystemManagement    system_management.RouterGroup
	Advertising         advertising.RouterGroup
	CronTask            cron_task.RouterGroup
	DataReport          data_report.RouterGroup
	Material            material.RouterGroup
	ChatMonitor         chat_monitor.RouterGroup
	LiveChat            live_chat.RouterGroup
	ActivityEngine      activity_engine.RouterGroup
}
