package system_management

import api2 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	SystemManagementRouter
}

var (
	searchApi	= api2.ApiGroupApp.SystemManagementApiGroup.SearchApi
)

