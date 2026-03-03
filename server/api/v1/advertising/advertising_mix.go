package advertising

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AdvertisingMixApi struct {
}

func (receiver *AdvertisingMixApi) List(ctx *gin.Context) {
	var req api2.AdvertisingMixListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respTotal, respErr := advertisingMixService.List(ctx, &req)
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

func (receiver *AdvertisingMixApi) Add(ctx *gin.Context) {
	var req api2.AdvertisingMixAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format(ctx)
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := advertisingMixService.Add(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}

func (receiver *AdvertisingMixApi) Modify(ctx *gin.Context) {
	var req api2.AdvertisingMixModifyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format(ctx)
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := advertisingMixService.Modify(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}
