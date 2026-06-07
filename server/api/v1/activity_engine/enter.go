package activity_engine

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ActivityApi
	TemplateApi
	RewardItemApi
	SandboxApi
}

var activityService = service.ServiceGroupApp.ActivityEngineServiceGroup.ActivityService
var templateService = service.ServiceGroupApp.ActivityEngineServiceGroup.TemplateService
var rewardItemService = service.ServiceGroupApp.ActivityEngineServiceGroup.RewardItemService
