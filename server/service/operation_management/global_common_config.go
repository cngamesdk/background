package operation_management

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type GlobalCommonConfigService struct {
}

func (receiver *GlobalCommonConfigService) List(ctx context.Context, req *api.GlobalCommonConfigListReq) (
	resp interface{}, total int64, err error) {
	model := operation_management.NewDimGlobalCommonConfigModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.PlatformId > 0 {
		tmpDb.Where("platform_id = ?", req.PlatformId)
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []operation_management.DimGlobalCommonConfigModel
	if listErr := tmpDb.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (receiver *GlobalCommonConfigService) Add(ctx context.Context, req *api.GlobalCommonConfigAddReq) (
	resp api.GlobalCommonConfigAddResp, err error) {
	model := operation_management.NewDimGlobalCommonConfigModel()
	req.DimGlobalCommonConfigModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存失败", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *GlobalCommonConfigService) Modify(ctx context.Context, req *api.GlobalCommonConfigModifyReq) (
	resp api.GlobalCommonConfigModifyResp, err error) {
	model := operation_management.NewDimGlobalCommonConfigModel()
	req.DimGlobalCommonConfigModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存失败", zap.Error(saveErr))
		return
	}
	return
}
