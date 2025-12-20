package data_report

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	api2 "github.com/flipped-aurora/gin-vue-admin/server/model/data_report/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaymentStatusApi struct {
}

func (receiver *PaymentStatusApi) List(ctx *gin.Context) {
	var req api2.PaymentStatusListReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	resp, respErr := paymentStatusService.List(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(respErr))
		response.FailWithDetailed(resp, "获取失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}
