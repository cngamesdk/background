package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	errors2 "github.com/pkg/errors"
	"strings"
)

type AdvertisingMediaListReq struct {
	advertising.DimAdvertisingMediaModel
	request.PageInfo
}

type AdvertisingMediaListResp struct {
}

type AdvertisingMediaAddReq struct {
	advertising.DimAdvertisingMediaModel
}

func (receiver *AdvertisingMediaAddReq) Format() {
	receiver.Id = 0
	receiver.MediaName = strings.TrimSpace(receiver.MediaName)
	receiver.Code = strings.TrimSpace(receiver.Code)
}

func (receiver *AdvertisingMediaAddReq) Validate() (err error) {
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if validateErr := validate.EmptyString(receiver.MediaName); validateErr != nil {
		err = errors2.Wrap(validateErr, "媒体名称")
		return
	}
	if validateErr := validate.EmptyString(receiver.Code); validateErr != nil {
		err = errors2.Wrap(validateErr, "媒体码")
		return
	}
	return
}

type AdvertisingMediaAddResp struct {
}

type AdvertisingMediaModifyReq struct {
	advertising.DimAdvertisingMediaModel
}

func (receiver *AdvertisingMediaModifyReq) Format() {
	receiver.MediaName = strings.TrimSpace(receiver.MediaName)
	receiver.Code = strings.TrimSpace(receiver.Code)
}

func (receiver *AdvertisingMediaModifyReq) Validate() (err error) {
	if receiver.Id <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if validateErr := validate.EmptyString(receiver.MediaName); validateErr != nil {
		err = errors2.Wrap(validateErr, "媒体名称")
		return
	}
	if validateErr := validate.EmptyString(receiver.Code); validateErr != nil {
		err = errors2.Wrap(validateErr, "媒体码")
		return
	}
	return
}

type AdvertisingMediaModifyResp struct {
}
