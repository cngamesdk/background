package operation_management

import (
	"context"
	"fmt"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type PayChannelSwitchService struct {
}

func (p *PayChannelSwitchService) List(ctx context.Context, req *api.PayChannelSwitchListReq) (
	resp interface{}, total int64, err error) {
	alias := "pay"
	model := operation_management.NewDimPayChannelSwitchModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	tmpDb.Where("status <> ?", sql.StatusDelete)
	if req.PayType != "" {
		tmpDb.Where("pay_type = ?", req.PayType)
	}
	if req.PayChannelId > 0 {
		tmpDb.Where("pay_channels LIKE ?", "%"+fmt.Sprintf("\"pay_channel_id\": %d", req.PayChannelId)+"%")
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	var list []operation_management.DimPayChannelSwitchModel
	if listErr := tmpDb.Select(alias + ".*,platform_name").
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

func (p *PayChannelSwitchService) Add(ctx context.Context, req *api.PayChannelSwitchAddReq) (
	resp api.PayChannelSwitchAddReq, err error) {
	model := operation_management.NewDimPayChannelSwitchModel()
	req.DimPayChannelSwitchModel.DimPayChannelSwitchModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (p *PayChannelSwitchService) Modify(ctx context.Context, req *api.PayChannelSwitchModifyReq) (
	resp api.PayChannelSwitchModifyResp, err error) {
	model := operation_management.NewDimPayChannelSwitchModel()
	req.DimPayChannelSwitchModel.DimPayChannelSwitchModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
