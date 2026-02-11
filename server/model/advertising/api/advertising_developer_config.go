package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	errors2 "github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

type AdvertisingDeveloperConfigListReq struct {
	advertising.DimAdvertisingDeveloperConfigModel
	request.PageInfo
}

type AdvertisingDeveloperConfigListResp struct {
	advertising.DimAdvertisingDeveloperConfigModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
	CompanyName  string `json:"company_name" gorm:"company_name"`
	MediaName    string `json:"media_name" gorm:"media_name"`
}

func (receiver *AdvertisingDeveloperConfigListResp) AfterFind(tx *gorm.DB) (err error) {
	receiver.DimAdvertisingDeveloperConfigModel.DimAdvertisingDeveloperConfigModel.AesKey = func() string {
		return global.GVA_CONFIG.Common.AesCryptKey
	}
	return receiver.DimAdvertisingDeveloperConfigModel.DimAdvertisingDeveloperConfigModel.AfterFind(tx)
}

type AdvertisingDeveloperConfigAddReq struct {
	advertising.DimAdvertisingDeveloperConfigModel
}

func (receiver *AdvertisingDeveloperConfigAddReq) Format() {
	receiver.Id = 0
	receiver.Name = strings.TrimSpace(receiver.Name)
	receiver.Code = strings.TrimSpace(receiver.Code)
	receiver.AppId = strings.TrimSpace(receiver.AppId)
	receiver.Secret = strings.TrimSpace(receiver.Secret)
}

func (receiver *AdvertisingDeveloperConfigAddReq) Validate() (err error) {
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if validateErr := validate.EmptyString(receiver.Name); validateErr != nil {
		err = errors2.Wrap(validateErr, "配置名称")
		return
	}
	if validateErr := validate.EmptyString(receiver.Code); validateErr != nil {
		err = errors2.Wrap(validateErr, "媒体码")
		return
	}
	if receiver.CompanyId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "主体")
		return
	}
	if validateErr := validate.EmptyString(receiver.AppId); validateErr != nil {
		err = errors2.Wrap(validateErr, "应用ID")
		return
	}
	if validateErr := validate.EmptyString(receiver.Secret); validateErr != nil {
		err = errors2.Wrap(validateErr, "应用密钥")
		return
	}
	return
}

type AdvertisingDeveloperConfigAddResp struct {
}

type AdvertisingDeveloperConfigModifyReq struct {
	advertising.DimAdvertisingDeveloperConfigModel
}

func (receiver *AdvertisingDeveloperConfigModifyReq) Format() {
	receiver.SecretCrypt = ""
	receiver.Name = strings.TrimSpace(receiver.Name)
	receiver.Code = strings.TrimSpace(receiver.Code)
	receiver.AppId = strings.TrimSpace(receiver.AppId)
	receiver.Secret = strings.TrimSpace(receiver.Secret)
}

func (receiver *AdvertisingDeveloperConfigModifyReq) Validate() (err error) {
	if receiver.Id <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if validateErr := validate.EmptyString(receiver.Name); validateErr != nil {
		err = errors2.Wrap(validateErr, "配置名称")
		return
	}
	if validateErr := validate.EmptyString(receiver.Code); validateErr != nil {
		err = errors2.Wrap(validateErr, "媒体码")
		return
	}
	if receiver.CompanyId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "主体")
		return
	}
	if validateErr := validate.EmptyString(receiver.AppId); validateErr != nil {
		err = errors2.Wrap(validateErr, "应用ID")
		return
	}
	if validateErr := validate.EmptyString(receiver.Secret); validateErr != nil {
		err = errors2.Wrap(validateErr, "应用密钥")
		return
	}
	return
}

type AdvertisingDeveloperConfigModifyResp struct {
}
