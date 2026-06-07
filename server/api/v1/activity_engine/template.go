package activity_engine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine/api"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type TemplateApi struct{}

func (a *TemplateApi) List(ctx *gin.Context) {
	var req api.TemplateListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := templateService.List(ctx.Request.Context(), &req)
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List: list, Total: total, Page: req.Page, PageSize: req.PageSize,
	}, ctx)
}

func (a *TemplateApi) Add(ctx *gin.Context) {
	var req api.TemplateAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := templateService.Add(ctx.Request.Context(), &req); err != nil {
		response.FailWithMessage("创建失败: "+err.Error(), ctx)
		return
	}
	response.OkWithMessage("创建成功", ctx)
}

func (a *TemplateApi) Clone(ctx *gin.Context) {
	var req api.TemplateCloneReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := templateService.Clone(ctx.Request.Context(), &req); err != nil {
		response.FailWithMessage("克隆失败: "+err.Error(), ctx)
		return
	}
	response.OkWithMessage("克隆成功", ctx)
}
