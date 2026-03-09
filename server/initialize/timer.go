package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
	"github.com/flipped-aurora/gin-vue-admin/server/task/media/oceanengine"
	"go.uber.org/zap"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		//游戏打包
		gamePackagingEntryID, gamePackagingErr := global.GVA_Timer.AddTaskByFunc("GamePackaging", "@every 1m", func() {
			taskErr := task.GamePackaging(global.GVA_DB)
			if taskErr != nil {
				global.GVA_LOG.Error("添加游戏打包任务执行异常", zap.Error(taskErr))
			}
		}, "游戏打包任务", option...)
		if gamePackagingErr != nil {
			global.GVA_LOG.Error("添加游戏打包定时任务失败", zap.Error(gamePackagingErr))
		} else {
			global.GVA_LOG.Info("添加游戏打包定时任务启动成功", zap.Any("id", gamePackagingEntryID))
		}

		//刷新媒体token
		refreshTokenEntryID, refreshTokenErr := global.GVA_Timer.AddTaskByFunc("RefreshToken", "@every 1h", func() {
			taskErr := task.RefreshToken(global.GVA_DB)
			if taskErr != nil {
				global.GVA_LOG.Error("执行任务异常", zap.Error(taskErr))
			}
		}, "刷新媒体Token任务", option...)
		if refreshTokenErr != nil {
			global.GVA_LOG.Error("添加刷新Token任务失败", zap.Error(refreshTokenErr))
		} else {
			global.GVA_LOG.Info("添加刷新Token成功", zap.Any("id", refreshTokenEntryID))
		}

		//巨量引擎 - 获取头条安卓应用列表
		oceanengineAppTask := &task.Task{}
		oceanengineAppEntryID, oceanengineAppErr := global.GVA_Timer.AddTaskByFunc("OceanengineApp", "@every 1h", func() {
			taskErr := oceanengineAppTask.Run(global.GVA_DB, oceanengine.OceanengineAppList)
			if taskErr != nil {
				global.GVA_LOG.Error("执行任务异常", zap.Error(taskErr))
			}
		}, "获取头条安卓应用列表", option...)
		if oceanengineAppErr != nil {
			global.GVA_LOG.Error("获取头条安卓应用列表异常", zap.Error(oceanengineAppErr))
		} else {
			global.GVA_LOG.Info("获取头条安卓应用列表成功", zap.Any("id", oceanengineAppEntryID))
		}

		//巨量引擎 - 获取升级版巨量引擎工作台下账户列表
		oceanengineEbpAdvertiserTask := &task.Task{}
		oceanengineEbpAdvertiserEntryID, oceanengineEbpAdvertiserErr := global.GVA_Timer.AddTaskByFunc("OceanengineEbpAdvertiserList", "@every 10m", func() {
			taskErr := oceanengineEbpAdvertiserTask.Run(global.GVA_DB, oceanengine.EbpAdvertiserList)
			if taskErr != nil {
				global.GVA_LOG.Error("执行任务异常", zap.Error(taskErr))
			}
		}, "获取升级版巨量引擎工作台下账户列表", option...)
		if oceanengineEbpAdvertiserErr != nil {
			global.GVA_LOG.Error("获取升级版巨量引擎工作台下账户列表异常", zap.Error(oceanengineEbpAdvertiserErr))
		} else {
			global.GVA_LOG.Info("获取升级版巨量引擎工作台下账户列表成功", zap.Any("id", oceanengineEbpAdvertiserEntryID))
		}

	}()
}
