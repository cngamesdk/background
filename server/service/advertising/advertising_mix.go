package advertising

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"go.uber.org/zap"
)

type AdvertisingMixService struct {
}

func (receiver *AdvertisingMixService) List(ctx context.Context, req *api2.AdvertisingMixListReq) (resp interface{}, total int64, err error) {
	alias := "mix"
	model := advertising.NewOdsAdvertisingMixLogModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.Id > 0 {
		tmpDb.Where("id = ?", req.Id)
	}
	if req.Name != "" {
		tmpDb.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Code != "" {
		tmpDb.Where("code = ?", req.Code)
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	model2.JoinMediaByCode(tmpDb, alias)
	var list []advertising.OdsAdvertisingMixLogModel
	if listErr := tmpDb.
		Select(alias + ".*,platform_name,media_name").
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

func (receiver *AdvertisingMixService) Add(ctx context.Context, req *api2.AdvertisingMixAddReq) (resp api2.AdvertisingMixAddResp, err error) {
	model := advertising.NewOdsAdvertisingMixLogModel()
	req.OdsAdvertisingMixLogModel.OdsAdvertisingMixLogModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *AdvertisingMixService) Modify(ctx context.Context, req *api2.AdvertisingMixModifyReq) (resp api2.AdvertisingMixModifyResp, err error) {
	model := advertising.NewOdsAdvertisingMixLogModel()
	req.OdsAdvertisingMixLogModel.OdsAdvertisingMixLogModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
