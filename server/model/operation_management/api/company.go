package api

import (
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"strings"
)

type CompanyListReq struct {
	operation_management.DimCompanyModel
	request.PageInfo
}

type CompanyListResp struct {
}

type CompanyAddReq struct {
	operation_management.DimCompanyModel
}

func (a *CompanyAddReq) Format() {
	a.Id = 0
	a.CompanyName = strings.TrimSpace(a.CompanyName)
	a.PrivacyPolicyUrl = strings.TrimSpace(a.PrivacyPolicyUrl)
	a.UserAgreementUrl = strings.TrimSpace(a.UserAgreementUrl)
}

func (a *CompanyAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.CompanyName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type CompanyAddResp struct {
}

type CompanyModifyReq struct {
	operation_management.DimCompanyModel
}

func (a *CompanyModifyReq) Format() {
	a.CompanyName = strings.TrimSpace(a.CompanyName)
	a.PrivacyPolicyUrl = strings.TrimSpace(a.PrivacyPolicyUrl)
	a.UserAgreementUrl = strings.TrimSpace(a.UserAgreementUrl)
}

func (a *CompanyModifyReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.CompanyName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type CompanyModifyResp struct {
}
