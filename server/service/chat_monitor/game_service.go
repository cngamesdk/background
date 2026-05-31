package chat_monitor

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chat_monitor"
	"gorm.io/gorm"
)

type GameService struct{}

func (s *GameService) CreateGame(game chat_monitor.Game) error {
	return global.GVA_DB.Create(&game).Error
}

func (s *GameService) UpdateGame(game chat_monitor.Game) error {
	return global.GVA_DB.Model(&chat_monitor.Game{}).Where("id = ?", game.ID).Updates(&game).Error
}

func (s *GameService) DeleteGame(id uint) error {
	return global.GVA_DB.Delete(&chat_monitor.Game{}, id).Error
}

func (s *GameService) GetGameList(search chat_monitor.GameSearch) ([]chat_monitor.Game, int64, error) {
	var list []chat_monitor.Game
	var total int64

	db := global.GVA_DB.Model(&chat_monitor.Game{})
	if search.Name != "" {
		db = db.Where("name LIKE ?", "%"+search.Name+"%")
	}
	if search.Status != nil {
		db = db.Where("status = ?", *search.Status)
	}
	db.Count(&total)

	err := db.Scopes(paginate(search.Page, search.PageSize)).Order("id DESC").Find(&list).Error
	return list, total, err
}

func (s *GameService) GetGameByAppID(appID string) (chat_monitor.Game, error) {
	var game chat_monitor.Game
	err := global.GVA_DB.Where("app_id = ?", appID).First(&game).Error
	return game, err
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
