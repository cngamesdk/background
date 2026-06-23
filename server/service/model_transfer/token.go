package model_transfer

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer/api"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TokenService struct{}

// Create 创建Token
func (s *TokenService) Create(ctx context.Context, req *api.TokenCreateReq) error {
	// 生成Token
	token := uuid.New().String()

	aiToken := &model_transfer.AiToken{
		Token:         token,
		Name:          req.Name,
		Type:          req.Type,
		TokenLimit:    req.TokenLimit,
		RequestLimit:  req.RequestLimit,
		ExpireAt:      req.ExpireAt,
		Status:        1, // 默认启用
		AllowedModels: req.AllowedModels,
		IPWhitelist:   req.IPWhitelist,
		Creator:       req.Creator,
	}

	if err := global.GVA_DB.WithContext(ctx).Create(aiToken).Error; err != nil {
		global.GVA_LOG.Error("创建Token失败", zap.Error(err))
		return err
	}

	return nil
}

// Update 更新Token
func (s *TokenService) Update(ctx context.Context, req *api.TokenUpdateReq) error {
	updates := make(map[string]interface{})

	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.TokenLimit >= 0 {
		updates["token_limit"] = req.TokenLimit
	}
	if req.RequestLimit >= 0 {
		updates["request_limit"] = req.RequestLimit
	}
	if req.ExpireAt != nil {
		updates["expire_at"] = req.ExpireAt
	}
	if req.Status > 0 {
		updates["status"] = req.Status
	}
	if req.AllowedModels != "" {
		updates["allowed_models"] = req.AllowedModels
	}
	if req.IPWhitelist != "" {
		updates["ip_whitelist"] = req.IPWhitelist
	}

	if len(updates) == 0 {
		return errors.New("没有要更新的字段")
	}

	if err := global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiToken{}).
		Where("id = ?", req.ID).
		Updates(updates).Error; err != nil {
		global.GVA_LOG.Error("更新Token失败", zap.Error(err))
		return err
	}

	return nil
}

// Delete 删除Token（软删除，实际是禁用）
func (s *TokenService) Delete(ctx context.Context, id int64) error {
	if err := global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiToken{}).
		Where("id = ?", id).
		Update("status", 2).Error; err != nil {
		global.GVA_LOG.Error("删除Token失败", zap.Error(err))
		return err
	}

	return nil
}

// List Token列表
func (s *TokenService) List(ctx context.Context, req *api.TokenListReq) ([]api.TokenListResp, int64, error) {
	var list []api.TokenListResp
	var total int64

	db := global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiToken{})

	// 筛选条件
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Type > 0 {
		db = db.Where("type = ?", req.Type)
	}
	if req.Status > 0 {
		db = db.Where("status = ?", req.Status)
	}

	// 获取总数
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("获取Token总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页查询
	if err := db.Scopes(func(d *gorm.DB) *gorm.DB {
		offset := (req.Page - 1) * req.PageSize
		return d.Offset(offset).Limit(req.PageSize)
	}).Order("id DESC").Find(&list).Error; err != nil {
		global.GVA_LOG.Error("查询Token列表失败", zap.Error(err))
		return nil, 0, err
	}

	// 计算使用率
	for i := range list {
		if list[i].TokenLimit > 0 {
			list[i].UsagePercent = float64(list[i].UsedTokens) * 100.0 / float64(list[i].TokenLimit)
		}
	}

	return list, total, nil
}

// Detail Token详情
func (s *TokenService) Detail(ctx context.Context, id int64) (*api.TokenDetailResp, error) {
	var detail api.TokenDetailResp

	if err := global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiToken{}).
		Where("id = ?", id).
		First(&detail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Token不存在")
		}
		global.GVA_LOG.Error("查询Token详情失败", zap.Error(err))
		return nil, err
	}

	// 查询今日统计
	var todayStats struct {
		Requests int64
		Tokens   int64
	}
	global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiUsageLog{}).
		Select("COUNT(*) as requests, COALESCE(SUM(total_tokens), 0) as tokens").
		Where("token_id = ? AND DATE(request_time) = CURDATE()", id).
		Scan(&todayStats)

	detail.TodayRequests = todayStats.Requests
	detail.TodayTokens = todayStats.Tokens

	// 计算使用率
	if detail.TokenLimit > 0 {
		detail.UsagePercent = float64(detail.UsedTokens) * 100.0 / float64(detail.TokenLimit)
	}

	return &detail, nil
}

// Regenerate 重新生成Token
func (s *TokenService) Regenerate(ctx context.Context, req *api.TokenRegenerateReq) (string, error) {
	newToken := uuid.New().String()

	if err := global.GVA_DB.WithContext(ctx).Model(&model_transfer.AiToken{}).
		Where("id = ?", req.ID).
		Update("token", newToken).Error; err != nil {
		global.GVA_LOG.Error("重新生成Token失败", zap.Error(err))
		return "", err
	}

	return newToken, nil
}
