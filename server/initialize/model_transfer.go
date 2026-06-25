package initialize

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
	"time"
)

// InitModelTransferMenu 初始化AI模型中转菜单
func InitModelTransferMenu() {
	ctx := context.Background()

	// 检查是否已经初始化
	var count int64
	if err := global.GVA_DB.Model(&system.SysBaseMenu{}).
		Where("name = ?", "modelTransfer").
		Count(&count).Error; err == nil && count > 0 {
		global.GVA_LOG.Info("AI模型中转菜单已存在，跳过初始化")
		return
	}

	// 一级菜单：AI模型中转
	parentMenu := system.SysBaseMenu{
		GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
		ParentId:  0,
		Path:      "modelTransfer",
		Name:      "modelTransfer",
		Hidden:    false,
		Component: "view/model_transfer/index.vue",
		Sort:      90,
		Meta: system.Meta{
			Title:       "AI模型中转",
			Icon:        "cpu",
			KeepAlive:   false,
			DefaultMenu: false,
			CloseTab:    false,
		},
	}

	if err := global.GVA_DB.WithContext(ctx).Create(&parentMenu).Error; err != nil {
		global.GVA_LOG.Error("创建AI模型中转一级菜单失败", zap.Error(err))
		return
	}

	global.GVA_LOG.Info("创建AI模型中转一级菜单成功", zap.Uint("id", parentMenu.ID))

	// 二级菜单列表
	subMenus := []system.SysBaseMenu{
		{
			GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			ParentId:  parentMenu.ID,
			Path:      "token",
			Name:      "token",
			Hidden:    false,
			Component: "view/model_transfer/token/index.vue",
			Sort:      1,
			Meta: system.Meta{
				Title:       "Token管理",
				Icon:        "key",
				KeepAlive:   true,
				DefaultMenu: false,
				CloseTab:    false,
			},
		},
		{
			GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			ParentId:  parentMenu.ID,
			Path:      "report",
			Name:      "report",
			Hidden:    false,
			Component: "view/model_transfer/report/index.vue",
			Sort:      2,
			Meta: system.Meta{
				Title:       "使用报表",
				Icon:        "data-analysis",
				KeepAlive:   true,
				DefaultMenu: false,
				CloseTab:    false,
			},
		},
		{
			GVA_MODEL: global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			ParentId:  parentMenu.ID,
			Path:      "daily",
			Name:      "daily",
			Hidden:    false,
			Component: "view/model_transfer/daily/index.vue",
			Sort:      3,
			Meta: system.Meta{
				Title:       "日报统计",
				Icon:        "pie-chart",
				KeepAlive:   true,
				DefaultMenu: false,
				CloseTab:    false,
			},
		},
	}

	// 批量创建二级菜单
	for _, menu := range subMenus {
		if err := global.GVA_DB.WithContext(ctx).Create(&menu).Error; err != nil {
			global.GVA_LOG.Error("创建子菜单失败", zap.String("name", menu.Name), zap.Error(err))
		} else {
			global.GVA_LOG.Info("创建子菜单成功", zap.String("name", menu.Name), zap.Uint("id", menu.ID))
		}
	}

	global.GVA_LOG.Info("AI模型中转菜单初始化完成")
}

// InitModelTransferAPI 初始化AI模型中转API
func InitModelTransferAPI() {
	ctx := context.Background()

	// 检查是否已经初始化
	var count int64
	if err := global.GVA_DB.Model(&system.SysApi{}).
		Where("api_group = ?", "AI模型中转").
		Count(&count).Error; err == nil && count > 0 {
		global.GVA_LOG.Info("AI模型中转API已存在，跳过初始化")
		return
	}

	// API列表
	apis := []system.SysApi{
		// Token管理相关API
		{
			GVA_MODEL:   global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/model-transfer/token/create",
			Description: "创建Token",
			ApiGroup:    "AI模型中转",
			Method:      "POST",
		},
		{
			GVA_MODEL:   global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/model-transfer/token/update",
			Description: "更新Token",
			ApiGroup:    "AI模型中转",
			Method:      "POST",
		},
		{
			GVA_MODEL:   global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/model-transfer/token/delete",
			Description: "删除Token",
			ApiGroup:    "AI模型中转",
			Method:      "POST",
		},
		{
			GVA_MODEL:   global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/model-transfer/token/list",
			Description: "Token列表",
			ApiGroup:    "AI模型中转",
			Method:      "POST",
		},
		{
			GVA_MODEL:   global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/model-transfer/token/detail",
			Description: "Token详情",
			ApiGroup:    "AI模型中转",
			Method:      "POST",
		},
		{
			GVA_MODEL:   global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/model-transfer/token/regenerate",
			Description: "重新生成Token",
			ApiGroup:    "AI模型中转",
			Method:      "POST",
		},
		// 报表相关API
		{
			GVA_MODEL:   global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/model-transfer/report/daily",
			Description: "日报查询",
			ApiGroup:    "AI模型中转",
			Method:      "POST",
		},
		{
			GVA_MODEL:   global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/model-transfer/report/token-usage",
			Description: "Token使用详情",
			ApiGroup:    "AI模型中转",
			Method:      "POST",
		},
		{
			GVA_MODEL:   global.GVA_MODEL{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Path:        "/model-transfer/report/summary",
			Description: "汇总报表",
			ApiGroup:    "AI模型中转",
			Method:      "POST",
		},
	}

	// 批量创建API
	for _, api := range apis {
		if err := global.GVA_DB.WithContext(ctx).Create(&api).Error; err != nil {
			global.GVA_LOG.Error("创建API失败", zap.String("path", api.Path), zap.Error(err))
		} else {
			global.GVA_LOG.Info("创建API成功", zap.String("path", api.Path), zap.Uint("id", api.ID))
		}
	}

	global.GVA_LOG.Info("AI模型中转API初始化完成")
}

// InitModelTransferData 初始化AI模型中转所有数据
func InitModelTransferData() {
	global.GVA_LOG.Info("开始初始化AI模型中转数据")

	// 初始化菜单
	InitModelTransferMenu()

	// 初始化API
	InitModelTransferAPI()

	global.GVA_LOG.Info("AI模型中转数据初始化完成")
}
