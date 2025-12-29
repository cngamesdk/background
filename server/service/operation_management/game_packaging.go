package operation_management

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type GamePackagingService struct {
}

func (g *GamePackagingService) LogList(ctx context.Context, req *api.GamePackagingLogListReq) (
	resp interface{}, total int64, err error) {
	model := operation_management.NewOdsGamePackagingLogModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.PlatformId > 0 {
		tmpDb.Where("platform_id = ?", req.PlatformId)
	}
	if req.GameId > 0 {
		tmpDb.Where("game_id = ?", req.GameId)
	}
	if req.AgentId > 0 {
		tmpDb.Where("agent_id = ?", req.AgentId)
	}
	if req.SiteId > 0 {
		tmpDb.Where("site_id = ?", req.SiteId)
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []operation_management.OdsGamePackagingLogModel
	if listErr := tmpDb.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Order("id DESC").Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (g *GamePackagingService) Add(ctx context.Context, req *api.GamePackagingAddReq) (
	resp api.GamePackagingConfigAddResp, err error) {
	model := operation_management.NewOdsGamePackagingLogModel()
	req.OdsGamePackagingLogModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("创建异常", zap.Error(saveErr))
		return
	}
	return
}
