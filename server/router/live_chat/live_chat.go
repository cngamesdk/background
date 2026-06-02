package live_chat

import (
	"context"
	"sync"

	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router/live_chat/initialize"
	"github.com/gin-gonic/gin"
)

type LiveChatRouter struct{}

var once sync.Once

func (r *LiveChatRouter) InitLiveChatRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	once.Do(func() {
		ctx := context.Background()
		initialize.Menu(ctx)
		initialize.Api(ctx)
	})

	liveChatApi := v1.ApiGroupApp.LiveChatApiGroup

	// Write operations (with operation record)
	recordRouter := Router.Group("liveChat").Use(middleware.OperationRecord())
	{
		recordRouter.POST("product/create", liveChatApi.ProductApi.Create)
		recordRouter.PUT("product/update", liveChatApi.ProductApi.Update)
		recordRouter.DELETE("product/delete", liveChatApi.ProductApi.Delete)

		recordRouter.POST("faq/create", liveChatApi.FaqApi.Create)
		recordRouter.PUT("faq/update", liveChatApi.FaqApi.Update)
		recordRouter.DELETE("faq/delete", liveChatApi.FaqApi.Delete)
		recordRouter.POST("faq/import", liveChatApi.FaqApi.Import)

		recordRouter.POST("agent/online", liveChatApi.AgentApi.Online)
		recordRouter.POST("agent/offline", liveChatApi.AgentApi.Offline)
		recordRouter.PUT("agent/update", liveChatApi.AgentApi.Update)

		recordRouter.POST("chat/assign", liveChatApi.ChatApi.Assign)
		recordRouter.POST("chat/reply", liveChatApi.ChatApi.Reply)
		recordRouter.POST("chat/close", liveChatApi.ChatApi.Close)
	}

	// Read operations
	privateRouter := Router.Group("liveChat")
	{
		privateRouter.GET("product/list", liveChatApi.ProductApi.List)
		privateRouter.GET("product/:id", liveChatApi.ProductApi.Detail)

		privateRouter.GET("faq/list", liveChatApi.FaqApi.List)
		privateRouter.GET("faq/categories", liveChatApi.FaqApi.Categories)

		privateRouter.GET("agent/list", liveChatApi.AgentApi.List)
		privateRouter.GET("agent/status", liveChatApi.AgentApi.Status)

		privateRouter.GET("chat/sessions", liveChatApi.ChatApi.Sessions)
		privateRouter.GET("chat/session/:id", liveChatApi.ChatApi.SessionDetail)

		privateRouter.GET("report/overview", liveChatApi.ReportApi.Overview)
		privateRouter.GET("report/trend", liveChatApi.ReportApi.Trend)
	}
}
