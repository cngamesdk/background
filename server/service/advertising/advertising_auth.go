package advertising

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ad_placement/ad_platform"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
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
	//授权回调入库
	return
}
