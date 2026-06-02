package live_chat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"
)

type ProductService struct{}

func (s *ProductService) Create(p *live_chat.Product) error {
	return global.GVA_DB.Create(p).Error
}

func (s *ProductService) Update(p *live_chat.Product) error {
	return global.GVA_DB.Save(p).Error
}

func (s *ProductService) Delete(id int64) error {
	return global.GVA_DB.Delete(&live_chat.Product{}, id).Error
}

func (s *ProductService) GetByID(id int64) (*live_chat.Product, error) {
	var p live_chat.Product
	err := global.GVA_DB.First(&p, id).Error
	return &p, err
}

func (s *ProductService) List(search live_chat.ProductSearch, page, pageSize int) ([]live_chat.Product, int64, error) {
	var list []live_chat.Product
	var total int64
	db := global.GVA_DB.Model(&live_chat.Product{})
	if search.Name != "" {
		db = db.Where("name LIKE ?", "%"+search.Name+"%")
	}
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	}
	db.Count(&total)
	return list, total, paginate(page, pageSize)(db).Order("id DESC").Find(&list).Error
}
