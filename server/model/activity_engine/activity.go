package activity_engine

import "time"

type OdsActivityConfig struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	PlatformID        int64     `json:"platform_id" gorm:"index:ix_platform_game;not null;default:0;comment:平台ID"`
	GameID            int64     `json:"game_id" gorm:"index:ix_platform_game;not null;default:0;comment:游戏ID"`
	ActivityName      string    `json:"activity_name" gorm:"size:200;not null;default:'';comment:活动名称"`
	ActivityType      string    `json:"activity_type" gorm:"size:50;index:ix_type;not null;default:'';comment:活动类型编码"`
	Description       string    `json:"description" gorm:"type:text;comment:活动描述"`
	StartTime         string    `json:"start_time" gorm:"size:30;not null;default:'';comment:开始时间"`
	EndTime           string    `json:"end_time" gorm:"size:30;not null;default:'';comment:结束时间"`
	TriggerConfig     string    `json:"trigger_config" gorm:"type:json;not null;comment:触发条件配置"`
	CalculationConfig string    `json:"calculation_config" gorm:"type:json;not null;comment:计算逻辑配置"`
	RewardConfig      string    `json:"reward_config" gorm:"type:json;not null;comment:奖励策略配置"`
	ConstraintConfig  string    `json:"constraint_config" gorm:"type:json;not null;comment:约束规则配置"`
	Priority          int       `json:"priority" gorm:"not null;default:0;comment:优先级"`
	Version           int       `json:"version" gorm:"not null;default:1;comment:配置版本号"`
	Status            string    `json:"status" gorm:"size:50;index:ix_status_time;not null;default:not-started;comment:状态"`
	GrayscaleRatio    int       `json:"grayscale_ratio" gorm:"not null;default:100;comment:灰度比例(1-100)"`
	CreatedBy         string    `json:"created_by" gorm:"size:100;not null;default:'';comment:创建人"`
	UpdatedBy         string    `json:"updated_by" gorm:"size:100;not null;default:'';comment:最后修改人"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (OdsActivityConfig) TableName() string { return "ods_activity_config" }

type OdsActivityTemplate struct {
	ID                int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TemplateName      string    `json:"template_name" gorm:"size:200;not null;default:'';comment:模板名称"`
	ActivityType      string    `json:"activity_type" gorm:"size:50;not null;default:'';comment:活动类型"`
	TriggerConfig     string    `json:"trigger_config" gorm:"type:json;not null;comment:触发条件模板"`
	CalculationConfig string    `json:"calculation_config" gorm:"type:json;not null;comment:计算逻辑模板"`
	RewardConfig      string    `json:"reward_config" gorm:"type:json;not null;comment:奖励策略模板"`
	ConstraintConfig  string    `json:"constraint_config" gorm:"type:json;not null;comment:约束规则模板"`
	Description       string    `json:"description" gorm:"type:text;comment:模板说明"`
	Status            string    `json:"status" gorm:"size:50;not null;default:normal"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (OdsActivityTemplate) TableName() string { return "ods_activity_template" }

type OdsActivityRewardItem struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	PlatformID int64     `json:"platform_id" gorm:"uniqueIndex:ix_platform_code;not null;default:0;comment:平台ID"`
	ItemCode   string    `json:"item_code" gorm:"size:100;uniqueIndex:ix_platform_code;not null;default:'';comment:道具编码"`
	ItemName   string    `json:"item_name" gorm:"size:200;not null;default:'';comment:道具名称"`
	ItemType   string    `json:"item_type" gorm:"size:50;not null;default:'';comment:道具类型"`
	Icon       string    `json:"icon" gorm:"size:512;not null;default:'';comment:图标"`
	Status     string    `json:"status" gorm:"size:50;not null;default:normal"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (OdsActivityRewardItem) TableName() string { return "ods_activity_reward_item" }
