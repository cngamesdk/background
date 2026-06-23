package model_transfer

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
	"time"
)

// AiToken Token管理表
type AiToken struct {
	ID            int64      `gorm:"primarykey;autoIncrement" json:"id"`
	Token         string     `gorm:"type:varchar(128);not null;uniqueIndex:uk_token" json:"token"`
	Name          string     `gorm:"type:varchar(255);not null;index:idx_name" json:"name"`
	Type          int8       `gorm:"type:tinyint;not null;default:1;comment:Token类型：1-企业 2-个人" json:"type"`
	TokenLimit    int64      `gorm:"type:bigint;not null;default:0;comment:Token数量限制（0=无限制）" json:"tokenLimit"`
	UsedTokens    int64      `gorm:"type:bigint;not null;default:0;comment:已使用Token数量" json:"usedTokens"`
	RequestLimit  int        `gorm:"type:int;not null;default:0;comment:请求频率限制（次/分钟，0=无限制）" json:"requestLimit"`
	ExpireAt      *time.Time `gorm:"type:datetime;comment:过期时间" json:"expireAt"`
	Status        int8       `gorm:"type:tinyint;not null;default:1;index:idx_status;comment:状态：1-启用 2-禁用" json:"status"`
	AllowedModels string     `gorm:"type:json;comment:允许的模型列表" json:"allowedModels"`
	IPWhitelist   string     `gorm:"type:json;comment:IP白名单" json:"ipWhitelist"`
	Creator       string     `gorm:"type:varchar(64)" json:"creator"`
	CreatedAt     time.Time  `gorm:"type:datetime;not null;autoCreateTime" json:"createdAt"`
	UpdatedAt     time.Time  `gorm:"type:datetime;not null;autoUpdateTime" json:"updatedAt"`
}

func (AiToken) TableName() string {
	return "dim_ai_token"
}

// NewAiTokenModel 创建模型实例
func NewAiTokenModel() *AiToken {
	return &AiToken{}
}

// Db 获取数据库连接
func (m *AiToken) Db() *gorm.DB {
	return global.GVA_DB
}
