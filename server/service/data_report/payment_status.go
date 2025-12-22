package data_report

import (
	"context"
	"fmt"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data_report/api"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"slices"
	"time"
)

type PaymentStatusService struct {
}

func (receiver *PaymentStatusService) List(ctx context.Context, req *api.PaymentStatusListReq) (resp response.PageResult, err error) {
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	subQuery, buildErr := req.BuildDb(global.GVA_DB.WithContext(ctx))
	if buildErr != nil {
		err = buildErr
		global.GVA_LOG.Error("构建DB异常", zap.Error(buildErr))
		return
	}

	var list []api.PaymentStatusListResp

	resp.Sql = sql.GetFindSql(subQuery)

	if listErr := subQuery.Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
		return
	}

	var listFormat []*api.PaymentStatusListRespFormat

	last := api.PaymentStatusListResp{}
	for _, item := range list {
		sameGroup := (fmt.Sprintf("%+v", item.BaseResp) + item.StatDate) == (fmt.Sprintf("%+v", last.BaseResp) + last.StatDate)
		last = item
		tmpRespFormat := &api.PaymentStatusListRespFormat{
			BaseResp:             item.BaseResp,
			Reg:                  item.Reg,
			Cost:                 item.Cost,
			CumulativePayCount:   item.CumulativePayCount,
			CumulativeActiveUser: item.CumulativeActiveUser,
			CumulativePayAmount:  item.CumulativePayAmount,
			RoiRateStr:           utils.FloatDecimal2Str(utils.Percent(item.CumulativePayAmount, item.Cost)),
			Ltv:                  utils.Percent(item.CumulativePayAmount, item.Reg),
		}

		if sameGroup && len(listFormat) > 0 {
			tmpRespFormat = listFormat[len(listFormat)-1]
		} else if item.ActiveDays > 0 {
			listFormat = append(listFormat, tmpRespFormat)
		}
		tmpRespFormat.NDayContainer = append(tmpRespFormat.NDayContainer, api.PaymentStatusNDayData{
			NDay:                  item.ActiveDays,
			DailyPayment:          item.PayAmount,
			DailyPaymentUsers:     item.ActiveUser,
			DailyPaymentFrequency: item.PayCount,
		})
	}
	maxDay := 0
	if len(req.ActiveDays) > 0 {
		maxDay = req.ActiveDays[len(req.ActiveDays)-1]
	}

	//补全天数
	for _, item := range listFormat {
		var tmpNDayContainer []api.PaymentStatusNDayData
		for index := 1; index <= maxDay; index++ {
			tmpNDayData := api.PaymentStatusNDayData{
				NDay: index,
			}
			for _, nDayItem := range item.NDayContainer {
				if nDayItem.NDay == index {
					tmpNDayData = nDayItem
					break
				}
			}
			//第一项
			tmpCumulativePayments := 0
			tmpCumulativePaymentUsers := 0
			tmpCumulativePaymentFrequency := 0

			tmpStatDate := req.EndTime
			if req.AggregationTime != api.AggregationTimeAll {
				tmpStatDate = item.StatDate
			}
			statDate, _ := datetime.FormatStrToTime(tmpStatDate, "yyyy-MM-dd")
			if len(tmpNDayContainer) > 0 && (statDate.Add(time.Hour*time.Duration(index*24)).Unix()-time.Now().Unix() < 24*3600) {
				lastItem := tmpNDayContainer[len(tmpNDayContainer)-1]
				tmpCumulativePayments = lastItem.CumulativePayments
				tmpCumulativePaymentUsers = lastItem.CumulativePaymentUsers
				tmpCumulativePaymentFrequency = lastItem.CumulativePaymentFrequency
			}
			tmpNDayData.CumulativePayments = tmpCumulativePayments + tmpNDayData.DailyPayment
			tmpNDayData.CumulativePaymentUsers = tmpCumulativePaymentUsers + tmpNDayData.DailyPaymentUsers
			tmpNDayData.CumulativePaymentFrequency = tmpCumulativePaymentFrequency + tmpNDayData.DailyPaymentFrequency
			tmpNDayData.RoiRateStr = utils.FloatDecimal2Str(utils.Percent(tmpNDayData.CumulativePayments, item.Cost))
			tmpNDayData.Ltv = utils.Percent(tmpNDayData.CumulativePayments, item.Reg)
			tmpNDayData.Show = slices.Contains(req.ActiveDays, tmpNDayData.NDay)
			tmpNDayContainer = append(tmpNDayContainer, tmpNDayData)
		}
		item.NDayContainer = tmpNDayContainer
	}
	resp.List = listFormat
	return
}
