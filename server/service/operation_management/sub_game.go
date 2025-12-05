package operation_management

import (
	"context"
	"fmt"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type SubGameService struct {

}

func (receiver *SubGameService) List(ctx context.Context, req *api.SubGameListReq) (resp interface{}, total int64, err error) {
	model := operation_management.NewDimGameModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.GameName != "" {
		if validator.IsNumberStr(req.GameName) {
			tmpDb.Where("id = ?", req.GameName)
		} else {
			tmpDb.Where("game_name like ?", "%"+ req.GameName + "%")
		}
	}
	if countErr := tmpDb.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	var list []operation_management.DimGameModel
	if listErr := tmpDb.Offset((req.Page - 1)*req.PageSize).Limit(req.PageSize).Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
		return
	}
	resp = list
	return
}

func (receiver *SubGameService) Add(ctx context.Context, req *api.SubGameAddReq) (resp api.SubGameAddResp, err error) {
	model := operation_management.NewDimGameModel()
	req.DimGameModel.DimGameModel.Db = model.Db
	if saveErr := req.Create(ctx); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

func (receiver *SubGameService) Modify(ctx context.Context, req *api.SubGameModifyReq) (resp api.SubGameModifyResp, err error) {
	model := operation_management.NewDimGameModel()
	req.DimGameModel.DimGameModel.Db = model.Db
	if saveErr := req.Updates(ctx, "id = ?", req.Id); saveErr != nil {
		err = saveErr
		global.GVA_LOG.Error("保存异常", zap.Error(saveErr))
		return
	}
	return
}

// Config 获取子游戏配置
func (receiver *SubGameService) Config(ctx context.Context, req *api.SubGameConfigReq) (resp api.SubGameConfigResp, err error) {
	if global.GVA_CONFIG.Common.GameHashKey == "" {
		err = errors.New("游戏加密密钥未配置(必须与API配置一致)")
		return
	}
	if global.GVA_CONFIG.Common.AesCryptKey == "" {
		err = errors.New("AES加密密钥未配置(必须与API配置一致)")
		return
	}
	model := operation_management.NewDimGameModel()
	if takeErr := model.Take(ctx, "*", "id = ?", req.Id); takeErr != nil {
		err = takeErr
		global.GVA_LOG.Warn("获取异常", zap.Error(takeErr))
		return
	}
	plaform := cryptor.AesEcbEncrypt([]byte(cast.ToString(model.PlatformId)), []byte(global.GVA_CONFIG.Common.AesCryptKey))
	content := fmt.Sprintf(`游戏名称: %s
平台(platform): %s
游戏ID(game_id): %d
API密钥(app_key): %s
二次认证密钥(login_key): %s
充值密钥(pay_key): %s`,
model.GameName,
cryptor.Base64StdEncode(string(plaform)),
req.Id,
common.GetGameAppKey(req.Id, global.GVA_CONFIG.Common.GameHashKey),
common.GetGameLoginKey(req.Id, global.GVA_CONFIG.Common.GameHashKey),
common.GetGamePayKey(req.Id, global.GVA_CONFIG.Common.GameHashKey),
	)
	resp.Content = content
	return
}