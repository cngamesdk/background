package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"strings"
)

type ProductCommonConfigurationListReq struct {
	operation_management.DimProductCommonConfigurationModel
	request.PageInfo
}

type ProductCommonConfigurationListResp struct {
}

type ProductCommonConfigurationAddReq struct {
	operation_management.DimProductCommonConfigurationModel
}

func (c *ProductCommonConfigurationAddReq) Format() {
	c.ConfigName = strings.TrimSpace(c.ConfigName)
	c.ShippingAddress = strings.TrimSpace(c.ShippingAddress)
}

func (c *ProductCommonConfigurationAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(c.ConfigName); validateErr != nil {
		err = errors.Wrap(error2.ErrorParamEmpty, "config_name")
		return
	}
	if validateErr := validate.EmptyString(c.ShippingAddress); validateErr != nil {
		err = errors.Wrap(error2.ErrorParamEmpty, "shipping_address")
		return
	}
	return
}

type ProductCommonConfigurationAddResp struct {
}

type ProductCommonConfigurationModifyReq struct {
	operation_management.DimProductCommonConfigurationModel
}

func (c *ProductCommonConfigurationModifyReq) Format() {
	c.ConfigName = strings.TrimSpace(c.ConfigName)
	c.ShippingAddress = strings.TrimSpace(c.ShippingAddress)
}

func (c *ProductCommonConfigurationModifyReq) Validate() (err error) {
	if validateErr := validate.EmptyString(c.ConfigName); validateErr != nil {
		err = errors.Wrap(error2.ErrorParamEmpty, "config_name")
		return
	}
	if validateErr := validate.EmptyString(c.ShippingAddress); validateErr != nil {
		err = errors.Wrap(error2.ErrorParamEmpty, "shipping_address")
		return
	}
	return
}

type ProductCommonConfigurationModifyResp struct {
}
