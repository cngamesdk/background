package task

import (
	"bytes"
	"context"
	"fmt"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/duke-git/lancet/v2/cryptor"
	datastructure "github.com/duke-git/lancet/v2/datastructure/queue"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// 游戏打包任务

const (
	batchSize = 100
)

var (
	isRunning = false
	buckets   = make(map[int]*datastructure.ArrayQueue[operation_management.OdsGamePackagingLogModel])
)

func init() {
	processNum := getProcessNum()
	for index := 0; index < processNum; index++ {
		queue := datastructure.NewArrayQueue[operation_management.OdsGamePackagingLogModel](batchSize)
		buckets[index] = queue
		go consumeGamePackaging(queue)
	}
}

func getProcessNum() int {
	processNum := global.GVA_CONFIG.Common.GamePackaging.ProcessCount
	if processNum <= 0 {
		processNum = runtime.NumCPU()
	}
	return processNum
}

func consumeGamePackaging(queue *datastructure.ArrayQueue[operation_management.OdsGamePackagingLogModel]) {
	for {
		if queue.IsEmpty() {
			time.Sleep(time.Second * 2)
			continue
		}
		gamePackingItem, dequeueOk := queue.Dequeue()
		if !dequeueOk {
			global.GVA_LOG.Error("出队列失败")
			continue
		}
		if err := dealGamePackaging(gamePackingItem); err != nil {
			global.GVA_LOG.Error("打包异常", zap.Error(err), zap.Any("data", gamePackingItem))
		}
	}
}

func dealGamePackaging(req operation_management.OdsGamePackagingLogModel) (err error) {
	ctx := context.Background()
	execScript := ""
	execResultStr := ""
	gamePackagingPath := ""

	//收尾工作
	defer func() {
		status := sql.StatusSuccess
		if err != nil {
			status = sql.StatusFail
		}
		updateModel := operation_management.NewOdsGamePackagingLogModel()
		updateModel.Status = status
		updateModel.ExecCmd = execScript
		updateModel.ExecCmdResult = execResultStr
		updateModel.GamePackagePath = gamePackagingPath
		if updateErr := updateModel.Updates(ctx, "id = ?", req.Id); updateErr != nil {
			err = updateErr
			global.GVA_LOG.Error("更新收尾数据异常", zap.Error(updateErr))
			return
		}
		return
	}()

	//更新状态
	updateModel := operation_management.NewOdsGamePackagingLogModel()
	updateModel.Status = sql.StatusProcessing
	if updateErr := updateModel.Updates(ctx, "id = ?", req.Id); updateErr != nil {
		err = updateErr
		global.GVA_LOG.Error("更新状态异常", zap.Error(updateErr))
		return
	}
	//获取全局配置
	globalConfig := operation_management.NewDimGlobalCommonConfigModel()
	if takeErr := globalConfig.Take(ctx, "*", "platform_id = ?", req.PlatformId); takeErr != nil {
		err = takeErr
		global.GVA_LOG.Error("更新状态异常", zap.Error(takeErr))
		return
	}

	//读取渠道所属媒体
	agentDetailModel := &advertising.DimAgentDetailInfoModel{}
	if agentDetailErr := agentDetailModel.GetAgentDetailInfoByAgentId(ctx, req.AgentId); agentDetailErr != nil {
		err = agentDetailErr
		global.GVA_LOG.Error("获取渠道详情异常", zap.Error(agentDetailErr))
		return
	}
	gamePackagingConfig := operation_management.NewDimGamePackagingConfigModel()
	if takeErr := gamePackagingConfig.Take(ctx, "*", "platform_id = ? and game_id = ? and common_media = ? and status = ? and use_status = ?",
		req.PlatformId,
		req.GameId,
		agentDetailModel.CommonMedia,
		sql.StatusNormal,
		common.UseStatusNormal,
	); takeErr != nil {
		err = takeErr
		global.GVA_LOG.Error("获取游戏母包异常", zap.Error(takeErr))
		return
	}
	gameModel := operation_management.NewDimGameModel()
	if takeErr := gameModel.Take(ctx, "*", "id = ?", req.GameId); takeErr != nil {
		err = takeErr
		global.GVA_LOG.Error("获取游戏详情异常", zap.Error(takeErr))
		return
	}
	srcGameApk := gamePackagingConfig.GamePackagePath
	dstGameApkDir := fmt.Sprintf("%s/%d/%s",
		global.GVA_CONFIG.Common.GamePackaging.Path,
		req.PlatformId,
		cryptor.Md5String(gameModel.PackageName),
	)
	if !fileutil.IsExist(dstGameApkDir) {
		if createDirErr := fileutil.CreateDir(dstGameApkDir); createDirErr != nil {
			err = createDirErr
			global.GVA_LOG.Error("创建目录异常", zap.Error(createDirErr))
			return
		}
	}
	dstGameApk := fmt.Sprintf("%s/%d.apk",
		dstGameApkDir,
		req.SiteId,
	)

	execArgs := []string{
		"-jar",
		"-noverify",
		globalConfig.GamePackagingToolPath,
		"-v",
		"v2",
		"-a",
		cast.ToString(req.AgentId),
		"-s",
		cast.ToString(req.SiteId),
		srcGameApk,
		dstGameApk,
	}
	execScript = globalConfig.JavaExecutionPath + " " + strings.Join(execArgs, " ")
	cmd := exec.Command(globalConfig.JavaExecutionPath, execArgs...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误
	execErr := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	execResultStr = fmt.Sprintf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if execErr != nil {
		err = execErr
		global.GVA_LOG.Error("执行命令异常", zap.Error(execErr))
		return
	}
	gamePackagingPath = dstGameApk
	return
}

func GamePackaging(db *gorm.DB) (err error) {
	ctx := context.Background()
	if isRunning {
		return
	}
	isRunning = true
	processNum := global.GVA_CONFIG.Common.GamePackaging.ProcessCount

	for {
		//准备打包的日志
		model := operation_management.NewOdsGamePackagingLogModel()
		var list []operation_management.OdsGamePackagingLogModel
		if listErr := db.WithContext(ctx).
			Table(model.TableName()).
			Where("status = ?", sql.StatusNotStarted).
			Order("sort DESC, id DESC").
			Limit(batchSize).Find(&list).Error; listErr != nil {
			if !errors.Is(listErr, gorm.ErrRecordNotFound) {
				global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
			}
			time.Sleep(time.Second * 5)
			continue
		}

		//放入队列中
		for index, item := range list {
			bucketIndex := index % processNum
			buckets[bucketIndex].Enqueue(item)
		}
		time.Sleep(time.Second * 5)
	}

}
