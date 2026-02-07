package material

import (
	"github.com/cngamesdk/go-core/model/sql/material"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type OdsMaterialLogModel struct {
	material.OdsMaterialLogModel
}

func NewOdsMaterialLogModel() *OdsMaterialLogModel {
	model := &OdsMaterialLogModel{}
	model.OdsMaterialLogModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
