package api

import (
	"github.com/cngamesdk/go-core/validate"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/cron_task"
	"github.com/pkg/errors"
	"strings"
)

type CronTaskListReq struct {
	cron_task.DimCronTaskConfigModel
	request.PageInfo
}

type CronTaskListResp struct {
}

type CronTaskAddReq struct {
	cron_task.DimCronTaskConfigModel
}

func (a *CronTaskAddReq) Format() {
	a.Id = 0
	a.Name = strings.TrimSpace(a.Name)
	a.Spec = strings.TrimSpace(a.Spec)
	a.Remark = strings.TrimSpace(a.Remark)
	a.Status = strings.TrimSpace(a.Status)
	a.Content = strings.TrimSpace(a.Content)
	a.Config = make(map[string]interface{})
	a.TaskType = strings.TrimSpace(a.TaskType)
}

func (a *CronTaskAddReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.Name); validateErr != nil {
		err = errors.Wrap(validateErr, "name")
		return
	}
	if validateErr := validate.EmptyString(a.Spec); validateErr != nil {
		err = errors.Wrap(validateErr, "spec")
		return
	}
	return
}

type CronTaskAddResp struct {
}

type CronTaskModifyReq struct {
	cron_task.DimCronTaskConfigModel
}

func (a *CronTaskModifyReq) Format() {
	a.Name = strings.TrimSpace(a.Name)
	a.Spec = strings.TrimSpace(a.Spec)
	a.Remark = strings.TrimSpace(a.Remark)
	a.Status = strings.TrimSpace(a.Status)
	a.Content = strings.TrimSpace(a.Content)
	a.TaskType = strings.TrimSpace(a.TaskType)
	a.Config = nil // 不需要更新
}

func (a *CronTaskModifyReq) Validate() (err error) {
	if validateErr := validate.EmptyString(a.Name); validateErr != nil {
		err = errors.Wrap(validateErr, "name")
		return
	}
	if validateErr := validate.EmptyString(a.Spec); validateErr != nil {
		err = errors.Wrap(validateErr, "spec")
		return
	}
	return
}

type CronTaskModifyResp struct {
}

type CronTaskLogListReq struct {
	cron_task.OdsCronTaskLogModel
	request.PageInfo
}
