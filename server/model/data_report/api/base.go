package api

import (
	errors2 "github.com/pkg/errors"
)

type DimensionFilterInterface interface {
}

const (
	StatisticalCaliberRootGameBack30 = "root-game-back-30" // 30天回流统计口径
)

const (
	AggregationTimeDay   = "day"         // 按日
	AggregationTimeMonth = "month"       // 按月
	AggregationTimeAll   = "aggregation" // 聚合
)

const (
	OperatorLt       = "lt"        // 小于
	OperatorLtEqual  = "lt-equal"  // 小于等于
	OperatorGt       = "gt"        // 大于
	OperatorGtEqual  = "gt-equal"  // 大于等于
	OperatorEqual    = "equal"     // 等于
	OperatorNotEqual = "not-equal" // 不等于
	OperatorNot      = "not"       // 非
	OperatorIn       = "in"        // in
	OperatorNotIn    = "not-in"    // not in
	OperatorBetween  = "between"   // 之间
)

type BaseDataReport struct {
	StatisticalCaliber string            `json:"statistical_caliber" form:"statistical_caliber"` // 口径
	DimensionsFilter   []DimensionFilter `json:"dimension_filter" form:"dimension_filter"`       // 维度筛选
	Dimensions         []string          `json:"dimensions" form:"dimensions"`                   // 维度选择
	Indicators         []string          `json:"indicators" form:"indicators"`                   // 指标
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
