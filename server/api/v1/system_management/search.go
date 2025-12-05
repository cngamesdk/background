package system_management

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system_management/api"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SearchApi struct {

}

func (receiver *SearchApi) Search(ctx *gin.Context) {
	var req api.SearchReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	req.Format()
	if validateErr := req.Validate(); validateErr != nil {
		response.FailWithMessage("验证失败:"+validateErr.Error(), ctx)
		return
	}
	resp, respErr := searchService.Search(ctx, &req)
	if respErr != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(respErr))
		response.FailWithMessage("获取失败:"+respErr.Error(), ctx)
		return
	}
	response.OkWithData(resp, ctx)
	return
}