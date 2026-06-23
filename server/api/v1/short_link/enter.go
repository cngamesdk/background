package short_link

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ShortLinkApi
}

var (
	shortLinkService = service.ServiceGroupApp.ShortLinkServiceGroup.ShortLinkService
)
