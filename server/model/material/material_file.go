package material

import (
	"github.com/cngamesdk/go-core/model/sql/material"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type OdsMaterialFileLogModel struct {
	material.OdsMaterialFileLogModel
}

func NewOdsMaterialFileLogModel() *OdsMaterialFileLogModel {
	model := &OdsMaterialFileLogModel{}
	model.OdsMaterialFileLogModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
