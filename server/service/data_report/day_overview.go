package data_report

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/data_report/api"
	"go.uber.org/zap"
)

type DayOverviewService struct {
}

func (receiver DayOverviewService) List(ctx context.Context, req *api.DayRootGameBackOverviewListReq) (resp response.PageResult, err error) {
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	subQuery, buildErr := req.BuildDb()
	if buildErr != nil {
		err = buildErr
		global.GVA_LOG.Error("构建DB异常", zap.Error(buildErr))
		return
	}

	resp.Sql = sql.GetTakeSql(subQuery)

	var total int64
	if countErr := subQuery.Count(&total).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	if total <= 0 {
		return
	}
	var list []api.DayRootGameBackOverviewListRespData
	tmpListDb := subQuery.Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize)

	resp.Sql = sql.GetFindSql(tmpListDb)

	if listErr := tmpListDb.Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
		return
	}
	resp.List = list
	return
}
