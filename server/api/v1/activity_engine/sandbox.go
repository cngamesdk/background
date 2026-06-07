package activity_engine

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/activity_engine/api"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type SandboxApi struct{}

func (a *SandboxApi) Simulate(ctx *gin.Context) {
	var req api.SandboxSimulateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	// 调用活动引擎运行时API进行模拟
	engineURL := global.GVA_CONFIG.Common.Endpoint
	if engineURL == "" {
		engineURL = "http://127.0.0.1:9090"
	}

	// 构造请求体
	eventPayload := map[string]interface{}{
		"platform_id": 1,
		"game_id":     0,
		"user_id":     req.UserID,
		"event_type":  req.EventType,
		"event_data":  req.EventData,
		"timestamp":   time.Now().Unix(),
	}
	body, _ := json.Marshal(eventPayload)

	// 调用引擎 /api/v1/event
	resp, err := http.Post(engineURL+"/api/v1/event", "application/json",
		newReaderFromBytes(body))
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("调用引擎失败: %v", err), ctx)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	response.OkWithData(result, ctx)
}

func newReaderFromBytes(b []byte) *bytesReader {
	return &bytesReader{data: b, pos: 0}
}

type bytesReader struct {
	data []byte
	pos  int
}

func (r *bytesReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, fmt.Errorf("EOF")
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}

// ValidateConfig 校验活动配置JSON合法性
func ValidateConfig(ctx context.Context, triggerCfg, calcCfg, rewardCfg, constraintCfg string) []string {
	var errors []string

	// 校验触发条件
	var trigger map[string]interface{}
	if err := json.Unmarshal([]byte(triggerCfg), &trigger); err != nil {
		errors = append(errors, "触发条件JSON格式错误: "+err.Error())
	} else {
		if _, ok := trigger["event_type"]; !ok {
			errors = append(errors, "触发条件缺少event_type字段")
		}
	}

	// 校验计算逻辑
	var calc map[string]interface{}
	if err := json.Unmarshal([]byte(calcCfg), &calc); err != nil {
		errors = append(errors, "计算逻辑JSON格式错误: "+err.Error())
	} else {
		mode, _ := calc["mode"].(string)
		validModes := map[string]bool{"accumulate": true, "daily_reset": true, "dedup": true, "time_limited": true}
		if !validModes[mode] {
			errors = append(errors, "计算逻辑mode无效，可选: accumulate/daily_reset/dedup/time_limited")
		}
	}

	// 校验奖励策略
	var reward map[string]interface{}
	if err := json.Unmarshal([]byte(rewardCfg), &reward); err != nil {
		errors = append(errors, "奖励策略JSON格式错误: "+err.Error())
	} else {
		strategy, _ := reward["strategy"].(string)
		validStrategies := map[string]bool{"fixed": true, "tiered": true, "probability": true}
		if !validStrategies[strategy] {
			errors = append(errors, "奖励策略strategy无效，可选: fixed/tiered/probability")
		}
		tiers, _ := reward["tiers"].([]interface{})
		if len(tiers) == 0 {
			errors = append(errors, "奖励策略至少需要一个档位")
		}
	}

	// 校验约束规则
	var constraint map[string]interface{}
	if err := json.Unmarshal([]byte(constraintCfg), &constraint); err != nil {
		errors = append(errors, "约束规则JSON格式错误: "+err.Error())
	}

	return errors
}
