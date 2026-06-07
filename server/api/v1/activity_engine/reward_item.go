package activity_engine

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine/api"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type RewardItemApi struct{}

func (a *RewardItemApi) Search(ctx *gin.Context) {
	var req api.RewardItemSearchReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	list, total, err := rewardItemService.Search(ctx.Request.Context(), &req)
	if err != nil {
		response.FailWithMessage("查询失败: "+err.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List: list, Total: total, Page: req.Page, PageSize: req.PageSize,
	}, ctx)
}
