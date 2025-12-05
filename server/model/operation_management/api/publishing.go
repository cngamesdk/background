package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"strings"
)

type PublishingChannelConfigListReq struct {
	operation_management.DimPublishingChannelConfigModel
	request.PageInfo
}

type PublishingChannelConfigListResp struct {
	
}

type PublishingChannelConfigAddReq struct {
	operation_management.DimPublishingChannelConfigModel
}

func (a *PublishingChannelConfigAddReq) Format() {
	a.Id = 0
	a.ChannelName = strings.TrimSpace(a.ChannelName)
}

func (a *PublishingChannelConfigAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.ChannelName); validateErr != nil {
		err = errors.Wrap(error2.ErrorParamEmpty, "渠道名称")
		return
	}
	return
}

type PublishingChannelConfigAddResp struct {

}

type PublishingChannelConfigModifyReq struct {
	operation_management.DimPublishingChannelConfigModel
}

func (a *PublishingChannelConfigModifyReq) Format() {
	a.ChannelName = strings.TrimSpace(a.ChannelName)
}

func (a *PublishingChannelConfigModifyReq) Validate() (err error) {
	if a.Id <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if validateErr := validate.EmptyString(a.ChannelName); validateErr != nil {
		err = errors.Wrap(error2.ErrorParamEmpty, "渠道名称")
		return
	}
	return
}

type PublishingChannelConfigModifyResp struct {

}

type PublishingChannelGameConfigListReq struct {
	operation_management.DimPublishingChannelGameConfigModel
	request.PageInfo
}

type PublishingChannelGameConfigListResp struct {
	
}

type PublishingChannelGameConfigAddReq struct {
	operation_management.DimPublishingChannelGameConfigModel
}

func (a *PublishingChannelGameConfigAddReq) Format() {
	a.Id = 0
}

func (a *PublishingChannelGameConfigAddReq) Validate() (err error) {
	if a.PlatformId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "平台ID")
		return
	}
	if a.GameId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "游戏ID")
		return
	}
	if a.ChannelId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "联运渠道")
		return
	}
	if a.AgentId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "渠道ID")
		return
	}
	if a.SiteId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "广告位ID")
		return
	}
	return
}

type PublishingChannelGameConfigAddResp struct {

}

type PublishingChannelGameConfigModifyReq struct {
	operation_management.DimPublishingChannelGameConfigModel
}

func (a *PublishingChannelGameConfigModifyReq) Format() {
}

func (a *PublishingChannelGameConfigModifyReq) Validate() (err error) {
	if a.Id <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if a.PlatformId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "平台ID")
		return
	}
	if a.GameId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "游戏ID")
		return
	}
	if a.ChannelId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "联运渠道ID")
		return
	}
	if a.AgentId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "渠道ID")
		return
	}
	if a.SiteId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "广告位ID")
		return
	}
	return
}

type PublishingChannelGameConfigModifyResp struct {

}