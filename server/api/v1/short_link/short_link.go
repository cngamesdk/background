package short_link

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/short_link/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ShortLinkApi struct{}

func (a *ShortLinkApi) Create(ctx *gin.Context) {
	var req api.ShortLinkCreateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	// 获取当前登录用户名作为creator
	creator := "admin"
	resp, err := shortLinkService.Create(ctx.Request.Context(), &req, creator)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
}

func (a *ShortLinkApi) List(ctx *gin.Context) {
	var req api.ShortLinkListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := shortLinkService.List(ctx.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, ctx)
}

func (a *ShortLinkApi) Detail(ctx *gin.Context) {
	var req api.ShortLinkDetailReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	record, err := shortLinkService.Detail(ctx.Request.Context(), req.Id)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
		return
	}
	response.OkWithData(record, ctx)
}

func (a *ShortLinkApi) Update(ctx *gin.Context) {
	var req api.ShortLinkUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := shortLinkService.Update(ctx.Request.Context(), &req); err != nil {
		global.GVA_LOG.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func (a *ShortLinkApi) Delete(ctx *gin.Context) {
	var req api.ShortLinkDetailReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := shortLinkService.Delete(ctx.Request.Context(), req.Id); err != nil {
		global.GVA_LOG.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

func (a *ShortLinkApi) ClickLogList(ctx *gin.Context) {
	var req api.ClickLogListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := shortLinkService.ClickLogList(ctx.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, ctx)
}
