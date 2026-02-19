package ad_platform

import (
	"context"
	"fmt"
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	TencentAdDevelopersUrl = "https://developers.e.qq.com"
	TencentAdApiUrl        = "https://api.e.qq.com"
)

type TencentAdAdapter struct {
	config   AdapterConfig
	client   *resty.Client
	logger   *zap.Logger
	token    string
	tokenExp time.Time
}

func NewTencentAdAdapter(logger *zap.Logger) *TencentAdAdapter {
	return &TencentAdAdapter{
		logger: logger,
	}
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
	url := fmt.Sprintf("%s/oauth/authorize?client_id=%s&state=%s&redirect_uri=%s", TencentAdDevelopersUrl, req.AppId, req.State, o.GetAuthCallbackUrl())
	resp.Url = url
	return
}

func (o *TencentAdAdapter) AuthCallback(ctx context.Context, req map[string]interface{}) (resp AuthCallbackResp, err error) {
	return
}

func (o *TencentAdAdapter) Init(config AdapterConfig) error {
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

func (o *TencentAdAdapter) RefreshToken(ctx context.Context) error {
	o.logger.Info("Refreshing token")
	return nil
}

func (o *TencentAdAdapter) ensureToken(ctx context.Context) error {
	if time.Now().Add(5 * time.Minute).After(o.tokenExp) {
		return o.RefreshToken(ctx)
	}
	return nil
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
