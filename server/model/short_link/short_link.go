package short_link

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type DimShortLink struct {
	Id          int64      `json:"id" gorm:"primaryKey;autoIncrement:false"`
	ShortCode   string     `json:"short_code" gorm:"type:varchar(16);uniqueIndex;not null"`
	OriginalUrl string     `json:"original_url" gorm:"type:text;not null"`
	Domain      string     `json:"domain" gorm:"type:varchar(255);not null"`
	Title       string     `json:"title" gorm:"type:varchar(255)"`
	Status      int8       `json:"status" gorm:"type:tinyint;default:1;not null;index"`
	ExpireAt    *time.Time `json:"expire_at" gorm:"type:datetime"`
	TotalClicks int64      `json:"total_clicks" gorm:"type:bigint;default:0"`
	Creator     string     `json:"creator" gorm:"type:varchar(64)"`
	CreatedAt   time.Time  `json:"created_at" gorm:"type:datetime;autoCreateTime"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"type:datetime;autoUpdateTime"`
}

func (DimShortLink) TableName() string {
	return "dim_short_link"
}

func NewDimShortLink() *DimShortLink {
	return &DimShortLink{}
}

func (d *DimShortLink) Db() *gorm.DB {
	return global.GVA_DB
}

type OdsClickLog struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	ShortCode string    `json:"short_code" gorm:"type:varchar(16);index;not null"`
	Ip        string    `json:"ip" gorm:"type:varchar(64)"`
	UserAgent string    `json:"user_agent" gorm:"type:varchar(512)"`
	Referer   string    `json:"referer" gorm:"type:varchar(1024)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime;autoCreateTime"`
}

func (OdsClickLog) TableName() string {
	return "ods_click_log"
}
