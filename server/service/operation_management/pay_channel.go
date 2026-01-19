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

type PayChannelService struct {
}

func (p *PayChannelService) List(ctx context.Context, req *api.PayChannelListReq) (resp interface{}, total int64, err error) {
	alias := "pay"
	model := operation_management.NewDimPayChannelModel()
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
	model2.JoinCompany(tmpDb, alias)
	var list []operation_management.DimPayChannelModel
	if listErr := tmpDb.
		Select(alias + ".*,platform_name,company_name").
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

func (p *PayChannelService) Add(ctx context.Context, req *api.PayChannelAddReq) (resp api.PayChannelAddResp, err error) {
	model := operation_management.NewDimPayChannelModel()
	req.DimPayChannelModel.DimPayChannelModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (p *PayChannelService) Modify(ctx context.Context, req *api.PayChannelModifyReq) (resp api.PayChannelModifyResp, err error) {
	model := operation_management.NewDimPayChannelModel()
	req.DimPayChannelModel.DimPayChannelModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
