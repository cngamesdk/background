package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimSiteModel struct {
	common.DimSiteModel
	PlatformName string `json:"platform_name" gorm:"platform_name"`
	AgentName    string `json:"agent_name" gorm:"agent_name"`
}

func NewDimSiteModel() *DimSiteModel {
	model := &DimSiteModel{}
	model.DimSiteModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
