package api

import (
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
	"strings"
)

type SubGameListReq struct {
	operation_management.DimGameModel
	request.PageInfo
}

type SubGameConfigReq struct {
	operation_management.DimGameModel
}

type SubGameConfigResp struct {
	Content string `json:"content"`
}

type SubGameAddReq struct {
	operation_management.DimGameModel
}

func (receiver *SubGameAddReq) Format() {
	receiver.Id = 0
	receiver.GameName = strings.TrimSpace(receiver.GameName)
}

func (receiver *SubGameAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(receiver.GameName); validateErr != nil {
		err = errors.Wrap(validateErr, "游戏名称")
		return
	}
	return
}

type SubGameAddResp struct {
}

type SubGameModifyReq struct {
	operation_management.DimGameModel
}

func (receiver *SubGameModifyReq) Format() {
	receiver.GameName = strings.TrimSpace(receiver.GameName)
}

func (receiver *SubGameModifyReq) Validate() (err error) {
	if validateErr := validate.EmptyString(receiver.GameName); validateErr != nil {
		err = errors.Wrap(validateErr, "游戏名称")
		return
	}
	if receiver.MainGameId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "主游戏ID")
		return
	}
	if receiver.CompanyId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "主体ID")
		return
	}
	return
}

type SubGameModifyResp struct {

}

type MainGameListReq struct {
	operation_management.DimMainGameModel
	request.PageInfo
}

type MainGameListResp struct {
}

type MainGameAddReq struct {
	operation_management.DimMainGameModel
}

func (a *MainGameAddReq) Format() {
	a.Id = 0
	a.GameName = strings.TrimSpace(a.GameName)
}

func (a *MainGameAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.GameName); validateErr != nil {
		err = errors.Wrap(validateErr, "主游戏名称")
		return
	}
	if a.RootGameId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "根游戏ID")
		return
	}
	return
}

type MainGameAddResp struct {
}

type MainGameModifyReq struct {
	operation_management.DimMainGameModel
}

func (a *MainGameModifyReq) Format() {
	a.GameName = strings.TrimSpace(a.GameName)
}

func (a *MainGameModifyReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.GameName); validateErr != nil {
		err = errors.Wrap(validateErr, "主游戏名称")
		return
	}
	if a.RootGameId <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "根游戏ID")
		return
	}
	return
}

type MainGameModifyResp struct {
}

type RootGameListReq struct {
	operation_management.DimRootGameModel
	request.PageInfo
}

type RootGameListResp struct {
}

type RootGameAddReq struct {
	operation_management.DimRootGameModel
}

func (a *RootGameAddReq) Format() {
	a.Id = 0
	a.GameName = strings.TrimSpace(a.GameName)
	a.ContractName = strings.TrimSpace(a.ContractName)
}

func (a *RootGameAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.GameName); validateErr != nil {
		err = errors.Wrap(validateErr, "根游戏名称")
		return
	}
	if validateErr := validate.EmptyString(a.ContractName); validateErr != nil {
		err = errors.Wrap(validateErr, "合同游戏名称")
		return
	}
	if a.ProfitSharingRatio <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "分成比例")
		return
	}
	return
}

type RootGameAddResp struct {
}

type RootGameModifyReq struct {
	operation_management.DimRootGameModel
}

func (a *RootGameModifyReq) Format() {
	a.GameName = strings.TrimSpace(a.GameName)
	a.ContractName = strings.TrimSpace(a.ContractName)
}

func (a *RootGameModifyReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.GameName); validateErr != nil {
		err = errors.Wrap(validateErr, "根游戏名称")
		return
	}
	if validateErr := validate.EmptyString(a.ContractName); validateErr != nil {
		err = errors.Wrap(validateErr, "合同游戏名称")
		return
	}
	if a.ProfitSharingRatio <= 0 {
		err = errors.Wrap(error2.ErrorParamEmpty, "分成比例")
		return
	}
	return
}

type RootGameModifyResp struct {
}