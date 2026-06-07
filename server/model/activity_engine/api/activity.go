package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ActivityListReq struct {
	activity_engine.OdsActivityConfig
	request.PageInfo
}

type ActivityAddReq struct {
	activity_engine.OdsActivityConfig
}

type ActivityModifyReq struct {
	activity_engine.OdsActivityConfig
}

type ActivityDetailReq struct {
	ID int64 `json:"id" binding:"required"`
}

type ActivityPublishReq struct {
	ID int64 `json:"id" binding:"required"`
}

type ActivityOfflineReq struct {
	ID int64 `json:"id" binding:"required"`
}

type TemplateListReq struct {
	activity_engine.OdsActivityTemplate
	request.PageInfo
}

type TemplateAddReq struct {
	activity_engine.OdsActivityTemplate
}

type TemplateCloneReq struct {
	TemplateID int64 `json:"template_id" binding:"required"`
	activity_engine.OdsActivityConfig
}

type RewardItemSearchReq struct {
	PlatformID int64  `json:"platform_id"`
	Keyword    string `json:"keyword"`
	request.PageInfo
}

type GrayscaleUpdateReq struct {
	ID             int64 `json:"id" binding:"required"`
	GrayscaleRatio int   `json:"grayscale_ratio" binding:"required,min=1,max=100"`
}

type SandboxSimulateReq struct {
	ActivityID int64                  `json:"activity_id" binding:"required"`
	UserID     int64                  `json:"user_id" binding:"required"`
	EventType  string                 `json:"event_type" binding:"required"`
	EventData  map[string]interface{} `json:"event_data"`
}
