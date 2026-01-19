package operation_management

import (
	"context"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type PublishingService struct {
}

func (p *PublishingService) ChannelConfigList(ctx context.Context, req *api.PublishingChannelConfigListReq) (
	resp interface{}, total int64, err error) {
	alias := "channel"
	model := operation_management.NewDimPublishingChannelConfigModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.ChannelName != "" {
		if validator.IsNumberStr(req.ChannelName) {
			tmpDb.Where("id = ?", req.ChannelName)
		} else {
			tmpDb.Where("channel_name like ?", "%"+req.ChannelName+"%")
		}
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	var list []operation_management.DimPublishingChannelConfigModel
	if listErr := tmpDb.
		Select(alias + ".*,platform_name").
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (p *PublishingService) ChannelConfigAdd(ctx context.Context, req *api.PublishingChannelConfigAddReq) (
	resp api.PublishingChannelConfigAddResp, err error) {
	model := operation_management.NewDimPublishingChannelConfigModel()
	req.DimPublishingChannelConfigModel.DimPublishingChannelConfigModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (p *PublishingService) ChannelConfigModify(ctx context.Context, req *api.PublishingChannelConfigModifyReq) (
	resp api.PublishingChannelConfigModifyResp, err error) {
	model := operation_management.NewDimPublishingChannelConfigModel()
	req.DimPublishingChannelConfigModel.DimPublishingChannelConfigModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (p *PublishingService) ChannelGameConfigList(ctx context.Context, req *api.PublishingChannelGameConfigListReq) (
	resp interface{}, total int64, err error) {
	alias := "config"
	model := operation_management.NewDimPublishingChannelGameConfigModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	model2.JoinGame(tmpDb, alias)
	model2.JoinCoPublishing(tmpDb, alias)
	model2.JoinAgent(tmpDb, alias)
	model2.JoinSite(tmpDb, alias)
	var list []operation_management.DimPublishingChannelGameConfigModel
	if listErr := tmpDb.
		Select(alias + ".*,platform_name,channel_name,agent_name,site_name,game_name").
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (p *PublishingService) ChannelGameConfigAdd(ctx context.Context, req *api.PublishingChannelGameConfigAddReq) (
	resp api.PublishingChannelGameConfigAddResp, err error) {
	model := operation_management.NewDimPublishingChannelGameConfigModel()
	req.DimPublishingChannelGameConfigModel.DimPublishingChannelGameConfigModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (p *PublishingService) ChannelGameConfigModify(ctx context.Context, req *api.PublishingChannelGameConfigModifyReq) (
	resp api.PublishingChannelGameConfigModifyResp, err error) {
	model := operation_management.NewDimPublishingChannelGameConfigModel()
	req.DimPublishingChannelGameConfigModel.DimPublishingChannelGameConfigModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
