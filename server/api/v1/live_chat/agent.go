package live_chat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	lcModel "github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
)

type AgentApi struct{}

func (a *AgentApi) Online(c *gin.Context) {
	var req lcModel.AgentOnlineReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := lcServiceGroup.AgentService.GoOnline(int64(userID), req.ProductID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("上线成功", c)
}

func (a *AgentApi) Offline(c *gin.Context) {
	var req lcModel.AgentOnlineReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	if err := lcServiceGroup.AgentService.GoOffline(int64(userID), req.ProductID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("下线成功", c)
}

func (a *AgentApi) Update(c *gin.Context) {
	var req lcModel.AgentUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lcServiceGroup.AgentService.Update(req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (a *AgentApi) List(c *gin.Context) {
	var pageInfo request.PageInfo
	productID := int64(0)
	if v, ok := c.GetQuery("product_id"); ok && v != "" {
		c.ShouldBindQuery(&struct{ ProductID int }{ProductID: int(productID)})
	}
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := lcServiceGroup.AgentService.List(productID, pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功", c)
}

func (a *AgentApi) Status(c *gin.Context) {
	userID := utils.GetUserID(c)
	productID := int64(0)
	if v, ok := c.GetQuery("product_id"); ok && v != "" {
		c.ShouldBindQuery(&struct{ ProductID int }{})
	}
	agent, err := lcServiceGroup.AgentService.GetStatus(int64(userID), productID)
	if err != nil {
		response.OkWithData(gin.H{"status": "offline"}, c)
		return
	}
	response.OkWithData(agent, c)
}
