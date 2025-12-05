package operation_management

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PublishingApi struct {
	
}

func (receiver *PublishingApi) ChannelConfigList(ctx *gin.Context) {
	var req api.PublishingChannelConfigListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respTotal, respErr := publishingService.ChannelConfigList(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(respErr))
		response.FailWithMessage("获取失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List: resp,
		Total: respTotal,
		Page: req.Page,
		PageSize: req.PageSize,
	}, ctx)
	return
}

func (receiver *PublishingApi) ChannelConfigAdd(ctx *gin.Context) {
	var req api.PublishingChannelConfigAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := publishingService.ChannelConfigAdd(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}

func (receiver *PublishingApi) ChannelConfigModify(ctx *gin.Context) {
	var req api.PublishingChannelConfigModifyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := publishingService.ChannelConfigModify(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}

func (receiver *PublishingApi) ChannelGameConfigList(ctx *gin.Context) {
	var req api.PublishingChannelGameConfigListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respTotal, respErr := publishingService.ChannelGameConfigList(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(respErr))
		response.FailWithMessage("获取失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List: resp,
		Total: respTotal,
		Page: req.Page,
		PageSize: req.PageSize,
	}, ctx)
	return
}

func (receiver *PublishingApi) ChannelGameConfigAdd(ctx *gin.Context) {
	var req api.PublishingChannelGameConfigAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := publishingService.ChannelGameConfigAdd(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}

func (receiver *PublishingApi) ChannelGameConfigModify(ctx *gin.Context) {
	var req api.PublishingChannelGameConfigModifyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := publishingService.ChannelGameConfigModify(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}