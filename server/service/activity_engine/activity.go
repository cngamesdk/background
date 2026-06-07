package activity_engine

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine"
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine/api"
)

type ActivityService struct{}

func (s *ActivityService) List(ctx context.Context, req *api.ActivityListReq) (list []activity_engine.OdsActivityConfig, total int64, err error) {
	db := global.GVA_DB.WithContext(ctx).Model(&activity_engine.OdsActivityConfig{})
	if req.ActivityName != "" {
		db = db.Where("activity_name LIKE ?", "%"+req.ActivityName+"%")
	}
	if req.ActivityType != "" {
		db = db.Where("activity_type = ?", req.ActivityType)
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.PlatformID > 0 {
		db = db.Where("platform_id = ?", req.PlatformID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id DESC").Scopes(req.Paginate()).Find(&list).Error
	return
}

func (s *ActivityService) Add(ctx context.Context, req *api.ActivityAddReq) error {
	return global.GVA_DB.WithContext(ctx).Create(&req.OdsActivityConfig).Error
}

func (s *ActivityService) Modify(ctx context.Context, req *api.ActivityModifyReq) error {
	req.Version = req.Version + 1
	return global.GVA_DB.WithContext(ctx).Save(&req.OdsActivityConfig).Error
}

func (s *ActivityService) Detail(ctx context.Context, id int64) (activity_engine.OdsActivityConfig, error) {
	var result activity_engine.OdsActivityConfig
	err := global.GVA_DB.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return result, err
}

func (s *ActivityService) Publish(ctx context.Context, id int64) error {
	err := global.GVA_DB.WithContext(ctx).Model(&activity_engine.OdsActivityConfig{}).
		Where("id = ?", id).Update("status", "normal").Error
	if err != nil {
		return err
	}
	// 发布配置变更通知到Redis
	s.notifyReload(ctx, id, "update")
	return nil
}

func (s *ActivityService) Offline(ctx context.Context, id int64) error {
	err := global.GVA_DB.WithContext(ctx).Model(&activity_engine.OdsActivityConfig{}).
		Where("id = ?", id).Update("status", "remove").Error
	if err != nil {
		return err
	}
	s.notifyReload(ctx, id, "delete")
	return nil
}

func (s *ActivityService) UpdateGrayscale(ctx context.Context, id int64, ratio int) error {
	err := global.GVA_DB.WithContext(ctx).Model(&activity_engine.OdsActivityConfig{}).
		Where("id = ?", id).Update("grayscale_ratio", ratio).Error
	if err != nil {
		return err
	}
	s.notifyReload(ctx, id, "update")
	return nil
}

func (s *ActivityService) notifyReload(ctx context.Context, activityID int64, action string) {
	if global.GVA_REDIS == nil {
		return
	}
	payload, _ := json.Marshal(map[string]interface{}{
		"action":      action,
		"activity_id": activityID,
	})
	err := global.GVA_REDIS.Publish(ctx, "ae:config:reload", string(payload)).Err()
	if err != nil {
		fmt.Printf("notify reload error: %v\n", err)
	}
}
