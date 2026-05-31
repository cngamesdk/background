package chat_monitor

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// Game 游戏接入配置
type Game struct {
	global.GVA_MODEL
	AppID       string `json:"appId" gorm:"column:app_id;type:varchar(64);uniqueIndex;not null;comment:游戏AppID"`
	AppSecret   string `json:"appSecret" gorm:"column:app_secret;type:varchar(128);not null;comment:签名密钥"`
	Name        string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:游戏名称"`
	CallbackURL string `json:"callbackUrl" gorm:"column:callback_url;type:varchar(512);comment:封禁回调地址"`
	Status      *int   `json:"status" gorm:"column:status;default:1;comment:1启用 0禁用"`
	ConfigJSON  string `json:"configJson" gorm:"column:config_json;type:json;comment:扩展配置"`
}

func (Game) TableName() string { return "cm_game" }

func (g *Game) BeforeSave(tx *gorm.DB) error {
	if g.ConfigJSON == "" {
		g.ConfigJSON = "{}"
	}
	return nil
}

func (g *Game) BeforeCreate(tx *gorm.DB) error {
	return g.BeforeSave(tx)
}

// ChatMessage 聊天消息
type ChatMessage struct {
	global.GVA_MODEL
	AppID       string `json:"appId" gorm:"column:app_id;type:varchar(64);not null;index:idx_app_time;comment:游戏ID"`
	Channel     string `json:"channel" gorm:"column:channel;type:varchar(64);comment:频道"`
	SenderUID   string `json:"senderUid" gorm:"column:sender_uid;type:varchar(128);not null;index:idx_sender;comment:发送者账号"`
	SenderName  string `json:"senderName" gorm:"column:sender_name;type:varchar(128);comment:发送者昵称"`
	SenderIP    string `json:"senderIp" gorm:"column:sender_ip;type:varchar(45);index:idx_ip;comment:发送者IP"`
	RoleID      string `json:"roleId" gorm:"column:role_id;type:varchar(128);comment:角色ID"`
	Content     string `json:"content" gorm:"column:content;type:text;not null;comment:消息内容"`
	MsgType     *int   `json:"msgType" gorm:"column:msg_type;default:1;comment:1文本 2图片 3语音"`
	IsSensitive *int   `json:"isSensitive" gorm:"column:is_sensitive;default:0;index:idx_sensitive;comment:是否敏感"`
	HitWords    string `json:"hitWords" gorm:"column:hit_words;type:varchar(512);comment:命中敏感词"`
	RiskLevel   *int   `json:"riskLevel" gorm:"column:risk_level;default:0;comment:风险等级"`
	ExtraJSON   string `json:"extraJson" gorm:"column:extra_json;type:json;comment:扩展字段"`
	SentAt      string `json:"sentAt" gorm:"column:sent_at;type:datetime;not null;index:idx_app_time;comment:发送时间"`
}

func (ChatMessage) TableName() string { return "cm_chat_message" }

// SensitiveWord 敏感词
type SensitiveWord struct {
	global.GVA_MODEL
	Word     string `json:"word" gorm:"column:word;type:varchar(256);not null;comment:敏感词"`
	Category string `json:"category" gorm:"column:category;type:varchar(64);default:default;comment:分类"`
	Level    *int   `json:"level" gorm:"column:level;default:2;comment:风险等级 1低 2中 3高"`
	AppID    string `json:"appId" gorm:"column:app_id;type:varchar(64);index:idx_app_status;comment:空=全局"`
	IsRegex  *int   `json:"isRegex" gorm:"column:is_regex;default:0;comment:是否正则"`
	Status   *int   `json:"status" gorm:"column:status;default:1;index:idx_app_status;comment:1启用 0禁用"`
}

func (SensitiveWord) TableName() string { return "cm_sensitive_word" }

// Whitelist 白名单
type Whitelist struct {
	global.GVA_MODEL
	Word  string `json:"word" gorm:"column:word;type:varchar(256);not null;comment:白名单词"`
	AppID string `json:"appId" gorm:"column:app_id;type:varchar(64);comment:空=全局"`
}

func (Whitelist) TableName() string { return "cm_whitelist" }

// BanRecord 封禁记录
type BanRecord struct {
	global.GVA_MODEL
	AppID          string `json:"appId" gorm:"column:app_id;type:varchar(64);not null;comment:游戏ID"`
	BanType        *int   `json:"banType" gorm:"column:ban_type;not null;comment:1账号 2角色 3IP"`
	Target         string `json:"target" gorm:"column:target;type:varchar(128);not null;index:idx_target;comment:封禁目标"`
	Reason         string `json:"reason" gorm:"column:reason;type:varchar(512);comment:封禁原因"`
	Duration       *int   `json:"duration" gorm:"column:duration;default:0;comment:时长秒 0永久"`
	StartAt        string `json:"startAt" gorm:"column:start_at;type:datetime;not null;comment:开始时间"`
	ExpireAt       string `json:"expireAt" gorm:"column:expire_at;type:datetime;comment:过期时间"`
	Status         *int   `json:"status" gorm:"column:status;default:1;index:idx_target;comment:1生效 2已解封 3已过期"`
	OperatorID     *uint  `json:"operatorId" gorm:"column:operator_id;comment:操作人ID"`
	CallbackStatus *int   `json:"callbackStatus" gorm:"column:callback_status;default:0;comment:回调状态"`
}

func (BanRecord) TableName() string { return "cm_ban_record" }

// DailyStats 每日统计
type DailyStats struct {
	global.GVA_MODEL
	AppID          string `json:"appId" gorm:"column:app_id;type:varchar(64);not null;uniqueIndex:uk_app_date;comment:游戏ID"`
	StatDate       string `json:"statDate" gorm:"column:stat_date;type:date;not null;uniqueIndex:uk_app_date;comment:统计日期"`
	TotalMessages  int64  `json:"totalMessages" gorm:"column:total_messages;default:0;comment:总消息数"`
	SensitiveCount int64  `json:"sensitiveCount" gorm:"column:sensitive_count;default:0;comment:敏感消息数"`
	ActiveUsers    int64  `json:"activeUsers" gorm:"column:active_users;default:0;comment:活跃用户数"`
	BanCount       int    `json:"banCount" gorm:"column:ban_count;default:0;comment:封禁数"`
}

func (DailyStats) TableName() string { return "cm_daily_stats" }
