package task

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	"github.com/duke-git/lancet/v2/random"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	advertising2 "github.com/flipped-aurora/gin-vue-admin/server/service/advertising"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var (
	refreshTokenRunning = false
)

// RefreshToken 刷新媒体token定时器
func RefreshToken(db *gorm.DB) (err error) {
	if refreshTokenRunning {
		return
	}
	requestId, _ := random.UUIdV4()
	pCtx := context.Background()
	ctx := context.WithValue(pCtx, "request_id", requestId)
	global.GVA_LOG.Info("开始执行", zap.Any("request_id", requestId))
	refreshTokenRunning = true
	defer func() {
		refreshTokenRunning = false
		global.GVA_LOG.Info("结束执行", zap.Any("request_id", requestId))
	}()

	model := advertising.NewDimAdvertisingMediaAuthModel()
	var count int64
	page := 1
	pageSize := 100
	tempDb := db.WithContext(ctx).
		Table(model.TableName()).
		Where("status = ?", sql.StatusNormal).
		Where("refresh_token_expires_at between ? and ?", time.Now().Add(time.Duration(-5)*time.Hour), time.Now())

	if countErr := tempDb.Count(&count).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	totalPage := count / int64(pageSize)
	if count%int64(pageSize) != 0 {
		totalPage += 1
	}
	for int64(page) <= totalPage {
		var result []advertising.DimAdvertisingMediaAuthModel
		if listErr := tempDb.Order("id DESC").Limit(pageSize).Offset(page - 1).Take(&result).Error; listErr != nil {
			err = listErr
			global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
			return
		}
		for _, item := range result {
			service := &advertising2.AdvertisingAuthService{}
			refreshErr := service.RefreshToken(ctx, item)
			if refreshErr != nil {
				global.GVA_LOG.Error("刷新Token异常", zap.Error(refreshErr), zap.Any("data", item))
			}
		}
		page++
	}
	return
}
