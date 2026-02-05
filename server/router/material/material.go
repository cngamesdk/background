package material

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MaterialRouter struct {
}

func (s *MaterialRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("material")
	apiRouter.Use(middleware.OperationRecord())
	{
		themeGroup := apiRouter.Group("theme")
		{
			themeGroup.POST("list", themeApi.List)
			themeGroup.POST("add", themeApi.Add)
			themeGroup.POST("modify", themeApi.Modify)
		}
	}
}
