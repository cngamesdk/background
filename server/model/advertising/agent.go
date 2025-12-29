package advertising

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DimAgentModel struct {
	common.DimAgentModel
}

func NewDimAgentModel() *DimAgentModel {
	model := &DimAgentModel{}
	model.DimAgentModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}

type DimAgentDetailInfoModel struct {
	DimAgentModel
	ChannelGroupName     string `json:"channel_group_name"`
	AdvertisingMediaName string `json:"advertising_media_name"`
	AdvertisingMediaId   int64  `json:"advertising_media_id"`
	CommonMedia          string `json:"common_media"`
}

func (receiver *DimAgentDetailInfoModel) GetAgentDetailInfoByAgentId(ctx context.Context, agentId int64) (err error) {
	agentModel := NewDimAgentModel()
	channelGroupModel := NewDimChannelGroupModel()
	mediaModel := NewDimAdvertisingMediaModel()
	takeErr := global.GVA_DB.WithContext(ctx).
		Select("agent.*,channel_group.channel_group_name as channel_group_name, media.id as advertising_media_id, media.belong_common_media as common_media").
		Table(agentModel.TableName()+" as agent").
		Joins("join "+channelGroupModel.TableName()+" as channel_group on agent.platform_id = channel_group.platform_id and agent.channel_group_id = channel_group.id").
		Joins("join "+mediaModel.TableName()+" as media on  channel_group.platform_id = media.platform_id and channel_group.advertising_media_id = media.id").
		Where("agent.id = ?", agentId).Take(receiver).Error
	if takeErr != nil {
		err = takeErr
		global.GVA_LOG.Error("获取详情异常", zap.Error(takeErr))
		return
	}
	return
}
