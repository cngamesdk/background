package oceanengine

import (
	"context"
	"github.com/cngamesdk/go-core/model/sql"
	advertising2 "github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/duke-git/lancet/v2/convertor"
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

// EbpAdvertiserList 获取升级版巨量引擎工作台下账户列表
func EbpAdvertiserList(db *gorm.DB) (err error) {
	uuid, _ := random.UUIdV4()
	ctxParent := context.Background()
	ctx := context.WithValue(ctxParent, "request_id", uuid)
	accountModel := advertising.NewDimAdvertisingMediaAccountModel()
	tmpAccountDb := db.WithContext(ctx).
		Table(accountModel.TableName()).
		Select("auth_id, account_id,code").
		Where("status = ?", sql.StatusNormal).
		Where("code = ?", advertising2.MediaCodeOceanengine).
		Where("role = ?", ocean_engine.AccountRolePlatformRoleEnterpriseBpAdmin)
	var countAccount int64
	if countErr := tmpAccountDb.Count(&countAccount).Error; countErr != nil {
		err = countErr
		global.GVA_LOG.Error("获取总数异常", zap.Error(countErr))
		return
	}
	pageAccount := 1
	pageSizeAccount := 100
	totalPageAccount := cast.ToInt(countAccount) / pageSizeAccount
	if cast.ToInt(countAccount)%pageSizeAccount != 0 {
		totalPageAccount += 1
	}
	for pageAccount <= totalPageAccount {
		var listAccount []advertising.DimAdvertisingMediaAccountModel
		if listErr := tmpAccountDb.Limit(pageSizeAccount).Offset((pageAccount - 1) * pageSizeAccount).Find(&listAccount).Error; listErr != nil {
			err = listErr
			global.GVA_LOG.Error("获取列表异常", zap.Error(listErr))
			return
		}
		for _, accountItem := range listAccount {
			authModel := advertising.NewDimAdvertisingMediaAuthModel()
			if takeErr := authModel.Take(ctx, "*", "id = ?", accountItem.AuthId); takeErr != nil {
				global.GVA_LOG.Error("获取AUTH异常", zap.Error(takeErr))
				return
			}
			apiPage := 1
			apiPageSize := 100
			mediaService := ad_platform.NewOceanEngineAdapter(global.GVA_LOG)
			mediaService.Init(ad_platform.AdapterConfig{Auth: *authModel})
			for {
				apiResult, apiErr := mediaService.EbpAdvertiserList(ctx, &ocean_engine.EbpAdvertiserListReq{
					EnterpriseOrganizationId: accountItem.AccountId,
					AccountSource:            ocean_engine.AccountSourceAd,
					Page:                     apiPage,
					PageSize:                 apiPageSize,
				})
				if apiErr != nil {
					global.GVA_LOG.Error("获取列表异常", zap.Error(apiErr))
					break
				}
				if apiPage >= apiResult.PageInfo.TotalPage {
					break
				}
				var sqlDataContainer []advertising.DimAdvertisingMediaAccountModel
				for _, apiAccountItem := range apiResult.AccountList {

					extension, extensionErr := convertor.StructToMap(apiAccountItem)
					if extensionErr != nil {
						global.GVA_LOG.Error("转换失败", zap.Error(extensionErr))
					}

					mediaAccountModel := advertising.DimAdvertisingMediaAccountModel{}
					mediaAccountModel.PlatformId = accountItem.PlatformId
					mediaAccountModel.AccountName = apiAccountItem.AccountName
					mediaAccountModel.AccountId = apiAccountItem.AccountId
					mediaAccountModel.Status = sql.StatusNormal
					mediaAccountModel.AuthId = accountItem.AuthId
					mediaAccountModel.Role = ocean_engine.AccountRoleAdvertiser
					mediaAccountModel.Extension = extension
					mediaAccountModel.Code = accountItem.Code

					sqlDataContainer = append(sqlDataContainer, mediaAccountModel)
				}

				tableName := (&advertising.DimAdvertisingMediaAccountModel{}).TableName()
				if batchErr := db.WithContext(ctx).Table(tableName).Clauses(clause.OnConflict{
					Columns:   []clause.Column{{Name: "account_id"}, {Name: "code"}},
					UpdateAll: true,
				}).Create(&sqlDataContainer).Error; batchErr != nil {
					global.GVA_LOG.Error("批量保存失败", zap.Error(batchErr), zap.Any("data", sqlDataContainer))
				}

				apiPage++
			}
		}
	}
	return
}
