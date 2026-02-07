package api

import (
	"github.com/cngamesdk/go-core/model/sql"
	material2 "github.com/cngamesdk/go-core/model/sql/material"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/material"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

type MaterialListReq struct {
	material.OdsMaterialLogModel
	request.PageInfo
}

type MaterialListResp struct {
	material.OdsMaterialLogModel
	PlatformName     string `json:"platform_name" gorm:"platform_name"`
	MaterialTypeName string `json:"material_type_name" gorm:"-"`
	ThemeName        string `json:"theme_name" gorm:"theme_name"`
	SourceName       string `json:"source_name" gorm:"-"`
	StatusName       string `json:"status_name" gorm:"-"`
	VisibilityName   string `json:"visibility_name" gorm:"-"`
}

func (receiver *MaterialListResp) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *MaterialListResp) findHook(tx *gorm.DB) (err error) {
	if receiver.MaterialType != "" && receiver.MaterialTypeName == "" {
		if name, ok := material2.MaterialTypes[receiver.MaterialType]; ok {
			receiver.MaterialTypeName = name
		}
	}
	if receiver.MaterialType != "" && receiver.MaterialTypeName == "" {
		if name, ok := material2.MaterialTypes[receiver.MaterialType]; ok {
			receiver.MaterialTypeName = name
		}
	}
	if receiver.Source != "" && receiver.SourceName == "" {
		if name, ok := material2.MaterialSources[receiver.Source]; ok {
			receiver.SourceName = name
		}
	}
	if receiver.Status != "" && receiver.StatusName == "" {
		if name, ok := sql.StatusMap[receiver.Status]; ok {
			receiver.StatusName = name
		}
	}
	if receiver.Visibility != "" && receiver.VisibilityName == "" {
		if name, ok := material2.MaterialVisibilities[receiver.Visibility]; ok {
			receiver.VisibilityName = name
		}
	}
	return
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
