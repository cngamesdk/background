package advertising

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"go.uber.org/zap"
)

type AdvertisingDeveloperConfigService struct {
}

func (receiver *AdvertisingDeveloperConfigService) List(ctx context.Context, req *api2.AdvertisingDeveloperConfigListReq) (resp interface{}, total int64, err error) {
	alias := "config"
	model := advertising.NewDimAdvertisingDeveloperConfigModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.Name != "" {
		tmpDb.Where("name like ?", "%"+req.Name+"%")
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	model2.JoinCompany(tmpDb, alias)
	model2.JoinMediaByCode(tmpDb, alias)
	var list []api2.AdvertisingDeveloperConfigListResp
	if listErr := tmpDb.
		Select(alias + ".*,platform_name,company_name,media_name").
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

func (receiver *AdvertisingDeveloperConfigService) Add(ctx context.Context, req *api2.AdvertisingDeveloperConfigAddReq) (resp api2.AdvertisingDeveloperConfigAddResp, err error) {
	model := advertising.NewDimAdvertisingDeveloperConfigModel()
	req.DimAdvertisingDeveloperConfigModel.DimAdvertisingDeveloperConfigModel.Db = model.Db
	req.DimAdvertisingDeveloperConfigModel.DimAdvertisingDeveloperConfigModel.AesKey = model.AesKey
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *AdvertisingDeveloperConfigService) Modify(ctx context.Context, req *api2.AdvertisingDeveloperConfigModifyReq) (resp api2.AdvertisingDeveloperConfigModifyResp, err error) {
	model := advertising.NewDimAdvertisingDeveloperConfigModel()
	req.DimAdvertisingDeveloperConfigModel.DimAdvertisingDeveloperConfigModel.Db = model.Db
	req.DimAdvertisingDeveloperConfigModel.DimAdvertisingDeveloperConfigModel.AesKey = model.AesKey
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
