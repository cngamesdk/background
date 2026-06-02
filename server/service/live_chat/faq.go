package live_chat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"
)

type FaqService struct{}

func (s *FaqService) Create(f *live_chat.Faq) error {
	return global.GVA_DB.Create(f).Error
}

func (s *FaqService) Update(f *live_chat.Faq) error {
	return global.GVA_DB.Save(f).Error
}

func (s *FaqService) Delete(id int64) error {
	return global.GVA_DB.Delete(&live_chat.Faq{}, id).Error
}

func (s *FaqService) GetByID(id int64) (*live_chat.Faq, error) {
	var f live_chat.Faq
	err := global.GVA_DB.First(&f, id).Error
	return &f, err
}

func (s *FaqService) List(search live_chat.FaqSearch, page, pageSize int) ([]live_chat.Faq, int64, error) {
	var list []live_chat.Faq
	var total int64
	db := global.GVA_DB.Model(&live_chat.Faq{})
	if search.ProductID > 0 {
		db = db.Where("product_id = ?", search.ProductID)
	}
	if search.Category != "" {
		db = db.Where("category = ?", search.Category)
	}
	if search.Keyword != "" {
		db = db.Where("question LIKE ? OR answer LIKE ? OR keywords LIKE ?",
			"%"+search.Keyword+"%", "%"+search.Keyword+"%", "%"+search.Keyword+"%")
	}
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	}
	db.Count(&total)
	return list, total, paginate(page, pageSize)(db).Order("priority DESC, id DESC").Find(&list).Error
}

func (s *FaqService) Categories(productID int64) ([]string, error) {
	var categories []string
	err := global.GVA_DB.Model(&live_chat.Faq{}).
		Where("product_id = ? AND status = ?", productID, "normal").
		Distinct("category").Pluck("category", &categories).Error
	return categories, err
}
