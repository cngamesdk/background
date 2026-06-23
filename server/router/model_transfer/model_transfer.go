package model_transfer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ModelTransferRouter struct{}

func (s *ModelTransferRouter) InitTokenRouter(Router *gin.RouterGroup) {
	tokenRouter := Router.Group("model-transfer/token").Use(middleware.OperationRecord())
	tokenApi := v1.ApiGroupApp.ModelTransferApiGroup.TokenApi
	{
		tokenRouter.POST("create", tokenApi.Create)   // 创建Token
		tokenRouter.POST("update", tokenApi.Update)   // 更新Token
		tokenRouter.POST("delete", tokenApi.Delete)   // 删除Token
		tokenRouter.POST("list", tokenApi.List)       // Token列表
		tokenRouter.POST("detail", tokenApi.Detail)   // Token详情
		tokenRouter.POST("regenerate", tokenApi.Regenerate) // 重新生成Token
	}
}

func (s *ModelTransferRouter) InitReportRouter(Router *gin.RouterGroup) {
	reportRouter := Router.Group("model-transfer/report").Use(middleware.OperationRecord())
	reportApi := v1.ApiGroupApp.ModelTransferApiGroup.ReportApi
	{
		reportRouter.POST("daily", reportApi.DailyReport)         // 日报表
		reportRouter.POST("token-usage", reportApi.TokenUsage)    // Token使用详情
		reportRouter.POST("summary", reportApi.SummaryReport)     // 汇总报表
	}
}
