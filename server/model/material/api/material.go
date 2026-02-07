package api

import (
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/material"
	"github.com/pkg/errors"
	"strings"
)

type MaterialListReq struct {
	material.OdsMaterialLogModel
	request.PageInfo
}

type MaterialListResp struct {
	material.OdsMaterialLogModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
}

type MaterialAddReq struct {
	material.OdsMaterialLogModel
}

func (a *MaterialAddReq) Format() {
	a.Id = 0
	a.MaterialName = strings.TrimSpace(a.MaterialName)
}

func (a *MaterialAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.MaterialName); validateErr != nil {
		err = errors.Wrap(validateErr, "素材名称")
		return
	}
	return
}

type MaterialAddResp struct {
}

type MaterialModifyReq struct {
	material.OdsMaterialLogModel
}

func (a *MaterialModifyReq) Format() {
	a.MaterialName = strings.TrimSpace(a.MaterialName)
}

func (a *MaterialModifyReq) Validate() (err error) {
	if a.Id <= 0 {
		err = errors.New("主键ID不能为空")
		return
	}
	if validateErr := validate.EmptyString(a.MaterialName); validateErr != nil {
		err = errors.Wrap(validateErr, "素材名称")
		return
	}
	return
}

type MaterialModifyResp struct {
}
