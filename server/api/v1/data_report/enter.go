package data_report

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DayOverviewApi
}

var (
	dayOverviewService = service.ServiceGroupApp.DataReportServiceGroup.DayOverviewService
)
