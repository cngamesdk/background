package api

import (
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data_report"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	aliasPaymentStatus = "payment"
)

var (
	paymentStatusFieldsMap = map[string]string{
		"platform_id":  "platform.id as platform_id, platform.platform_name as platform_name",
		"root_game_id": "root_game.id as root_game_id, root_game.game_name as root_game_name",
		"main_game_id": "main_game.id as main_game_id, main_game.game_name as main_game_name",
		"game_id":      "game.id as game_id, game.game_name as game_name",
		"agent_id":     "agent.id as agent_id, agent.agent_name as agent_name",
		"site_id":      "site.id as site_id, site.site_name as site_name",
		"active_days":  aliasPaymentStatus + ".active_days AS active_days",
		"pay_count":    "SUM( " + aliasPaymentStatus + ".pay_count ) AS pay_count",
		"active_user":  "SUM( " + aliasPaymentStatus + ".active_user ) AS active_user",
		"pay_amount":   "SUM( " + aliasPaymentStatus + ".pay_amount ) AS pay_amount",
	}
	paymentStatusWheresMap = map[string]string{
		"platform_id":  aliasPaymentStatus + ".platform_id",
		"root_game_id": "root_game.id",
		"main_game_id": "main_game.id",
		"game_id":      "game.id",
		"agent_id":     "agent.id",
		"site_id":      "site.id",
	}
	paymentStatusGroupsMap = map[string]string{
		"platform_id":  "platform.id",
		"root_game_id": "root_game.id",
		"main_game_id": "main_game.id",
		"game_id":      "game.id",
		"agent_id":     "agent.id",
		"site_id":      "site.id",
		"active_days":  aliasPaymentStatus + ".active_days",
	}
	paymentStatusJoinsMap = map[string]func(tx *gorm.DB){
		"platform_id": func(tx *gorm.DB) {
			tx.Joins("join dim_platform as platform on platform.id = " + aliasPaymentStatus + ".platform_id")
		},
		"agent_id": func(tx *gorm.DB) {
			tx.Joins("join dim_agent as agent on agent.platform_id = " + aliasPaymentStatus + ".platform_id and agent.id = " + aliasPaymentStatus + ".agent_id")
		},
		"site_id": func(tx *gorm.DB) {
			tx.Joins("join dim_site as site on site.platform_id = " + aliasPaymentStatus + ".platform_id and site.id = " + aliasPaymentStatus + ".site_id")
		},
	}
	paymentStatusOrdersMap = map[string]string{
		"active_days": "active_days",
		"stat_date":   "stat_date",
	}
)

type PaymentStatusListReq struct {
	BaseDataReport
	request.PageInfo
	ActiveDays []int `json:"-" form:"-"`
}

func (receiver *PaymentStatusListReq) Format() {
	var tmpIndicators []string
	tmpIndicators = append(tmpIndicators, "pay_count", "active_user", "pay_amount")
	if len(receiver.Indicators) > 0 {
		for _, indicator := range receiver.Indicators {
			isDel := false
			if validator.IsNumberStr(indicator) {
				isDel = true
				receiver.ActiveDays = append(receiver.ActiveDays, cast.ToInt(indicator))
			} else {
				pattern := `\d+-\d+`
				if regexp.MustCompile(pattern).MatchString(indicator) {
					isDel = true
					dayRange := strings.Split(indicator, "-")
					start := cast.ToInt(dayRange[0])
					end := cast.ToInt(dayRange[1])
					for start <= end {
						receiver.ActiveDays = append(receiver.ActiveDays, start)
						start++
					}
				}
				if !isDel {
					tmpIndicators = append(tmpIndicators, indicator)
				}
			}
		}
		receiver.Indicators = tmpIndicators
		//升序
		sort.Ints(receiver.ActiveDays)
	}

	if receiver.AggregationTime == AggregationTimeDay {
		receiver.Dimensions = append(receiver.Dimensions, "stat_date")
		receiver.Orders = append(receiver.Orders, "stat_date")
		retentionStatusFieldsMap["stat_date"] = "DATE_FORMAT(" + aliasPaymentStatus + ".reg_date, '%Y-%m-%d') as stat_date"
		retentionStatusGroupsMap["stat_date"] = "DATE_FORMAT(" + aliasPaymentStatus + ".reg_date, '%Y-%m-%d')"

	} else if receiver.AggregationTime == AggregationTimeMonth {
		receiver.Dimensions = append(receiver.Dimensions, "stat_date")
		receiver.Orders = append(receiver.Orders, "stat_date")
		retentionStatusFieldsMap["stat_date"] = "DATE_FORMAT(" + aliasPaymentStatus + ".reg_date, '%Y-%m') as stat_date"
		retentionStatusGroupsMap["stat_date"] = "DATE_FORMAT(" + aliasPaymentStatus + ".reg_date, '%Y-%m')"
	}
	receiver.Dimensions = append(receiver.Dimensions, "active_days")
	receiver.Orders = append(receiver.Orders, "active_days")
}

func (receiver *PaymentStatusListReq) BuildDb(tx *gorm.DB) (resp *gorm.DB, err error) {
	tmpDb := tx
	if receiver.StatisticalCaliber == StatisticalCaliberRootGameBack30 {
		tmpDb = tmpDb.Table(data_report.NewDwsDayRootGameBackPayActiveLogModel().TableName() + " as " + aliasPaymentStatus)
	}

	tmpDb.Joins("join dim_game as game on game.platform_id = " + aliasPaymentStatus + ".platform_id and game.id = " + aliasPaymentStatus + ".game_id")
	tmpDb.Joins("join dim_main_game as main_game on main_game.platform_id = game.platform_id and main_game.id = game.main_id")
	tmpDb.Joins("join dim_root_game as root_game on root_game.platform_id = main_game.platform_id and root_game.id = main_game.root_game_id")

	tmpDb.Where(aliasPaymentStatus+".reg_date BETWEEN ? AND ?", receiver.StartTime, receiver.EndTime)
	loginEndDate, _ := datetime.FormatStrToTime(receiver.EndTime, "yyyy-MM-dd")
	maxDay := 0
	if len(receiver.ActiveDays) > 0 {
		maxDay = receiver.ActiveDays[len(receiver.ActiveDays)-1]
	}
	loginEndDateAdd := loginEndDate.Add(time.Duration(maxDay) * 24 * time.Hour)
	tmpDb.Where(aliasPaymentStatus+".login_date BETWEEN ? AND ?", receiver.StartTime, datetime.FormatTimeToStr(loginEndDateAdd, "yyyy-MM-dd"))

	dbBuilder := &DbBuilder{Db: tmpDb, BaseDataReport: receiver.BaseDataReport}
	dbBuilder.
		SetFieldsMap(paymentStatusFieldsMap).
		SetJoinsMap(paymentStatusJoinsMap).
		SetWheresMap(paymentStatusWheresMap).
		SetGroupsMap(paymentStatusGroupsMap).
		SetOrdersMap(paymentStatusOrdersMap)

	resp = dbBuilder.Build()
	return
}

type PaymentStatusListResp struct {
	BaseResp
	Reg         int `json:"reg"`
	Cost        int `json:"cost"`
	ActiveDays  int `json:"active_days"`
	ActiveCount int `json:"active_count"`
}

type PaymentStatusListRespFormat struct {
	BaseResp
	Reg           int                     `json:"reg"`
	Cost          int                     `json:"cost"`
	NDayContainer []PaymentStatusNDayData `json:"n_day_container"`
}

type PaymentStatusNDayData struct {
	NDay                          int     `json:"n_day"`
	RoiRateStr                    string  `json:"roi_rate_str"`                      //回本率。如：60%;用户在首日新增后的累计付费总额/广告消耗成本
	Ltv                           float64 `json:"ltv"`                               //ltv。如：9.8;用户在首日新增后的累计付费总额/新增用户总户数
	CumulativePayments            int     `json:"cumulative_payments"`               //累计付费。用户在首日新增后的累计付费总额
	DailyRoiStr                   string  `json:"daily_roi_str"`                     //每日回本率。用户在首日新增后，在接下来的后推第N天的当日付费总额/广告消耗成本
	DailyLtv                      float64 `json:"daily_ltv"`                         //每日LTV。用户在首日新增后，在接下来的后推第N天的当日付费总额/新增用用数
	DailyPayment                  int     `json:"daily_payment"`                     //每日付费。用户在首日新增后，在接下来的后推第N天的当日付费总额
	DailyPaymentUsers             int     `json:"daily_payment_users"`               //每日付费人数。用户在首日新增后，在接下来的后推第N天里当日有产生付费行为的用户数量
	DailyArpu                     float64 `json:"daily_arpu"`                        //每日ARPU。户在首日新增后，在接下来的后推第N天的当日付费总额/用户在首日新增后，当天有登录的用户数
	DailyPaymentArpu              float64 `json:"daily_payment_arpu"`                //每日付费ARPU。用户在首日新增后，在接下来的后推第N天的当日付费总额/用户在首日新增后，当天有付费的用户数量
	CumulativePaymentUsers        int     `json:"cumulative_payment_users"`          //累计付费人数。用户在首日新增后，在接下来的后推第N天的累计付费用户总数
	CumulativePaymentUsersRateStr string  `json:"cumulative_payment_users_rate_str"` //累计付费率。用户在首日新增后，在接下来的后推第N天的累计付费用户总数/新增用户总数
	DailyPaymentUsersRateStr      string  `json:"daily_payment_users_rate_str"`      //每日付费率。用户在首日新增后，在接下来的后推第N天当天有付费的用户数量/用户在首日新增后，当天有登录的用户数量
	CumulativePaymentArpu         float64 `json:"cumulative_payment_arpu"`           //累计付费ARPU。用户在首日新增后的累计付费总额/付费用户总数
	CostOfPayment                 float64 `json:"cost_of_payment"`                   //付费成本。广告消耗成本/用户在首日新增后，在接下来的后推第N天的累计付费用户总数
	DailyPaymentFrequency         int     `json:"daily_payment_frequency"`           //每日付费次数。用户在首日新增后，在接下来的后推第N天里当日进行充值的总次数
	CumulativePaymentFrequency    int     `json:"cumulative_payment_frequency"`      //累计付费次数。用户在首日新增后，在接下来的后推第N天的累计充值次数
}
