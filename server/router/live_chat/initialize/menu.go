package initialize

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []system.SysBaseMenu{
		{
			ParentId:  0,
			Path:      "liveChat",
			Name:      "liveChat",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      32,
			Meta:      system.Meta{Title: "在线客服", Icon: "chat-line-round"},
		},
		{
			ParentId:  0,
			Path:      "liveChatProduct",
			Name:      "liveChatProduct",
			Hidden:    false,
			Component: "view/liveChat/product/index.vue",
			Sort:      1,
			Meta:      system.Meta{Title: "产品管理", Icon: "goods"},
		},
		{
			ParentId:  0,
			Path:      "liveChatFaq",
			Name:      "liveChatFaq",
			Hidden:    false,
			Component: "view/liveChat/faq/index.vue",
			Sort:      2,
			Meta:      system.Meta{Title: "问题库", Icon: "collection-tag"},
		},
		{
			ParentId:  0,
			Path:      "liveChatAgent",
			Name:      "liveChatAgent",
			Hidden:    false,
			Component: "view/liveChat/agent/index.vue",
			Sort:      3,
			Meta:      system.Meta{Title: "客服管理", Icon: "avatar"},
		},
		{
			ParentId:  0,
			Path:      "liveChatChat",
			Name:      "liveChatChat",
			Hidden:    false,
			Component: "view/liveChat/chat/index.vue",
			Sort:      4,
			Meta:      system.Meta{Title: "会话管理", Icon: "chat-dot-round"},
		},
		{
			ParentId:  0,
			Path:      "liveChatReport",
			Name:      "liveChatReport",
			Hidden:    false,
			Component: "view/liveChat/report/index.vue",
			Sort:      5,
			Meta:      system.Meta{Title: "数据报表", Icon: "data-line"},
		},
	}
	utils.RegisterMenus(entities...)
}
