package model_transfer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TokenApi struct{}

// Create 创建Token
// @Tags ModelTransfer
// @Summary 创建Token
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.TokenCreateReq true "创建Token"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /model-transfer/token/create [post]
func (a *TokenApi) Create(c *gin.Context) {
	var req api.TokenCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tokenService.Create(c.Request.Context(), &req); err != nil {
		global.GVA_LOG.Error("创建Token失败", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// Update 更新Token
// @Tags ModelTransfer
// @Summary 更新Token
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.TokenUpdateReq true "更新Token"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /model-transfer/token/update [post]
func (a *TokenApi) Update(c *gin.Context) {
	var req api.TokenUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tokenService.Update(c.Request.Context(), &req); err != nil {
		global.GVA_LOG.Error("更新Token失败", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// Delete 删除Token
// @Tags ModelTransfer
// @Summary 删除Token
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除Token"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /model-transfer/token/delete [post]
func (a *TokenApi) Delete(c *gin.Context) {
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tokenService.Delete(c.Request.Context(), req.ID); err != nil {
		global.GVA_LOG.Error("删除Token失败", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// List Token列表
// @Tags ModelTransfer
// @Summary Token列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.TokenListReq true "Token列表"
// @Success 200 {object} response.Response{data=response.PageResult} "获取成功"
// @Router /model-transfer/token/list [post]
func (a *TokenApi) List(c *gin.Context) {
	var req api.TokenListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := tokenService.List(c.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("查询Token列表失败", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, c)
}

// Detail Token详情
// @Tags ModelTransfer
// @Summary Token详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "Token详情"
// @Success 200 {object} response.Response{data=api.TokenDetailResp} "获取成功"
// @Router /model-transfer/token/detail [post]
func (a *TokenApi) Detail(c *gin.Context) {
	var req struct {
		ID int64 `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := tokenService.Detail(c.Request.Context(), req.ID)
	if err != nil {
		global.GVA_LOG.Error("查询Token详情失败", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}

	response.OkWithData(detail, c)
}

// Regenerate 重新生成Token
// @Tags ModelTransfer
// @Summary 重新生成Token
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.TokenRegenerateReq true "重新生成Token"
// @Success 200 {object} response.Response{data=api.TokenRegenerateResp} "生成成功"
// @Router /model-transfer/token/regenerate [post]
func (a *TokenApi) Regenerate(c *gin.Context) {
	var req api.TokenRegenerateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	token, err := tokenService.Regenerate(c.Request.Context(), &req)
	if err != nil {
		global.GVA_LOG.Error("重新生成Token失败", zap.Error(err))
		response.FailWithMessage("生成失败:"+err.Error(), c)
		return
	}

	response.OkWithData(api.TokenRegenerateResp{Token: token}, c)
}
