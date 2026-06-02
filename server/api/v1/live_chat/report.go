package live_chat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	lcModel "github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"

	"github.com/gin-gonic/gin"
)

type ReportApi struct{}

func (a *ReportApi) Overview(c *gin.Context) {
	var search lcModel.ReportSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	overview, err := lcServiceGroup.ReportService.GetOverview(search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(overview, c)
}

func (a *ReportApi) Trend(c *gin.Context) {
	var search lcModel.ReportSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	items, err := lcServiceGroup.ReportService.GetTrend(search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(items, c)
}
