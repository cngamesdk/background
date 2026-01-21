package advertising

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"go.uber.org/zap"
)

type AdvertisingMediaService struct {
}

func (receiver *AdvertisingMediaService) List(ctx context.Context, req *api2.AdvertisingMediaListReq) (resp interface{}, total int64, err error) {
	alias := "ad"
	model := advertising.NewDimAdvertisingMediaModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.AdvertisingMediaName != "" {
		tmpDb.Where("id = ? or advertising_media_name like ?", req.AdvertisingMediaName, "%"+req.AdvertisingMediaName+"%")
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	var list []advertising.DimAdvertisingMediaModel
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

func (receiver *AdvertisingMediaService) Add(ctx context.Context, req *api2.AdvertisingMediaAddReq) (resp api2.AdvertisingMediaAddResp, err error) {
	model := advertising.NewDimAdvertisingMediaModel()
	req.DimAdvertisingMediaModel.DimAdvertisingMediaModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *AdvertisingMediaService) Modify(ctx context.Context, req *api2.AdvertisingMediaModifyReq) (resp api2.AdvertisingMediaModifyResp, err error) {
	model := advertising.NewDimAdvertisingMediaModel()
	req.DimAdvertisingMediaModel.DimAdvertisingMediaModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
