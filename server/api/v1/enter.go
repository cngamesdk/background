package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/activity_engine"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/chat_monitor"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/data_report"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/live_chat"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/material"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/short_link"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system_management"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/model_transfer"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup              system.ApiGroup
	ExampleApiGroup             example.ApiGroup
	OperationManagementApiGroup operation_management.ApiGroup
	SystemManagementApiGroup    system_management.ApiGroup
	AdvertisingApiGroup         advertising.ApiGroup
	CronTaskApiGroup            cron_task.ApiGroup
	DataReportApiGroup          data_report.ApiGroup
	MaterialApiGroup            material.ApiGroup
	ChatMonitorApiGroup         chat_monitor.ApiGroup
	LiveChatApiGroup            live_chat.ApiGroup
	ActivityEngineApiGroup      activity_engine.ApiGroup
	ShortLinkApiGroup           short_link.ApiGroup
	ModelTransferApiGroup       model_transfer.ApiGroup
}
