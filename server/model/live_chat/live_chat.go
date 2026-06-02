package live_chat

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ---- Product ----

type Product struct {
	global.GVA_MODEL
	ProductCode    string `json:"product_code" gorm:"column:product_code;type:varchar(64);uniqueIndex;not null;comment:产品编码"`
	Name           string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:产品名称"`
	Logo           string `json:"logo" gorm:"column:logo;type:varchar(512);comment:Logo URL"`
	WelcomeTitle   string `json:"welcome_title" gorm:"column:welcome_title;type:varchar(256);comment:欢迎标题"`
	WelcomeMessage string `json:"welcome_message" gorm:"column:welcome_message;type:text;comment:欢迎消息"`
	Status         string `json:"status" gorm:"column:status;type:varchar(16);default:normal;comment:状态 normal/disabled"`
}

func (Product) TableName() string { return "lc_product" }

// ---- FAQ ----

type Faq struct {
	global.GVA_MODEL
	ProductID  int64  `json:"product_id" gorm:"column:product_id;type:bigint;index;not null;comment:产品ID"`
	Category   string `json:"category" gorm:"column:category;type:varchar(64);comment:分类"`
	Question   string `json:"question" gorm:"column:question;type:varchar(512);not null;comment:问题"`
	Answer     string `json:"answer" gorm:"column:answer;type:text;not null;comment:答案"`
	Keywords   string `json:"keywords" gorm:"column:keywords;type:varchar(512);comment:关键词"`
	Priority   int    `json:"priority" gorm:"column:priority;type:int;default:0;comment:优先级"`
	MatchCount int64  `json:"match_count" gorm:"column:match_count;type:bigint;default:0;comment:匹配次数"`
	Status     string `json:"status" gorm:"column:status;type:varchar(16);default:normal;comment:状态"`
}

func (Faq) TableName() string { return "lc_faq" }

// ---- Agent ----

type Agent struct {
	global.GVA_MODEL
	UserID          int64      `json:"user_id" gorm:"column:user_id;type:bigint;index;not null;comment:系统用户ID"`
	ProductID       int64      `json:"product_id" gorm:"column:product_id;type:bigint;index;not null;comment:产品ID"`
	Status          string     `json:"status" gorm:"column:status;type:varchar(16);default:offline;comment:状态"`
	MaxConcurrent   int        `json:"max_concurrent" gorm:"column:max_concurrent;type:int;default:5;comment:最大并发"`
	CurrentSessions int        `json:"current_sessions" gorm:"column:current_sessions;type:int;default:0;comment:当前会话数"`
	TotalServed     int64      `json:"total_served" gorm:"column:total_served;type:bigint;default:0;comment:总服务次数"`
	LastOnlineAt    *time.Time `json:"last_online_at" gorm:"column:last_online_at;type:datetime(0);comment:最后上线时间"`
}

func (Agent) TableName() string { return "lc_agent" }

// ---- ChatSession ----

type ChatSession struct {
	global.GVA_MODEL
	ProductID int64  `json:"product_id" gorm:"column:product_id;type:bigint;index;not null;comment:产品ID"`
	UserID    string `json:"user_id" gorm:"column:user_id;type:varchar(128);index;not null;comment:用户ID"`
	UserName  string `json:"user_name" gorm:"column:user_name;type:varchar(128);comment:用户名称"`
	AgentID   *int64 `json:"agent_id" gorm:"column:agent_id;type:bigint;index;comment:客服ID"`
	Status    string `json:"status" gorm:"column:status;type:varchar(16);default:waiting;index;comment:状态"`
	Source    string `json:"source" gorm:"column:source;type:varchar(32);default:h5;comment:来源"`
}

func (ChatSession) TableName() string { return "lc_chat_session" }

// ---- ChatMessage ----

type ChatMessage struct {
	ID            int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	SessionID     int64     `json:"session_id" gorm:"column:session_id;type:bigint;index;not null;comment:会话ID"`
	SenderType    string    `json:"sender_type" gorm:"column:sender_type;type:varchar(16);not null;comment:发送者类型"`
	SenderID      string    `json:"sender_id" gorm:"column:sender_id;type:varchar(128);comment:发送者ID"`
	SenderName    string    `json:"sender_name" gorm:"column:sender_name;type:varchar(128);comment:发送者名称"`
	Content       string    `json:"content" gorm:"column:content;type:text;comment:内容"`
	MsgType       string    `json:"msg_type" gorm:"column:msg_type;type:varchar(16);default:text;comment:消息类型"`
	AttachmentURL string    `json:"attachment_url" gorm:"column:attachment_url;type:varchar(512);comment:附件URL"`
	IsFaqReply    bool      `json:"is_faq_reply" gorm:"column:is_faq_reply;type:tinyint(1);default:0;comment:FAQ回复"`
	FaqID         *int64    `json:"faq_id" gorm:"column:faq_id;type:bigint;comment:FAQ ID"`
	CreatedAt     time.Time `json:"created_at" gorm:"type:datetime(0);autoCreateTime"`
}

func (ChatMessage) TableName() string { return "lc_chat_message" }

// ---- DailyReport ----

type DailyReport struct {
	global.GVA_MODEL
	ProductID       int64  `json:"product_id" gorm:"column:product_id;type:bigint;index;not null;comment:产品ID"`
	ReportDate      string `json:"report_date" gorm:"column:report_date;type:date;index;not null;comment:日期"`
	TotalSessions   int    `json:"total_sessions" gorm:"column:total_sessions;type:int;default:0;comment:总会话数"`
	FaqResolved     int    `json:"faq_resolved" gorm:"column:faq_resolved;type:int;default:0;comment:FAQ解决数"`
	AgentResolved   int    `json:"agent_resolved" gorm:"column:agent_resolved;type:int;default:0;comment:人工解决数"`
	AvgResponseTime int    `json:"avg_response_time" gorm:"column:avg_response_time;type:int;default:0;comment:平均响应时间"`
	AvgSessionTime  int    `json:"avg_session_time" gorm:"column:avg_session_time;type:int;default:0;comment:平均会话时长"`
}

func (DailyReport) TableName() string { return "lc_daily_report" }
