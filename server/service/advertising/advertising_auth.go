package advertising

import (
	"context"
	"github.com/cngamesdk/go-core/goroutine"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ad_placement/ad_platform"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
	"time"
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

	//授权入库
	model := advertising.NewDimAdvertisingMediaAuthModel()
	model.DimAdvertisingMediaAuthModel = respAuthCallback.DimAdvertisingMediaAuthModel
	if createErr := model.Create(ctx); createErr != nil {
		err = createErr
		global.GVA_LOG.Error("保存token异常", zap.Error(createErr))
		return
	}

	goroutine.CreateGoroutine(func() {
		err = a.GetAdvertiser(ctx, *model)
		if err != nil {
			global.GVA_LOG.Error("获取帐户异常", zap.Error(err))
			return
		}
		return
	}, func(any2 any) {
		global.GVA_LOG.Error("创建协程异常", zap.Any("err", any2))
	})

	return
}

func (a *AdvertisingAuthService) GetAdvertiser(ctx context.Context, req advertising.DimAdvertisingMediaAuthModel) (err error) {
	adapter, adapterErr := ad_platform.GetAdapterFactory(req.Code, global.GVA_LOG)
	if adapterErr != nil {
		err = adapterErr
		global.GVA_LOG.Error("获取媒体适配器异常", zap.Error(adapterErr))
		return
	}
	developer := advertising.NewDimAdvertisingDeveloperConfigModel()
	if takeErr := developer.Take(ctx, "*", "id = ?", req.Id); takeErr != nil {
		err = takeErr
		global.GVA_LOG.Error("获取开发者信息异常", zap.Error(takeErr))
		return
	}
	//拉取帐户信息
	adapter.Init(ad_platform.AdapterConfig{
		Developer:    *developer,
		AdvertiserID: cast.ToString(req.AccountId),
		Auth:         req,
		Timeout:      5 * time.Second,
	})
	accounts, accountsErr := adapter.AuthAdvertiserGet(ctx)
	if accountsErr != nil {
		err = accountsErr
		global.GVA_LOG.Error("拉取帐户列表异常", zap.Error(accountsErr))
		return
	}
	for index, item := range accounts {
		accounts[index].AuthId = req.Id
		if item.AccountId == cast.ToInt64(req.AccountId) {
			if roleType, roleTypeOk := req.Extension["role_type"]; roleTypeOk {
				accounts[index].Role = cast.ToString(roleType)
			}
		}
	}
	batchErr := global.GVA_DB.WithContext(ctx).
		Table((&advertising.DimAdvertisingMediaAccountModel{}).TableName()).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "code"}, {Name: "account_id"}},
			UpdateAll: true, // 更新除主键外的所有字段
		}).CreateInBatches(&accounts, 100).Error
	if batchErr != nil {
		err = batchErr
		global.GVA_LOG.Error("批量操作异常", zap.Error(batchErr), zap.Any("data", accounts))
		return
	}
	return
}

func (a *AdvertisingAuthService) RefreshToken(ctx context.Context) (err error) {
	return
}
