package ad_platform

import (
	"context"
	"encoding/json"
	"fmt"
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/random"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	advertising2 "github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	url2 "net/url"
	"strings"
	"time"
)

const (
	TencentAdDevelopersUrl = "https://developers.e.qq.com"
	TencentAdApiUrl        = "https://api.e.qq.com"
)

type TencentAdAdapter struct {
	baseAd
}

func NewTencentAdAdapter(logger *zap.Logger) *TencentAdAdapter {
	return &TencentAdAdapter{baseAd{logger: logger}}
}

func (o *TencentAdAdapter) Name() string {
	return "腾讯广告"
}

func (o *TencentAdAdapter) Code() string {
	return advertising.MediaCodeTencent
}

func (o *TencentAdAdapter) GetAuthCallbackUrl() string {
	return global.GVA_CONFIG.Common.Endpoint + "/advertising/auth/callback" + o.Code()
}

func (o *TencentAdAdapter) AuthRedirect(ctx context.Context, req *api.AdvertisingAuthRedirectReq) (resp api.AdvertisingAuthRedirectResp, err error) {
	url := fmt.Sprintf("%s/oauth/authorize?client_id=%s&state=%s&redirect_uri=%s", TencentAdDevelopersUrl, req.AppId, req.State, url2.QueryEscape(o.GetAuthCallbackUrl()))
	resp.Url = url
	return
}

func (o *TencentAdAdapter) AuthCallback(ctx context.Context, req map[string]interface{}) (resp advertising2.DimAdvertisingMediaAuthModel, err error) {
	authorizationCode, ok := req["authorization_code"]
	if !ok {
		err = errors.Wrap(error2.ErrorParamEmpty, "authorization_code为空")
		return
	}
	state, stateOk := req["state"]
	if !stateOk {
		err = errors.Wrap(error2.ErrorParamEmpty, "state为空")
		return
	}
	stateResult, stateResultErr := o.formatState(ctx, cast.ToString(state))
	if stateResultErr != nil {
		err = stateResultErr
		o.logger.Error("格式化state异常", zap.Error(stateResultErr))
		return
	}
	response, responseErr := o.getNewRestyClient().
		SetBaseURL(TencentAdApiUrl).
		SetQueryParams(map[string]string{
			"client_id":          stateResult.DeveloperInfo.AppId,
			"client_secret":      stateResult.DeveloperInfo.Secret,
			"grant_type":         "authorization_code",
			"authorization_code": cast.ToString(authorizationCode),
			"redirect_uri":       url2.QueryEscape(o.GetAuthCallbackUrl()),
		}).R().Get("/oauth/token")
	if responseErr != nil {
		err = responseErr
		o.logger.Error("获取token异常", zap.Error(responseErr))
		return
	}
	var tokenResponse struct {
		AuthorizerInfo struct {
			AccountUin      int64    `json:"account_uin"`
			AccountId       int64    `json:"account_id"`
			ScopeList       []string `json:"scope_list"`
			WechatAccountId string   `json:"wechat_account_id"`
			AccountRoleType string   `json:"account_role_type"`
			AccountType     string   `json:"account_type"`
			RoleType        string   `json:"role_type"`
		} `json:"authorizer_info"`
		AccessToken           string `json:"access_token"`
		RefreshToken          string `json:"refresh_token"`
		AccessTokenExpiresIn  int    `json:"access_token_expires_in"`
		RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	}
	dealErr := o.dealResponse(response, &tokenResponse)
	if dealErr != nil {
		err = dealErr
		o.logger.Error("处理返回异常", zap.Error(dealErr))
		return
	}
	responseJson, responseJsonErr := json.Marshal(tokenResponse)
	if responseJsonErr != nil {
		err = responseJsonErr
		o.logger.Error("json response异常", zap.Error(responseJsonErr))
		return
	}
	var extensionMap sql.CustomMapType
	if responseJsonUnErr := json.Unmarshal(responseJson, &extensionMap); responseJsonUnErr != nil {
		err = responseJsonUnErr
		o.logger.Error("json response反序列化异常", zap.Error(responseJsonUnErr))
		return
	}
	resp.PlatformId = stateResult.PlatformId
	resp.AccountId = tokenResponse.AuthorizerInfo.AccountId
	resp.Status = sql.StatusNormal
	resp.Extension = extensionMap
	resp.AccessToken = tokenResponse.AccessToken
	resp.RefreshToken = tokenResponse.RefreshToken
	resp.ExpiresAt = sql.MyCustomDatetime(time.Now().Add(time.Duration(tokenResponse.AccessTokenExpiresIn) * time.Second))
	resp.RefreshTokenExpiresAt = sql.MyCustomDatetime(time.Now().Add(time.Duration(tokenResponse.RefreshTokenExpiresIn) * time.Second))
	resp.DeveloperId = stateResult.DeveloperId
	resp.Code = o.Code()
	return
}

func (o *TencentAdAdapter) dealResponse(req *resty.Response, dst interface{}) (err error) {
	if req == nil {
		err = errors.New("the response is nil")
		return
	}
	var result struct {
		Code      int                    `json:"code"`
		Message   string                 `json:"message"`
		MessageCn string                 `json:"message_cn"`
		RequestId string                 `json:"request_id"`
		Data      map[string]interface{} `json:"data"`
	}
	if errJson := json.Unmarshal(req.Body(), &result); errJson != nil {
		err = errJson
		return
	}
	if result.Code != 0 {
		err = errors.New(fmt.Sprintf("code: %d, message: %s, request_id: %s", result.Code, result.Message, result.RequestId))
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

func (o *TencentAdAdapter) buildGlobalParams() string {
	nonce, _ := random.UUIdV4()
	globalParams := []string{
		"access_token=" + o.config.Auth.AccessToken,
		"timestamp=" + cast.ToString(time.Now().Unix()),
		"nonce=" + cryptor.Md5String(nonce),
	}
	return strings.Join(globalParams, "&")
}

// AuthAdvertiserGet 授权后获取广告主
func (o *TencentAdAdapter) AuthAdvertiserGet(ctx context.Context) (resp []advertising2.DimAdvertisingMediaAccountModel, err error) {
	fields := []string{
		"account_id",
		"daily_budget",
		"registration_type",
		"corporation_name",
		"corporation_licence",
		"certification_image_id",
		"certification_image",
		"individual_qualification",
		"area_code",
		"mdm_id",
		"mdm_name",
		"system_industry_id",
		"customized_industry",
		"introduction_url",
		"corporate_brand_name",
		"memo",
		"system_status",
		"reject_message",
		"is_adx",
		"business_alias",
		"contact_person",
		"contact_person_email",
		"contact_person_telephone",
		"contact_person_mobile",
		"websites",
		"agency_account_id",
		"operators",
	}
	fieldsJson, _ := json.Marshal(fields)
	type advertiserGetResponse struct {
		List []struct {
			AccountId               int64  `json:"account_id"`
			DailyBudget             int64  `json:"daily_budget"`
			RegistrationType        string `json:"registration_type"`
			CorporationName         string `json:"corporation_name"`
			CorporationLicence      string `json:"corporation_licence"`
			CertificationImageId    string `json:"certification_image_id"`
			CertificationImage      string `json:"certification_image"`
			IndividualQualification struct {
				Name                       string `json:"name"`
				IdentificationNumber       string `json:"identification_number"`
				IdentificationFrontImageId string `json:"identification_front_image_id"`
				IdentificationBackImageId  string `json:"identification_back_image_id"`
			} `json:"individual_qualification"`
			AreaCode               int    `json:"area_code"`
			MdmId                  int    `json:"mdm_id"`
			MdmName                string `json:"mdm_name"`
			SystemIndustryId       int    `json:"system_industry_id"`
			CustomizedIndustry     string `json:"customized_industry"`
			IntroductionUrl        string `json:"introduction_url"`
			CorporateBrandName     string `json:"corporate_brand_name"`
			Memo                   string `json:"memo"`
			SystemStatus           string `json:"system_status"`
			RejectMessage          string `json:"reject_message"`
			IsAdx                  bool   `json:"is_adx"`
			BusinessAlias          string `json:"business_alias"`
			ContactPerson          string `json:"contact_person"`
			ContactPersonEmail     string `json:"contact_person_email"`
			ContactPersonTelephone string `json:"contact_person_telephone"`
			ContactPersonMobile    string `json:"contact_person_mobile"`
			Websites               struct {
				WebsiteDomain string `json:"website_domain"`
				IcpImageId    string `json:"icp_image_id"`
				SystemStatus  string `json:"system_status"`
				RejectMessage string `json:"reject_message"`
			} `json:"websites"`
			AgencyAccountId string `json:"agency_account_id"`
			Operators       struct {
				OperatorId      int    `json:"operator_id"`
				OperatorName    string `json:"operator_name"`
				Qq              int64  `json:"qq"`
				WechatAccountId string `json:"wechat_account_id"`
				IsMaster        bool   `json:"is_master"`
			} `json:"operators"`
		} `json:"list"`
		PageInfo struct {
			Page        int `json:"page"`
			PageSize    int `json:"page_size"`
			TotalNumber int `json:"total_number"`
			TotalPage   int `json:"total_page"`
		} `json:"page_info"`
		CursorPageInfo struct {
			PageSize    int  `json:"page_size"`
			TotalNumber int  `json:"total_number"`
			HasMore     bool `json:"has_more"`
			Cursor      int  `json:"cursor"`
		} `json:"cursor_page_info"`
	}
	cursor := 1
	for {
		response, responseErr := o.getRestyClient().
			SetBaseURL(TencentAdApiUrl).
			R().
			SetQueryParams(map[string]string{
				"account_id":      o.config.AdvertiserID,
				"fields":          string(fieldsJson),
				"pagination_mode": "PAGINATION_MODE_CURSOR",
				"page_size":       cast.ToString(100),
				"cursor":          cast.ToString(cursor),
			}).
			Get("/v3.0/advertiser/get?" + o.buildGlobalParams())
		if responseErr != nil {
			err = responseErr
			o.logger.Error("获取列表异常", zap.Error(responseErr))
			return
		}
		var tmpResponse advertiserGetResponse
		if err = o.dealResponse(response, &tmpResponse); err != nil {
			o.logger.Error("处理返回异常", zap.Error(err))
			return
		}
		for _, item := range tmpResponse.List {

			extension, extensionErr := json.Marshal(item)
			if extensionErr != nil {
				o.logger.Error("item序列化异常", zap.Error(extensionErr))
			}
			var extensionMap sql.CustomMapType
			if extensionMapUnJsonErr := json.Unmarshal(extension, &extensionMap); extensionMapUnJsonErr != nil {
				o.logger.Error("item反序列化异常", zap.Error(extensionMapUnJsonErr))
			}
			tempModel := advertising2.DimAdvertisingMediaAccountModel{}
			tempModel.AccountName = cast.ToString(item.AccountId)
			tempModel.AccountId = item.AccountId
			tempModel.Status = o.convertStatus(item.SystemStatus)
			tempModel.AuthId = o.config.Auth.Id
			tempModel.Role = ""
			tempModel.Extension = extensionMap
			tempModel.Code = o.Code()
			resp = append(resp, tempModel)
		}

		if tmpResponse.CursorPageInfo.HasMore {
			cursor = tmpResponse.CursorPageInfo.Cursor
		} else {
			break
		}
	}
	return
}

func (o *TencentAdAdapter) convertStatus(status string) string {
	// https://developers.e.qq.com/v3.0/docs/enums#api_customer_system_status
	if status == "CUSTOMER_STATUS_NORMAL" {
		return sql.StatusNormal
	}
	return sql.StatusFail
}

func (o *TencentAdAdapter) RefreshToken(ctx context.Context) (resp advertising2.DimAdvertisingMediaAuthModel, err error) {
	o.logger.Info("Refreshing token")

	return
}

func (o *TencentAdAdapter) CreateCampaign(ctx context.Context, req *CreateCampaignRequest) (*CampaignResponse, error) {
	return nil, nil
}

func (o *TencentAdAdapter) convertToOceanEngine(req *CreateCampaignRequest) map[string]interface{} {
	return nil
}

func (o *TencentAdAdapter) convertLocation(locations []string) []map[string]interface{} {
	return nil
}

func (o *TencentAdAdapter) UpdateCampaign(ctx context.Context, externalID string, req *UpdateCampaignRequest) (resp *CampaignResponse, err error) {
	return
}

func (o *TencentAdAdapter) PauseCampaign(ctx context.Context, externalID string) (err error) {
	return
}

func (o *TencentAdAdapter) ResumeCampaign(ctx context.Context, externalID string) (err error) {
	return
}

func (o *TencentAdAdapter) DeleteCampaign(ctx context.Context, externalID string) (err error) {
	return
}

func (o *TencentAdAdapter) GetCampaign(ctx context.Context, externalID string) (resp *CampaignResponse, err error) {
	return
}

func (o *TencentAdAdapter) SyncCampaigns(ctx context.Context, startTime, endTime time.Time) (resp []*CampaignResponse, err error) {
	return
}

func (o *TencentAdAdapter) GetMetrics(ctx context.Context, externalID string, startDate, endDate string) (resp []*DailyMetrics, err error) {
	return
}

func (o *TencentAdAdapter) GetBalance(ctx context.Context) (resp float64, err error) {
	return
}

func (o TencentAdAdapter) GetDailyReport(ctx context.Context, date string) (resp []*DailyMetrics, err error) {
	return
}
