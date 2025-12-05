package system_management

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SystemManagementRouter struct {

}

func (s *SystemManagementRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("system_management")
	apiRouter.Use(middleware.OperationRecord())
	{
		//搜索
		searchGroup := apiRouter.Group("search")
		{
			searchGroup.POST("search", searchApi.Search)     // 所有维度搜索
		}
	}
}
