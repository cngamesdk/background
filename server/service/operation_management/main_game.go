package operation_management

import (
	"context"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type MainGameService struct {
	
}

// List 主游戏列表
func (receiver *MainGameService) List(ctx context.Context, req *api.MainGameListReq) (resp interface{}, total int64, err error) {
	model := operation_management.NewDimMainGameModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.GameName != "" {
		if validator.IsNumberStr(req.GameName) {
			tmpDb.Where("id = ?", req.GameName)
		} else {
			tmpDb.Where("game_name like ?", "%"+ req.GameName +"%")
		}
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []operation_management.DimMainGameModel
	if listErr := tmpDb.Limit(req.PageSize).Offset((req.Page -1)*req.PageSize).Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取列表失败", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (receiver *MainGameService) Add(ctx context.Context, req *api.MainGameAddReq) (resp api.MainGameAddResp, err error) {
	model := operation_management.NewDimMainGameModel()
	req.DimMainGameModel.DimMainGameModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("创建异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *MainGameService) Modify(ctx context.Context, req *api.MainGameModifyReq) (resp api.MainGameModifyResp, err error) {
	model := operation_management.NewDimMainGameModel()
	req.DimMainGameModel.DimMainGameModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}