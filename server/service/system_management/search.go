package system_management

import (
	"context"
	error2 "github.com/cngamesdk/go-core/model/error"
	advertising2 "github.com/cngamesdk/go-core/model/sql/advertising"
	"github.com/cngamesdk/go-core/model/sql/common"
	"github.com/duke-git/lancet/v2/validator"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/advertising"
	"github.com/flipped-aurora/gin-vue-admin/server/model/material"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system_management/api"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type data struct {
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

type dataNew struct {
	Label interface{} `json:"label"`
	Value interface{} `json:"value"`
}

type SearchService struct {
}

func (receiver *SearchService) Search(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	var dimTypes = map[string]func(ctx context.Context, req *api.SearchReq) (resp interface{}, err error){
		"game-type":              receiver.searchGameType,
		"game-os":                receiver.searchGameOs,
		"game-status":            receiver.searchGameStatus,
		"game-cooperation-model": receiver.searchGameCooperationModel,
		"company":                receiver.searchCompany,
		"platform":               receiver.searchPlatform,
		"root-game":              receiver.searchRootGame,
		"main-game":              receiver.searchMainGame,
		"sub-game":               receiver.searchSubGame,
		"product-common-config":  receiver.searchProductCommonConfig,
		"pay-type":               receiver.searchPayType,
		"pay-status":             receiver.searchPayStatus,
		"pay-channel":            receiver.searchPayChannel,
		"publishing-channel":     receiver.searchPublishingChannel,
		"agent":                  receiver.searchAgent,
		"site":                   receiver.searchSite,
		"common-media":           receiver.searchCommonMedia,
		"media":                  receiver.searchMedia,
		"channel-group":          receiver.searchChannelGroup,
		"settlement-type":        receiver.searchSettlementType,
		"material-theme":         receiver.searchMaterialTheme,
	}
	searchFun, ok := dimTypes[req.DimType]
	if !ok {
		err = errors.Wrap(error2.ErrorParamEmpty, "dim_type未知")
		return
	}
	return searchFun(ctx, req)
}

// searchGameType 搜索游戏类型
func (receiver *SearchService) searchPlatform(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := operation_management.NewDimPlatformModel()
	tmpDb := model.Db().WithContext(ctx).Select("*").Table(model.TableName())
	if req.Keyword != "" {
		if validator.IsNumberStr(req.Keyword) {
			tmpDb.Where("id = ?", req.Keyword)
		} else {
			tmpDb.Where("platform_name like ?", "%"+req.Keyword+"%")
		}
	}
	var list []operation_management.DimPlatformModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.PlatformName})
	}
	resp = respList
	return
}

// searchGameType 搜索游戏类型
func (receiver *SearchService) searchGameType(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	var list []data
	for key, item := range common.GameTypes {
		list = append(list, data{Key: key, Value: item})
	}
	resp = list
	return
}

// searchGameOs 搜索游戏系统
func (receiver *SearchService) searchGameOs(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	var list []data
	for key, item := range common.GameOss {
		list = append(list, data{Key: key, Value: item})
	}
	resp = list
	return
}

// searchGameOs 搜索游戏状态
func (receiver *SearchService) searchGameStatus(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	var list []data
	for key, item := range common.GameStatuss {
		list = append(list, data{Key: key, Value: item})
	}
	resp = list
	return
}

// searchGameCooperationModel 搜索游戏合作模式
func (receiver *SearchService) searchGameCooperationModel(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	var list []data
	for key, item := range common.CooperationModels {
		list = append(list, data{Key: key, Value: item})
	}
	resp = list
	return
}

// searchCompany 主体搜索
func (receiver *SearchService) searchCompany(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := operation_management.NewDimCompanyModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.Keyword != "" {
		if validator.IsNumberStr(req.Keyword) {
			tmpDb.Where("id = ?", req.Keyword)
		} else {
			tmpDb.Where("company_name like ?", "%"+req.Keyword+"%")
		}
	}
	var list []operation_management.DimCompanyModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.CompanyName})
	}
	resp = respList
	return
}

// searchRootGame 搜索根游戏
func (receiver *SearchService) searchRootGame(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := operation_management.NewDimRootGameModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.Keyword != "" {
		if validator.IsNumberStr(req.Keyword) {
			tmpDb.Where("id = ?", req.Keyword)
		} else {
			tmpDb.Where("game_name like ?", "%"+req.Keyword+"%")
		}
	}
	var list []operation_management.DimRootGameModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.GameName})
	}
	resp = respList
	return
}

// searchMainGame 搜索主游戏
func (receiver *SearchService) searchMainGame(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := operation_management.NewDimMainGameModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.Keyword != "" {
		if validator.IsNumberStr(req.Keyword) {
			tmpDb.Where("id = ?", req.Keyword)
		} else {
			tmpDb.Where("game_name like ?", "%"+req.Keyword+"%")
		}
	}
	var list []operation_management.DimMainGameModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.GameName})
	}
	resp = respList
	return
}

// searchSubGame 搜索子游戏
func (receiver *SearchService) searchSubGame(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := operation_management.NewDimGameModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.Keyword != "" {
		if validator.IsNumberStr(req.Keyword) {
			tmpDb.Where("id = ?", req.Keyword)
		} else {
			tmpDb.Where("game_name like ?", "%"+req.Keyword+"%")
		}
	}
	var list []operation_management.DimGameModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.GameName})
	}
	resp = respList
	return
}

// searchSubGame 搜索子游戏
func (receiver *SearchService) searchProductCommonConfig(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := operation_management.NewDimProductCommonConfigurationModel()
	tmpDb := model.Db().WithContext(ctx).Table(model.TableName())
	if req.Keyword != "" {
		if validator.IsNumberStr(req.Keyword) {
			tmpDb.Where("id = ?", req.Keyword)
		} else {
			tmpDb.Where("config_name like ?", "%"+req.Keyword+"%")
		}
	}
	var list []operation_management.DimProductCommonConfigurationModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.ConfigName})
	}
	resp = respList
	return
}

// searchPayType 搜索支付网关
func (receiver *SearchService) searchPayType(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	var respList []data
	for key, item := range common.PayTypes {
		respList = append(respList, data{Key: key, Value: item})
	}
	resp = respList
	return
}

// searchPayStatus 搜索支付状态
func (receiver *SearchService) searchPayStatus(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	var respList []data
	for key, item := range common.PayStatuss {
		respList = append(respList, data{Key: key, Value: item})
	}
	resp = respList
	return
}

// searchPayChannel 搜索支付渠道
func (receiver *SearchService) searchPayChannel(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := operation_management.NewDimPayChannelModel()
	tempDb := model.Db().WithContext(ctx).Table(model.TableName())
	var list []operation_management.DimPayChannelModel
	if !validator.IsEmptyString(req.Keyword) {
		if validator.IsNumberStr(req.Keyword) {
			tempDb.Where("id = ?", req.Keyword)
		} else {
			tempDb.Where("channel_name like ?", "%"+req.Keyword+"%")
		}
	}

	if !validator.IsEmptyString(req.PayType) {
		tempDb.Where("pay_type = ?", req.PayType)
	}

	if validator.IsEmptyString(req.Status) {
		tempDb.Where("status = ?", req.Status)
	}
	if listErr := tempDb.Limit(50).Order("id DESC").Take(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取异常", zap.Error(listErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.ChannelName})
	}
	resp = respList
	return
}

// searchPublishingChannel 搜索发行渠道
func (receiver *SearchService) searchPublishingChannel(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := operation_management.NewDimPublishingChannelConfigModel()
	tmpDb := model.Db().WithContext(ctx).Select("*")
	if req.PlatformId > 0 {
		tmpDb.Where("platform_id = ?", req.PlatformId)
	}
	if req.Keyword != "" {
		tmpDb.Where(" id = ? or channel_name like ?", req.Keyword, "%"+req.Keyword+"%")
	}
	var list []operation_management.DimPublishingChannelConfigModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.ChannelName})
	}
	resp = respList
	return
}

// searchAgent 搜索渠道
func (receiver *SearchService) searchAgent(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := advertising.NewDimAgentModel()
	tmpDb := model.Db().WithContext(ctx).Select("*")
	if req.PlatformId > 0 {
		tmpDb.Where("platform_id = ?", req.PlatformId)
	}
	if req.Keyword != "" {
		tmpDb.Where(" id = ? or agent_name like ?", req.Keyword, "%"+req.Keyword+"%")
	}
	var list []advertising.DimAgentModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.AgentName})
	}
	resp = respList
	return
}

// searchAgent 搜索渠道
func (receiver *SearchService) searchSite(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := advertising.NewDimSiteModel()
	tmpDb := model.Db().WithContext(ctx).Select("*")
	if req.PlatformId > 0 {
		tmpDb.Where("platform_id = ?", req.PlatformId)
	}
	if req.Keyword != "" {
		tmpDb.Where(" id = ? or site_name like ?", req.Keyword, "%"+req.Keyword+"%")
	}
	var list []advertising.DimSiteModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.SiteName})
	}
	resp = respList
	return
}

// searchAgent 搜索渠道
func (receiver *SearchService) searchCommonMedia(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	var respList []data
	for key, item := range advertising2.MediaCodesMap {
		respList = append(respList, data{Key: key, Value: item})
	}
	resp = respList
	return
}

func (receiver *SearchService) searchMedia(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := advertising.NewDimAdvertisingMediaModel()
	tmpDb := model.Db().WithContext(ctx).Select("*")
	if req.PlatformId > 0 {
		tmpDb.Where("platform_id = ?", req.PlatformId)
	}
	if req.Keyword != "" {
		tmpDb.Where(" id = ? or media_name like ?", req.Keyword, "%"+req.Keyword+"%")
	}
	var list []advertising.DimAdvertisingMediaModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.MediaName})
	}
	resp = respList
	return
}

func (receiver *SearchService) searchChannelGroup(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := advertising.NewDimChannelGroupModel()
	tmpDb := model.Db().WithContext(ctx).Select("*")
	if req.PlatformId > 0 {
		tmpDb.Where("platform_id = ?", req.PlatformId)
	}
	if req.Keyword != "" {
		tmpDb.Where(" id = ? or channel_group_name like ?", req.Keyword, "%"+req.Keyword+"%")
	}
	var list []advertising.DimChannelGroupModel
	if findErr := tmpDb.Limit(50).Find(&list).Error; findErr != nil {
		err = findErr
		global.GVA_LOG.Error("获取异常", zap.Error(findErr))
		return
	}
	var respList []data
	for _, item := range list {
		respList = append(respList, data{Key: item.Id, Value: item.ChannelGroupName})
	}
	resp = respList
	return
}

func (receiver *SearchService) searchSettlementType(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	var respList []data
	for key, item := range common.SettlementTypes {
		respList = append(respList, data{Key: key, Value: item})
	}
	resp = respList
	return
}

// 题材
func (receiver *SearchService) searchMaterialTheme(ctx context.Context, req *api.SearchReq) (resp interface{}, err error) {
	model := material.NewDimMaterialThemeModel()
	tmpDb := model.Db().
		WithContext(ctx).
		Select("id,theme_name,parent_id").
		Where("platform_id = ?", req.PlatformId)
	if req.Keyword != "" {
		tmpDb.Where("theme_name like ?", "%"+req.Keyword+"%")
	}

	var list []material.DimMaterialThemeModel
	if listErr := tmpDb.Order("id DESC").Find(&list).Error; listErr != nil {
		err = listErr
		global.GVA_LOG.Error("获取异常", zap.Error(listErr))
		return
	}

	type formatMaterialThemeStruct struct {
		dataNew
		Children []dataNew `json:"children"`
	}

	var formatData []formatMaterialThemeStruct

	//先找父节点
	for _, item := range list {
		if item.ParentId == 0 {
			myFormat := formatMaterialThemeStruct{}
			myFormat.Label = item.ThemeName
			myFormat.Value = item.Id
			formatData = append(formatData, myFormat)
		}
	}

	for index, item := range formatData {
		for _, itemList := range list {
			if cast.ToInt64(item.Value) == itemList.ParentId {
				formatData[index].Children = append(formatData[index].Children, dataNew{Value: itemList.Id, Label: itemList.ThemeName})
			}
		}
	}
	resp = formatData
	return
}
