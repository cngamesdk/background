package chat_monitor

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat_monitor"
)

type ChatService struct{}

func (s *ChatService) GetChatHistory(search chat_monitor.ChatMessageSearch) ([]chat_monitor.ChatMessage, int64, error) {
	var list []chat_monitor.ChatMessage
	var total int64

	db := global.GVA_DB.Model(&chat_monitor.ChatMessage{})
	if search.AppID != "" {
		db = db.Where("app_id = ?", search.AppID)
	}
	if search.Channel != "" {
		db = db.Where("channel = ?", search.Channel)
	}
	if search.SenderUID != "" {
		db = db.Where("sender_uid = ?", search.SenderUID)
	}
	if search.Keyword != "" {
		db = db.Where("content LIKE ?", "%"+search.Keyword+"%")
	}
	if search.StartTime != "" {
		db = db.Where("sent_at >= ?", search.StartTime)
	}
	if search.EndTime != "" {
		db = db.Where("sent_at <= ?", search.EndTime)
	}
	if search.OnlySensitive {
		db = db.Where("is_sensitive = 1")
	}
	db.Count(&total)

	err := db.Scopes(paginate(search.Page, search.PageSize)).Order("sent_at DESC").Find(&list).Error
	return list, total, err
}

type SensitiveService struct{}

func (s *SensitiveService) CreateSensitiveWord(word chat_monitor.SensitiveWord) error {
	return global.GVA_DB.Create(&word).Error
}

func (s *SensitiveService) ImportSensitiveWords(req chat_monitor.SensitiveImportReq) (int64, error) {
	words := make([]chat_monitor.SensitiveWord, 0, len(req.Words))
	for _, w := range req.Words {
		level := req.Level
		if level == 0 {
			level = 2
		}
		category := req.Category
		if category == "" {
			category = "default"
		}
		words = append(words, chat_monitor.SensitiveWord{
			Word:     w,
			Category: category,
			Level:    &level,
			AppID:    req.AppID,
			Status:   intPtr(1),
		})
	}
	result := global.GVA_DB.CreateInBatches(words, 100)
	return result.RowsAffected, result.Error
}

func (s *SensitiveService) UpdateSensitiveWord(word chat_monitor.SensitiveWord) error {
	return global.GVA_DB.Model(&chat_monitor.SensitiveWord{}).Where("id = ?", word.ID).Updates(&word).Error
}

func (s *SensitiveService) DeleteSensitiveWord(id uint) error {
	return global.GVA_DB.Delete(&chat_monitor.SensitiveWord{}, id).Error
}

func (s *SensitiveService) GetSensitiveWordList(search chat_monitor.SensitiveWordSearch) ([]chat_monitor.SensitiveWord, int64, error) {
	var list []chat_monitor.SensitiveWord
	var total int64

	db := global.GVA_DB.Model(&chat_monitor.SensitiveWord{})
	if search.AppID != "" {
		db = db.Where("app_id = ? OR app_id = ''", search.AppID)
	}
	if search.Category != "" {
		db = db.Where("category = ?", search.Category)
	}
	db.Count(&total)

	err := db.Scopes(paginate(search.Page, search.PageSize)).Order("id DESC").Find(&list).Error
	return list, total, err
}

type WhitelistService struct{}

func (s *WhitelistService) CreateWhitelist(word chat_monitor.Whitelist) error {
	return global.GVA_DB.Create(&word).Error
}

func (s *WhitelistService) DeleteWhitelist(id uint) error {
	return global.GVA_DB.Delete(&chat_monitor.Whitelist{}, id).Error
}

func (s *WhitelistService) GetWhitelistList(search chat_monitor.WhitelistSearch) ([]chat_monitor.Whitelist, int64, error) {
	var list []chat_monitor.Whitelist
	var total int64

	db := global.GVA_DB.Model(&chat_monitor.Whitelist{})
	if search.AppID != "" {
		db = db.Where("app_id = ? OR app_id = ''", search.AppID)
	}
	db.Count(&total)

	err := db.Scopes(paginate(search.Page, search.PageSize)).Order("id DESC").Find(&list).Error
	return list, total, err
}
