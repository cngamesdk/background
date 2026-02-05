package operation_management

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PlatformApi struct {
}

func (receiver *PlatformApi) List(ctx *gin.Context) {
	var req api.PlatformListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respTotal, respErr := platformService.List(ctx, &req)
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

func (receiver *PlatformApi) Add(ctx *gin.Context) {
	var req api.PlatformAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := platformService.Add(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}

func (receiver *PlatformApi) Modify(ctx *gin.Context) {
	var req api.PlatformModifyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := platformService.Modify(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}
