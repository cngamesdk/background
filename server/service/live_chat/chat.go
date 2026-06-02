package live_chat

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ChatService struct{}

func (s *ChatService) GetSessions(search live_chat.ChatSessionSearch, page, pageSize int) ([]live_chat.ChatSession, int64, error) {
	var list []live_chat.ChatSession
	var total int64
	db := global.GVA_DB.Model(&live_chat.ChatSession{})
	if search.ProductID > 0 {
		db = db.Where("product_id = ?", search.ProductID)
	}
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	} else {
		db = db.Where("status != ?", "closed")
	}
	if search.AgentID != nil {
		db = db.Where("agent_id = ?", *search.AgentID)
	}
	if search.UserID != "" {
		db = db.Where("user_id LIKE ?", "%"+search.UserID+"%")
	}
	db.Count(&total)
	return list, total, paginate(page, pageSize)(db).Order("created_at DESC").Find(&list).Error
}

func (s *ChatService) GetSession(id int64) (*live_chat.ChatSession, error) {
	var session live_chat.ChatSession
	err := global.GVA_DB.First(&session, id).Error
	return &session, err
}

func (s *ChatService) GetMessages(sessionID int64, page, pageSize int) ([]live_chat.ChatMessage, int64, error) {
	var list []live_chat.ChatMessage
	var total int64
	db := global.GVA_DB.Model(&live_chat.ChatMessage{}).Where("session_id = ?", sessionID)
	db.Count(&total)
	return list, total, paginate(page, pageSize)(db).Order("created_at ASC").Find(&list).Error
}

func (s *ChatService) AssignAgent(sessionID, agentID int64) error {
	return global.GVA_DB.Model(&live_chat.ChatSession{}).Where("id = ?", sessionID).Updates(map[string]interface{}{
		"agent_id": agentID,
		"status":   "active",
	}).Error
}

func (s *ChatService) AgentReply(sessionID int64, agentID int64, agentName, content, msgType string) error {
	// 不在这里保存消息，直接调用 live-chat 服务
	// live-chat 服务会保存消息并通过 WebSocket 广播
	return s.notifyLiveChatService(sessionID, agentID, agentName, content, msgType)
}

func (s *ChatService) notifyLiveChatService(sessionID int64, agentID int64, agentName, content, msgType string) error {
	// Get live-chat service URL from config or use default
	liveChatURL := global.GVA_CONFIG.System.LiveChatURL
	if liveChatURL == "" {
		liveChatURL = "http://localhost:8890" // default
	}

	// Prepare request
	payload := map[string]interface{}{
		"session_id": sessionID,
		"agent_id":   agentID,
		"agent_name": agentName,
		"content":    content,
		"msg_type":   msgType,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		global.GVA_LOG.Error("Failed to marshal agent reply payload", zap.Error(err))
		return err
	}

	// Send HTTP request
	resp, err := http.Post(
		liveChatURL+"/api/v1/admin/agent-reply",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		global.GVA_LOG.Error("Failed to notify live-chat service", zap.Error(err))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		global.GVA_LOG.Error("Live-chat service returned error", zap.Int("status", resp.StatusCode))
		return err
	}

	return nil
}

func (s *ChatService) CloseSession(sessionID int64) error {
	now := time.Now()
	err := global.GVA_DB.Model(&live_chat.ChatSession{}).Where("id = ?", sessionID).Updates(map[string]interface{}{
		"status":    "closed",
		"closed_at": now,
	}).Error
	if err != nil {
		return err
	}

	// Notify live-chat service via HTTP to broadcast via WebSocket
	go s.notifySessionClosed(sessionID)

	return nil
}

func (s *ChatService) notifySessionClosed(sessionID int64) {
	// Get live-chat service URL from config or use default
	liveChatURL := global.GVA_CONFIG.System.LiveChatURL
	if liveChatURL == "" {
		liveChatURL = "http://localhost:8890" // default
	}

	// Prepare request
	payload := map[string]interface{}{
		"session_id": sessionID,
		"reason":     "客服已关闭会话",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		global.GVA_LOG.Error("Failed to marshal close session payload", zap.Error(err))
		return
	}

	// Send HTTP request
	resp, err := http.Post(
		liveChatURL+"/api/v1/admin/close-session",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		global.GVA_LOG.Error("Failed to notify live-chat service about session close", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		global.GVA_LOG.Error("Live-chat service returned error on close session", zap.Int("status", resp.StatusCode))
	}
}

func (s *ChatService) GetWaitingCount(productID int64) (int64, error) {
	var count int64
	err := global.GVA_DB.Model(&live_chat.ChatSession{}).
		Where("product_id = ? AND status = ?", productID, "waiting").Count(&count).Error
	return count, err
}

func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
