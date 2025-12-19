package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data_report"
	"gorm.io/gorm"
)

const (
	aliasOverview = "overview"
)

var (
	dayOverviewFieldsMap = map[string]string{
		"platform_id":       "platform.id as platform_id, platform.platform_name as platform_name",
		"root_game_id":      "root_game.id as root_game_id, root_game.game_name as root_game_name",
		"main_game_id":      "main_game.id as main_game_id, main_game.game_name as main_game_name",
		"game_id":           "game.id as game_id, game.game_name as game_name",
		"agent_id":          "agent.id as agent_id, agent.agent_name as agent_name",
		"site_id":           "site.id as site_id, site.site_name as site_name",
		"activation":        "SUM( " + aliasOverview + ".activation ) AS activation",
		"activation_device": "SUM( " + aliasOverview + ".activation_device ) AS activation_device",
		"launch":            "SUM( " + aliasOverview + ".launch ) AS launch",
		"launch_device":     "SUM( " + aliasOverview + ".launch_device ) AS launch_device",
		"reg":               "SUM( " + aliasOverview + ".reg ) AS reg",
		"reg_device":        "SUM( " + aliasOverview + ".reg_device ) AS reg_device",
		"login":             "SUM( " + aliasOverview + ".reg_device ) AS login",
		"login_user":        "SUM( " + aliasOverview + ".login_user ) AS login_user",
		"login_device":      "SUM( " + aliasOverview + ".login_device ) AS login_device",
		"role":              "SUM( " + aliasOverview + ".login_device ) AS role",
		"role_user":         "SUM( " + aliasOverview + ".role_user ) AS role_user",
		"role_device":       "SUM( " + aliasOverview + ".role_device ) AS role_device",
		"pay":               "SUM( " + aliasOverview + ".pay ) AS pay",
		"pay_user":          "SUM( " + aliasOverview + ".pay_user ) AS pay_user",
		"pay_device":        "SUM( " + aliasOverview + ".pay_device ) AS pay_device",
		"pay_money":         "SUM( " + aliasOverview + ".pay_money ) AS pay_money",
	}
	dayOverviewWheresMap = map[string]string{
		"platform_id":  aliasOverview + ".platform_id",
		"root_game_id": "root_game.id",
		"main_game_id": "main_game.id",
		"game_id":      "game.id",
		"agent_id":     "agent.id",
		"site_id":      "site.id",
	}
	dayOverviewGroupsMap = map[string]string{
		"platform_id":  "platform.id",
		"root_game_id": "root_game.id",
		"main_game_id": "main_game.id",
		"game_id":      "game.id",
		"agent_id":     "agent.id",
		"site_id":      "site.id",
	}
	dayOverviewJoinsMap = map[string]func(tx *gorm.DB){
		"platform_id": func(tx *gorm.DB) {
			tx.Joins("join dim_platform as platform on platform.id = " + aliasOverview + ".platform_id")
		},
		"agent_id": func(tx *gorm.DB) {
			tx.Joins("join dim_agent as agent on agent.platform_id = " + aliasOverview + ".platform_id and agent.id = " + aliasOverview + ".agent_id")
		},
		"site_id": func(tx *gorm.DB) {
			tx.Joins("join dim_site as site on site.platform_id = " + aliasOverview + ".platform_id and site.id = " + aliasOverview + ".site_id")
		},
	}
)

type DayOverviewListReq struct {
	BaseDataReport
	request.PageInfo
}

func (receiver *DayOverviewListReq) Format() {
	if len(receiver.Indicators) <= 0 {
		receiver.Indicators = append(receiver.Indicators, "reg")
	}
}

func (receiver *DayOverviewListReq) BuildDb(tx *gorm.DB) (resp *gorm.DB, err error) {
	tmpDb := tx
	tmpDb = tmpDb.Where(aliasOverview+".stat_date BETWEEN ? AND ?", receiver.StartTime, receiver.EndTime)
	if receiver.StatisticalCaliber == StatisticalCaliberRootGameBack30 {
		tmpDb = tmpDb.Table(data_report.NewDwsDayRootGameBackOverviewLogModel().TableName() + " as " + aliasOverview)
	}

	tmpDb.Joins("join dim_game as game on game.platform_id = " + aliasOverview + ".platform_id and game.id = " + aliasOverview + ".game_id")
	tmpDb.Joins("join dim_main_game as main_game on main_game.platform_id = game.platform_id and main_game.id = game.main_id")
	tmpDb.Joins("join dim_root_game as root_game on root_game.platform_id = main_game.platform_id and root_game.id = main_game.root_game_id")

	dbBuilder := &DbBuilder{Db: tmpDb, BaseDataReport: receiver.BaseDataReport}
	dbBuilder.
		SetFieldsMap(dayOverviewFieldsMap).
		SetJoinsMap(dayOverviewJoinsMap).
		SetWheresMap(dayOverviewWheresMap).
		SetGroupsMap(dayOverviewGroupsMap).
		BuildAggregationTime(func(tx *gorm.DB) {
			if receiver.AggregationTime == AggregationTimeDay {
				dbBuilder.SelectsContainer = append(dbBuilder.SelectsContainer, "DATE_FORMAT("+aliasOverview+".stat_date, '%Y-%m-%d') as stat_date")
				dbBuilder.GroupsContainer = append(dbBuilder.GroupsContainer, "DATE_FORMAT("+aliasOverview+".stat_date, '%Y-%m-%d')")
				dbBuilder.OrdersContainer = append(dbBuilder.OrdersContainer, "stat_date")
			} else if receiver.AggregationTime == AggregationTimeMonth {
				dbBuilder.SelectsContainer = append(dbBuilder.SelectsContainer, "DATE_FORMAT("+aliasOverview+".stat_date, '%Y-%m') as stat_date")
				dbBuilder.GroupsContainer = append(dbBuilder.GroupsContainer, "DATE_FORMAT("+aliasOverview+".stat_date, '%Y-%m')")
				dbBuilder.OrdersContainer = append(dbBuilder.OrdersContainer, "stat_date")
			}
		})
	resp = BuildTemporaryTable("tmp", dbBuilder.Build())
	return
}

type DayOverviewListRespData struct {
	BaseResp
	Activation       int `json:"activation,omitempty"`
	ActivationDevice int `json:"activation_device,omitempty"`
	Launch           int `json:"launch,omitempty"`
	LaunchDevice     int `json:"launch_device,omitempty"`
	Reg              int `json:"reg,omitempty"`
	RegDevice        int `json:"reg_device,omitempty"`
	Login            int `json:"login,omitempty"`
	LoginUser        int `json:"login_user,omitempty"`
	LoginDevice      int `json:"login_device,omitempty"`
	Role             int `json:"role,omitempty"`
	RoleUser         int `json:"role_user,omitempty"`
	RoleDevice       int `json:"role_device,omitempty"`
	Pay              int `json:"pay,omitempty"`
	PayUser          int `json:"pay_user,omitempty"`
	PayDevice        int `json:"pay_device,omitempty"`
	PayMoney         int `json:"pay_money,omitempty"`
}
