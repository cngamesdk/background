package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []system.SysApi{
		// Product
		{Path: "/liveChat/product/create", Description: "创建产品", ApiGroup: "在线客服", Method: "POST"},
		{Path: "/liveChat/product/update", Description: "更新产品", ApiGroup: "在线客服", Method: "PUT"},
		{Path: "/liveChat/product/delete", Description: "删除产品", ApiGroup: "在线客服", Method: "DELETE"},
		{Path: "/liveChat/product/list", Description: "产品列表", ApiGroup: "在线客服", Method: "GET"},
		{Path: "/liveChat/product/:id", Description: "产品详情", ApiGroup: "在线客服", Method: "GET"},

		// FAQ
		{Path: "/liveChat/faq/create", Description: "创建FAQ", ApiGroup: "在线客服", Method: "POST"},
		{Path: "/liveChat/faq/update", Description: "更新FAQ", ApiGroup: "在线客服", Method: "PUT"},
		{Path: "/liveChat/faq/delete", Description: "删除FAQ", ApiGroup: "在线客服", Method: "DELETE"},
		{Path: "/liveChat/faq/list", Description: "FAQ列表", ApiGroup: "在线客服", Method: "GET"},
		{Path: "/liveChat/faq/import", Description: "批量导入FAQ", ApiGroup: "在线客服", Method: "POST"},
		{Path: "/liveChat/faq/categories", Description: "FAQ分类", ApiGroup: "在线客服", Method: "GET"},

		// Agent
		{Path: "/liveChat/agent/online", Description: "客服上线", ApiGroup: "在线客服", Method: "POST"},
		{Path: "/liveChat/agent/offline", Description: "客服下线", ApiGroup: "在线客服", Method: "POST"},
		{Path: "/liveChat/agent/update", Description: "更新客服配置", ApiGroup: "在线客服", Method: "PUT"},
		{Path: "/liveChat/agent/list", Description: "客服列表", ApiGroup: "在线客服", Method: "GET"},
		{Path: "/liveChat/agent/status", Description: "客服状态", ApiGroup: "在线客服", Method: "GET"},

		// Chat
		{Path: "/liveChat/chat/sessions", Description: "会话列表", ApiGroup: "在线客服", Method: "GET"},
		{Path: "/liveChat/chat/session/:id", Description: "会话详情", ApiGroup: "在线客服", Method: "GET"},
		{Path: "/liveChat/chat/assign", Description: "分配会话", ApiGroup: "在线客服", Method: "POST"},
		{Path: "/liveChat/chat/reply", Description: "客服回复", ApiGroup: "在线客服", Method: "POST"},
		{Path: "/liveChat/chat/close", Description: "关闭会话", ApiGroup: "在线客服", Method: "POST"},

		// Report
		{Path: "/liveChat/report/overview", Description: "数据概览", ApiGroup: "在线客服", Method: "GET"},
		{Path: "/liveChat/report/trend", Description: "趋势数据", ApiGroup: "在线客服", Method: "GET"},
	}
	utils.RegisterApis(entities...)
}
