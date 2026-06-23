package short_link

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/short_link"
	"github.com/flipped-aurora/gin-vue-admin/server/model/short_link/api"
	"go.uber.org/zap"
)

type ShortLinkService struct{}

// 雪花算法（轻量内嵌）
var (
	snowOnce   sync.Once
	snowflake  *snowflakeGen
)

const (
	sfEpoch         = int64(1700000000000)
	sfMachineBits   = 10
	sfSequenceBits  = 12
	sfMaxSequence   = -1 ^ (-1 << sfSequenceBits)
	sfMachineShift  = sfSequenceBits
	sfTimestampShift = sfMachineBits + sfSequenceBits
)

type snowflakeGen struct {
	mu        sync.Mutex
	timestamp int64
	machineID int64
	sequence  int64
}

func getSnowflake() *snowflakeGen {
	snowOnce.Do(func() {
		snowflake = &snowflakeGen{machineID: 2} // background使用machineID=2区分
	})
	return snowflake
}

func (s *snowflakeGen) generate() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UnixMilli() - sfEpoch
	if now == s.timestamp {
		s.sequence = (s.sequence + 1) & sfMaxSequence
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixMilli() - sfEpoch
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now
	return now<<sfTimestampShift | s.machineID<<sfMachineShift | s.sequence
}

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func base62Encode(num int64) string {
	if num == 0 {
		return "0"
	}
	var result []byte
	for num > 0 {
		result = append([]byte{base62Chars[num%62]}, result...)
		num /= 62
	}
	return string(result)
}

func generateShortCode() string {
	id := getSnowflake().generate()
	code := base62Encode(id)
	if len(code) > 8 {
		code = code[len(code)-8:]
	}
	return code
}

func (s *ShortLinkService) Create(ctx context.Context, req *api.ShortLinkCreateReq, creator string) (resp api.ShortLinkCreateResp, err error) {
	shortCode := generateShortCode()
	domain := "http://127.0.0.1:8090" // 默认域名，实际从配置读取
	// 尝试从已有记录获取域名
	var existing short_link.DimShortLink
	if global.GVA_DB.WithContext(ctx).First(&existing).Error == nil && existing.Domain != "" {
		domain = existing.Domain
	}

	record := short_link.DimShortLink{
		ShortCode:   shortCode,
		OriginalUrl: req.OriginalUrl,
		Domain:      domain,
		Title:       req.Title,
		Status:      1,
		Creator:     creator,
	}

	if req.ExpireDays > 0 {
		expire := time.Now().AddDate(0, 0, req.ExpireDays)
		record.ExpireAt = &expire
	}

	if err = global.GVA_DB.WithContext(ctx).Create(&record).Error; err != nil {
		global.GVA_LOG.Error("创建短链接失败", zap.Error(err))
		return
	}

	resp.ShortCode = shortCode
	resp.ShortUrl = fmt.Sprintf("%s/%s", domain, shortCode)
	return
}

func (s *ShortLinkService) List(ctx context.Context, req *api.ShortLinkListReq) (list []short_link.DimShortLink, total int64, err error) {
	db := global.GVA_DB.WithContext(ctx).Model(&short_link.DimShortLink{})
	if req.Title != "" {
		db = db.Where("title LIKE ?", "%"+req.Title+"%")
	}
	if req.ShortCode != "" {
		db = db.Where("short_code = ?", req.ShortCode)
	}
	if req.Status != 0 {
		db = db.Where("status = ?", req.Status)
	}

	if err = db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("获取总数异常", zap.Error(err))
		return
	}

	if err = db.Order("id DESC").
		Scopes(req.Paginate()).
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&list).Error; err != nil {
		global.GVA_LOG.Error("获取列表异常", zap.Error(err))
		return
	}
	return
}

func (s *ShortLinkService) Detail(ctx context.Context, id int64) (record short_link.DimShortLink, err error) {
	if err = global.GVA_DB.WithContext(ctx).Where("id = ?", id).First(&record).Error; err != nil {
		global.GVA_LOG.Error("获取详情异常", zap.Error(err))
	}
	return
}

func (s *ShortLinkService) Update(ctx context.Context, req *api.ShortLinkUpdateReq) (err error) {
	updates := map[string]interface{}{}
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if err = global.GVA_DB.WithContext(ctx).Model(&short_link.DimShortLink{}).
		Where("id = ?", req.Id).Updates(updates).Error; err != nil {
		global.GVA_LOG.Error("更新异常", zap.Error(err))
	}
	return
}

func (s *ShortLinkService) Delete(ctx context.Context, id int64) (err error) {
	if err = global.GVA_DB.WithContext(ctx).Delete(&short_link.DimShortLink{}, "id = ?", id).Error; err != nil {
		global.GVA_LOG.Error("删除异常", zap.Error(err))
	}
	return
}

func (s *ShortLinkService) ClickLogList(ctx context.Context, req *api.ClickLogListReq) (list []short_link.OdsClickLog, total int64, err error) {
	db := global.GVA_DB.WithContext(ctx).Model(&short_link.OdsClickLog{}).
		Where("short_code = ?", req.ShortCode)

	if err = db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("获取点击日志总数异常", zap.Error(err))
		return
	}

	if err = db.Order("id DESC").
		Limit(req.PageSize).
		Offset((req.Page - 1) * req.PageSize).
		Find(&list).Error; err != nil {
		global.GVA_LOG.Error("获取点击日志列表异常", zap.Error(err))
		return
	}
	return
}
