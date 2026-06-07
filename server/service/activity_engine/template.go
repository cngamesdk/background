package activity_engine

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine"
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine/api"
)

type TemplateService struct{}

func (s *TemplateService) List(ctx context.Context, req *api.TemplateListReq) (list []activity_engine.OdsActivityTemplate, total int64, err error) {
	db := global.GVA_DB.WithContext(ctx).Model(&activity_engine.OdsActivityTemplate{})
	if req.TemplateName != "" {
		db = db.Where("template_name LIKE ?", "%"+req.TemplateName+"%")
	}
	if req.ActivityType != "" {
		db = db.Where("activity_type = ?", req.ActivityType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id DESC").Scopes(req.Paginate()).Find(&list).Error
	return
}

func (s *TemplateService) Add(ctx context.Context, req *api.TemplateAddReq) error {
	return global.GVA_DB.WithContext(ctx).Create(&req.OdsActivityTemplate).Error
}

func (s *TemplateService) Clone(ctx context.Context, req *api.TemplateCloneReq) error {
	var tpl activity_engine.OdsActivityTemplate
	if err := global.GVA_DB.WithContext(ctx).Where("id = ?", req.TemplateID).First(&tpl).Error; err != nil {
		return err
	}
	activity := req.OdsActivityConfig
	activity.TriggerConfig = tpl.TriggerConfig
	activity.CalculationConfig = tpl.CalculationConfig
	activity.RewardConfig = tpl.RewardConfig
	activity.ConstraintConfig = tpl.ConstraintConfig
	activity.ActivityType = tpl.ActivityType
	activity.Status = "not-started"
	activity.Version = 1
	return global.GVA_DB.WithContext(ctx).Create(&activity).Error
}
