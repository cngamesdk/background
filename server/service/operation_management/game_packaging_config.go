package operation_management

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	model2 "github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GamePackagingConfigService struct {
}

func (receiver *GamePackagingConfigService) List(ctx context.Context, req *api.GamePackagingConfigListReq) (
	resp interface{}, total int64, err error) {
	alias := "config"
	model := operation_management.NewDimGamePackagingConfigModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName() + " as " + alias)
	if req.PlatformId > 0 {
		tmpDb.Where("platform_id = ?", req.PlatformId)
	}
	if req.GameId > 0 {
		tmpDb.Where("game_id = ?", req.GameId)
	}
	if req.MediaCode != "" {
		tmpDb.Where("media_code = ?", req.MediaCode)
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	model2.JoinPlatform(tmpDb, alias)
	model2.JoinGame(tmpDb, alias)
	var list []operation_management.DimGamePackagingConfigModel
	if listErr := tmpDb.
		Select(alias + ".*,platform_name,game_name").
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (receiver *GamePackagingConfigService) Add(ctx context.Context, req *api.GamePackagingConfigAddReq) (
	resp api.GamePackagingConfigAddResp, err error) {
	transactionErr := global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		//当前为使用时，该游戏的其他版本下架，只保留一条使用中
		if req.UseStatus == common.UseStatusNormal {
			updateModel := operation_management.NewDimGamePackagingConfigModel()
			updateModel.UseStatus = common.UseStatusRemove
			updateModel.DimGamePackagingConfigModel.Db = func() *gorm.DB {
				return tx
			}
			if updateErr := updateModel.Updates(ctx, "platform_id = ? and game_id = ? and media_code = ? and use_status = ? ",
				req.PlatformId, req.GameId, req.MediaCode, common.UseStatusNormal); updateErr != nil {
				global.GVA_LOG.Error("更新异常", zap.Error(updateErr))
				return updateErr
			}
		}

		req.DimGamePackagingConfigModel.Db = func() *gorm.DB {
			return tx
		}
		if saveErr := req.Create(ctx); saveErr != nil {
			global.GVA_LOG.Error("保存失败", zap.Error(saveErr))
			return saveErr
		}

		return nil
	})
	err = transactionErr
	global.GVA_LOG.Error("事务异常", zap.Error(transactionErr))
	return
}

func (receiver *GamePackagingConfigService) Modify(ctx context.Context, req *api.GamePackagingConfigModifyReq) (
	resp api.GamePackagingConfigModifyResp, err error) {

	transactionErr := global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		//当前为使用时，该游戏的其他版本下架，只保留一条使用中
		if req.UseStatus == common.UseStatusNormal {
			updateModel := operation_management.NewDimGamePackagingConfigModel()
			updateModel.UseStatus = common.UseStatusRemove
			updateModel.DimGamePackagingConfigModel.Db = func() *gorm.DB {
				return tx
			}
			if updateErr := updateModel.Updates(ctx, "platform_id = ? and game_id = ? and media_code = ? and use_status = ? ",
				req.PlatformId, req.GameId, req.MediaCode, common.UseStatusNormal); updateErr != nil {
				global.GVA_LOG.Error("更新异常", zap.Error(updateErr))
				return updateErr
			}
		}

		req.DimGamePackagingConfigModel.Db = func() *gorm.DB {
			return tx
		}
		if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
			global.GVA_LOG.Error("保存失败", zap.Error(saveErr))
			return saveErr
		}
		return nil
	})
	err = transactionErr
	global.GVA_LOG.Error("事务异常", zap.Error(transactionErr))
	return
}
