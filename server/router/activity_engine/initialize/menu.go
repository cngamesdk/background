package initialize

import (
	"context"

	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  0,
			Path:      "activityEngine",
			Name:      "activityEngine",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      35,
			Meta:      model.Meta{Title: "活动引擎", Icon: "trophy"},
		},
		{
			Path:      "activityList",
			Name:      "activityList",
			Hidden:    false,
			Component: "view/activityEngine/activityList/list.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "活动管理", Icon: "list"},
		},
		{
			Path:      "activityEdit",
			Name:      "activityEdit",
			Hidden:    true,
			Component: "view/activityEngine/activityEdit/index.vue",
			Sort:      2,
			Meta:      model.Meta{Title: "活动编辑", Icon: "edit", ActiveName: "activityList"},
		},
		{
			Path:      "activityTemplate",
			Name:      "activityTemplate",
			Hidden:    false,
			Component: "view/activityEngine/templateList/list.vue",
			Sort:      3,
			Meta:      model.Meta{Title: "活动模板", Icon: "document-copy"},
		},
		{
			Path:      "activitySandbox",
			Name:      "activitySandbox",
			Hidden:    false,
			Component: "view/activityEngine/sandbox/index.vue",
			Sort:      4,
			Meta:      model.Meta{Title: "沙箱测试", Icon: "cpu"},
		},
		{
			Path:      "activityGrayscale",
			Name:      "activityGrayscale",
			Hidden:    false,
			Component: "view/activityEngine/grayscale/index.vue",
			Sort:      5,
			Meta:      model.Meta{Title: "灰度发布", Icon: "data-analysis"},
		},
	}
	utils.RegisterMenus(entities...)
}
