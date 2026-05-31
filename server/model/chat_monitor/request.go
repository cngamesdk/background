package chat_monitor

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

// GameSearch 游戏列表查询
type GameSearch struct {
	Name   string `json:"name" form:"name"`
	Status *int   `json:"status" form:"status"`
	request.PageInfo
}

// ChatMessageSearch 聊天记录查询
type ChatMessageSearch struct {
	AppID         string `json:"appId" form:"appId"`
	Channel       string `json:"channel" form:"channel"`
	SenderUID     string `json:"senderUid" form:"senderUid"`
	Keyword       string `json:"keyword" form:"keyword"`
	StartTime     string `json:"startTime" form:"startTime"`
	EndTime       string `json:"endTime" form:"endTime"`
	OnlySensitive bool   `json:"onlySensitive" form:"onlySensitive"`
	request.PageInfo
}

// SensitiveWordSearch 敏感词查询
type SensitiveWordSearch struct {
	AppID    string `json:"appId" form:"appId"`
	Category string `json:"category" form:"category"`
	request.PageInfo
}

// WhitelistSearch 白名单查询
type WhitelistSearch struct {
	AppID string `json:"appId" form:"appId"`
	request.PageInfo
}

// BanRecordSearch 封禁记录查询
type BanRecordSearch struct {
	AppID   string `json:"appId" form:"appId"`
	BanType *int   `json:"banType" form:"banType"`
	Status  *int   `json:"status" form:"status"`
	request.PageInfo
}

// BanCreateReq 创建封禁请求
type BanCreateReq struct {
	AppID    string `json:"appId" binding:"required"`
	BanType  int    `json:"banType" binding:"required"`
	Target   string `json:"target" binding:"required"`
	Reason   string `json:"reason"`
	Duration int    `json:"duration"`
}

// SensitiveImportReq 批量导入敏感词
type SensitiveImportReq struct {
	Words    []string `json:"words" binding:"required"`
	Category string   `json:"category"`
	Level    int      `json:"level"`
	AppID    string   `json:"appId"`
}

// StatsOverviewReq 统计概览请求
type StatsOverviewReq struct {
	AppID string `json:"appId" form:"appId"`
}

// StatsTrendReq 趋势请求
type StatsTrendReq struct {
	AppID string `json:"appId" form:"appId"`
	Days  int    `json:"days" form:"days"`
}

// ViolatorsReq 违规用户请求
type ViolatorsReq struct {
	AppID string `json:"appId" form:"appId"`
	Limit int    `json:"limit" form:"limit"`
}
