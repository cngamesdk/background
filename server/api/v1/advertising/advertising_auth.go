package advertising

import (
	"github.com/cngamesdk/go-core/translator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdvertisingAuthApi struct {
}

func (receiver *AdvertisingAuthApi) Redirect(ctx *gin.Context) {
	var req api2.AdvertisingAuthRedirectReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(translator.DealErr(err).Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(ctx); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := advertisingAuthService.Redirect(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(respErr))
		response.FailWithMessage("获取失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}

// Callback 授权回调
func (receiver *AdvertisingAuthApi) Callback(ctx *gin.Context) {
	req := make(map[string]interface{})
	// 获取所有查询参数
	queryParams := ctx.Request.URL.Query()
	for key, values := range queryParams {
		if len(values) == 1 {
			req[key] = values[0]
		} else {
			req[key] = values
		}
	}
	resp, respErr := advertisingAuthService.Callback(ctx, req)
	if respErr != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(respErr))
		response.FailWithMessage("获取失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}
