package chat_monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat_monitor"
)

type BanService struct{}

var callbackService = CallbackService{}

func (s *BanService) CreateBan(req chat_monitor.BanCreateReq, operatorID uint) error {
	now := time.Now()
	record := chat_monitor.BanRecord{
		AppID:    req.AppID,
		BanType:  &req.BanType,
		Target:   req.Target,
		Reason:   req.Reason,
		Duration: &req.Duration,
		StartAt:  now.Format("2006-01-02 15:04:05"),
		Status:   intPtr(1),
	}
	record.OperatorID = &operatorID

	if req.Duration > 0 {
		expire := now.Add(time.Duration(req.Duration) * time.Second)
		record.ExpireAt = expire.Format("2006-01-02 15:04:05")
	}

	if err := global.GVA_DB.Create(&record).Error; err != nil {
		return err
	}

	// 写入Redis（如果Redis可用）
	if global.GVA_REDIS != nil {
		key := banRedisKey(req.AppID, req.BanType, req.Target)
		if req.Duration > 0 {
			global.GVA_REDIS.Set(context.Background(), key, "1", time.Duration(req.Duration)*time.Second)
		} else {
			global.GVA_REDIS.Set(context.Background(), key, "1", 0)
		}
	}

	// 异步回调通知游戏研发
	go callbackService.NotifyGame(record, "ban")

	return nil
}

func (s *BanService) RevokeBan(id uint) error {
	var record chat_monitor.BanRecord
	if err := global.GVA_DB.First(&record, id).Error; err != nil {
		return err
	}

	if err := global.GVA_DB.Model(&record).Update("status", 2).Error; err != nil {
		return err
	}

	if global.GVA_REDIS != nil {
		key := banRedisKey(record.AppID, *record.BanType, record.Target)
		global.GVA_REDIS.Del(context.Background(), key)
	}

	// 异步回调通知游戏研发
	go callbackService.NotifyGame(record, "unban")

	return nil
}

func (s *BanService) GetBanList(search chat_monitor.BanRecordSearch) ([]chat_monitor.BanRecord, int64, error) {
	var list []chat_monitor.BanRecord
	var total int64

	db := global.GVA_DB.Model(&chat_monitor.BanRecord{})
	if search.AppID != "" {
		db = db.Where("app_id = ?", search.AppID)
	}
	if search.BanType != nil {
		db = db.Where("ban_type = ?", *search.BanType)
	}
	if search.Status != nil {
		db = db.Where("status = ?", *search.Status)
	}
	db.Count(&total)

	err := db.Scopes(paginate(search.Page, search.PageSize)).Order("id DESC").Find(&list).Error
	return list, total, err
}

func banRedisKey(appID string, banType int, target string) string {
	typeStr := "account"
	switch banType {
	case 2:
		typeStr = "role"
	case 3:
		typeStr = "ip"
	}
	return fmt.Sprintf("cm:ban:%s:%s:%s", typeStr, appID, target)
}

func intPtr(v int) *int { return &v }
