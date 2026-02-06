package material

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/material"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/material/api"
	"go.uber.org/zap"
)

type ThemeService struct {
}

func (receiver *ThemeService) List(ctx context.Context, req *api2.MaterialThemeListReq) (resp interface{}, total int64, err error) {
	model := material.NewDimMaterialThemeModel()
	alias := "theme"
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.ThemeName != "" {
		tmpDb.Where("theme_name like ?", "%"+req.ThemeName+"%")
	}
	tmpDb.Where("parent_id = ?", req.ParentId)
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	var list []api2.MaterialThemeListResp
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

func (receiver *ThemeService) Add(ctx context.Context, req *api2.MaterialThemeAddReq) (resp api2.MaterialThemeAddResp, err error) {
	model := material.NewDimMaterialThemeModel()
	req.DimMaterialThemeModel.DimMaterialThemeModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *ThemeService) Modify(ctx context.Context, req *api2.MaterialThemeModifyReq) (resp api2.MaterialThemeModifyResp, err error) {
	model := material.NewDimMaterialThemeModel()
	req.DimMaterialThemeModel.DimMaterialThemeModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}
