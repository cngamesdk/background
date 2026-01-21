package advertising

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"go.uber.org/zap"
)

type AgentService struct {
}

func (receiver *AgentService) List(ctx context.Context, req *api2.AgentListReq) (resp interface{}, total int64, err error) {
	alias := "agent"
	model := advertising.NewDimAgentModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.AgentName != "" {
		tmpDb.Where("id = ? or agent_name like ?", req.AgentName, "%"+req.AgentName+"%")
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	model2.JoinChannelGroup(tmpDb, alias)
	var list []advertising.DimAgentModel
	if listErr := tmpDb.
		Select(alias + ".*,platform_name,channel_group_name").
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

func (receiver *AgentService) Add(ctx context.Context, req *api2.AgentAddReq) (resp api2.AgentAddResp, err error) {
	model := advertising.NewDimAgentModel()
	req.DimAgentModel.DimAgentModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *AgentService) Modify(ctx context.Context, req *api2.AgentModifyReq) (resp api2.AgentModifyResp, err error) {
	model := advertising.NewDimAgentModel()
	req.DimAgentModel.DimAgentModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
