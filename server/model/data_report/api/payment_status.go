package api

import (
	"fmt"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/slice"
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
	aliasPaymentStatus     = "payment"
	aliasCumulativePayment = "cumulative_payment"
)

var (
	paymentStatusFieldsMap = map[string]string{
		"platform_id":                      "platform.id as platform_id, platform.platform_name as platform_name",
		"root_game_id":                     "root_game.id as root_game_id, root_game.game_name as root_game_name",
		"main_game_id":                     "main_game.id as main_game_id, main_game.game_name as main_game_name",
		"game_id":                          "game.id as game_id, game.game_name as game_name",
		"agent_id":                         "agent.id as agent_id, agent.agent_name as agent_name",
		"site_id":                          "site.id as site_id, site.site_name as site_name",
		"active_days":                      aliasPaymentStatus + ".active_days AS active_days",
		"reg":                              "SUM( " + aliasOverview + ".reg ) AS reg",
		"cost":                             "SUM( " + aliasOverview + ".cost ) AS cost",
		"pay_count":                        "SUM( " + aliasPaymentStatus + ".pay_count ) AS pay_count",
		"active_user":                      "SUM( " + aliasPaymentStatus + ".active_user ) AS active_user",
		"pay_amount":                       "SUM( " + aliasPaymentStatus + ".pay_amount ) AS pay_amount",
		"join_all":                         aliasOverview + ".*,IFNULL(" + aliasPaymentStatus + ".active_days,0) AS active_days,IFNULL(" + aliasPaymentStatus + ".pay_count,0) AS pay_count,IFNULL(" + aliasPaymentStatus + ".active_user,0) AS active_user,IFNULL(" + aliasPaymentStatus + ".pay_amount,0) AS pay_amount",
		"join_all_cumulative_payment":      "IFNULL(" + aliasCumulativePayment + ".pay_count,0) AS cumulative_pay_count,IFNULL(" + aliasCumulativePayment + ".active_user,0) AS cumulative_active_user,IFNULL(" + aliasCumulativePayment + ".pay_amount,0) AS cumulative_pay_amount",
		StatDateDay + aliasOverview:        "DATE_FORMAT(" + aliasOverview + ".stat_date, '%Y-%m-%d') AS stat_date",
		StatDateMonth + aliasOverview:      "DATE_FORMAT(" + aliasOverview + ".stat_date, '%Y-%m') AS stat_date",
		StatDateDay + aliasPaymentStatus:   "DATE_FORMAT(" + aliasPaymentStatus + ".reg_date, '%Y-%m-%d') AS stat_date",
		StatDateMonth + aliasPaymentStatus: "DATE_FORMAT(" + aliasPaymentStatus + ".reg_date, '%Y-%m') AS stat_date",
	}
	paymentStatusGroupsMap = map[string]string{
		"platform_id":                      "platform.id",
		"root_game_id":                     "root_game.id",
		"main_game_id":                     "main_game.id",
		"game_id":                          "game.id",
		"agent_id":                         "agent.id",
		"site_id":                          "site.id",
		"active_days":                      aliasPaymentStatus + ".active_days",
		StatDateDay + aliasOverview:        "DATE_FORMAT(" + aliasOverview + ".stat_date, '%Y-%m-%d')",
		StatDateMonth + aliasOverview:      "DATE_FORMAT(" + aliasOverview + ".stat_date, '%Y-%m')",
		StatDateDay + aliasPaymentStatus:   "DATE_FORMAT(" + aliasPaymentStatus + ".reg_date, '%Y-%m-%d')",
		StatDateMonth + aliasPaymentStatus: "DATE_FORMAT(" + aliasPaymentStatus + ".reg_date, '%Y-%m')",
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
}

func (receiver *PaymentStatusListReq) getJoinsMap(alias string) map[string]func(tx *gorm.DB) {
	return map[string]func(tx *gorm.DB){
		"platform_id": func(tx *gorm.DB) {
			tx.Joins("join dim_platform as platform on platform.id = " + alias + ".platform_id")
		},
		"agent_id": func(tx *gorm.DB) {
			tx.Joins("join dim_agent as agent on agent.platform_id = " + alias + ".platform_id and agent.id = " + alias + ".agent_id")
		},
		"site_id": func(tx *gorm.DB) {
			tx.Joins("join dim_site as site on site.platform_id = " + alias + ".platform_id and site.id = " + alias + ".site_id")
		},
	}
}

func (receiver *PaymentStatusListReq) getWheresMap(alias string) map[string]string {
	return map[string]string{
		"platform_id":  alias + ".platform_id",
		"root_game_id": "root_game.id",
		"main_game_id": "main_game.id",
		"game_id":      "game.id",
		"agent_id":     "agent.id",
		"site_id":      "site.id",
	}
}

func (receiver *PaymentStatusListReq) BuildOverviewDb(tx *gorm.DB) (resp *gorm.DB) {
	tmpDb := tx
	if receiver.StatisticalCaliber == StatisticalCaliberRootGameBack30 {
		tmpDb = tmpDb.Table(data_report.NewDwsDayRootGameBackOverviewLogModel().TableName() + " as " + aliasOverview)
	}

	if slice.ContainAny(GameRelationDimensions, receiver.Dimensions) {
		tmpDb.Joins("join dim_game as game on game.platform_id = " + aliasOverview + ".platform_id and game.id = " + aliasOverview + ".game_id")
		tmpDb.Joins("join dim_main_game as main_game on main_game.platform_id = game.platform_id and main_game.id = game.main_id")
		tmpDb.Joins("join dim_root_game as root_game on root_game.platform_id = main_game.platform_id and root_game.id = main_game.root_game_id")
	}

	tmpDb.Where(aliasOverview+".stat_date BETWEEN ? AND ?", receiver.StartTime, receiver.EndTime)

	commonReq := receiver.BaseDataReport
	commonReq.Indicators = append(commonReq.Indicators, "reg")
	if receiver.AggregationTime == AggregationTimeDay {
		commonReq.Dimensions = append(commonReq.Dimensions, StatDateDay+aliasOverview)
	} else if receiver.AggregationTime == AggregationTimeMonth {
		commonReq.Dimensions = append(commonReq.Dimensions, StatDateMonth+aliasOverview)
	}
	dbBuilder := &DbBuilder{Db: tmpDb, BaseDataReport: commonReq}
	dbBuilder.
		SetFieldsMap(paymentStatusFieldsMap).
		SetJoinsMap(receiver.getJoinsMap(aliasOverview)).
		SetWheresMap(receiver.getWheresMap(aliasOverview)).
		SetGroupsMap(paymentStatusGroupsMap)
	resp = dbBuilder.Build()
	return
}

// BuildCumulativePayment 构建累计付费
func (receiver *PaymentStatusListReq) BuildCumulativePayment(tx *gorm.DB) (resp *gorm.DB) {
	tmpDb := tx
	if receiver.StatisticalCaliber == StatisticalCaliberRootGameBack30 {
		tmpDb = tmpDb.Table(data_report.NewDwsDayRootGameBackPayActiveLogModel().TableName() + " as " + aliasPaymentStatus)
	}
	commonBaseReq := receiver.BaseDataReport
	commonBaseReq.Indicators = append(commonBaseReq.Indicators, "pay_count", "active_user", "pay_amount")

	if slice.ContainAny(GameRelationDimensions, receiver.Dimensions) {
		tmpDb.Joins("join dim_game as game on game.platform_id = " + aliasPaymentStatus + ".platform_id and game.id = " + aliasPaymentStatus + ".game_id")
		tmpDb.Joins("join dim_main_game as main_game on main_game.platform_id = game.platform_id and main_game.id = game.main_id")
		tmpDb.Joins("join dim_root_game as root_game on root_game.platform_id = main_game.platform_id and root_game.id = main_game.root_game_id")
	}

	tmpDb.Where(aliasPaymentStatus+".reg_date BETWEEN ? AND ?", receiver.StartTime, receiver.EndTime)

	if receiver.AggregationTime == AggregationTimeDay {
		commonBaseReq.Dimensions = append(commonBaseReq.Dimensions, StatDateDay+aliasPaymentStatus)
	} else if receiver.AggregationTime == AggregationTimeMonth {
		commonBaseReq.Dimensions = append(commonBaseReq.Dimensions, StatDateMonth+aliasPaymentStatus)
	}

	dbBuilder := &DbBuilder{Db: tmpDb, BaseDataReport: commonBaseReq}
	dbBuilder.
		SetFieldsMap(paymentStatusFieldsMap).
		SetJoinsMap(receiver.getJoinsMap(aliasPaymentStatus)).
		SetWheresMap(receiver.getWheresMap(aliasPaymentStatus)).
		SetGroupsMap(paymentStatusGroupsMap)

	paymentStatusDb := dbBuilder.Build()
	resp = paymentStatusDb
	return
}

func (receiver *PaymentStatusListReq) BuildPaymentDb(tx *gorm.DB) (resp *gorm.DB) {
	tmpDb := tx
	if receiver.StatisticalCaliber == StatisticalCaliberRootGameBack30 {
		tmpDb = tmpDb.Table(data_report.NewDwsDayRootGameBackPayActiveLogModel().TableName() + " as " + aliasPaymentStatus)
	}

	commonBaseReq := receiver.BaseDataReport
	commonBaseReq.Indicators = append(commonBaseReq.Indicators, "pay_count", "active_user", "pay_amount")
	commonBaseReq.Dimensions = append(commonBaseReq.Dimensions, "active_days")

	if slice.ContainAny(GameRelationDimensions, receiver.Dimensions) {
		tmpDb.Joins("join dim_game as game on game.platform_id = " + aliasPaymentStatus + ".platform_id and game.id = " + aliasPaymentStatus + ".game_id")
		tmpDb.Joins("join dim_main_game as main_game on main_game.platform_id = game.platform_id and main_game.id = game.main_id")
		tmpDb.Joins("join dim_root_game as root_game on root_game.platform_id = main_game.platform_id and root_game.id = main_game.root_game_id")
	}

	tmpDb.Where(aliasPaymentStatus+".reg_date BETWEEN ? AND ?", receiver.StartTime, receiver.EndTime)

	if receiver.AggregationTime == AggregationTimeDay {
		commonBaseReq.Dimensions = append(commonBaseReq.Dimensions, StatDateDay+aliasPaymentStatus)
	} else if receiver.AggregationTime == AggregationTimeMonth {
		commonBaseReq.Dimensions = append(commonBaseReq.Dimensions, StatDateMonth+aliasPaymentStatus)
	}

	loginEndDate, _ := datetime.FormatStrToTime(receiver.EndTime, "yyyy-MM-dd")
	maxDay := 0
	if len(receiver.ActiveDays) > 0 {
		maxDay = receiver.ActiveDays[len(receiver.ActiveDays)-1]
	}
	loginEndDateAdd := loginEndDate.Add(time.Duration(maxDay-1) * 24 * time.Hour)
	tmpDb.Where(aliasPaymentStatus+".pay_date BETWEEN ? AND ?", receiver.StartTime, datetime.FormatTimeToStr(loginEndDateAdd, "yyyy-MM-dd"))

	dbBuilder := &DbBuilder{Db: tmpDb, BaseDataReport: commonBaseReq}
	dbBuilder.
		SetFieldsMap(paymentStatusFieldsMap).
		SetJoinsMap(receiver.getJoinsMap(aliasPaymentStatus)).
		SetWheresMap(receiver.getWheresMap(aliasPaymentStatus)).
		SetGroupsMap(paymentStatusGroupsMap)

	paymentStatusDb := dbBuilder.Build()
	resp = paymentStatusDb
	return
}

func (receiver *PaymentStatusListReq) BuildCombineJoinOn(alias1, alias2 string) string {
	if len(receiver.Dimensions) <= 0 && receiver.AggregationTime != AggregationTimeAll {
		receiver.Dimensions = append(receiver.Dimensions, "stat_date")
	}
	var combineOn []string
	for _, item := range receiver.Dimensions {
		combineOn = append(combineOn, fmt.Sprintf("%s.%s = %s.%s", alias1, item, alias2, item))
	}

	if len(combineOn) <= 0 {
		combineOn = append(combineOn, " 1 = 1 ")
	}
	return strings.Join(combineOn, " AND ")
}

func (receiver *PaymentStatusListReq) BuildDb(tx *gorm.DB) (resp *gorm.DB, err error) {

	overviewDb := receiver.BuildOverviewDb(tx)
	paymentDb := receiver.BuildPaymentDb(tx)
	cumulativePaymentDb := receiver.BuildCumulativePayment(tx)

	var joinAllOrders []string
	if receiver.AggregationTime != AggregationTimeAll {
		joinAllOrders = append(joinAllOrders, "stat_date")
	}
	joinAllOrders = append(joinAllOrders, "active_days")

	joinAllDb := BuildTemporaryTable(aliasOverview, overviewDb).
		Joins("LEFT JOIN (?) as "+aliasPaymentStatus+" ON "+receiver.BuildCombineJoinOn(aliasOverview, aliasPaymentStatus), paymentDb).
		Joins("LEFT JOIN (?) as "+aliasCumulativePayment+" ON "+receiver.BuildCombineJoinOn(aliasOverview, aliasCumulativePayment), cumulativePaymentDb)

	dbBuilder := &DbBuilder{
		Db:             joinAllDb,
		BaseDataReport: BaseDataReport{Indicators: []string{"join_all", "join_all_cumulative_payment"}, Orders: joinAllOrders}}
	dbBuilder.SetFieldsMap(paymentStatusFieldsMap)
	dbBuilder.SetOrdersMap(paymentStatusOrdersMap)
	resp = dbBuilder.Build()
	return
}

type PaymentStatusListResp struct {
	BaseResp
	Reg                  int `json:"reg"`
	Cost                 int `json:"cost"`
	ActiveDays           int `json:"active_days"`
	CumulativePayCount   int `json:"cumulative_pay_count"`
	CumulativeActiveUser int `json:"cumulative_active_user"`
	CumulativePayAmount  int `json:"cumulative_pay_amount"`
	PayCount             int `json:"pay_count"`
	ActiveUser           int `json:"active_user"`
	PayAmount            int `json:"pay_amount"`
}

type PaymentStatusListRespFormat struct {
	BaseResp
	Reg                  int                     `json:"reg"`
	Cost                 int                     `json:"cost"`
	CumulativePayCount   int                     `json:"cumulative_pay_count"`
	CumulativeActiveUser int                     `json:"cumulative_active_user"`
	CumulativePayAmount  int                     `json:"cumulative_pay_amount"`
	RoiRateStr           string                  `json:"roi_rate_str"` // 当前回本率
	Ltv                  float64                 `json:"ltv"`          //当前LTV
	NDayContainer        []PaymentStatusNDayData `json:"n_day_container"`
}

type PaymentStatusNDayData struct {
	Show                          bool    `json:"show"`                              //是否展示
	NDay                          int     `json:"n_day"`                             // N日
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
