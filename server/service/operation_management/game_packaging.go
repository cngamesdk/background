package operation_management

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/model/operation_management/api"
)

type GamePackagingService struct {
}

func (g *GamePackagingService) LogList(ctx context.Context, req *api.GamePackagingLogListReq) (
	resp interface{}, total int64, err error) {
	return
}
