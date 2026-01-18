package operation_management

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type GameAppVersionConfigurationService struct {
}

func (p *GameAppVersionConfigurationService) List(ctx context.Context, req *api.GameAppVersionConfigurationListReq) (
	resp interface{}, count int64, err error) {
	alias := "config"
	model := operation_management.NewDimGameAppVersionConfiguration()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if countErr := tmpDb.Count(&count).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	model2.JoinGame(tmpDb, alias)
	model2.JoinProductCommonConfig(tmpDb, alias)
	var list []operation_management.DimGameAppVersionConfiguration
	if findErr := tmpDb.
		Select(alias + ".*,platform.platform_name,game.game_name,common_config.config_name").
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&list).Error; findErr != nil {
		err = findErr
		return
	}
	resp = list
	return
}

func (p *GameAppVersionConfigurationService) Add(ctx context.Context, req *api.GameAppVersionConfigurationAddReq) (
	resp api.GameAppVersionConfigurationAddResp, err error) {
	model := operation_management.NewDimGameAppVersionConfiguration()
	req.DimGameAppVersionConfiguration.DimGameAppVersionConfiguration.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (p *GameAppVersionConfigurationService) Modify(ctx context.Context, req *api.GameAppVersionConfigurationModifyReq) (
	resp api.GameAppVersionConfigurationModifyResp, err error) {
	model := operation_management.NewDimGameAppVersionConfiguration()
	req.DimGameAppVersionConfiguration.DimGameAppVersionConfiguration.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
