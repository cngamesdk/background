package material

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	ThemeApi
}

var (
	materialThemeService = service.ServiceGroupApp.MaterialGroup.ThemeService
)
