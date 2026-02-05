package advertising

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	AdvertisingMediaApi
	ChannelGroupApi
	AgentApi
	SiteApi
}

var (
	advertisingMediaService = service.ServiceGroupApp.AdvertisingServiceGroup.AdvertisingMediaService
	channelGroupService     = service.ServiceGroupApp.AdvertisingServiceGroup.ChannelGroupService
	agentService            = service.ServiceGroupApp.AdvertisingServiceGroup.AgentService
	siteService             = service.ServiceGroupApp.AdvertisingServiceGroup.SiteService
)
