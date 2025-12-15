package api

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data_report"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"slices"
	"strings"
)

const (
	aliasOverview = "overview"
)

var fieldsMap = map[string]string{
	"platform_id":  "platform.id as platform_id, platform.platform_name as platform_name",
	"root_game_id": "root_game.id as root_game_id, root_game.game_name as root_game_name",
	"main_game_id": "main_game.id as main_game_id, main_game.game_name as main_game_name",
	"game_id":      "game.id as game_id, game.game_name as game_name",
	"agent_id":     "agent.id as agent_id, agent.agent_name as agent_name",
	"site_id":      "site.id as site_id, site.site_name as site_name",
}

var wheresMap = map[string]string{
	"platform_id":  aliasOverview + ".platform_id",
	"root_game_id": "root_game.id",
	"main_game_id": "main_game.id",
	"game_id":      "game.id",
	"agent_id":     "agent.id",
	"site_id":      "site.id",
}

var groupsMap = map[string]string{
	"platform_id":  "platform.id",
	"root_game_id": "root_game.id",
	"main_game_id": "main_game.id",
	"game_id":      "game.id",
	"agent_id":     "agent.id",
	"site_id":      "site.id",
}

var joinTablesMap = map[string]func(tx *gorm.DB){
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

type DayRootGameBackOverviewListReq struct {
	BaseDataReport
	request.PageInfo
}

func (receiver *DayRootGameBackOverviewListReq) Format() {
	if len(receiver.Indicators) <= 0 {
		receiver.Indicators = append(receiver.Indicators, "reg")
	}
}

func (receiver *DayRootGameBackOverviewListReq) BuildDb() (resp *gorm.DB, err error) {
	var selects []string
	var groups []string
	tmpDb := global.GVA_DB
	tmpDb = tmpDb.Where(aliasOverview+".stat_date BETWEEN ? AND ?", receiver.StartTime, receiver.EndTime)
	if receiver.StatisticalCaliber == StatisticalCaliberRootGameBack30 {
		tmpDb = tmpDb.Table(data_report.NewDwsDayRootGameBackOverviewLogModel().TableName() + " as " + aliasOverview)
	}

	tmpDb.Joins("join dim_game as game on game.platform_id = " + aliasOverview + ".platform_id and game.id = " + aliasOverview + ".game_id")
	tmpDb.Joins("join dim_main_game as main_game on main_game.platform_id = game.platform_id and main_game.id = game.main_id")
	tmpDb.Joins("join dim_root_game as root_game on root_game.platform_id = main_game.platform_id and root_game.id = main_game.root_game_id")

	for item, joinFun := range joinTablesMap {
		if slices.Contains(receiver.Dimensions, item) {
			joinFun(tmpDb)
		}
	}

	if receiver.AggregationTime == AggregationTimeDay {
		selects = append(selects, "DATE_FORMAT("+aliasOverview+".stat_date, '%Y-%m-%d') as stat_date")
		groups = append(groups, "DATE_FORMAT("+aliasOverview+".stat_date, '%Y-%m-%d')")
	} else if receiver.AggregationTime == AggregationTimeMonth {
		selects = append(selects, "DATE_FORMAT("+aliasOverview+".stat_date, '%Y-%m') as stat_date")
		groups = append(groups, "DATE_FORMAT("+aliasOverview+".stat_date, '%Y-%m')")
	} else if receiver.AggregationTime == AggregationTimeAll {

	}
	for _, item := range receiver.Dimensions {
		fieldTmp, fieldTmpOk := fieldsMap[item]
		if fieldTmpOk {
			selects = append(selects, fieldTmp)
		} else {
			selects = append(selects, item)
		}
		groupTmp, groupTmpOk := groupsMap[item]
		if groupTmpOk {
			groups = append(groups, groupTmp)
		} else {
			groups = append(groups, item)
		}
	}
	for _, item := range receiver.Indicators {
		fieldTmp, fieldTmpOk := fieldsMap[item]
		if fieldTmpOk {
			selects = append(selects, fieldTmp)
		} else {
			selects = append(selects, fmt.Sprintf("SUM( %s ) as %s", item, item))
		}
	}
	for _, item := range receiver.DimensionsFilter {
		whereColumn, whereColumnOK := wheresMap[item.Key]
		tmpColumn := item.Key
		if whereColumnOK {
			tmpColumn = whereColumn
		}
		args, buildErr := item.GetSqlOperator()
		if buildErr != nil {
			err = buildErr
			global.GVA_LOG.Error("维度筛选异常", zap.Error(buildErr))
			return
		}
		tmpDb.Where(fmt.Sprintf("%s %s", tmpColumn, args), item.Value)
	}
	if len(selects) > 0 {
		tmpDb.Select(strings.Join(selects, ","))
	}
	if len(groups) > 0 {
		tmpDb.Group(strings.Join(groups, ","))
		tmpDb = global.GVA_DB.Table("(?) as tmp", tmpDb)
	}
	resp = tmpDb
	return
}

type DayRootGameBackOverviewListRespData struct {
	PlatformId       int64  `json:"platform_id,omitempty"`
	PlatformName     string `json:"platform_name,omitempty"`
	StatDate         string `json:"stat_date,omitempty"`
	RootGameId       int64  `json:"root_game_id,omitempty"`
	RootGameName     string `json:"root_game_name,omitempty"`
	MainGameId       int64  `json:"main_game_id,omitempty"`
	MainGameName     string `json:"main_game_name,omitempty"`
	GameId           int64  `json:"game_id,omitempty"`
	GameName         string `json:"game_name,omitempty"`
	AgentId          int64  `json:"agent_id,omitempty"`
	AgentName        string `json:"agent_name,omitempty"`
	SiteId           int64  `json:"site_id,omitempty"`
	SiteName         string `json:"site_name,omitempty"`
	Ad3Id            int64  `json:"ad3_id,omitempty"`
	Activation       int64  `json:"activation,omitempty"`
	ActivationDevice int64  `json:"activation_device,omitempty"`
	Launch           int64  `json:"launch,omitempty"`
	LaunchDevice     int64  `json:"launch_device,omitempty"`
	Reg              int64  `json:"reg,omitempty"`
	RegDevice        int64  `json:"reg_device,omitempty"`
	Login            int64  `json:"login,omitempty"`
	LoginUser        int64  `json:"login_user,omitempty"`
	LoginDevice      int64  `json:"login_device,omitempty"`
	Role             int64  `json:"role,omitempty"`
	RoleUser         int64  `json:"role_user,omitempty"`
	RoleDevice       int64  `json:"role_device,omitempty"`
	Pay              int64  `json:"pay,omitempty"`
	PayUser          int64  `json:"pay_user,omitempty"`
	PayDevice        int64  `json:"pay_device,omitempty"`
	PayMoney         int64  `json:"pay_money,omitempty"`
}
