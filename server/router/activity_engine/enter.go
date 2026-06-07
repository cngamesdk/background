package activity_engine

import api2 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ActivityEngineRouter
}

var (
	activityApi   = api2.ApiGroupApp.ActivityEngineApiGroup.ActivityApi
	templateApi   = api2.ApiGroupApp.ActivityEngineApiGroup.TemplateApi
	rewardItemApi = api2.ApiGroupApp.ActivityEngineApiGroup.RewardItemApi
	sandboxApi    = api2.ApiGroupApp.ActivityEngineApiGroup.SandboxApi
)
