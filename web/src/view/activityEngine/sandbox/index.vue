<template>
  <div class="sandbox-container">
    <el-card header="沙箱测试">
      <el-form :model="form" label-width="120px">
        <el-form-item label="活动ID" required>
          <el-input-number v-model="form.activity_id" :min="1" controls-position="right" />
        </el-form-item>
        <el-form-item label="用户ID" required>
          <el-input-number v-model="form.user_id" :min="1" controls-position="right" />
        </el-form-item>
        <el-form-item label="事件类型" required>
          <el-select v-model="form.event_type" placeholder="请选择">
            <el-option label="充值" value="recharge" />
            <el-option label="登录" value="login" />
            <el-option label="签到" value="signin" />
            <el-option label="分享" value="share" />
            <el-option label="自定义" value="custom" />
          </el-select>
        </el-form-item>
        <el-form-item label="事件数据">
          <el-input v-model="eventDataStr" type="textarea" :rows="4" placeholder='{"amount": 100, "channel": "wechat"}' />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSimulate" :loading="loading">模拟执行</el-button>
        </el-form-item>
      </el-form>
    </el-card>

    <el-card header="执行结果" style="margin-top: 20px" v-if="result">
      <el-descriptions :column="1" border>
        <el-descriptions-item label="状态">
          <el-tag :type="result.code === 0 ? 'success' : 'danger'">{{ result.code === 0 ? '成功' : '失败' }}</el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="消息">{{ result.msg }}</el-descriptions-item>
      </el-descriptions>
      <el-divider content-position="left">返回数据</el-divider>
      <el-input type="textarea" :rows="10" :model-value="JSON.stringify(result.data, null, 2)" readonly />
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { sandboxSimulate } from '@/api/activityEngine'

const form = ref({ activity_id: null, user_id: null, event_type: '', event_data: {} })
const eventDataStr = ref('')
const loading = ref(false)
const result = ref(null)

const handleSimulate = async () => {
  if (!form.value.activity_id || !form.value.user_id || !form.value.event_type) {
    ElMessage.warning('请填写必填项')
    return
  }
  try {
    form.value.event_data = eventDataStr.value ? JSON.parse(eventDataStr.value) : {}
  } catch {
    ElMessage.error('事件数据JSON格式错误')
    return
  }
  loading.value = true
  try {
    const res = await sandboxSimulate(form.value)
    result.value = res
  } finally {
    loading.value = false
  }
}
</script>
