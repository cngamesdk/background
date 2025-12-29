package operation_management

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OperationManagementRouter struct {
}

func (s *OperationManagementRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("operation_management")
	apiRouter.Use(middleware.OperationRecord())
	{
		//游戏管理
		gameManageRouter := apiRouter.Group("game_manage")
		subGameGroup := gameManageRouter.Group("sub_game")
		{
			subGameGroup.POST("list", subGameApi.List)     // 子游戏列表
			subGameGroup.POST("config", subGameApi.Config) // 子游戏配置
			subGameGroup.POST("add", subGameApi.Add)       // 子游戏添加
			subGameGroup.POST("modify", subGameApi.Modify) // 子游戏修改
		}

		mainGameGroup := gameManageRouter.Group("main_game")
		{
			mainGameGroup.POST("list", mainGameApi.List)     // 主游戏列表
			mainGameGroup.POST("add", mainGameApi.Add)       // 主游戏添加
			mainGameGroup.POST("modify", mainGameApi.Modify) // 主游戏修改
		}

		rootGameGroup := gameManageRouter.Group("root_game")
		{
			rootGameGroup.POST("list", rootGameApi.List)     // 根游戏列表
			rootGameGroup.POST("add", rootGameApi.Add)       // 根游戏添加
			rootGameGroup.POST("modify", rootGameApi.Modify) // 根游戏修改
		}

		//平台管理
		platformManageRouter := apiRouter.Group("platform")
		{
			platformManageRouter.POST("list", platformApi.List)     // 列表
			platformManageRouter.POST("add", platformApi.Add)       // 添加
			platformManageRouter.POST("modify", platformApi.Modify) // 修改
		}

		//主体管理
		companyManageRouter := apiRouter.Group("company")
		{
			companyManageRouter.POST("list", companyApi.List)     // 主体列表
			companyManageRouter.POST("add", companyApi.Add)       // 主体添加
			companyManageRouter.POST("modify", companyApi.Modify) // 主体修改
		}

		//产品通用配置管理
		productCommonConfigurationManageRouter := apiRouter.Group("product_common_configuration")
		{
			productCommonConfigurationManageRouter.POST("list", productCommonConfigurationApi.List)     // 列表
			productCommonConfigurationManageRouter.POST("add", productCommonConfigurationApi.Add)       // 添加
			productCommonConfigurationManageRouter.POST("modify", productCommonConfigurationApi.Modify) // 修改
		}

		//游戏版本配置管理
		gameAppVersionConfigurationManageRouter := apiRouter.Group("game_app_version_configuration")
		{
			gameAppVersionConfigurationManageRouter.POST("list", gameAppVersionConfigurationApi.List)     // 列表
			gameAppVersionConfigurationManageRouter.POST("add", gameAppVersionConfigurationApi.Add)       // 添加
			gameAppVersionConfigurationManageRouter.POST("modify", gameAppVersionConfigurationApi.Modify) // 修改
		}

		//支付渠道管理
		payChannelManageRouter := apiRouter.Group("pay_channel")
		{
			payChannelManageRouter.POST("list", payChannelApi.List)     // 支付渠道列表
			payChannelManageRouter.POST("add", payChannelApi.Add)       // 支付渠道添加
			payChannelManageRouter.POST("modify", payChannelApi.Modify) // 支付渠道修改
		}

		//发行管理
		publishingManageRouter := apiRouter.Group("publishing")
		{
			//渠道管理
			channelManageRouter := publishingManageRouter.Group("channel")
			{
				channelManageRouter.POST("list", publishingApi.ChannelConfigList)     // 发行渠道列表
				channelManageRouter.POST("add", publishingApi.ChannelConfigAdd)       // 发行渠道添加
				channelManageRouter.POST("modify", publishingApi.ChannelConfigModify) // 发行渠道修改
			}
			//渠道游戏管理
			channelGameManageRouter := publishingManageRouter.Group("channel-game")
			{
				channelGameManageRouter.POST("list", publishingApi.ChannelGameConfigList)     // 发行渠道游戏列表
				channelGameManageRouter.POST("add", publishingApi.ChannelGameConfigAdd)       // 发行渠道游戏添加
				channelGameManageRouter.POST("modify", publishingApi.ChannelGameConfigModify) // 发行渠道游戏修改
			}
		}

		globalCommonConfigRouter := apiRouter.Group("global-common-config")
		{
			globalCommonConfigRouter.POST("list", globalCommonConfigApi.List)     // 全局配置列表
			globalCommonConfigRouter.POST("add", globalCommonConfigApi.Add)       // 全局配置添加
			globalCommonConfigRouter.POST("modify", globalCommonConfigApi.Modify) // 全局配置修改
		}

		gamePackagingConfigRouter := apiRouter.Group("game-packaging-config")
		{
			gamePackagingConfigRouter.POST("list", gamePackagingConfigApi.List)     // 游戏打包配置列表
			gamePackagingConfigRouter.POST("add", gamePackagingConfigApi.Add)       // 游戏打包配置添加
			gamePackagingConfigRouter.POST("modify", gamePackagingConfigApi.Modify) // 游戏打包配置修改
		}

		gamePackagingRouter := apiRouter.Group("game-packaging")
		{
			gamePackagingRouter.POST("list", gamePackagingApi.List) // 游戏打包日志列表
			gamePackagingRouter.POST("add", gamePackagingApi.Add)   // 游戏打包添加
		}
	}
}
