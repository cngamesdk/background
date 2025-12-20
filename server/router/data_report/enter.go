package data_report

import (
	api2 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	DataReportRouter
}

var (
	dayOverviewApi     = api2.ApiGroupApp.DataReportApiGroup.DayOverviewApi
	retentionStatusApi = api2.ApiGroupApp.DataReportApiGroup.RetentionStatusApi
	paymentStatusApi   = api2.ApiGroupApp.DataReportApiGroup.PaymentStatusApi
)
