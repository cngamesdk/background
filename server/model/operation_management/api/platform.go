package api

import (
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"strings"
)

type PlatformListReq struct {
	operation_management.DimPlatformModel
	request.PageInfo
}

type PlatformListResp struct {
}

type PlatformAddReq struct {
	operation_management.DimPlatformModel
}

func (a *PlatformAddReq) Format() {
	a.Id = 0
	a.PlatformName = strings.TrimSpace(a.PlatformName)
}

func (a *PlatformAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.PlatformName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type PlatformAddResp struct {
}

type PlatformModifyReq struct {
	operation_management.DimPlatformModel
}

func (a *PlatformModifyReq) Format() {
	a.PlatformName = strings.TrimSpace(a.PlatformName)
}

func (a *PlatformModifyReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.PlatformName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type PlatformModifyResp struct {
}
