package data_report

import (
	"context"
	"fmt"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data_report/api"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

type PaymentStatusService struct {
}

func (receiver PaymentStatusService) List(ctx context.Context, req *api.PaymentStatusListReq) (resp response.PageResult, err error) {
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
			BaseResp: item.BaseResp,
			Reg:      item.Reg,
		}

		if sameGroup && len(listFormat) > 0 {
			tmpRespFormat = listFormat[len(listFormat)-1]

			//补全数据
			lenNDayContainer := len(tmpRespFormat.NDayContainer)
			if lenNDayContainer >= 1 {
				topItem := tmpRespFormat.NDayContainer[lenNDayContainer-1]
				existsLastDayRetention := topItem.NDay == (item.ActiveDays - 1)
				if !existsLastDayRetention {
					tmpRespFormat.NDayContainer = append(tmpRespFormat.NDayContainer, api.PaymentStatusNDayData{
						NDay:       item.ActiveDays - 1,
						RoiRateStr: utils.FloatDecimal2Str(0),
					})
				}
			}

		} else {
			listFormat = append(listFormat, tmpRespFormat)
		}
		if item.ActiveDays != 1 {
			rate := utils.Percent(item.PayAmount, tmpRespFormat.Reg)
			tmpRespFormat.NDayContainer = append(tmpRespFormat.NDayContainer, api.PaymentStatusNDayData{
				NDay:         item.ActiveDays,
				DailyPayment: item.PayAmount,
				Ltv:          rate,
				RoiRateStr:   utils.FloatDecimal2Str(rate),
			})
		}
	}

	resp.List = listFormat
	return
}
