package live_chat

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/live_chat"
)

type ReportService struct{}

func (s *ReportService) GetOverview(req live_chat.ReportSearch) (*live_chat.ReportOverview, error) {
	overview := &live_chat.ReportOverview{}

	// Total sessions
	var totalSessions int64
	global.GVA_DB.Model(&live_chat.ChatSession{}).
		Where("created_at BETWEEN ? AND ?", req.StartDate, req.EndDate+" 23:59:59").
		Where("product_id = ?", req.ProductID).Count(&totalSessions)
	overview.TotalSessions = int(totalSessions)

	// FAQ resolved
	var faqResolved int64
	global.GVA_DB.Model(&live_chat.ChatMessage{}).
		Joins("JOIN lc_chat_session ON lc_chat_message.session_id = lc_chat_session.id").
		Where("lc_chat_message.is_faq_reply = 1").
		Where("lc_chat_session.product_id = ?", req.ProductID).
		Where("lc_chat_message.created_at BETWEEN ? AND ?", req.StartDate, req.EndDate+" 23:59:59").
		Distinct("lc_chat_message.session_id").Count(&faqResolved)
	overview.FaqResolved = int(faqResolved)

	// Agent resolved
	var agentResolved int64
	global.GVA_DB.Model(&live_chat.ChatSession{}).
		Where("status = ? AND agent_id IS NOT NULL", "closed").
		Where("product_id = ?", req.ProductID).
		Where("created_at BETWEEN ? AND ?", req.StartDate, req.EndDate+" 23:59:59").
		Count(&agentResolved)
	overview.AgentResolved = int(agentResolved)

	// FAQ rate
	if overview.TotalSessions > 0 {
		overview.FaqRate = float64(overview.FaqResolved) / float64(overview.TotalSessions) * 100
	}

	// Online agents
	var onlineAgents int64
	global.GVA_DB.Model(&live_chat.Agent{}).
		Where("product_id = ? AND status = ?", req.ProductID, "online").Count(&onlineAgents)
	overview.OnlineAgents = int(onlineAgents)

	// Waiting sessions
	var waitingSessions int64
	global.GVA_DB.Model(&live_chat.ChatSession{}).
		Where("product_id = ? AND status = ?", req.ProductID, "waiting").Count(&waitingSessions)
	overview.WaitingSessions = int(waitingSessions)

	return overview, nil
}

func (s *ReportService) GetTrend(req live_chat.ReportSearch) ([]live_chat.ReportTrendItem, error) {
	var items []live_chat.ReportTrendItem

	start, _ := time.Parse("2006-01-02", req.StartDate)
	end, _ := time.Parse("2006-01-02", req.EndDate)

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dateStr := d.Format("2006-01-02")
		item := live_chat.ReportTrendItem{Date: dateStr}

		var ts int64
		global.GVA_DB.Model(&live_chat.ChatSession{}).
			Where("DATE(created_at) = ? AND product_id = ?", dateStr, req.ProductID).
			Count(&ts)
		item.TotalSessions = int(ts)

		items = append(items, item)
	}
	return items, nil
}

func (s *ReportService) AggregateDailyStats(date string) error {
	var products []live_chat.Product
	global.GVA_DB.Where("status = ?", "normal").Find(&products)

	for _, p := range products {
		var totalSessions, faqResolved, agentResolved int64

		global.GVA_DB.Model(&live_chat.ChatSession{}).
			Where("DATE(created_at) = ? AND product_id = ?", date, p.ID).
			Count(&totalSessions)

		global.GVA_DB.Model(&live_chat.ChatMessage{}).
			Where("DATE(created_at) = ? AND is_faq_reply = 1", date).
			Count(&faqResolved)

		global.GVA_DB.Model(&live_chat.ChatSession{}).
			Where("DATE(closed_at) = ? AND product_id = ? AND agent_id IS NOT NULL", date, p.ID).
			Count(&agentResolved)

		report := live_chat.DailyReport{
			ProductID:     int64(p.ID),
			ReportDate:    date,
			TotalSessions: int(totalSessions),
			FaqResolved:   int(faqResolved),
			AgentResolved: int(agentResolved),
		}

		global.GVA_DB.Where("product_id = ? AND report_date = ?", p.ID, date).
			Assign(report).FirstOrCreate(&report)
	}
	return nil
}
