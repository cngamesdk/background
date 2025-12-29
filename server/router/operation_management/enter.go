package operation_management

import (
	api2 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	OperationManagementRouter
}

var (
	subGameApi                     = api2.ApiGroupApp.OperationManagementApiGroup.SubGameApi
	mainGameApi                    = api2.ApiGroupApp.OperationManagementApiGroup.MainGameApi
	rootGameApi                    = api2.ApiGroupApp.OperationManagementApiGroup.RootGameApi
	companyApi                     = api2.ApiGroupApp.OperationManagementApiGroup.CompanyApi
	payChannelApi                  = api2.ApiGroupApp.OperationManagementApiGroup.PayChannelApi
	publishingApi                  = api2.ApiGroupApp.OperationManagementApiGroup.PublishingApi
	platformApi                    = api2.ApiGroupApp.OperationManagementApiGroup.PlatformApi
	gameAppVersionConfigurationApi = api2.ApiGroupApp.OperationManagementApiGroup.GameAppVersionConfigurationApi
	productCommonConfigurationApi  = api2.ApiGroupApp.OperationManagementApiGroup.ProductCommonConfigurationApi
	globalCommonConfigApi          = api2.ApiGroupApp.OperationManagementApiGroup.GlobalCommonConfigApi
	gamePackagingConfigApi         = api2.ApiGroupApp.OperationManagementApiGroup.GamePackagingConfigApi
	gamePackagingApi               = api2.ApiGroupApp.OperationManagementApiGroup.GamePackagingApi
)
