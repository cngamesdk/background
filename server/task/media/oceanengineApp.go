package media

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	advertising2 "github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/duke-git/lancet/v2/random"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ad_placement/ad_platform"
	"github.com/flipped-aurora/gin-vue-admin/server/service/ad_placement/ad_platform/model/ocean_engine"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 获取安卓应用列表
func OceanengineAppList(db *gorm.DB) (err error) {
	ctxParent := context.Background()
	uuid, _ := random.UUIdV4()
	ctx := context.WithValue(ctxParent, "request_id", uuid)
	//获取EBP帐户
	accountModel := advertising.NewDimAdvertisingMediaAccountModel()
	tmpAccountDb := accountModel.Db().
		WithContext(ctx).
		Select("account_id,auth_id").
		Table(accountModel.TableName()).
		Where("status = ? and code = ? and role = ?",
			sql.StatusNormal,
			advertising2.MediaCodeOceanengine,
			ocean_engine.AccountRolePlatformRoleEnterpriseBpAdmin)
	var count int64
	if countErr := tmpAccountDb.Count(&count).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取帐户总数异常", zap.Error(countErr))
		return
	}
	page := 1
	pageSize := 100
	totalPage := cast.ToInt(count) / pageSize
	if cast.ToInt(count)%pageSize != 0 {
		totalPage += 1
	}
	//分页获取EBP帐户
	for page <= totalPage {
		var accountList []advertising.DimAdvertisingMediaAccountModel
		listErr := tmpAccountDb.Limit(pageSize).Offset((page - 1) * pageSize).
			Order("id DESC").Find(&accountList).Error
		page++
		if listErr != nil {
			err = listErr
			global.GVA_LOG.Error("获取帐户列表异常", zap.Error(listErr))
			continue
		}
		//遍历EBP帐户
		for _, accountItem := range accountList {
			//获取token
			tokenModel := advertising.NewDimAdvertisingMediaAuthModel()
			if takeErr := tokenModel.Take(ctx, "*", "id = ?", accountItem.AuthId); takeErr != nil {
				global.GVA_LOG.Error("获取token异常", zap.Error(takeErr), zap.Any("data", accountItem))
				continue
			}

			pageAppList := 1
			pageSizeAppList := 100
			totalPageAppList := 0
			oceanEngine := ad_platform.NewOceanEngineAdapter(global.GVA_LOG)
			oceanEngine.Init(ad_platform.AdapterConfig{Auth: *tokenModel})
			for {
				//分页获取应用列表
				appListReq := ocean_engine.AppListReq{
					AccountId:            accountItem.AccountId,
					AccountType:          ocean_engine.AccountTypeEbp,
					AssetManagementScope: ocean_engine.AssetManagementScopeTraverse,
					Page:                 pageAppList,
					PageSize:             pageSizeAppList,
				}
				appListResult, appListErr := oceanEngine.GetAppList(ctx, appListReq)
				if appListErr != nil {
					global.GVA_LOG.Error("获取列表异常", zap.Error(appListErr), zap.Any("data", appListReq))
					break
				}
				totalPageAppList = appListResult.PageInfo.TotalPage
				if pageAppList >= totalPageAppList {
					break
				}
				pageAppList++
				// 补全信息
				for appIndex, _ := range appListResult.BasicAppList {
					appListResult.BasicAppList[appIndex].PlatformId = accountItem.PlatformId
					appListResult.BasicAppList[appIndex].AccountId = accountItem.AccountId
				}
				appLogModel := &advertising.OdsAdvertisingOceanengineAppLogModel{}
				batchCreateErr := db.WithContext(ctx).Table(appLogModel.TableName()).Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "basic_package_id"}},
					UpdateAll: true,
				}).Create(&appListResult).Error
				if batchCreateErr != nil {
					global.GVA_LOG.Error("批量插入异常", zap.Error(batchCreateErr), zap.Any("data", appListResult))
					continue
				}
			}
		}
	}
	return
}
