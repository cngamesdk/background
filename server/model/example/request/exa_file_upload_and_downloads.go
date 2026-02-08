package request

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	time2 "time"
)

// 业务对应的存储目录映射
var BizStoreDirMap = map[string]func(req UploadFileReq) string{
	"material": func(req UploadFileReq) string {
		time := time2.Now()
		return fmt.Sprintf("material/platform-%d/%s/%s/%s",
			req.PlatformId,
			time.Format("2006"),
			time.Format("01"),
			time.Format("02"))
	},
}

type ExaAttachmentCategorySearch struct {
	ClassId int `json:"classId" form:"classId"`
	request.PageInfo
}

type UploadFileReq struct {
	PlatformId int64  `json:"platform_id"`
	Biz        string `json:"biz"`
}
