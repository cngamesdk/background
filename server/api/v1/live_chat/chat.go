package live_chat

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	lcModel "github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
)

type ChatApi struct{}

func (a *ChatApi) Sessions(c *gin.Context) {
	var search lcModel.ChatSessionSearch
	var pageInfo request.PageInfo
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := c.ShouldBindQuery(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := lcServiceGroup.ChatService.GetSessions(search, pageInfo.Page, pageInfo.PageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{List: list, Total: total, Page: pageInfo.Page, PageSize: pageInfo.PageSize}, "获取成功", c)
}

func (a *ChatApi) SessionDetail(c *gin.Context) {
	id := c.Param("id")
	sessionID := int64(0)
	fmt.Sscanf(id, "%d", &sessionID)
	session, err := lcServiceGroup.ChatService.GetSession(sessionID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	messages, _, _ := lcServiceGroup.ChatService.GetMessages(sessionID, 1, 100)
	response.OkWithData(gin.H{"session": session, "messages": messages}, c)
}

func (a *ChatApi) Assign(c *gin.Context) {
	var req lcModel.AssignSessionReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lcServiceGroup.ChatService.AssignAgent(req.SessionID, req.AgentID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("分配成功", c)
}

func (a *ChatApi) Reply(c *gin.Context) {
	var req lcModel.AgentReplyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.MsgType == "" {
		req.MsgType = "text"
	}
	userID := utils.GetUserID(c)
	userName := utils.GetUserName(c)
	if err := lcServiceGroup.ChatService.AgentReply(req.SessionID, int64(userID), userName, req.Content, req.MsgType); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("回复成功", c)
}

func (a *ChatApi) Close(c *gin.Context) {
	var req request.GetById
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lcServiceGroup.ChatService.CloseSession(int64(req.ID)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("关闭成功", c)
}
