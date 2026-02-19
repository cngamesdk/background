package api

import (
	"context"
	"encoding/json"
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	errors2 "github.com/pkg/errors"
	"net/url"
)

type AuthStateData struct {
	UserId      uint  `json:"user_id"`
	DeveloperId int64 `json:"developer_id"`
}

type AdvertisingAuthRedirectReq struct {
	advertising.DimAdvertisingDeveloperConfigModel
	State    string `json:"state"`
	AuthType string `json:"auth_type"` // 磁力引擎引擎存在授权类型
}

func (receiver *AdvertisingAuthRedirectReq) Format() {
	receiver.State = ""
}

func (receiver *AdvertisingAuthRedirectReq) Validate(ctx context.Context) (err error) {
	if receiver.Id <= 0 {
		err = errors2.Wrap(error2.ErrorParamEmpty, "id")
		return
	}
	ginCtx, ok := ctx.(*gin.Context)
	if !ok {
		err = errors2.Wrap(error2.ErrorParamEmpty, "上下文必须为gin")
		return
	}
	claims, claimsErr := utils.GetClaims(ginCtx)
	if claimsErr != nil {
		err = errors2.Wrap(claimsErr, "获取登录信息异常")
		return
	}

	model := advertising.NewDimAdvertisingDeveloperConfigModel()
	if takeErr := model.Take(ctx, "*", "id = ?", receiver.Id); takeErr != nil {
		err = errors2.Wrap(error2.ErrorRecordIsNotFind, "配置不存在")
		return
	}
	stateData := AuthStateData{
		UserId:      claims.BaseClaims.ID,
		DeveloperId: receiver.Id,
	}
	stateDataByte, stateDataErr := json.Marshal(stateData)
	if stateDataErr != nil {
		err = errors2.Wrap(stateDataErr, "JSON state异常")
		return
	}
	cryptStateData := cryptor.AesEcbEncrypt(stateDataByte, []byte(global.GVA_CONFIG.Common.AesCryptKey))
	receiver.State = url.QueryEscape(string(cryptStateData))

	receiver.DimAdvertisingDeveloperConfigModel = *model

	return
}

type AdvertisingAuthRedirectResp struct {
	Url string `json:"url"`
}
