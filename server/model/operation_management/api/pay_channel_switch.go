package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"slices"
	"strings"
)

type PayChannelSwitchListReq struct {
	operation_management.DimPayChannelSwitchModel
	request.PageInfo
	PayChannelId int `json:"pay_channel_id"`
}

type PayChannelSwitchListResp struct {
}

type PayChannelSwitchAddReq struct {
	operation_management.DimPayChannelSwitchModel
}

func (c *PayChannelSwitchAddReq) Format() {
	c.Id = 0
	c.PayType = strings.TrimSpace(c.PayType)
}

func (c *PayChannelSwitchAddReq) Validate() (err error) {
	if c.PlatformId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "平台ID")
		return
	}
	if common.GetPayTypeName(c.PayType) == "" {
		err = errors.Wrap(error2.ErrorParamEmpty, "支付网关为空或者非法")
		return
	}
	if len(c.Rules) <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "请配置规则")
		return
	}
	for _, item := range c.Rules {
		if item.Key == "" || item.Operator == "" || len(item.Value) <= 0 {
			err = errors.Wrap(error2.ErrorParamEmpty, "规则中存在空项，请重新配置")
			return
		}
	}
	if len(c.PayChannels) <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "请配置充值渠道")
		return
	}
	var tempPayChannelIds []int64
	var existsActivePayChannel bool
	for _, item := range c.PayChannels {
		if item.PayChannelId <= 0 || item.Weight <= 0 || item.Mode == "" {
			err = errors.Wrap(error2.ErrorParamEmpty, "充值渠道中存在空项，请重新配置")
			return
		}
		if slices.Contains(tempPayChannelIds, item.PayChannelId) {
			err = errors.New("支付渠道中存在重复项，请重新配置")
			return
		}
		if item.Mode == common.PayChannelSwitchModeActive {
			existsActivePayChannel = true
		}
		tempPayChannelIds = append(tempPayChannelIds, item.PayChannelId)
	}
	if !existsActivePayChannel {
		err = errors.New("支付渠道中必须设置主模式，请重新配置")
		return
	}
	return
}

type PayChannelSwitchAddResp struct {
}

type PayChannelSwitchModifyReq struct {
	operation_management.DimPayChannelSwitchModel
}

func (c *PayChannelSwitchModifyReq) Format() {
	c.PayType = strings.TrimSpace(c.PayType)
}

func (c *PayChannelSwitchModifyReq) Validate() (err error) {
	if c.Id <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if c.PlatformId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "平台ID")
		return
	}
	if common.GetPayTypeName(c.PayType) == "" {
		err = errors.Wrap(error2.ErrorParamEmpty, "支付网关为空或者非法")
		return
	}
	if len(c.Rules) <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "请配置规则")
		return
	}
	for _, item := range c.Rules {
		if item.Key == "" || item.Operator == "" || len(item.Value) <= 0 {
			err = errors.Wrap(error2.ErrorParamEmpty, "规则中存在空项，请重新配置")
			return
		}
	}
	if len(c.PayChannels) <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "请配置充值渠道")
		return
	}
	var existsActivePayChannel bool
	var tempPayChannelIds []int64
	for _, item := range c.PayChannels {
		if item.PayChannelId <= 0 || item.Weight <= 0 || item.Mode == "" {
			err = errors.Wrap(error2.ErrorParamEmpty, "充值渠道中存在空项，请重新配置")
			return
		}
		if slices.Contains(tempPayChannelIds, item.PayChannelId) {
			err = errors.New("支付渠道中存在重复项，请重新配置")
			return
		}
		if item.Mode == common.PayChannelSwitchModeActive {
			existsActivePayChannel = true
		}
		tempPayChannelIds = append(tempPayChannelIds, item.PayChannelId)
	}
	if !existsActivePayChannel {
		err = errors.New("支付渠道中必须设置主模式，请重新配置")
		return
	}
	return
}

type PayChannelSwitchModifyResp struct {
}
