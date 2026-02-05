package cron_task

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CronTaskApi
}

var (
	cronTaskService = service.ServiceGroupApp.CronTaskServiceGroup.CronTaskService
)
