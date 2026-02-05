package api

import (
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/material"
	"github.com/pkg/errors"
	"strings"
)

type MaterialThemeListReq struct {
	material.DimMaterialThemeModel
	request.PageInfo
}

type MaterialThemeListResp struct {
}

type MaterialThemeAddReq struct {
	material.DimMaterialThemeModel
}

func (a *MaterialThemeAddReq) Format() {
	a.Id = 0
	a.ThemeName = strings.TrimSpace(a.ThemeName)
}

func (a *MaterialThemeAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.ThemeName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type MaterialThemeAddResp struct {
}

type MaterialThemeModifyReq struct {
	material.DimMaterialThemeModel
}

func (a *MaterialThemeModifyReq) Format() {
	a.ThemeName = strings.TrimSpace(a.ThemeName)
}

func (a *MaterialThemeModifyReq) Validate() (err error) {
	if a.Id <= 0 {
		err = errors.New("主键ID不能为空")
		return
	}
	if validateErr := validate.EmptyString(a.ThemeName); validateErr != nil {
		err = validateErr
		return
	}
	return
}

type MaterialThemeModifyResp struct {
}
