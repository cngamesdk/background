package initialize

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/task"
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

	}()
}
