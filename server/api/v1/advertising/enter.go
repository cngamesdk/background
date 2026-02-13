package advertising

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AdvertisingMediaApi
	ChannelGroupApi
	AgentApi
	SiteApi
	AdvertisingDeveloperConfigApi
	AdvertisingAuthApi
}

var (
	advertisingMediaService           = service.ServiceGroupApp.AdvertisingServiceGroup.AdvertisingMediaService
	channelGroupService               = service.ServiceGroupApp.AdvertisingServiceGroup.ChannelGroupService
	agentService                      = service.ServiceGroupApp.AdvertisingServiceGroup.AgentService
	siteService                       = service.ServiceGroupApp.AdvertisingServiceGroup.SiteService
	advertisingDeveloperConfigService = service.ServiceGroupApp.AdvertisingServiceGroup.AdvertisingDeveloperConfigService
	advertisingAuthService            = service.ServiceGroupApp.AdvertisingServiceGroup.AdvertisingAuthService
)
