package operation_management

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service"
)

type ApiGroup struct {
	SubGameApi
	MainGameApi
	RootGameApi
	CompanyApi
	PayChannelApi
	PublishingApi
	PlatformApi
	GameAppVersionConfigurationApi
	ProductCommonConfigurationApi
	GlobalCommonConfigApi
	GamePackagingConfigApi
}

var (
	subGameService                     = service.ServiceGroupApp.OperationManagementServiceGroup.SubGameService
	mainGameService                    = service.ServiceGroupApp.OperationManagementServiceGroup.MainGameService
	rootGameService                    = service.ServiceGroupApp.OperationManagementServiceGroup.RootGameService
	companyService                     = service.ServiceGroupApp.OperationManagementServiceGroup.CompanyService
	payChannelService                  = service.ServiceGroupApp.OperationManagementServiceGroup.PayChannelService
	publishingService                  = service.ServiceGroupApp.OperationManagementServiceGroup.PublishingService
	platformService                    = service.ServiceGroupApp.OperationManagementServiceGroup.PlatformService
	gameAppVersionConfigurationService = service.ServiceGroupApp.OperationManagementServiceGroup.GameAppVersionConfigurationService
	productCommonConfigurationService  = service.ServiceGroupApp.OperationManagementServiceGroup.ProductCommonConfigurationService
	globalCommonConfigService          = service.ServiceGroupApp.OperationManagementServiceGroup.GlobalCommonConfigService
	gamePackagingConfigService         = service.ServiceGroupApp.OperationManagementServiceGroup.GamePackagingConfigService
)
