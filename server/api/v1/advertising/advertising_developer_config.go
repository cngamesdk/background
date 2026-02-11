package advertising

import (
	"github.com/cngamesdk/go-core/translator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdvertisingDeveloperConfigApi struct {
}

func (receiver *AdvertisingDeveloperConfigApi) List(ctx *gin.Context) {
	var req api2.AdvertisingDeveloperConfigListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(translator.DealErr(err).Error(), ctx)
		return
	}
	resp, respTotal, respErr := advertisingDeveloperConfigService.List(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(respErr))
		response.FailWithMessage("获取失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List:     resp,
		Total:    respTotal,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, ctx)
	return
}

func (receiver *AdvertisingDeveloperConfigApi) Add(ctx *gin.Context) {
	var req api2.AdvertisingDeveloperConfigAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(translator.DealErr(err).Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := advertisingDeveloperConfigService.Add(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}

func (receiver *AdvertisingDeveloperConfigApi) Modify(ctx *gin.Context) {
	var req api2.AdvertisingDeveloperConfigModifyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(translator.DealErr(err).Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := advertisingDeveloperConfigService.Modify(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}
