package live_chat

// ---- Product ----

type ProductSearch struct {
	Name   string `json:"name" form:"name"`
	Status string `json:"status" form:"status"`
}

// ---- FAQ ----

type FaqSearch struct {
	ProductID int64  `json:"product_id" form:"product_id"`
	Category  string `json:"category" form:"category"`
	Keyword   string `json:"keyword" form:"keyword"`
	Status    string `json:"status" form:"status"`
}

type FaqImportReq struct {
	ProductID int64 `json:"product_id"`
	Items     []Faq `json:"items"`
}

// ---- Agent ----

type AgentOnlineReq struct {
	ProductID int64 `json:"product_id"`
}

type AgentUpdateReq struct {
	ID            int64 `json:"id"`
	MaxConcurrent int   `json:"max_concurrent"`
}

// ---- Chat ----

type ChatSessionSearch struct {
	ProductID int64  `json:"product_id" form:"product_id"`
	AgentID   *int64 `json:"agent_id" form:"agent_id"`
	Status    string `json:"status" form:"status"`
	UserID    string `json:"user_id" form:"user_id"`
}

type AssignSessionReq struct {
	SessionID int64 `json:"session_id"`
	AgentID   int64 `json:"agent_id"`
}

type AgentReplyReq struct {
	SessionID int64  `json:"session_id"`
	Content   string `json:"content"`
	MsgType   string `json:"msg_type"`
}

type TransferSessionReq struct {
	SessionID int64 `json:"session_id"`
	ToAgentID int64 `json:"to_agent_id"`
}

// ---- Report ----

type ReportSearch struct {
	ProductID int64  `json:"product_id" form:"product_id"`
	StartDate string `json:"start_date" form:"start_date"`
	EndDate   string `json:"end_date" form:"end_date"`
	AgentID   int64  `json:"agent_id" form:"agent_id"`
}

type ReportOverview struct {
	TotalSessions   int     `json:"total_sessions"`
	FaqResolved     int     `json:"faq_resolved"`
	AgentResolved   int     `json:"agent_resolved"`
	FaqRate         float64 `json:"faq_rate"`
	OnlineAgents    int     `json:"online_agents"`
	WaitingSessions int     `json:"waiting_sessions"`
	AvgResponseTime float64 `json:"avg_response_time"`
}

type ReportTrendItem struct {
	Date          string `json:"date"`
	TotalSessions int    `json:"total_sessions"`
	FaqResolved   int    `json:"faq_resolved"`
	AgentResolved int    `json:"agent_resolved"`
}
