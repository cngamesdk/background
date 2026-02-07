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

type MaterialService struct {
}

func (receiver *MaterialService) List(ctx context.Context, req *api2.MaterialListReq) (resp interface{}, total int64, err error) {
	model := material.NewOdsMaterialLogModel()
	alias := "theme"
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	tmpDb.Where("status != ?", sql.StatusDelete)
	if req.MaterialName != "" {
		tmpDb.Where("theme_name like ?", "%"+req.MaterialName+"%")
	}
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

func (receiver *MaterialService) Add(ctx context.Context, req *api2.MaterialAddReq) (resp api2.MaterialAddResp, err error) {
	model := material.NewOdsMaterialLogModel()
	req.OdsMaterialLogModel.OdsMaterialLogModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *MaterialService) Modify(ctx context.Context, req *api2.MaterialModifyReq) (resp api2.MaterialModifyResp, err error) {
	model := material.NewOdsMaterialLogModel()
	req.OdsMaterialLogModel.OdsMaterialLogModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
