package api

import (
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"strings"
)

type GamePackagingConfigListReq struct {
	request.PageInfo
	operation_management.DimGamePackagingConfigModel
}

type GamePackagingConfigListResp struct {
}

type GamePackagingConfigAddReq struct {
	operation_management.DimGamePackagingConfigModel
}

func (receiver *GamePackagingConfigAddReq) Format() {
	receiver.Id = 0
	receiver.MediaCode = strings.TrimSpace(receiver.MediaCode)
	receiver.GamePackagePath = strings.TrimSpace(receiver.GamePackagePath)
	receiver.Status = strings.TrimSpace(receiver.Status)
	receiver.UseStatus = strings.TrimSpace(receiver.UseStatus)
}

func (receiver *GamePackagingConfigAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(receiver.MediaCode); validateErr != nil {
		err = errors.Wrap(validateErr, "媒体码")
		return
	}
	if validateErr := validate.EmptyString(receiver.GamePackagePath); validateErr != nil {
		err = errors.Wrap(validateErr, "母包地址")
		return
	}
	if validateErr := validate.EmptyString(receiver.GamePackageHash); validateErr != nil {
		err = errors.Wrap(validateErr, "母包校验HASH")
		return
	}
	if validateErr := validate.EmptyString(receiver.Status); validateErr != nil {
		err = errors.Wrap(validateErr, "状态")
		return
	}
	if validateErr := validate.EmptyString(receiver.UseStatus); validateErr != nil {
		err = errors.Wrap(validateErr, "使用状态")
		return
	}
	return
}

type GamePackagingConfigAddResp struct {
}

type GamePackagingConfigModifyReq struct {
	operation_management.DimGamePackagingConfigModel
}

func (receiver *GamePackagingConfigModifyReq) Format() {
	receiver.MediaCode = strings.TrimSpace(receiver.MediaCode)
	receiver.GamePackagePath = strings.TrimSpace(receiver.GamePackagePath)
	receiver.Status = strings.TrimSpace(receiver.Status)
	receiver.UseStatus = strings.TrimSpace(receiver.UseStatus)
}

func (receiver *GamePackagingConfigModifyReq) Validate() (err error) {
	if receiver.Id <= 0 {
		err = errors.New("主键ID不能为空")
		return
	}
	if validateErr := validate.EmptyString(receiver.MediaCode); validateErr != nil {
		err = errors.Wrap(validateErr, "媒体码")
		return
	}
	if validateErr := validate.EmptyString(receiver.GamePackagePath); validateErr != nil {
		err = errors.Wrap(validateErr, "母包地址")
		return
	}
	if validateErr := validate.EmptyString(receiver.GamePackageHash); validateErr != nil {
		err = errors.Wrap(validateErr, "母包校验HASH")
		return
	}
	if validateErr := validate.EmptyString(receiver.Status); validateErr != nil {
		err = errors.Wrap(validateErr, "状态")
		return
	}
	if validateErr := validate.EmptyString(receiver.UseStatus); validateErr != nil {
		err = errors.Wrap(validateErr, "使用状态")
		return
	}
	return
}

type GamePackagingConfigModifyResp struct {
}
