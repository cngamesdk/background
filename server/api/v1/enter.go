package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system_management"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup              system.ApiGroup
	ExampleApiGroup             example.ApiGroup
	OperationManagementApiGroup operation_management.ApiGroup
	SystemManagementApiGroup    system_management.ApiGroup
	AdvertisingApiGroup         advertising.ApiGroup
	CronTaskApiGroup            cron_task.ApiGroup
}
