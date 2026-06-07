package activity_engine

import (
	"context"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router/activity_engine/initialize"
	"github.com/gin-gonic/gin"
)

type ActivityEngineRouter struct{}

var once sync.Once

func (r *ActivityEngineRouter) InitApiRouter(Router *gin.RouterGroup) {
	// 注册菜单和API（只执行一次）
	once.Do(func() {
		ctx := context.Background()
		initialize.Menu(ctx)
		initialize.Api(ctx)
	})

	apiRouter := Router.Group("activity_engine")
	apiRouter.Use(middleware.OperationRecord())
	{
		// 活动管理
		activityGroup := apiRouter.Group("activity")
		{
			activityGroup.POST("list", activityApi.List)
			activityGroup.POST("add", activityApi.Add)
			activityGroup.POST("modify", activityApi.Modify)
			activityGroup.POST("detail", activityApi.Detail)
			activityGroup.POST("publish", activityApi.Publish)
			activityGroup.POST("offline", activityApi.Offline)
		}

		// 模板管理
		templateGroup := apiRouter.Group("template")
		{
			templateGroup.POST("list", templateApi.List)
			templateGroup.POST("add", templateApi.Add)
			templateGroup.POST("clone", templateApi.Clone)
		}

		// 奖励道具
		rewardItemGroup := apiRouter.Group("reward_item")
		{
			rewardItemGroup.POST("search", rewardItemApi.Search)
		}

		// 灰度管理
		grayscaleGroup := apiRouter.Group("grayscale")
		{
			grayscaleGroup.POST("update", activityApi.UpdateGrayscale)
		}

		// 沙箱测试
		sandboxGroup := apiRouter.Group("sandbox")
		{
			sandboxGroup.POST("simulate", sandboxApi.Simulate)
		}
	}
}
