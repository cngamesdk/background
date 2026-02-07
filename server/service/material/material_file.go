package material

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/material"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/material/api"
	"go.uber.org/zap"
)

type MaterialFileService struct {
}

func (receiver *MaterialFileService) List(ctx context.Context, req *api2.MaterialFileListReq) (resp interface{}, total int64, err error) {
	model := material.NewOdsMaterialFileLogModel()
	alias := "theme"
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	tmpDb.Where("status != ?", sql.StatusDelete)
	tmpDb.Where("material_id = ?", req.MaterialId)
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	var list []api2.MaterialListResp
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

func (receiver *MaterialFileService) Add(ctx context.Context, req *api2.MaterialFileAddReq) (resp api2.MaterialFileAddResp, err error) {
	model := material.NewOdsMaterialFileLogModel()
	req.OdsMaterialFileLogModel.OdsMaterialFileLogModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *MaterialFileService) Modify(ctx context.Context, req *api2.MaterialFileModifyReq) (resp api2.MaterialFileModifyResp, err error) {
	model := material.NewOdsMaterialFileLogModel()
	req.OdsMaterialFileLogModel.OdsMaterialFileLogModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
