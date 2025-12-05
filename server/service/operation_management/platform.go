package operation_management

import (
	"context"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type PlatformService struct {

}

func (receiver *PlatformService) List(ctx context.Context, req *api.PlatformListReq) (resp interface{}, total int64, err error) {
	model := operation_management.NewDimPlatformModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.PlatformName != "" {
		if validator.IsNumberStr(req.PlatformName) {
			tmpDb.Where("id = ?", req.PlatformName)
		} else {
			tmpDb.Where("platform_name like ?", "%"+ req.PlatformName +"%")
		}
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []operation_management.DimPlatformModel
	if listErr := tmpDb.Limit(req.PageSize).Offset((req.Page - 1)*req.PageSize).Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (receiver *PlatformService) Add(ctx context.Context, req *api.PlatformAddReq) (resp api.PlatformAddResp, err error) {
	model := operation_management.NewDimPlatformModel()
	req.DimPlatformModel.DimPlatformModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *PlatformService) Modify(ctx context.Context, req *api.PlatformModifyReq) (resp api.PlatformModifyResp, err error) {
	model := operation_management.NewDimPlatformModel()
	req.DimPlatformModel.DimPlatformModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
