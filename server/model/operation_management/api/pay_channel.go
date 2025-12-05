package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"strings"
)

type PayChannelListReq struct {
	operation_management.DimPayChannelModel
	request.PageInfo
}

type PayChannelListResp struct {
	
}

type PayChannelAddReq struct {
	operation_management.DimPayChannelModel
}

func (c *PayChannelAddReq) Format() {
	c.Id = 0
	c.ChannelName = strings.TrimSpace(c.ChannelName)
}

func (c *PayChannelAddReq) Validate() (err error) {
	if c.PlatformId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "平台ID")
		return
	}
	if c.CompanyId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "支付主体")
		return
	}
	if validateErr := validate.EmptyString(c.ChannelName); validateErr != nil {
		err = errors.Wrap(validateErr, "渠道名称")
		return
	}
	payTypeName := common.GetPayTypeName(c.PayType)
	if payTypeName == "" {
		err = errors.Wrap(error2.ErrorParamEmpty, "支付方式未知")
		return
	}
	if common.GetPayStatusName(c.Status) == "" {
		err = errors.New("支付状态未知" + c.Status)
		return
	}
	if c.Rate <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "费率")
		return
	}
	return
}

type PayChannelAddResp struct {
	
}

type PayChannelModifyReq struct {
	operation_management.DimPayChannelModel
}

func (c *PayChannelModifyReq) Format() {
	c.ChannelName = strings.TrimSpace(c.ChannelName)
}

func (c *PayChannelModifyReq) Validate() (err error) {
	if c.Id <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if c.PlatformId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "平台ID")
		return
	}
	if c.CompanyId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "支付主体")
		return
	}
	if validateErr := validate.EmptyString(c.ChannelName); validateErr != nil {
		err = errors.Wrap(validateErr, "渠道名称")
		return
	}
	payTypeName := common.GetPayTypeName(c.PayType)
	if payTypeName == "" {
		err = errors.Wrap(error2.ErrorParamEmpty, "支付方式未知")
		return
	}
	if common.GetPayStatusName(c.Status) == "" {
		err = errors.New("支付状态未知" + c.Status)
		return
	}
	if c.Rate <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "费率")
		return
	}
	return
}

type PayChannelModifyResp struct {

}
