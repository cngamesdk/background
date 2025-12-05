package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/router/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
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
}
