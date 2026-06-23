package short_link

import api2 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	ShortLinkRouter
}

var (
	shortLinkApi = api2.ApiGroupApp.ShortLinkApiGroup.ShortLinkApi
)
