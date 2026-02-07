package material

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ThemeApi
	MaterialApi
}

var (
	materialThemeService = service.ServiceGroupApp.MaterialGroup.ThemeService
	materialService      = service.ServiceGroupApp.MaterialGroup.MaterialService
)
