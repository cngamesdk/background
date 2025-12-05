package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	errors2 "github.com/pkg/errors"
	"strings"
)

type AgentListReq struct {
	advertising.DimAgentModel
	request.PageInfo
}

type AgentListResp struct {
}

type AgentAddReq struct {
	advertising.DimAgentModel
}

func (receiver *AgentAddReq) Format() {
	receiver.Id = 0
	receiver.AgentName = strings.TrimSpace(receiver.AgentName)
	receiver.SettlementType = strings.TrimSpace(receiver.SettlementType)
}

func (receiver *AgentAddReq) Validate() (err error) {
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if receiver.ChannelGroupId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "channel_group_id")
		return
	}
	if validateErr := validate.EmptyString(receiver.AgentName); validateErr != nil {
		err = validateErr
		return
	}
	if validateErr := validate.EmptyString(receiver.SettlementType); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type AgentAddResp struct {

}

type AgentModifyReq struct {
	advertising.DimAgentModel
}

func (receiver *AgentModifyReq) Format() {
	receiver.AgentName = strings.TrimSpace(receiver.AgentName)
	receiver.SettlementType = strings.TrimSpace(receiver.SettlementType)
}

func (receiver *AgentModifyReq) Validate() (err error) {
	if receiver.Id <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if receiver.ChannelGroupId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "channel_group_id")
		return
	}
	if validateErr := validate.EmptyString(receiver.AgentName); validateErr != nil {
		err = validateErr
		return
	}
	if validateErr := validate.EmptyString(receiver.SettlementType); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type AgentModifyResp struct {

}
