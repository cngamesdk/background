package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{
		{Path: "/short-link/create", Description: "创建短链接", ApiGroup: "短链接管理", Method: "POST"},
		{Path: "/short-link/list", Description: "短链接列表", ApiGroup: "短链接管理", Method: "POST"},
		{Path: "/short-link/detail", Description: "短链接详情", ApiGroup: "短链接管理", Method: "POST"},
		{Path: "/short-link/update", Description: "更新短链接", ApiGroup: "短链接管理", Method: "POST"},
		{Path: "/short-link/delete", Description: "删除短链接", ApiGroup: "短链接管理", Method: "POST"},
		{Path: "/short-link/click-log/list", Description: "点击日志列表", ApiGroup: "短链接管理", Method: "POST"},
	}
	utils.RegisterApis(entities...)

	// 为超级管理员角色(888)添加casbin策略
	registerCasbinPolicies()
}

func registerCasbinPolicies() {
	policies := [][]interface{}{
		{"p", "888", "/short-link/create", "POST"},
		{"p", "888", "/short-link/list", "POST"},
		{"p", "888", "/short-link/detail", "POST"},
		{"p", "888", "/short-link/update", "POST"},
		{"p", "888", "/short-link/delete", "POST"},
		{"p", "888", "/short-link/click-log/list", "POST"},
	}

	for _, policy := range policies {
		var count int64
		global.GVA_DB.Table("casbin_rule").
			Where("ptype = ? AND v0 = ? AND v1 = ? AND v2 = ?", policy[0], policy[1], policy[2], policy[3]).
			Count(&count)
		if count > 0 {
			continue
		}
		if err := global.GVA_DB.Exec(
			"INSERT INTO casbin_rule (ptype, v0, v1, v2) VALUES (?, ?, ?, ?)",
			policy[0], policy[1], policy[2], policy[3],
		).Error; err != nil {
			fmt.Println("添加casbin策略失败:", err)
		}
	}
}
