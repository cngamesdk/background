package cron_task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CronTaskRouter struct {
}

func (s *CronTaskRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("cron-task")
	apiRouter.Use(middleware.OperationRecord())
	{
		configGroup := apiRouter.Group("config")
		{
			configGroup.POST("list", cronTaskApi.ConfigList)
			configGroup.POST("add", cronTaskApi.ConfigAdd)
			configGroup.POST("modify", cronTaskApi.ConfigModify)
		}

		logGroup := apiRouter.Group("log")
		{
			logGroup.POST("list", cronTaskApi.LogList)
		}
	}
}
