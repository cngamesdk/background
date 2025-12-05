package operation_management

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SubGameApi struct{}

// List 获取列表
func (receiver *SubGameApi) List(ctx *gin.Context) {
	var req api.SubGameListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respTotal, respErr := subGameService.List(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(respErr))
		response.FailWithMessage("获取失败", ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List: resp,
		Total: respTotal,
		Page: req.Page,
		PageSize: req.PageSize,
	}, ctx)
}

func (receiver *SubGameApi) Add(ctx *gin.Context) {
	var req api.SubGameAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		response.FailWithMessage("验证失败:" + validateErr.Error(), ctx)
		return
	}
	resp, respErr := subGameService.Add(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("操作失败", zap.Error(respErr))
		response.FailWithMessage("添加失败:" + respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
}

func (receiver *SubGameApi) Modify(ctx *gin.Context) {
	var req api.SubGameModifyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		response.FailWithMessage("验证失败:" + validateErr.Error(), ctx)
		return
	}
	resp, respErr := subGameService.Modify(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("修改失败", zap.Error(respErr))
		response.FailWithMessage("修改失败:" + respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
}

// Config 获取子游戏配置
func (receiver *SubGameApi) Config(ctx *gin.Context) {
	var req api.SubGameConfigReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := subGameService.Config(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(respErr))
		response.FailWithMessage("获取失败:" + respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
}