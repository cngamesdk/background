package api

import (
	"github.com/cngamesdk/go-core/validate"
	"github.com/pkg/errors"
	"strings"
)

type SearchReq struct {
	PlatformId int64 `json:"platform_id" form:"platform_id"`
	DimType string `json:"dim_type" form:"dim_type" binding:"required"`
	Keyword string `json:"keyword" form:"keyword"`
}

func (receiver *SearchReq) Format() {
	receiver.DimType = strings.TrimSpace(receiver.DimType)
	receiver.Keyword = strings.TrimSpace(receiver.Keyword)
}

func (receiver *SearchReq) Validate() (err error) {
	if validateErr := validate.EmptyString(receiver.DimType); validateErr != nil {
		err = errors.Wrap(validateErr, "dim_type")
		return
	}
	return
}

type SearchResp struct {

}
