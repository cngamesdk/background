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

type MaterialFileListReq struct {
	material.OdsMaterialFileLogModel
	request.PageInfo
}

type MaterialFileListResp struct {
	material.OdsMaterialFileLogModel
	PlatformName   string `json:"platform_name" gorm:"platform_name"`
	FileTypeName   string `json:"file_type_name" gorm:"-"`
	SourceName     string `json:"source_name" gorm:"-"`
	StatusName     string `json:"status_name" gorm:"-"`
	VisibilityName string `json:"visibility_name" gorm:"-"`
}

func (receiver *MaterialFileListResp) AfterFind(tx *gorm.DB) (err error) {
	return receiver.findHook(tx)
}

func (receiver *MaterialFileListResp) findHook(tx *gorm.DB) (err error) {
	if receiver.FileType != "" && receiver.FileTypeName == "" {
		if name, ok := material2.MaterialTypes[receiver.FileType]; ok {
			receiver.FileTypeName = name
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

type MaterialFileAddReq struct {
	material.OdsMaterialFileLogModel
}

func (a *MaterialFileAddReq) Format() {
	a.Id = 0
	a.Url = strings.TrimSpace(a.Url)
	a.FileName = strings.TrimSpace(a.FileName)
	a.Signature = strings.TrimSpace(a.Signature)
	if a.Status == "" {
		a.Status = sql.StatusNormal
	}
	if a.Visibility == "" {
		a.Visibility = material2.MaterialVisibilityPublic
	}
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
