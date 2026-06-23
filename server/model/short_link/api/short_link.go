package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/short_link"
)

type ShortLinkListReq struct {
	short_link.DimShortLink
	request.PageInfo
}

type ShortLinkDetailReq struct {
	Id int64 `json:"id" binding:"required"`
}

type ShortLinkCreateReq struct {
	OriginalUrl string `json:"original_url" binding:"required"`
	Title       string `json:"title"`
	ExpireDays  int    `json:"expire_days"`
}

type ShortLinkCreateResp struct {
	ShortCode string `json:"short_code"`
	ShortUrl  string `json:"short_url"`
}

type ShortLinkUpdateReq struct {
	Id     int64  `json:"id" binding:"required"`
	Title  string `json:"title"`
	Status *int8  `json:"status"`
}

type ClickLogListReq struct {
	ShortCode string `json:"short_code" binding:"required"`
	request.PageInfo
}
