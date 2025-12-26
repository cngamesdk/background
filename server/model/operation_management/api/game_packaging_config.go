package api

import (
	"github.com/cngamesdk/go-core/model/sql/advertising"
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
	receiver.CommonMedia = strings.TrimSpace(receiver.CommonMedia)
	receiver.GamePackagePath = strings.TrimSpace(receiver.GamePackagePath)
	receiver.Status = strings.TrimSpace(receiver.Status)
	receiver.UseStatus = strings.TrimSpace(receiver.UseStatus)
}

func (receiver *GamePackagingConfigAddReq) Validate() (err error) {
	if _, ok := advertising.CommonMediasMap[receiver.CommonMedia]; !ok {
		err = errors.New("媒体未知")
		return
	}
	if receiver.GamePackagePath == "" {
		err = errors.New(receiver.CommonMedia + " 对应的游戏母包地址不能为空")
		return
	}
	if receiver.GamePackageHash == "" {
		err = errors.New(receiver.CommonMedia + " 对应的游戏母包校验HASH不能为空")
		return
	}
	if receiver.Status == "" {
		err = errors.New("状态不能为空")
		return
	}
	if receiver.UseStatus == "" {
		err = errors.New("使用状态不能为空")
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
	receiver.CommonMedia = strings.TrimSpace(receiver.CommonMedia)
	receiver.GamePackagePath = strings.TrimSpace(receiver.GamePackagePath)
	receiver.Status = strings.TrimSpace(receiver.Status)
	receiver.UseStatus = strings.TrimSpace(receiver.UseStatus)
}

func (receiver *GamePackagingConfigModifyReq) Validate() (err error) {
	if receiver.Id <= 0 {
		err = errors.New("主键ID不能为空")
		return
	}
	if _, ok := advertising.CommonMediasMap[receiver.CommonMedia]; !ok {
		err = errors.New("媒体未知")
		return
	}
	if receiver.GamePackagePath == "" {
		err = errors.New(receiver.CommonMedia + " 对应的游戏母包地址不能为空")
		return
	}
	if receiver.GamePackageHash == "" {
		err = errors.New(receiver.CommonMedia + " 对应的游戏母包校验HASH不能为空")
		return
	}
	if receiver.Status == "" {
		err = errors.New("状态不能为空")
		return
	}
	if receiver.UseStatus == "" {
		err = errors.New("使用状态不能为空")
		return
	}
	return
}

type GamePackagingConfigModifyResp struct {
}
