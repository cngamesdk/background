package advertising

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"go.uber.org/zap"
)

type ChannelGroupService struct {
}

func (receiver *ChannelGroupService) List(ctx context.Context, req *api2.ChannelGroupListReq) (resp interface{}, total int64, err error) {
	alias := "c_g"
	model := advertising.NewDimChannelGroupModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.ChannelGroupName != "" {
		tmpDb.Where("id = ? or channel_group_name like ?", req.ChannelGroupName, "%"+req.ChannelGroupName+"%")
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	model2.JoinMedia(tmpDb, alias)
	var list []advertising.DimChannelGroupModel
	if listErr := tmpDb.
		Select(alias + ".*,platform_name,advertising_media_name as media_name").
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

func (receiver *ChannelGroupService) Add(ctx context.Context, req *api2.ChannelGroupAddReq) (resp api2.ChannelGroupAddResp, err error) {
	model := advertising.NewDimChannelGroupModel()
	req.DimChannelGroupModel.DimChannelGroupModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *ChannelGroupService) Modify(ctx context.Context, req *api2.ChannelGroupModifyReq) (resp api2.ChannelGroupModifyResp, err error) {
	model := advertising.NewDimChannelGroupModel()
	req.DimChannelGroupModel.DimChannelGroupModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
