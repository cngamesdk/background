package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/pkg/errors"
)

type GamePackagingLogListReq struct {
	request.PageInfo
	operation_management.OdsGamePackagingLogModel
}

type GamePackagingLogListResp struct {
}

type GamePackagingAddReq struct {
	operation_management.OdsGamePackagingLogModel
}

func (receiver *GamePackagingAddReq) Format() {
	receiver.Id = 0
}

func (receiver *GamePackagingAddReq) Validate() (err error) {
	if receiver.PlatformId <= 0 {
		err = errors.New("平台不能为空")
		return
	}
	if receiver.GameId <= 0 {
		err = errors.New("游戏ID不能为空")
		return
	}
	if receiver.SiteId <= 0 {
		err = errors.New("广告位ID不能为空")
		return
	}
	return
}
