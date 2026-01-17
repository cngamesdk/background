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

type CompanyService struct {
}

func (receiver *CompanyService) List(ctx context.Context, req *api.CompanyListReq) (resp interface{}, total int64, err error) {
	model := operation_management.NewDimCompanyModel()
	alias := "company"
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.CompanyName != "" {
		if validator.IsNumberStr(req.CompanyName) {
			tmpDb.Where("id = ?", req.CompanyName)
		} else {
			tmpDb.Where("company_name like ?", "%"+req.CompanyName+"%")
		}
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	var list []operation_management.DimCompanyModel
	if listErr := tmpDb.
		Select(alias + ".*,platform.platform_name").
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

func (receiver *CompanyService) Add(ctx context.Context, req *api.CompanyAddReq) (resp api.CompanyAddResp, err error) {
	model := operation_management.NewDimCompanyModel()
	req.DimCompanyModel.DimCompanyModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *CompanyService) Modify(ctx context.Context, req *api.CompanyModifyReq) (resp api.CompanyModifyResp, err error) {
	model := operation_management.NewDimCompanyModel()
	req.DimCompanyModel.DimCompanyModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
