package api

import (
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"strings"
)

type GlobalCommonConfigListReq struct {
	request.PageInfo
	operation_management.DimGlobalCommonConfigModel
}

type GlobalCommonConfigListResp struct {
}

type GlobalCommonConfigAddReq struct {
	operation_management.DimGlobalCommonConfigModel
}

func (receiver *GlobalCommonConfigAddReq) Format() {
	receiver.Id = 0
	receiver.JavaExecutionPath = strings.TrimSpace(receiver.JavaExecutionPath)
	receiver.GamePackagingToolPath = strings.TrimSpace(receiver.GamePackagingToolPath)
}

func (receiver *GlobalCommonConfigAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(receiver.JavaExecutionPath); validateErr != nil {
		err = errors.Wrap(validateErr, "java执行路径")
		return
	}
	if validateErr := validate.EmptyString(receiver.GamePackagingToolPath); validateErr != nil {
		err = errors.Wrap(validateErr, "打包工具")
		return
	}
	return
}

type GlobalCommonConfigAddResp struct {
}

type GlobalCommonConfigModifyReq struct {
	operation_management.DimGlobalCommonConfigModel
}

func (receiver *GlobalCommonConfigModifyReq) Format() {
	receiver.JavaExecutionPath = strings.TrimSpace(receiver.JavaExecutionPath)
	receiver.GamePackagingToolPath = strings.TrimSpace(receiver.GamePackagingToolPath)
}

func (receiver *GlobalCommonConfigModifyReq) Validate() (err error) {
	if receiver.Id <= 0 {
		err = errors.New("主键ID为空")
		return
	}
	if validateErr := validate.EmptyString(receiver.JavaExecutionPath); validateErr != nil {
		err = errors.Wrap(validateErr, "Java执行路径")
		return
	}
	if validateErr := validate.EmptyString(receiver.GamePackagingToolPath); validateErr != nil {
		err = errors.Wrap(validateErr, "打包工具")
		return
	}
	return
}

type GlobalCommonConfigModifyResp struct {
}
