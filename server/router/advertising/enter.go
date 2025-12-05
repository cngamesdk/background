package advertising

import v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	AdvertisingRouter
}

var (
	advertisingMediaApi = v1.ApiGroupApp.AdvertisingApiGroup.AdvertisingMediaApi
	channelGroupApi     = v1.ApiGroupApp.AdvertisingApiGroup.ChannelGroupApi
	agentApi            = v1.ApiGroupApp.AdvertisingApiGroup.AgentApi
	siteApi             = v1.ApiGroupApp.AdvertisingApiGroup.SiteApi
)
