package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	errors2 "github.com/pkg/errors"
	"strings"
)

type SiteListReq struct {
	advertising.DimSiteModel
	request.PageInfo
}

type SiteListResp struct {
}

type SiteAddReq struct {
	advertising.DimSiteModel
}

func (receiver *SiteAddReq) Format() {
	receiver.Id = 0
	receiver.SiteName = strings.TrimSpace(receiver.SiteName)
}

func (receiver *SiteAddReq) Validate() (err error) {
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if receiver.AgentId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "agent_id")
		return
	}
	if validateErr := validate.EmptyString(receiver.SiteName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type SiteAddResp struct {
}

type SiteModifyReq struct {
	advertising.DimSiteModel
}

func (receiver *SiteModifyReq) Format() {
	receiver.SiteName = strings.TrimSpace(receiver.SiteName)
}

func (receiver *SiteModifyReq) Validate() (err error) {
	if receiver.Id <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if receiver.AgentId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "agent_id")
		return
	}
	if validateErr := validate.EmptyString(receiver.SiteName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type SiteModifyResp struct {
}
