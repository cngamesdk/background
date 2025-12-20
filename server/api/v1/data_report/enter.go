package data_report

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DayOverviewApi
	RetentionStatusApi
	PaymentStatusApi
}

var (
	dayOverviewService     = service.ServiceGroupApp.DataReportServiceGroup.DayOverviewService
	retentionStatusService = service.ServiceGroupApp.DataReportServiceGroup.RetentionStatusService
	paymentStatusService   = service.ServiceGroupApp.DataReportServiceGroup.PaymentStatusService
)
