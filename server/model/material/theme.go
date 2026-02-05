package material

import (
	"github.com/cngamesdk/go-core/model/sql/material"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimMaterialThemeModel struct {
	material.DimMaterialThemeModel
}

func NewDimMaterialThemeModel() *DimMaterialThemeModel {
	model := &DimMaterialThemeModel{}
	model.DimMaterialThemeModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
