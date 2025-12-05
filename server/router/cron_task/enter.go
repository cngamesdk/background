package cron_task

import api2 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	CronTaskRouter
}

var (
	cronTaskApi                     = api2.ApiGroupApp.CronTaskApiGroup.CronTaskApi
)