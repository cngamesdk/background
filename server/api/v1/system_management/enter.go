package system_management

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	SearchApi
}

var (
	searchService = service.ServiceGroupApp.SystemManagementServiceGroup.SearchService
)
