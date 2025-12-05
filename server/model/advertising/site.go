package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimSiteModel struct {
	common.DimSiteModel
}

func NewDimSiteModel() *DimSiteModel {
	model := &DimSiteModel{}
	model.DimSiteModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}