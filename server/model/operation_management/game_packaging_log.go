package operation_management

import (
	"github.com/cngamesdk/go-core/model/sql/log"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type OdsGamePackagingLogModel struct {
	log.OdsGamePackagingLogModel
}

func NewOdsGamePackagingLogModel() *OdsGamePackagingLogModel {
	model := &OdsGamePackagingLogModel{}
	model.OdsGamePackagingLogModel.Db = func() *gorm.DB {
		return global.GVA_DB
	}
	return model
}
