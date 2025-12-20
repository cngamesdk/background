package api

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	errors2 "github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"slices"
	"strings"
)

type DimensionFilterInterface interface {
}

const (
	StatisticalCaliberRootGameBack30 = "root-game-back-30" // 30天回流统计口径
)

const (
	AggregationTimeDay   = "day"                   // 按日
	AggregationTimeMonth = "month"                 // 按月
	AggregationTimeAll   = "aggregation"           // 聚合
	StatDateDay          = "stat_date_day"         //按日统计
	StatDateWeek         = "stat_date_week"        //按周统计
	StatDateMonth        = "stat_date_month"       //按月统计
	StatDateYear         = "stat_date_year"        //按年统计
	StatDateAggregation  = "stat_date_aggregation" //聚合
)

const (
	OperatorLt       = "lt"        // 小于
	OperatorLtEqual  = "lt-equal"  // 小于等于
	OperatorGt       = "gt"        // 大于
	OperatorGtEqual  = "gt-equal"  // 大于等于
	OperatorEqual    = "equal"     // 等于
	OperatorNotEqual = "not-equal" // 不等于
	OperatorIn       = "in"        // in
	OperatorNotIn    = "not-in"    // not in
	OperatorBetween  = "between"   // 之间
)

var (
	GameRelationDimensions = []string{"root_game_id", "main_game_id", "game_id"} // 游戏相关维度
)

type DbBuilderWhere struct {
	Query interface{}
	Args  []interface{}
}

type DbBuilder struct {
	fieldsMap map[string]string
	wheresMap map[string]string
	groupsMap map[string]string
	joinsMap  map[string]func(tx *gorm.DB)
	ordersMap map[string]string
	Db        *gorm.DB
	BaseDataReport
	SelectsContainer []string
	GroupsContainer  []string
	OrdersContainer  []string
	WheresContainer  []DbBuilderWhere
}

func (receiver *DbBuilder) SetFieldsMap(req map[string]string) *DbBuilder {
	receiver.fieldsMap = req
	return receiver
}

func (receiver *DbBuilder) SetWheresMap(req map[string]string) *DbBuilder {
	receiver.wheresMap = req
	return receiver
}

func (receiver *DbBuilder) SetGroupsMap(req map[string]string) *DbBuilder {
	receiver.groupsMap = req
	return receiver
}

func (receiver *DbBuilder) SetJoinsMap(req map[string]func(tx *gorm.DB)) *DbBuilder {
	receiver.joinsMap = req
	return receiver
}

func (receiver *DbBuilder) SetOrdersMap(req map[string]string) *DbBuilder {
	receiver.ordersMap = req
	return receiver
}

func (receiver *DbBuilder) BuildJoins() *DbBuilder {
	for item, joinFun := range receiver.joinsMap {
		if slices.Contains(receiver.Dimensions, item) {
			joinFun(receiver.Db)
		}
	}
	return receiver
}

func (receiver *DbBuilder) BuildDimensions() *DbBuilder {
	for _, item := range receiver.Dimensions {
		fieldTmp, fieldTmpOk := receiver.fieldsMap[item]
		if !fieldTmpOk {
			fieldTmp = item
		}
		receiver.SelectsContainer = append(receiver.SelectsContainer, fieldTmp)

		groupTmp, groupTmpOk := receiver.groupsMap[item]
		if !groupTmpOk {
			groupTmp = item
		}
		receiver.GroupsContainer = append(receiver.GroupsContainer, groupTmp)
	}
	return receiver
}

func (receiver *DbBuilder) BuildIndicators() *DbBuilder {
	for _, item := range receiver.Indicators {
		fieldTmp, fieldTmpOk := receiver.fieldsMap[item]
		if !fieldTmpOk {
			fieldTmp = item
		}
		receiver.SelectsContainer = append(receiver.SelectsContainer, fieldTmp)
	}
	return receiver
}

func (receiver *DbBuilder) BuildOrders() *DbBuilder {
	for _, item := range receiver.Orders {
		tmp, tmpOk := receiver.ordersMap[item]
		if !tmpOk {
			tmp = item
		}
		receiver.OrdersContainer = append(receiver.OrdersContainer, tmp)
	}
	return receiver
}

func (receiver *DbBuilder) BuildDimensionsFilter() *DbBuilder {
	for _, item := range receiver.DimensionsFilter {
		whereColumn, whereColumnOK := receiver.wheresMap[item.Key]
		tmpColumn := item.Key
		if whereColumnOK {
			tmpColumn = whereColumn
		}
		args, buildErr := item.GetSqlOperator()
		if buildErr != nil {
			global.GVA_LOG.Error("维度筛选异常", zap.Error(buildErr))
		} else {
			receiver.Db.Where(fmt.Sprintf("%s %s", tmpColumn, args), item.Value)
		}
	}
	return receiver
}

func (receiver *DbBuilder) BuildAggregationTime(req func(tx *gorm.DB)) *DbBuilder {
	req(receiver.Db)
	return receiver
}

func (receiver *DbBuilder) Build() *gorm.DB {
	receiver.BuildDimensions()
	receiver.BuildIndicators()
	receiver.BuildDimensionsFilter()
	receiver.BuildJoins()
	receiver.BuildOrders()

	if len(receiver.SelectsContainer) > 0 {
		receiver.Db.Select(strings.Join(receiver.SelectsContainer, ","))
	}
	if len(receiver.WheresContainer) > 0 {
		for _, item := range receiver.WheresContainer {
			receiver.Db.Where(item.Query, item.Args)
		}
	}
	if len(receiver.GroupsContainer) > 0 {
		receiver.Db.Group(strings.Join(receiver.GroupsContainer, ","))
	}
	if len(receiver.OrdersContainer) > 0 {
		receiver.Db.Order(strings.Join(receiver.OrdersContainer, ","))
	}
	return receiver.Db
}

// BuildTemporaryTable 构建临时表
func BuildTemporaryTable(name string, args ...interface{}) *gorm.DB {
	return global.GVA_DB.Table("(?) as "+name, args)
}

type BaseDataReport struct {
	StatisticalCaliber string            `json:"statistical_caliber" form:"statistical_caliber"` // 口径
	DimensionsFilter   []DimensionFilter `json:"dimension_filter" form:"dimension_filter"`       // 维度筛选
	Dimensions         []string          `json:"dimensions" form:"dimensions"`                   // 维度选择
	Indicators         []string          `json:"indicators" form:"indicators"`                   // 指标
	Orders             []string          `json:"orders" form:"orders"`                           // 排序字段
	AggregationTime    string            `json:"aggregation_time" form:"aggregation_time"`       // 聚合时间
	StartTime          string            `json:"start_time" form:"start_time"`                   // 开始时间
	EndTime            string            `json:"end_time" form:"start_time"`                     // 结束时间
}

type DimensionFilter struct {
	Key      string        `json:"key" form:"key"`
	Operator string        `json:"operator" form:"operator"`
	Value    []interface{} `json:"value" form:"value"`
}

func (receiver *DimensionFilter) GetSqlOperator() (resp string, err error) {
	switch receiver.Operator {
	case OperatorLt:
		resp = "< ?"
		return
	case OperatorLtEqual:
		resp = "<= ?"
		return
	case OperatorGt:
		resp = "> ?"
		return
	case OperatorGtEqual:
		resp = ">= ?"
		return
	case OperatorEqual:
		resp = "= ?"
		return
	case OperatorNotEqual:
		resp = "!= ?"
		return
	case OperatorIn:
		resp = "in ?"
		return
	case OperatorNotIn:
		resp = "not in ?"
		return
	case OperatorBetween:
		resp = "between ? and ?"
		return
	default:
		err = errors2.New("未知操作符")
		return
	}
}

type BaseResp struct {
	PlatformId   int64  `json:"platform_id,omitempty"`
	PlatformName string `json:"platform_name,omitempty"`
	StatDate     string `json:"stat_date,omitempty"`
	RootGameId   int64  `json:"root_game_id,omitempty"`
	RootGameName string `json:"root_game_name,omitempty"`
	MainGameId   int64  `json:"main_game_id,omitempty"`
	MainGameName string `json:"main_game_name,omitempty"`
	GameId       int64  `json:"game_id,omitempty"`
	GameName     string `json:"game_name,omitempty"`
	AgentId      int64  `json:"agent_id,omitempty"`
	AgentName    string `json:"agent_name,omitempty"`
	SiteId       int64  `json:"site_id,omitempty"`
	SiteName     string `json:"site_name,omitempty"`
	Ad3Id        int64  `json:"ad3_id,omitempty"`
}
