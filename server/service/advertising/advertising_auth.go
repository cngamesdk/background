package advertising

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ad_placement/ad_platform"
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
