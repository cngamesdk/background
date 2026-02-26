package ad_platform

import (
	"context"
	"encoding/json"
	"fmt"
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	advertising2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	url2 "net/url"
	"slices"
	"strings"
	"time"
)

const (
	KuaiShouDevelopersUrl = "https://developers.e.kuaishou.com"
	KuaiShouAdUrl         = "https://ad.e.kuaishou.com"

	KuaiShouScopeReportService  = "report_service"  // 获取广告账户报表信息&建站信息
	KuaiShouScopeAccountService = "account_service" // 获取账户信息、余额、流水
	KuaiShouScopeAdQuery        = "ad_query"        // 获取广告计划、广告组、广告创意信息
	KuaiShouScopeAdManage       = "ad_manage"       // 创建&修改广告计划、广告组、广告创意&人群管理
	KuaiShouScopeAccountCert    = "account_cert"    // 代理商开通账户，此功能需要加白后使用
	KuaiShouScopeAdSeries       = "ad_series"       // 短剧业务使用

	KuaiShouAuthTypeAdvertiser = "advertiser" //广告主 / 服务商
	KuaiShouAuthTypeAgent      = "agent"      //代理商
	KuaiShouAuthTypeAdSocial   = "ad_social"  //聚星达人
	KuaiShouAuthTypeSeries     = "series"     //短剧广告主
)

type tokenResp struct {
	AccessToken           string `json:"access_token"`
	AccessTokenExpiresIn  int64  `json:"access_token_expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
}

type KuaiShouAdapter struct {
	baseAd
}

func NewKuaiShouAdapter(logger *zap.Logger) *KuaiShouAdapter {
	return &KuaiShouAdapter{baseAd{logger: logger}}
}

func (o *KuaiShouAdapter) Name() string {
	return "磁力引擎"
}

func (o *KuaiShouAdapter) Code() string {
	return advertising.MediaCodeKuaishou
}

func (o *KuaiShouAdapter) GetAuthCallbackUrl() string {
	return global.GVA_CONFIG.Common.Endpoint + "/advertising/auth/callback" + o.Code()
}

func (o *KuaiShouAdapter) AuthRedirect(ctx context.Context, req *api.AdvertisingAuthRedirectReq) (resp api.AdvertisingAuthRedirectResp, err error) {
	scopes := []string{
		KuaiShouScopeReportService,
		KuaiShouScopeAccountService,
		KuaiShouScopeAdQuery,
		KuaiShouScopeAdManage,
		KuaiShouScopeAccountCert,
		KuaiShouScopeAdSeries,
	}
	if !slices.Contains([]string{KuaiShouAuthTypeAdvertiser, KuaiShouAuthTypeAgent, KuaiShouAuthTypeAdSocial, KuaiShouAuthTypeSeries}, req.AuthType) {
		err = errors.New("授权类型未知" + req.AuthType)
		return
	}
	scopesEscape := url2.QueryEscape("[" + strings.Join(scopes, ",") + "]")
	url := fmt.Sprintf("%s/tools/authorize?app_id=%s&scope=%s&redirect_uri=%s&state=%s&oauth_type=%s",
		KuaiShouDevelopersUrl, req.AppId, scopesEscape, url2.QueryEscape(o.GetAuthCallbackUrl()+"?state="+req.State), req.State, req.AuthType)
	resp.Url = url
	return
}

func (o *KuaiShouAdapter) dealResponse(req *resty.Response, dst interface{}) (err error) {
	if req == nil {
		err = errors.New("the response is nil")
		return
	}
	var result struct {
		Code    int                    `json:"code"`
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data"`
	}
	if errJson := json.Unmarshal(req.Body(), &result); errJson != nil {
		err = errJson
		return
	}
	if result.Code != 0 {
		err = errors.New(fmt.Sprintf("code: %d, message: %s", result.Code, result.Message))
		o.logger.Error("return error",
			zap.Any("header", req.Request.Header),
			zap.Any("url", req.Request.URL),
			zap.Any("body", req.Request.Body),
			zap.Any("response", req.Body()))
		return
	}
	tempData, tempErr := json.Marshal(result.Data)
	if tempErr != nil {
		err = errors.Wrap(tempErr, "json marshal error")
		return
	}
	if unmarshalErr := json.Unmarshal(tempData, dst); unmarshalErr != nil {
		err = unmarshalErr
		return
	}
	return
}

func (o *KuaiShouAdapter) AuthCallback(ctx context.Context, req map[string]interface{}) (resp advertising2.DimAdvertisingMediaAuthModel, err error) {
	authCode, authCodeOk := req["auth_code"]
	if !authCodeOk {
		err = errors.Wrap(error2.ErrorParamEmpty, "auth_code为空")
		return
	}
	state, stateOk := req["state"]
	if !stateOk {
		err = errors.Wrap(error2.ErrorParamEmpty, "state为空")
		return
	}
	stateData, stateErr := o.formatState(ctx, cast.ToString(state))
	if stateErr != nil {
		err = stateErr
		o.logger.Error("解析state异常", zap.Error(stateErr))
		return
	}
	response, responseErr := o.getNewRestyClient().
		SetBaseURL(KuaiShouAdUrl).
		R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"app_id":    stateData.DeveloperInfo.AppId,
			"secret":    stateData.DeveloperInfo.Secret,
			"auth_code": cast.ToString(authCode),
		}).
		Post("/rest/openapi/oauth2/authorize/access_token")
	if responseErr != nil {
		err = responseErr
		o.logger.Error("请求token异常", zap.Error(responseErr))
		return
	}
	var result tokenResp
	handleResponseErr := o.dealResponse(response, &result)
	if handleResponseErr != nil {
		err = handleResponseErr
		o.logger.Error("解析response异常", zap.Error(handleResponseErr))
		return
	}
	resp.PlatformId = stateData.PlatformId
	resp.Status = sql.StatusNormal
	resp.AccessToken = result.AccessToken
	resp.RefreshToken = result.RefreshToken
	resp.ExpiresAt = sql.MyCustomDatetime(time.Now().Add(time.Duration(result.AccessTokenExpiresIn) * time.Second))
	resp.RefreshTokenExpiresAt = sql.MyCustomDatetime(time.Now().Add(time.Duration(result.RefreshTokenExpiresIn) * time.Second))
	resp.DeveloperId = stateData.DeveloperId
	resp.Code = o.Code()
	return
}

// AuthAdvertiserGet 授权后获取广告主
func (o *KuaiShouAdapter) AuthAdvertiserGet(ctx context.Context) (resp []advertising2.DimAdvertisingMediaAccountModel, err error) {
	page := 1
	pageSize := 100
	type result struct {
		IsEnd   bool    `json:"isEnd"`
		Details []int64 `json:"details"`
	}
	for {
		response, responseErr := o.getRestyClient().
			SetBaseURL(KuaiShouAdUrl).
			R().
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]interface{}{
				"app_id":       o.config.Developer.AppId,
				"secret":       o.config.Developer.Secret,
				"access_token": o.config.Auth.AccessToken,
				"page_no":      page,
				"page_size":    pageSize,
			}).
			Post("/rest/openapi/oauth2/authorize/approval/list")
		if responseErr != nil {
			err = responseErr
			o.logger.Error("获取列表异常", zap.Error(responseErr))
			return
		}
		var responseResult result
		if err = o.dealResponse(response, &responseResult); err != nil {
			o.logger.Error("解析response异常", zap.Error(err))
			return
		}
		for _, item := range responseResult.Details {
			tempModel := advertising2.DimAdvertisingMediaAccountModel{}
			tempModel.PlatformId = o.config.Auth.PlatformId
			tempModel.AccountName = cast.ToString(item)
			tempModel.AccountId = item
			tempModel.AuthId = o.config.Auth.Id
			tempModel.Status = sql.StatusNormal
			tempModel.Code = o.Code()
			resp = append(resp, tempModel)
		}
		if responseResult.IsEnd {
			break
		}
		page++
	}
	return
}

func (o *KuaiShouAdapter) Init(config AdapterConfig) error {
	o.config = config
	o.client = o.getNewRestyClient()
	return nil
}

func (o *KuaiShouAdapter) RefreshToken(ctx context.Context) (resp advertising2.DimAdvertisingMediaAuthModel, err error) {
	o.logger.Info("Refreshing token")
	response, respErr := o.getRestyClient().
		SetBaseURL(KuaiShouAdUrl).
		R().
		SetContext(ctx).
		SetQueryParams(map[string]string{
			"app_id":        o.config.Developer.AppId,
			"secret":        o.config.Developer.Secret,
			"refresh_token": o.config.Auth.RefreshToken,
		}).
		Get("/rest/openapi/oauth2/authorize/refresh_token")

	if respErr != nil {
		o.logger.Error("刷新token异常", zap.Error(respErr))
		err = fmt.Errorf("refresh token failed: %v", respErr)
		return
	}
	var result tokenResp
	handleErr := o.dealResponse(response, &result)
	if handleErr != nil {
		o.logger.Error("解析response异常", zap.Error(handleErr))
		err = fmt.Errorf("refresh token failed: %v", respErr)
		return
	}
	resp.AccessToken = result.AccessToken
	resp.RefreshToken = result.RefreshToken
	resp.ExpiresAt = sql.MyCustomDatetime(time.Now().Add(time.Duration(result.AccessTokenExpiresIn) * time.Second))
	resp.RefreshTokenExpiresAt = sql.MyCustomDatetime(time.Now().Add(time.Duration(result.RefreshTokenExpiresIn) * time.Second))
	return
}

func (o *KuaiShouAdapter) CreateCampaign(ctx context.Context, req *CreateCampaignRequest) (*CampaignResponse, error) {
	return nil, nil
}

func (o *KuaiShouAdapter) convertToOceanEngine(req *CreateCampaignRequest) map[string]interface{} {
	return nil
}

func (o *KuaiShouAdapter) convertLocation(locations []string) []map[string]interface{} {
	return nil
}

func (o *KuaiShouAdapter) UpdateCampaign(ctx context.Context, externalID string, req *UpdateCampaignRequest) (resp *CampaignResponse, err error) {
	return
}

func (o *KuaiShouAdapter) PauseCampaign(ctx context.Context, externalID string) (err error) {
	return
}

func (o *KuaiShouAdapter) ResumeCampaign(ctx context.Context, externalID string) (err error) {
	return
}

func (o *KuaiShouAdapter) DeleteCampaign(ctx context.Context, externalID string) (err error) {
	return
}

func (o *KuaiShouAdapter) GetCampaign(ctx context.Context, externalID string) (resp *CampaignResponse, err error) {
	return
}

func (o *KuaiShouAdapter) SyncCampaigns(ctx context.Context, startTime, endTime time.Time) (resp []*CampaignResponse, err error) {
	return
}

func (o *KuaiShouAdapter) GetMetrics(ctx context.Context, externalID string, startDate, endDate string) (resp []*DailyMetrics, err error) {
	return
}

func (o *KuaiShouAdapter) GetBalance(ctx context.Context) (resp float64, err error) {
	return
}

func (o KuaiShouAdapter) GetDailyReport(ctx context.Context, date string) (resp []*DailyMetrics, err error) {
	return
}
