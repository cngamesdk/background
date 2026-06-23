package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{
		{
			ParentId:  0,
			Path:      "shortLink",
			Name:      "shortLink",
			Hidden:    false,
			Component: "view/routerHolder.vue",
			Sort:      40,
			Meta:      model.Meta{Title: "短链接管理", Icon: "link"},
		},
		{
			Path:      "shortLinkList",
			Name:      "shortLinkList",
			Hidden:    false,
			Component: "view/shortLink/list.vue",
			Sort:      1,
			Meta:      model.Meta{Title: "短链接列表", Icon: "list"},
		},
	}
	utils.RegisterMenus(entities...)

	// 将菜单分配给超级管理员角色(888)
	assignMenuToAuthority()
}

func assignMenuToAuthority() {
	var menus []model.SysBaseMenu
	if err := global.GVA_DB.Where("name IN ?", []string{"shortLink", "shortLinkList"}).Find(&menus).Error; err != nil {
		fmt.Println("查询短链接菜单失败:", err)
		return
	}
	if len(menus) == 0 {
		return
	}

	// 检查是否已分配
	var count int64
	global.GVA_DB.Table("sys_authority_menus").
		Where("sys_authority_authority_id = ? AND sys_base_menu_id = ?", "888", menus[0].ID).
		Count(&count)
	if count > 0 {
		return
	}

	// 批量插入关联关系
	for _, menu := range menus {
		global.GVA_DB.Exec(
			"INSERT IGNORE INTO sys_authority_menus (sys_authority_authority_id, sys_base_menu_id) VALUES (?, ?)",
			"888", menu.ID,
		)
	}
}
