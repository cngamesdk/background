package operation_management

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type GamePackagingConfigService struct {
}

func (receiver *GamePackagingConfigService) List(ctx context.Context, req *api.GamePackagingConfigListReq) (
	resp interface{}, total int64, err error) {
	model := operation_management.NewDimGamePackagingConfigModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.PlatformId > 0 {
		tmpDb.Where("platform_id = ?", req.PlatformId)
	}
	if req.GameId > 0 {
		tmpDb.Where("game_id = ?", req.GameId)
	}
	if req.CommonMedia != "" {
		tmpDb.Where("common_media = ?", req.CommonMedia)
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []common.DimGamePackagingConfigModel
	if listErr := tmpDb.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (receiver *GamePackagingConfigService) Add(ctx context.Context, req *api.GamePackagingConfigAddReq) (
	resp api.GamePackagingConfigAddResp, err error) {
	model := operation_management.NewDimGamePackagingConfigModel()
	req.DimGamePackagingConfigModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存失败", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *GamePackagingConfigService) Modify(ctx context.Context, req *api.GamePackagingConfigModifyReq) (
	resp api.GamePackagingConfigModifyResp, err error) {
	model := operation_management.NewDimGamePackagingConfigModel()
	req.DimGamePackagingConfigModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存失败", zap.Error(saveErr))
		return
	}
	return
}
