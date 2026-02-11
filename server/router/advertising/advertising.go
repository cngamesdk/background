package advertising

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdvertisingRouter struct {
}

func (s *AdvertisingRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("advertising")
	apiRouter.Use(middleware.OperationRecord())
	{
		mediaGroup := apiRouter.Group("media")
		{
			mediaGroup.POST("list", advertisingMediaApi.List)
			mediaGroup.POST("add", advertisingMediaApi.Add)
			mediaGroup.POST("modify", advertisingMediaApi.Modify)
		}
		channelGroupGroup := apiRouter.Group("channel-group")
		{
			channelGroupGroup.POST("list", channelGroupApi.List)
			channelGroupGroup.POST("add", channelGroupApi.Add)
			channelGroupGroup.POST("modify", channelGroupApi.Modify)
		}
		agentGroup := apiRouter.Group("agent")
		{
			agentGroup.POST("list", agentApi.List)
			agentGroup.POST("add", agentApi.Add)
			agentGroup.POST("modify", agentApi.Modify)
		}
		siteGroup := apiRouter.Group("site")
		{
			siteGroup.POST("list", siteApi.List)
			siteGroup.POST("add", siteApi.Add)
			siteGroup.POST("modify", siteApi.Modify)
		}

		//媒体开发者配置分组
		developerConfigGroup := apiRouter.Group("developer-config")
		{
			developerConfigGroup.POST("list", advertisingDeveloperConfigApi.List)
			developerConfigGroup.POST("add", advertisingDeveloperConfigApi.Add)
			developerConfigGroup.POST("modify", advertisingDeveloperConfigApi.Modify)
		}
	}
}
