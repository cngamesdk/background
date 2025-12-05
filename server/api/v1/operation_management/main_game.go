package operation_management

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MainGameApi struct {
	
}

func (receiver *MainGameApi) List(ctx *gin.Context) {
	var req api.MainGameListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respTotal, respErr := mainGameService.List(ctx, &req)
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

func (receiver *MainGameApi) Add(ctx *gin.Context) {
	var req api.MainGameAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := mainGameService.Add(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}

func (receiver *MainGameApi) Modify(ctx *gin.Context) {
	var req api.MainGameModifyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := mainGameService.Modify(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}