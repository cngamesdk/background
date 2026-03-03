package api

import (
	"context"
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	errors2 "github.com/pkg/errors"
	"github.com/spf13/cast"
	"strings"
)

type AdvertisingMixListReq struct {
	advertising.OdsAdvertisingMixLogModel
	request.PageInfo
}

type AdvertisingMixListResp struct {
}

type AdvertisingMixAddReq struct {
	advertising.OdsAdvertisingMixLogModel
}

func (receiver *AdvertisingMixAddReq) Format(ctx context.Context) {
	receiver.Id = 0
	receiver.Code = strings.TrimSpace(receiver.Code)
	receiver.Name = strings.TrimSpace(receiver.Name)
	receiver.Status = sql.StatusNormal
	if ginCtx, ok := ctx.(*gin.Context); ok {
		receiver.UserId = cast.ToInt64(utils.GetUserID(ginCtx))
	}
}

func (receiver *AdvertisingMixAddReq) Validate() (err error) {
	if receiver.UserId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "user_id")
		return
	}
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if receiver.Code == "" {
		err = errors2.Wrap(error2.ErrorParamEmpty, "code")
		return
	}
	if receiver.CommonConfig.GameId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "子游戏")
		return
	}
	if len(receiver.AccountConfig.List) <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "帐户")
		return
	}
	if validateErr := validate.EmptyString(receiver.Name); validateErr != nil {
		err = errors2.Wrap(validateErr, "组合名称")
		return
	}
	return
}

type AdvertisingMixAddResp struct {
}

type AdvertisingMixModifyReq struct {
	advertising.OdsAdvertisingMixLogModel
}

func (receiver *AdvertisingMixModifyReq) Format(ctx context.Context) {
	receiver.Code = strings.TrimSpace(receiver.Code)
	receiver.Name = strings.TrimSpace(receiver.Name)
}

func (receiver *AdvertisingMixModifyReq) Validate() (err error) {
	if receiver.Id <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	if receiver.PlatformId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "platform_id")
		return
	}
	if receiver.Code == "" {
		err = errors2.Wrap(error2.ErrorParamEmpty, "code")
		return
	}
	if receiver.CommonConfig.GameId <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "子游戏")
		return
	}
	if len(receiver.AccountConfig.List) <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "帐户")
		return
	}
	if validateErr := validate.EmptyString(receiver.Name); validateErr != nil {
		err = errors2.Wrap(validateErr, "组合名称")
		return
	}
	return
}

type AdvertisingMixModifyResp struct {
}
