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
		dayOverviewGroup := apiRouter.Group("day-overview") // 每日总览
		{
			dayOverviewGroup.POST("list", dayOverviewApi.List)
		}

		retentionStatusGroup := apiRouter.Group("retention-status") // 留存情况
		{
			retentionStatusGroup.POST("list", retentionStatusApi.List)
		}
	}
}
