package ad_platform

import (
	"context"
	error2 "github.com/cngamesdk/go-core/model/error"
	"github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

// AdapterConfig 广告平台适配器配置
type AdapterConfig struct {
	AppID        string        `json:"app_id"`
	AppSecret    string        `json:"app_secret"`
	AdvertiserID string        `json:"advertiser_id"`
	AccessToken  string        `json:"access_token"`
	BaseURL      string        `json:"base_url"`
	Timeout      time.Duration `json:"timeout"`
}

// GetAdapterFactory 获取适配器工厂
func GetAdapterFactory(code string, logger *zap.Logger) (resp Adapter, err error) {
	switch code {
	case advertising.MediaCodeOceanengine:
		resp = NewOceanEngineAdapter(logger)
		return
	case advertising.MediaCodeTencent:
		resp = NewTencentAdAdapter(logger)
		return
	case advertising.MediaCodeKuaishou:
		resp = NewKuaiShouAdapter(logger)
		return
	default:
		err = errors.Wrap(error2.ErrorParamEmpty, code)
		return
	}
}

// Adapter 平台适配器接口
type Adapter interface {
	// 平台信息
	Name() string
	Code() string

	//授权
	AuthRedirect(ctx context.Context, req *api.AdvertisingAuthRedirectReq) (resp api.AdvertisingAuthRedirectResp, err error)
	AuthCallback(ctx context.Context, req map[string]interface{}) (resp AuthCallbackResp, err error)

	// 初始化
	Init(config AdapterConfig) error
	RefreshToken(ctx context.Context) error

	// 投放管理
	CreateCampaign(ctx context.Context, req *CreateCampaignRequest) (*CampaignResponse, error)
	UpdateCampaign(ctx context.Context, externalID string, req *UpdateCampaignRequest) (*CampaignResponse, error)
	PauseCampaign(ctx context.Context, externalID string) error
	ResumeCampaign(ctx context.Context, externalID string) error
	DeleteCampaign(ctx context.Context, externalID string) error
	GetCampaign(ctx context.Context, externalID string) (*CampaignResponse, error)

	// 数据同步
	SyncCampaigns(ctx context.Context, startTime, endTime time.Time) ([]*CampaignResponse, error)
	GetMetrics(ctx context.Context, externalID string, startDate, endDate string) ([]*DailyMetrics, error)

	// 账户管理
	GetBalance(ctx context.Context) (float64, error)
	GetDailyReport(ctx context.Context, date string) ([]*DailyMetrics, error)
}

type AuthCallbackResp struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
}

// CreateCampaignRequest 创建投放请求
type CreateCampaignRequest struct {
	Name        string                `json:"name"`
	Budget      float64               `json:"budget"`
	DailyBudget float64               `json:"daily_budget"`
	StartTime   time.Time             `json:"start_time"`
	EndTime     time.Time             `json:"end_time"`
	Targeting   advertising.Targeting `json:"targeting"`
	Creative    advertising.Creative  `json:"creative"`
}

type UpdateCampaignRequest struct {
}

// CampaignResponse 投放响应
type CampaignResponse struct {
	ExternalID string                 `json:"external_id"`
	Status     string                 `json:"status"`
	Message    string                 `json:"message"`
	Data       map[string]interface{} `json:"data"`
}

// DailyMetrics 每日指标
type DailyMetrics struct {
	Date        string  `json:"date"`
	Impressions int64   `json:"impressions"`
	Clicks      int64   `json:"clicks"`
	Conversions int64   `json:"conversions"`
	Spend       float64 `json:"spend"`
	CTR         float64 `json:"ctr"`
	CPC         float64 `json:"cpc"`
	CPA         float64 `json:"cpa"`
	ROI         float64 `json:"roi"`
}
