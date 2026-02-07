package api

import (
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/material"
	"github.com/pkg/errors"
	"strings"
)

type MaterialFileListReq struct {
	material.OdsMaterialFileLogModel
	request.PageInfo
}

type MaterialFileListResp struct {
	material.OdsMaterialFileLogModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
}

type MaterialFileAddReq struct {
	material.OdsMaterialFileLogModel
}

func (a *MaterialFileAddReq) Format() {
	a.Id = 0
	a.Url = strings.TrimSpace(a.Url)
	a.FileName = strings.TrimSpace(a.FileName)
	a.Signature = strings.TrimSpace(a.Signature)
}

func (a *MaterialFileAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.Url); validateErr != nil {
		err = errors.Wrap(validateErr, "文件地址")
		return
	}
	if validateErr := validate.EmptyString(a.FileName); validateErr != nil {
		err = errors.Wrap(validateErr, "文件名称")
		return
	}
	if validateErr := validate.EmptyString(a.Signature); validateErr != nil {
		err = errors.Wrap(validateErr, "文件签名")
		return
	}
	return
}

type MaterialFileAddResp struct {
}

type MaterialFileModifyReq struct {
	material.OdsMaterialFileLogModel
}

func (a *MaterialFileModifyReq) Format() {
	a.Url = strings.TrimSpace(a.Url)
	a.FileName = strings.TrimSpace(a.FileName)
	a.Signature = strings.TrimSpace(a.Signature)
}

func (a *MaterialFileModifyReq) Validate() (err error) {
	if a.Id <= 0 {
		err = errors.New("主键ID不能为空")
		return
	}
	if validateErr := validate.EmptyString(a.Url); validateErr != nil {
		err = errors.Wrap(validateErr, "文件地址")
		return
	}
	if validateErr := validate.EmptyString(a.FileName); validateErr != nil {
		err = errors.Wrap(validateErr, "文件名称")
		return
	}
	if validateErr := validate.EmptyString(a.Signature); validateErr != nil {
		err = errors.Wrap(validateErr, "文件签名")
		return
	}
	return
}

type MaterialFileModifyResp struct {
}
