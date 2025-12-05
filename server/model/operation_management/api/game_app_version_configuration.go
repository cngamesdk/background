package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"strings"
)

type GameAppVersionConfigurationListReq struct {
	operation_management.DimGameAppVersionConfiguration
	request.PageInfo
}

type GameAppVersionConfigurationListResp struct {

}

type GameAppVersionConfigurationAddReq struct {
	operation_management.DimGameAppVersionConfiguration
}

func (c *GameAppVersionConfigurationAddReq) Format() {
	c.AppVersionName = strings.TrimSpace(c.AppVersionName)
}

func (c *GameAppVersionConfigurationAddReq) Validate() (err error) {
	if c.PlatformId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if c.GameId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "game_id")
		return
	}
	if c.AppVersionCode <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "app_version_code")
		return
	}
	if validateErr := validate.EmptyString(c.AppVersionName); validateErr != nil {
		err = errors.Wrap(error2.ErrorParamEmpty, "app_version_name")
		return
	}
	if c.ProductConfigId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "product_config_id")
		return
	}
	return
}

type GameAppVersionConfigurationAddResp struct {

}

type GameAppVersionConfigurationModifyReq struct {
	operation_management.DimGameAppVersionConfiguration
}

func (c *GameAppVersionConfigurationModifyReq) Format() {
	c.AppVersionName = strings.TrimSpace(c.AppVersionName)
}

func (c *GameAppVersionConfigurationModifyReq) Validate() (err error) {
	if c.Id <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if c.PlatformId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if c.GameId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "game_id")
		return
	}
	if c.AppVersionCode <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "app_version_code")
		return
	}
	if validateErr := validate.EmptyString(c.AppVersionName); validateErr != nil {
		err = errors.Wrap(error2.ErrorParamEmpty, "app_version_name")
		return
	}
	if c.ProductConfigId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "product_config_id")
		return
	}
	return
}

type GameAppVersionConfigurationModifyResp struct {
}
