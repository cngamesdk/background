package material

import api2 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	MaterialRouter
}

var (
	themeApi = api2.ApiGroupApp.MaterialApiGroup.ThemeApi
)
