package ad_platform

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type baseAd struct {
	config   AdapterConfig
	client   *resty.Client
	logger   *zap.Logger
	token    string
	tokenExp time.Time
}

func (b *baseAd) getNewRestyClient() *resty.Client {
	return resty.New().
		SetBaseURL(b.config.BaseURL).
		SetTimeout(b.config.Timeout).
		SetRetryCount(3).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return r.StatusCode() >= http.StatusInternalServerError
		})
}

func (b *baseAd) getRestyClient() *resty.Client {
	if b.client != nil {
		return b.client
	}
	return b.getNewRestyClient()
}

func (b *baseAd) Init(config AdapterConfig) (err error) {
	b.config = config
	b.client = b.getNewRestyClient()
	return
}

func (b *baseAd) formatState(ctx context.Context, state string) (resp api.AuthStateData, err error) {
	err = json.Unmarshal([]byte(state), &resp)
	if err != nil {
		return
	}
	developerInfo, infoErr := b.getDeveloperInfo(ctx, resp.DeveloperId)
	if infoErr != nil {
		err = infoErr
		return
	}
	resp.DeveloperInfo = developerInfo
	return
}

func (b *baseAd) getDeveloperInfo(ctx context.Context, id int64) (resp advertising.DimAdvertisingDeveloperConfigModel, err error) {
	err = resp.Take(ctx, "*", "id = ?", id)
	return
}
