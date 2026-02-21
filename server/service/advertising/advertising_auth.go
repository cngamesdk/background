package advertising

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ad_placement/ad_platform"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

type AdvertisingAuthService struct {
}

func (a *AdvertisingAuthService) Redirect(ctx context.Context, req *api.AdvertisingAuthRedirectReq) (
	resp api.AdvertisingAuthRedirectResp, err error) {
	adapter, adapterErr := ad_platform.GetAdapterFactory(req.Code, global.GVA_LOG)
	if adapterErr != nil {
		err = adapterErr
		global.GVA_LOG.Error("获取媒体适配器异常", zap.Error(adapterErr))
		return
	}
	resp, err = adapter.AuthRedirect(ctx, req)
	return
}

func (a *AdvertisingAuthService) Callback(ctx context.Context, req map[string]interface{}) (
	resp string, err error) {
	code, ok := req["code"]
	if !ok {
		err = errors.New("code is not exists")
		return
	}
	adapter, adapterErr := ad_platform.GetAdapterFactory(cast.ToString(code), global.GVA_LOG)
	if adapterErr != nil {
		err = adapterErr
		global.GVA_LOG.Error("获取媒体适配器异常", zap.Error(adapterErr))
		return
	}
	respAuthCallback, respAuthCallbackErr := adapter.AuthCallback(ctx, req)
	if respAuthCallbackErr != nil {
		err = respAuthCallbackErr
		global.GVA_LOG.Error("授权回调异常", zap.Error(respAuthCallbackErr))
		return
	}

	//拉取帐户信息
	accounts, accountsErr := adapter.AuthAdvertiserGet(ctx, respAuthCallback.AccessToken)
	if accountsErr != nil {
		err = accountsErr
		global.GVA_LOG.Error("拉取帐户列表异常", zap.Error(accountsErr))
		return
	}

	var sqlData []advertising.DimAdvertisingMediaAccountModel
	for _, item := range accounts {
		insertModel := advertising.DimAdvertisingMediaAccountModel{}
		insertModel.Code = cast.ToString(code)
		insertModel.PlatformId = respAuthCallback.State.PlatformId
		insertModel.AccessToken = respAuthCallback.AccessToken
		insertModel.RefreshToken = respAuthCallback.RefreshToken
		insertModel.ExpiresAt = sql.MyCustomDatetime(respAuthCallback.ExpiresAt)
		insertModel.RefreshTokenExpiresAt = sql.MyCustomDatetime(respAuthCallback.RefreshTokenExpiresAt)
		insertModel.DeveloperId = respAuthCallback.State.DeveloperId
		insertModel.AccountId = item.AccountId
		insertModel.AccountName = item.AccountName
		insertModel.Role = item.Role
		insertModel.Status = sql.StatusNormal
		insertModel.Extension = item.Extension
		sqlData = append(sqlData, insertModel)
	}
	batchErr := global.GVA_DB.WithContext(ctx).
		Table((&advertising.DimAdvertisingMediaAccountModel{}).TableName()).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "code"}, {Name: "account_id"}},
			UpdateAll: true, // 更新除主键外的所有字段
		}).CreateInBatches(&sqlData, 100).Error
	if batchErr != nil {
		err = batchErr
		global.GVA_LOG.Error("批量操作异常", zap.Error(batchErr), zap.Any("data", sqlData))
		return
	}

	return
}
