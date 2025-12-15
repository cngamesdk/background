package data_report

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataReportRouter struct {
}

func (s *DataReportRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("data-report")
	apiRouter.Use(middleware.OperationRecord())
	{
		configGroup := apiRouter.Group("day-overview") // 每日总览
		{
			configGroup.POST("list", dayOverviewApi.List)
		}
	}
}
