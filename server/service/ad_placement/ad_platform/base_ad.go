package ad_platform

import (
	"context"
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising/api"
)

type baseAd struct {
}

func (b *baseAd) formatState(state string) (resp api.AuthStateData, err error) {
	err = json.Unmarshal([]byte(state), &resp)
	return
}

func (b *baseAd) getDeveloperInfo(ctx context.Context, id int64) (resp advertising.DimAdvertisingDeveloperConfigModel, err error) {
	err = resp.Take(ctx, "*", "id = ?", id)
	return
}
