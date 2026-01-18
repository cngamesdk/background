package operation_management

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type ProductCommonConfigurationService struct {
}

func (p *ProductCommonConfigurationService) List(ctx context.Context, req *api.ProductCommonConfigurationListReq) (
	resp interface{}, count int64, err error) {
	alias := "config"
	model := operation_management.NewDimProductCommonConfigurationModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.ConfigName != "" {
		tmpDb.Where("config_name like ?", "%"+req.ConfigName+"%")
	}
	if countErr := tmpDb.Count(&count).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	var list []operation_management.DimProductCommonConfigurationModel
	if findErr := tmpDb.
		Select(alias + ".*,platform.platform_name").
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&list).Error; findErr != nil {
		err = findErr
		return
	}
	resp = list
	return
}

func (p *ProductCommonConfigurationService) Add(ctx context.Context, req *api.ProductCommonConfigurationAddReq) (
	resp api.ProductCommonConfigurationAddResp, err error) {
	model := operation_management.NewDimProductCommonConfigurationModel()
	req.DimProductCommonConfigurationModel.DimProductCommonConfigurationModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (p *ProductCommonConfigurationService) Modify(ctx context.Context, req *api.ProductCommonConfigurationModifyReq) (
	resp api.ProductCommonConfigurationModifyResp, err error) {
	model := operation_management.NewDimProductCommonConfigurationModel()
	req.DimProductCommonConfigurationModel.DimProductCommonConfigurationModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
