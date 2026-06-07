package activity_engine

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine"
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine/api"
)

type RewardItemService struct{}

func (s *RewardItemService) Search(ctx context.Context, req *api.RewardItemSearchReq) (list []activity_engine.OdsActivityRewardItem, total int64, err error) {
	db := global.GVA_DB.WithContext(ctx).Model(&activity_engine.OdsActivityRewardItem{}).Where("status = ?", "normal")
	if req.PlatformID > 0 {
		db = db.Where("platform_id = ?", req.PlatformID)
	}
	if req.Keyword != "" {
		db = db.Where("item_name LIKE ? OR item_code LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id DESC").Scopes(req.Paginate()).Find(&list).Error
	return
}
