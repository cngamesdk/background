package cron_task

import (
	"context"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cron_task"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cron_task/api"
	"go.uber.org/zap"
)

type CronTaskService struct {

}

func (receiver *CronTaskService) ConfigList(ctx context.Context, req *api.CronTaskListReq) (resp interface{}, total int64, err error) {
	model := cron_task.NewDimCronTaskConfigModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.Name != "" {
		if validator.IsNumberStr(req.Name) {
			tmpDb.Where("id = ?", req.Name)
		} else {
			tmpDb.Where("name like ?", "%"+ req.Name +"%")
		}
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []cron_task.DimCronTaskConfigModel
	if listErr := tmpDb.
		Limit(req.PageSize).
		Offset((req.Page - 1)*req.PageSize).
		Order("id DESC").
		Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (receiver *CronTaskService) ConfigAdd(ctx context.Context, req *api.CronTaskAddReq) (resp api.CronTaskAddResp, err error) {
	model := cron_task.NewDimCronTaskConfigModel()
	req.DimCronTaskConfigModel.DimCronTaskConfigModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *CronTaskService) ConfigModify(ctx context.Context, req *api.CronTaskModifyReq) (resp api.CronTaskModifyResp, err error) {
	model := cron_task.NewDimCronTaskConfigModel()
	req.DimCronTaskConfigModel.DimCronTaskConfigModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *CronTaskService) LogList(ctx context.Context, req *api.CronTaskLogListReq) (resp interface{}, total int64, err error) {
	model := cron_task.NewOdsCronTaskLogModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	tmpDb.Where("config_id = ?", req.ConfigId)
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []cron_task.OdsCronTaskLogModel
	if listErr := tmpDb.Limit(req.PageSize).Offset((req.Page - 1)*req.PageSize).Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}