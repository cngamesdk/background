package material

import (
	"github.com/cngamesdk/go-core/translator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/material/api"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MaterialApi struct {
}

func (receiver *MaterialApi) List(ctx *gin.Context) {
	var req api2.MaterialListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(translator.DealErr(err).Error(), ctx)
		return
	}
	resp, respTotal, respErr := materialService.List(ctx, &req)
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

func (receiver *MaterialApi) Add(ctx *gin.Context) {
	var req api2.MaterialAddReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(translator.DealErr(err).Error(), ctx)
		return
	}
	if req.Author == "" {
		claim, claimErr := utils.GetClaims(ctx)
		if claimErr != nil {
			response.FailWithMessage(translator.DealErr(claimErr).Error(), ctx)
			return
		}
		req.Author = claim.NickName
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := materialService.Add(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}

func (receiver *MaterialApi) Modify(ctx *gin.Context) {
	var req api2.MaterialModifyReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(translator.DealErr(err).Error(), ctx)
		return
	}
	req.Format()
	if err := req.Validate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	resp, respErr := materialService.Modify(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("保存失败", zap.Error(respErr))
		response.FailWithMessage("保存失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}
