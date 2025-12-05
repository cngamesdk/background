package advertising

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"go.uber.org/zap"
)

type SiteService struct {
}

func (receiver *SiteService) List(ctx context.Context, req *api2.SiteListReq) (resp interface{}, total int64, err error) {
	model := advertising.NewDimSiteModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.SiteName != "" {
		tmpDb.Where("id = ? or site_name like ?", req.SiteName, "%"+ req.SiteName +"%")
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []advertising.DimSiteModel
	if listErr := tmpDb.Limit(req.PageSize).Offset((req.Page - 1)*req.PageSize).Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (receiver *SiteService) Add(ctx context.Context, req *api2.SiteAddReq) (resp api2.SiteAddResp, err error) {
	model := advertising.NewDimSiteModel()
	req.DimSiteModel.DimSiteModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *SiteService) Modify(ctx context.Context, req *api2.SiteModifyReq) (resp api2.SiteModifyResp, err error) {
	model := advertising.NewDimSiteModel()
	req.DimSiteModel.DimSiteModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}