package ad_platform

import (
	"context"
	"fmt"
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
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

type KuaiShouAdapter struct {
	config   AdapterConfig
	client   *resty.Client
	logger   *zap.Logger
	token    string
	tokenExp time.Time
}

func NewKuaiShouAdapter(logger *zap.Logger) *KuaiShouAdapter {
	return &KuaiShouAdapter{
		logger: logger,
	}
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
	url := fmt.Sprintf("%s/tools/authorize?app_id=%s&scope=%s&redirect_uri=%s&state=%s&oauth_type=%s",
		KuaiShouDevelopersUrl, req.AppId, "["+strings.Join(scopes, ",")+"]", o.GetAuthCallbackUrl(), req.State, req.AuthType)
	resp.Url = url
	return
}

func (o *KuaiShouAdapter) AuthCallback(ctx context.Context, req map[string]interface{}) (resp AuthCallbackResp, err error) {
	return
}

func (o *KuaiShouAdapter) Init(config AdapterConfig) error {
	o.config = config
	o.client = resty.New().
		SetBaseURL(config.BaseURL).
		SetTimeout(config.Timeout).
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return r.StatusCode() >= http.StatusInternalServerError
		})

	return o.RefreshToken(context.Background())
}

func (o *KuaiShouAdapter) RefreshToken(ctx context.Context) error {
	o.logger.Info("Refreshing token")
	return nil
}

func (o *KuaiShouAdapter) ensureToken(ctx context.Context) error {
	if time.Now().Add(5 * time.Minute).After(o.tokenExp) {
		return o.RefreshToken(ctx)
	}
	return nil
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
