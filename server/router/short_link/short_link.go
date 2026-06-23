package short_link

import (
	"context"
	"sync"

	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router/short_link/initialize"
	"github.com/gin-gonic/gin"
)

type ShortLinkRouter struct{}

var once sync.Once

func (r *ShortLinkRouter) InitApiRouter(Router *gin.RouterGroup) {
	once.Do(func() {
		ctx := context.Background()
		initialize.Menu(ctx)
		initialize.Api(ctx)
	})

	apiRouter := Router.Group("short-link")
	apiRouter.Use(middleware.OperationRecord())
	{
		apiRouter.POST("create", shortLinkApi.Create)
		apiRouter.POST("list", shortLinkApi.List)
		apiRouter.POST("detail", shortLinkApi.Detail)
		apiRouter.POST("update", shortLinkApi.Update)
		apiRouter.POST("delete", shortLinkApi.Delete)
		apiRouter.POST("click-log/list", shortLinkApi.ClickLogList)
	}
}
