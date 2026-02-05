package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	errors2 "github.com/pkg/errors"
	"strings"
)

type ChannelGroupListReq struct {
	advertising.DimChannelGroupModel
	request.PageInfo
}

type ChannelGroupListResp struct {
}

type ChannelGroupAddReq struct {
	advertising.DimChannelGroupModel
}

func (receiver *ChannelGroupAddReq) Format() {
	receiver.Id = 0
	receiver.ChannelGroupName = strings.TrimSpace(receiver.ChannelGroupName)
}

func (receiver *ChannelGroupAddReq) Validate() (err error) {
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if receiver.AdvertisingMediaId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "媒体ID")
		return
	}
	if validateErr := validate.EmptyString(receiver.ChannelGroupName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type ChannelGroupAddResp struct {
}

type ChannelGroupModifyReq struct {
	advertising.DimChannelGroupModel
}

func (receiver *ChannelGroupModifyReq) Format() {
	receiver.ChannelGroupName = strings.TrimSpace(receiver.ChannelGroupName)
}

func (receiver *ChannelGroupModifyReq) Validate() (err error) {
	if receiver.Id <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if receiver.AdvertisingMediaId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "媒体ID")
		return
	}
	if validateErr := validate.EmptyString(receiver.ChannelGroupName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type ChannelGroupModifyResp struct {
}
