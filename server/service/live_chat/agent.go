package live_chat

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"
	"gorm.io/gorm"
)

type AgentService struct{}

func (s *AgentService) GoOnline(userID, productID int64) error {
	var agent live_chat.Agent
	err := global.GVA_DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&agent).Error
	now := time.Now()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			agent = live_chat.Agent{
				UserID:        userID,
				ProductID:     productID,
				Status:        "online",
				MaxConcurrent: 5,
				LastOnlineAt:  &now,
			}
			return global.GVA_DB.Create(&agent).Error
		}
		return err
	}
	return global.GVA_DB.Model(&agent).Updates(map[string]interface{}{
		"status":         "online",
		"last_online_at": now,
	}).Error
}

func (s *AgentService) GoOffline(userID, productID int64) error {
	return global.GVA_DB.Model(&live_chat.Agent{}).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Update("status", "offline").Error
}

func (s *AgentService) Update(req live_chat.AgentUpdateReq) error {
	return global.GVA_DB.Model(&live_chat.Agent{}).Where("id = ?", req.ID).
		Update("max_concurrent", req.MaxConcurrent).Error
}

func (s *AgentService) List(productID int64, page, pageSize int) ([]live_chat.Agent, int64, error) {
	var list []live_chat.Agent
	var total int64
	db := global.GVA_DB.Model(&live_chat.Agent{})
	if productID > 0 {
		db = db.Where("product_id = ?", productID)
	}
	db.Count(&total)
	return list, total, paginate(page, pageSize)(db).Order("id DESC").Find(&list).Error
}

func (s *AgentService) GetStatus(userID, productID int64) (*live_chat.Agent, error) {
	var agent live_chat.Agent
	err := global.GVA_DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&agent).Error
	if err != nil {
		return nil, err
	}
	return &agent, nil
}
