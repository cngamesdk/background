package advertising

import (
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimAgentModel struct {
	common.DimAgentModel
}

func NewDimAgentModel() *DimAgentModel {
	model := &DimAgentModel{}
	model.DimAgentModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
