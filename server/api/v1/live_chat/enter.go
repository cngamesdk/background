package live_chat

import lcService "github.com/flipped-aurora/gin-vue-admin/server/service/live_chat"

type ApiGroup struct {
	ProductApi
	FaqApi
	AgentApi
	ChatApi
	ReportApi
}

var lcServiceGroup = new(lcService.ServiceGroup)
