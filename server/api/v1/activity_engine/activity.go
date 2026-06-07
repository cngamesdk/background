package activity_engine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine/api"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type ActivityApi struct{}

func (a *ActivityApi) List(ctx *gin.Context) {
	var req api.ActivityListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := activityService.List(ctx.Request.Context(), &req)
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List: list, Total: total, Page: req.Page, PageSize: req.PageSize,
	}, ctx)
}

func (a *ActivityApi) Add(ctx *gin.Context) {
	var req api.ActivityAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := activityService.Add(ctx.Request.Context(), &req); err != nil {
		response.FailWithMessage("创建失败: "+err.Error(), ctx)
		return
	}
	response.OkWithMessage("创建成功", ctx)
}

func (a *ActivityApi) Modify(ctx *gin.Context) {
	var req api.ActivityModifyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := activityService.Modify(ctx.Request.Context(), &req); err != nil {
		response.FailWithMessage("修改失败: "+err.Error(), ctx)
		return
	}
	response.OkWithMessage("修改成功", ctx)
}

func (a *ActivityApi) Detail(ctx *gin.Context) {
	var req api.ActivityDetailReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	result, err := activityService.Detail(ctx.Request.Context(), req.ID)
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), ctx)
		return
	}
	response.OkWithData(result, ctx)
}

func (a *ActivityApi) Publish(ctx *gin.Context) {
	var req api.ActivityPublishReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := activityService.Publish(ctx.Request.Context(), req.ID); err != nil {
		response.FailWithMessage("发布失败: "+err.Error(), ctx)
		return
	}
	response.OkWithMessage("发布成功", ctx)
}

func (a *ActivityApi) Offline(ctx *gin.Context) {
	var req api.ActivityOfflineReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := activityService.Offline(ctx.Request.Context(), req.ID); err != nil {
		response.FailWithMessage("下线失败: "+err.Error(), ctx)
		return
	}
	response.OkWithMessage("下线成功", ctx)
}

func (a *ActivityApi) UpdateGrayscale(ctx *gin.Context) {
	var req api.GrayscaleUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := activityService.UpdateGrayscale(ctx.Request.Context(), req.ID, req.GrayscaleRatio); err != nil {
		response.FailWithMessage("更新失败: "+err.Error(), ctx)
		return
	}
	response.OkWithMessage("更新成功", ctx)
}
