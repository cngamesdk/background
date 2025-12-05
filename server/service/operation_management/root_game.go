package operation_management

import (
	"context"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"go.uber.org/zap"
)

type RootGameService struct {

}

// List 主游戏列表
func (receiver *RootGameService) List(ctx context.Context, req *api.RootGameListReq) (resp interface{}, total int64, err error) {
	model := operation_management.NewDimRootGameModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.GameName != "" {
		if validator.IsNumberStr(req.GameName) {
			tmpDb.Where("id = ?", req.GameName)
		} else {
			tmpDb.Where("game_name like ?", "%"+ req.GameName +"%")
		}
	}
	if req.ContractName != "" {
		tmpDb.Where("contract_name like ?", "%"+ req.ContractName +"%")
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []operation_management.DimRootGameModel
	if listErr := tmpDb.Limit(req.PageSize).Offset((req.Page -1)*req.PageSize).Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取列表失败", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (receiver *RootGameService) Add(ctx context.Context, req *api.RootGameAddReq) (resp api.RootGameAddResp, err error) {
	model := operation_management.NewDimRootGameModel()
	req.DimRootGameModel.DimRootGameModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("创建异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *RootGameService) Modify(ctx context.Context, req *api.RootGameModifyReq) (resp api.RootGameModifyResp, err error) {
	model := operation_management.NewDimRootGameModel()
	req.DimRootGameModel.DimRootGameModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}