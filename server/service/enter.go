package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/service/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system_management"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup              system.ServiceGroup
	ExampleServiceGroup             example.ServiceGroup
	OperationManagementServiceGroup operation_management.ServiceGroup
	SystemManagementServiceGroup    system_management.ServiceGroup
	AdvertisingServiceGroup         advertising.ServiceGroup
	CronTaskServiceGroup         cron_task.ServiceGroup
}
