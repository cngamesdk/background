package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/model_transfer"
	"time"
)

// TokenCreateReq 创建Token请求
type TokenCreateReq struct {
	Name          string     `json:"name" binding:"required"`
	Type          int8       `json:"type" binding:"required,oneof=1 2"` // 1-企业 2-个人
	TokenLimit    int64      `json:"tokenLimit"`
	RequestLimit  int        `json:"requestLimit"`
	ExpireAt      *time.Time `json:"expireAt"`
	AllowedModels string     `json:"allowedModels"` // JSON字符串
	IPWhitelist   string     `json:"ipWhitelist"`   // JSON字符串
	Creator       string     `json:"creator"`
}

// TokenUpdateReq 更新Token请求
type TokenUpdateReq struct {
	ID            int64      `json:"id" binding:"required"`
	Name          string     `json:"name"`
	TokenLimit    int64      `json:"tokenLimit"`
	RequestLimit  int        `json:"requestLimit"`
	ExpireAt      *time.Time `json:"expireAt"`
	Status        int8       `json:"status" binding:"oneof=1 2"` // 1-启用 2-禁用
	AllowedModels string     `json:"allowedModels"`
	IPWhitelist   string     `json:"ipWhitelist"`
}

// TokenListReq Token列表请求
type TokenListReq struct {
	request.PageInfo
	Name   string `json:"name" form:"name"`
	Type   int8   `json:"type" form:"type"`
	Status int8   `json:"status" form:"status"`
}

// TokenListResp Token列表响应
type TokenListResp struct {
	model_transfer.AiToken
	UsagePercent float64 `json:"usagePercent" gorm:"-"` // 使用率
}

// TokenDetailResp Token详情响应
type TokenDetailResp struct {
	model_transfer.AiToken
	TodayRequests int64 `json:"todayRequests" gorm:"-"` // 今日请求数
	TodayTokens   int64 `json:"todayTokens" gorm:"-"`   // 今日Token数
	UsagePercent  float64 `json:"usagePercent" gorm:"-"` // 使用率
}

// TokenRegenerateReq 重新生成Token请求
type TokenRegenerateReq struct {
	ID int64 `json:"id" binding:"required"`
}

// TokenRegenerateResp 重新生成Token响应
type TokenRegenerateResp struct {
	Token string `json:"token"`
}
