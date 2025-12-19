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
	aliasActive = "active"
)

var (
	retentionStatusFieldsMap = map[string]string{
		"platform_id":  "platform.id as platform_id, platform.platform_name as platform_name",
		"root_game_id": "root_game.id as root_game_id, root_game.game_name as root_game_name",
		"main_game_id": "main_game.id as main_game_id, main_game.game_name as main_game_name",
		"game_id":      "game.id as game_id, game.game_name as game_name",
		"agent_id":     "agent.id as agent_id, agent.agent_name as agent_name",
		"site_id":      "site.id as site_id, site.site_name as site_name",
		"active_days":  aliasActive + ".active_days AS active_days",
		"active_count": "SUM( " + aliasActive + ".active_count ) AS active_count",
	}
	retentionStatusWheresMap = map[string]string{
		"platform_id":  aliasActive + ".platform_id",
		"root_game_id": "root_game.id",
		"main_game_id": "main_game.id",
		"game_id":      "game.id",
		"agent_id":     "agent.id",
		"site_id":      "site.id",
	}
	retentionStatusGroupsMap = map[string]string{
		"platform_id":  "platform.id",
		"root_game_id": "root_game.id",
		"main_game_id": "main_game.id",
		"game_id":      "game.id",
		"agent_id":     "agent.id",
		"site_id":      "site.id",
		"active_days":  aliasActive + ".active_days",
	}
	retentionStatusJoinsMap = map[string]func(tx *gorm.DB){
		"platform_id": func(tx *gorm.DB) {
			tx.Joins("join dim_platform as platform on platform.id = " + aliasActive + ".platform_id")
		},
		"agent_id": func(tx *gorm.DB) {
			tx.Joins("join dim_agent as agent on agent.platform_id = " + aliasActive + ".platform_id and agent.id = " + aliasActive + ".agent_id")
		},
		"site_id": func(tx *gorm.DB) {
			tx.Joins("join dim_site as site on site.platform_id = " + aliasActive + ".platform_id and site.id = " + aliasActive + ".site_id")
		},
	}
	retentionStatusOrdersMap = map[string]string{
		"active_days": "active_days",
		"stat_date":   "stat_date",
	}
)

type RetentionStatusListReq struct {
	BaseDataReport
	request.PageInfo
	ActiveDays []int `json:"-" form:"-"`
}

func (receiver *RetentionStatusListReq) Format() {

	var tmpIndicators []string
	tmpIndicators = append(tmpIndicators, "active_count")
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
		retentionStatusFieldsMap["stat_date"] = "DATE_FORMAT(" + aliasActive + ".reg_date, '%Y-%m-%d') as stat_date"
		retentionStatusGroupsMap["stat_date"] = "DATE_FORMAT(" + aliasActive + ".reg_date, '%Y-%m-%d')"
	} else if receiver.AggregationTime == AggregationTimeMonth {
		receiver.Dimensions = append(receiver.Dimensions, "stat_date")
		receiver.Orders = append(receiver.Orders, "stat_date")
		retentionStatusFieldsMap["stat_date"] = "DATE_FORMAT(" + aliasActive + ".reg_date, '%Y-%m') as stat_date"
		retentionStatusGroupsMap["stat_date"] = "DATE_FORMAT(" + aliasActive + ".reg_date, '%Y-%m')"
	}
	receiver.Dimensions = append(receiver.Dimensions, "active_days")
	receiver.Orders = append(receiver.Orders, "active_days")
}

func (receiver *RetentionStatusListReq) BuildDb(tx *gorm.DB) (resp *gorm.DB, err error) {
	tmpDb := tx
	if receiver.StatisticalCaliber == StatisticalCaliberRootGameBack30 {
		tmpDb = tmpDb.Table(data_report.NewDwsDayRootGameBackLoginActiveLogModel().TableName() + " as " + aliasActive)
	}

	tmpDb.Joins("join dim_game as game on game.platform_id = " + aliasActive + ".platform_id and game.id = " + aliasActive + ".game_id")
	tmpDb.Joins("join dim_main_game as main_game on main_game.platform_id = game.platform_id and main_game.id = game.main_id")
	tmpDb.Joins("join dim_root_game as root_game on root_game.platform_id = main_game.platform_id and root_game.id = main_game.root_game_id")

	tmpDb.Where(aliasActive+".reg_date BETWEEN ? AND ?", receiver.StartTime, receiver.EndTime)
	loginEndDate, _ := datetime.FormatStrToTime(receiver.EndTime, "yyyy-MM-dd")
	maxDay := 0
	if len(receiver.ActiveDays) > 0 {
		maxDay = receiver.ActiveDays[len(receiver.ActiveDays)-1]
	}
	loginEndDateAdd := loginEndDate.Add(time.Duration(maxDay) * 24 * time.Hour)
	tmpDb.Where(aliasActive+".login_date BETWEEN ? AND ?", receiver.StartTime, datetime.FormatTimeToStr(loginEndDateAdd, "yyyy-MM-dd"))

	dbBuilder := &DbBuilder{Db: tmpDb, BaseDataReport: receiver.BaseDataReport}
	dbBuilder.
		SetFieldsMap(retentionStatusFieldsMap).
		SetJoinsMap(retentionStatusJoinsMap).
		SetWheresMap(retentionStatusWheresMap).
		SetGroupsMap(retentionStatusGroupsMap).
		SetOrdersMap(retentionStatusOrdersMap)

	resp = dbBuilder.Build()
	return
}

type RetentionStatusListResp struct {
	BaseResp
	Reg         int `json:"reg"`
	ActiveDays  int `json:"active_days"`
	ActiveCount int `json:"active_count"`
}

type RetentionStatusListRespFormat struct {
	BaseResp
	Reg           int                       `json:"reg"`
	NDayContainer []RetentionStatusNDayData `json:"n_day_container"`
}

type RetentionStatusNDayData struct {
	NDay             int     `json:"n_day"`
	RetentionData    int     `json:"retention_data"`
	RetentionRate    float64 `json:"retention_rate"`
	RetentionRateStr string  `json:"retention_rate_str"`
}
