package ad_platform

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type OceanEngineAdapter struct {
	config   AdapterConfig
	client   *resty.Client
	logger   *zap.Logger
	token    string
	tokenExp time.Time
}

func NewOceanEngineAdapter(logger *zap.Logger) *OceanEngineAdapter {
	return &OceanEngineAdapter{
		logger: logger,
	}
}

func (o *OceanEngineAdapter) Name() string {
	return "巨量引擎"
}

func (o *OceanEngineAdapter) Code() string {
	return advertising.MediaCodeOceanengine
}

func (o *OceanEngineAdapter) AuthRedirect(ctx context.Context, req api.AdvertisingAuthRedirectReq) (resp api.AdvertisingAuthRedirectResp, err error) {
	return
}

func (o *OceanEngineAdapter) Init(config AdapterConfig) error {
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

func (o *OceanEngineAdapter) RefreshToken(ctx context.Context) error {
	o.logger.Info("Refreshing OceanEngine token")

	resp, err := o.client.R().
		SetHeader("Content-Type", "application/json").
		SetContext(ctx).
		SetBody(map[string]string{
			"app_id": o.config.AppID,
			"secret": o.config.AppSecret,
		}).
		Post("/oauth2/app_access_token")

	if err != nil {
		return fmt.Errorf("refresh token failed: %v", err)
	}

	var result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int    `json:"expires_in"`
		} `json:"data"`
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return err
	}

	if result.Code != 0 {
		return fmt.Errorf("refresh token error: %s", result.Message)
	}

	o.token = result.Data.AccessToken
	o.tokenExp = time.Now().Add(time.Duration(result.Data.ExpiresIn) * time.Second)

	o.logger.Info("OceanEngine token refreshed",
		zap.Time("expires_at", o.tokenExp))

	return nil
}

func (o *OceanEngineAdapter) ensureToken(ctx context.Context) error {
	if time.Now().Add(5 * time.Minute).After(o.tokenExp) {
		return o.RefreshToken(ctx)
	}
	return nil
}

func (o *OceanEngineAdapter) CreateCampaign(ctx context.Context, req *CreateCampaignRequest) (*CampaignResponse, error) {
	if err := o.ensureToken(ctx); err != nil {
		return nil, err
	}

	// 转换为巨量引擎格式
	oeCampaign := o.convertToOceanEngine(req)

	resp, err := o.client.R().
		SetContext(ctx).
		SetHeader("Access-Token", o.token).
		SetBody(oeCampaign).
		Post("/2/campaign/create")

	if err != nil {
		return nil, fmt.Errorf("create campaign request failed: %v", err)
	}

	var result struct {
		Code    int                    `json:"code"`
		Message string                 `json:"message"`
		Data    map[string]interface{} `json:"data"`
	}

	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	if result.Code != 0 {
		return nil, fmt.Errorf("create campaign failed: %s", result.Message)
	}

	externalID, _ := result.Data["campaign_id"].(string)

	return &CampaignResponse{
		ExternalID: externalID,
		Status:     advertising.CampaignStatusRunning,
		Data:       result.Data,
	}, nil
}

func (o *OceanEngineAdapter) convertToOceanEngine(req *CreateCampaignRequest) map[string]interface{} {
	// 年龄转换
	age := []string{}
	if req.Targeting.AgeRange[0] >= 18 && req.Targeting.AgeRange[1] <= 23 {
		age = append(age, "AGE_BETWEEN_18_23")
	}
	if req.Targeting.AgeRange[0] >= 24 && req.Targeting.AgeRange[1] <= 30 {
		age = append(age, "AGE_BETWEEN_24_30")
	}
	if req.Targeting.AgeRange[0] >= 31 && req.Targeting.AgeRange[1] <= 40 {
		age = append(age, "AGE_BETWEEN_31_40")
	}
	if req.Targeting.AgeRange[0] >= 41 && req.Targeting.AgeRange[1] <= 50 {
		age = append(age, "AGE_BETWEEN_41_50")
	}

	// 性别转换
	var gender string
	switch req.Targeting.Gender {
	case "male":
		gender = "GENDER_MALE"
	case "female":
		gender = "GENDER_FEMALE"
	default:
		gender = "GENDER_UNLIMITED"
	}

	return map[string]interface{}{
		"advertiser_id": o.config.AdvertiserID,
		"campaign_name": req.Name,
		"budget":        req.Budget,
		"budget_mode":   "BUDGET_MODE_DAY",
		"campaign_type": "FEED",
		"delivery_mode": "STANDARD",
		"delivery_range": map[string]interface{}{
			"location": o.convertLocation(req.Targeting.Location),
			"age":      age,
			"gender":   []string{gender},
			//"platform":        o.convertPlatform(req.Targeting.Platforms),
			//"ac":              o.convertInterests(req.Targeting.Interests),
			//"device_type":     o.convertDeviceType(req.Targeting.DeviceTypes),
			//"network_type":    o.convertNetworkType(req.Targeting.NetworkType),
		},
		"schedule_type": "SCHEDULE_START_END",
		"start_time":    req.StartTime.Format("2006-01-02 15:04:05"),
		"end_time":      req.EndTime.Format("2006-01-02 15:04:05"),
	}
}

func (o *OceanEngineAdapter) convertLocation(locations []string) []map[string]interface{} {
	var result []map[string]interface{}
	for _, loc := range locations {
		result = append(result, map[string]interface{}{
			"name":      loc,
			"level":     "CITY",
			"longitude": 0,
			"latitude":  0,
		})
	}
	return result
}

func (o *OceanEngineAdapter) UpdateCampaign(ctx context.Context, externalID string, req *UpdateCampaignRequest) (resp *CampaignResponse, err error) {
	return
}

func (o *OceanEngineAdapter) PauseCampaign(ctx context.Context, externalID string) (err error) {
	return
}

func (o *OceanEngineAdapter) ResumeCampaign(ctx context.Context, externalID string) (err error) {
	return
}

func (o *OceanEngineAdapter) DeleteCampaign(ctx context.Context, externalID string) (err error) {
	return
}

func (o *OceanEngineAdapter) GetCampaign(ctx context.Context, externalID string) (resp *CampaignResponse, err error) {
	return
}

func (o *OceanEngineAdapter) SyncCampaigns(ctx context.Context, startTime, endTime time.Time) (resp []*CampaignResponse, err error) {
	return
}

func (o *OceanEngineAdapter) GetMetrics(ctx context.Context, externalID string, startDate, endDate string) (resp []*DailyMetrics, err error) {
	return
}

func (o *OceanEngineAdapter) GetBalance(ctx context.Context) (resp float64, err error) {
	return
}

func (o OceanEngineAdapter) GetDailyReport(ctx context.Context, date string) (resp []*DailyMetrics, err error) {
	return
}
