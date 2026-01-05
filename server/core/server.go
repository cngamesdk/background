package core

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
	"time"
)

func RunServer() {
	if global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
		if global.GVA_CONFIG.System.UseMultipoint {
			initialize.RedisList()
		}
	}

	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	fmt.Printf(`
	欢迎使用 「梅花互娱」后台管理系统
	当前版本: %s
	联系方式： 微信号：doudoualvin，QQ群：827883750
	项目地址：https://github.com/cngamesdk
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认MCP SSE地址:http://127.0.0.1%s%s
	默认MCP Message地址:http://127.0.0.1%s%s
	默认前端文件运行地址:http://127.0.0.1:8080
`, global.Version, address, address, global.GVA_CONFIG.MCP.SSEPath, address, global.GVA_CONFIG.MCP.MessagePath)
	initServer(address, Router, 10*time.Minute, 10*time.Minute)
}
