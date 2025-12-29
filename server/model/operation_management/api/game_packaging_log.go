package api

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type GamePackagingLogListReq struct {
	request.PageInfo
	operation_management.OdsGamePackagingLogModel
}

type GamePackagingLogListResp struct {
}

type GamePackagingAddReq struct {
	operation_management.OdsGamePackagingLogModel
}

func (receiver *GamePackagingAddReq) Format() {
	receiver.Id = 0
	receiver.Status = sql.StatusNotStarted
	if receiver.AgentId <= 0 {
		model := advertising.NewDimSiteModel()
		siteInfoErr := model.Take(context.Background(), "*", "id = ?", receiver.SiteId)
		if siteInfoErr != nil {
			global.GVA_LOG.Error("获取广告位信息异常", zap.Error(siteInfoErr))
		}
		receiver.AgentId = model.AgentId
	}
}

func (receiver *GamePackagingAddReq) Validate(ctx context.Context) (err error) {
	if receiver.PlatformId <= 0 {
		err = errors.New("平台不能为空")
		return
	}
	if receiver.GameId <= 0 {
		err = errors.New("游戏ID不能为空")
		return
	}
	if receiver.AgentId <= 0 {
		err = errors.New("渠道ID不能为空")
		return
	}
	if receiver.SiteId <= 0 {
		err = errors.New("广告位ID不能为空")
		return
	}
	//验证是否配置打包工具
	globalCommonConfigModel := operation_management.NewDimGlobalCommonConfigModel()
	configInfoErr := globalCommonConfigModel.Take(ctx, "*", "platform_id = ?", receiver.PlatformId)
	if configInfoErr != nil {
		err = configInfoErr
		global.GVA_LOG.Error("获取配置异常", zap.Error(configInfoErr))
		return
	}
	if globalCommonConfigModel.GamePackagingToolPath == "" {
		err = errors.New("请先配置打包工具.")
		global.GVA_LOG.Warn("未获取到打包工具", zap.Any("data", globalCommonConfigModel))
		return
	}

	//获取所属媒体
	agentModel := &advertising.DimAgentDetailInfoModel{}
	agentErr := agentModel.GetAgentDetailInfoByAgentId(ctx, receiver.AgentId)
	if agentErr != nil {
		err = agentErr
		global.GVA_LOG.Error("获取信息异常", zap.Error(agentErr))
		return
	}
	//验证游戏对应的媒体是否存在可使用的母包
	gamePackagingConfigModel := operation_management.NewDimGamePackagingConfigModel()
	if takeErr := gamePackagingConfigModel.Take(ctx, "*", "platform_id = ? and game_id = ? and common_media = ? and use_status = ?",
		receiver.PlatformId, receiver.GameId, agentModel.CommonMedia, common.UseStatusNormal); takeErr != nil {
		err = takeErr
		global.GVA_LOG.Error("获取信息异常", zap.Error(takeErr))
		return
	}
	return
}
