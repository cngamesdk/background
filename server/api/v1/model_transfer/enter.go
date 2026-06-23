package model_transfer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	TokenApi
	ReportApi
}

var (
	tokenService  = service.ServiceGroupApp.ModelTransferServiceGroup.TokenService
	reportService = service.ServiceGroupApp.ModelTransferServiceGroup.ReportService
)
