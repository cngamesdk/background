package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		// 活动管理
		{Path: "/activity_engine/activity/list", Description: "活动列表", ApiGroup: "活动引擎", Method: "POST"},
		{Path: "/activity_engine/activity/add", Description: "新增活动", ApiGroup: "活动引擎", Method: "POST"},
		{Path: "/activity_engine/activity/modify", Description: "修改活动", ApiGroup: "活动引擎", Method: "POST"},
		{Path: "/activity_engine/activity/detail", Description: "活动详情", ApiGroup: "活动引擎", Method: "POST"},
		{Path: "/activity_engine/activity/publish", Description: "发布活动", ApiGroup: "活动引擎", Method: "POST"},
		{Path: "/activity_engine/activity/offline", Description: "下线活动", ApiGroup: "活动引擎", Method: "POST"},
		// 模板管理
		{Path: "/activity_engine/template/list", Description: "模板列表", ApiGroup: "活动引擎", Method: "POST"},
		{Path: "/activity_engine/template/add", Description: "新增模板", ApiGroup: "活动引擎", Method: "POST"},
		{Path: "/activity_engine/template/clone", Description: "克隆模板", ApiGroup: "活动引擎", Method: "POST"},
		// 奖励道具
		{Path: "/activity_engine/reward_item/search", Description: "搜索奖励道具", ApiGroup: "活动引擎", Method: "POST"},
		// 灰度管理
		{Path: "/activity_engine/grayscale/update", Description: "更新灰度比例", ApiGroup: "活动引擎", Method: "POST"},
		// 沙箱测试
		{Path: "/activity_engine/sandbox/simulate", Description: "沙箱模拟", ApiGroup: "活动引擎", Method: "POST"},
	}
	utils.RegisterApis(entities...)
}
