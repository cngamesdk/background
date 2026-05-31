<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="游戏">
          <el-select v-model="searchInfo.appId" clearable placeholder="全部游戏" style="width: 180px" @change="getList">
            <el-option v-for="g in gameList" :key="g.appId" :label="g.name" :value="g.appId" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-switch v-model="onlySensitive" active-text="仅敏感消息" @change="getList" />
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="chat-stream">
        <div v-for="msg in messages" :key="msg.ID" class="chat-item" :class="{ sensitive: msg.isSensitive }">
          <div class="chat-meta">
            <el-tag size="small">{{ msg.appId }}</el-tag>
            <span class="channel">[{{ msg.channel || '全服' }}]</span>
            <span class="sender">{{ msg.senderName || msg.senderUid }}</span>
            <span class="time">{{ msg.sentAt }}</span>
            <span class="ip">{{ msg.senderIp }}</span>
          </div>
          <div class="chat-content">
            <span v-if="msg.isSensitive" class="hit-words">[{{ msg.hitWords }}]</span>
            {{ msg.content }}
          </div>
          <div class="chat-actions">
            <el-button size="small" type="danger" @click="handleBan(msg, 1)">封号</el-button>
            <el-button size="small" type="warning" @click="handleBan(msg, 2)">禁言</el-button>
            <el-button size="small" type="info" @click="handleBan(msg, 3)">封IP</el-button>
          </div>
        </div>
        <div v-if="messages.length === 0" class="empty">暂无实时消息，请确认 api-server SSE 服务已启动</div>
      </div>
    </div>

    <el-dialog v-model="banDialog.visible" title="执行封禁" width="400px">
      <el-form :model="banDialog.form" label-width="80px">
        <el-form-item label="类型">
          <el-tag>{{ { 1: '账号封禁', 2: '角色禁言', 3: 'IP封禁' }[banDialog.form.banType] }}</el-tag>
        </el-form-item>
        <el-form-item label="目标">
          <el-input v-model="banDialog.form.target" :placeholder="targetPlaceholder" />
        </el-form-item>
        <el-form-item label="时长">
          <el-select v-model="banDialog.form.duration" style="width: 100%">
            <el-option label="1小时" :value="3600" />
            <el-option label="6小时" :value="21600" />
            <el-option label="1天" :value="86400" />
            <el-option label="7天" :value="604800" />
            <el-option label="永久" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item label="原因">
          <el-input v-model="banDialog.form.reason" type="textarea" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="banDialog.visible = false">取消</el-button>
        <el-button type="primary" @click="submitBan">确认封禁</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, onUnmounted } from 'vue'
import { getGameList, createBan } from '@/api/chatMonitor'
import { ElMessage } from 'element-plus'

defineOptions({ name: 'ChatMonitorRealtime' })

const messages = ref([])
const gameList = ref([])
const searchInfo = reactive({ appId: '' })
const onlySensitive = ref(false)
let eventSource = null

const banDialog = reactive({
  visible: false,
  form: { appId: '', banType: 1, target: '', duration: 3600, reason: '' },
})

const targetPlaceholder = computed(() => {
  const map = { 1: '请输入账号ID', 2: '请输入角色ID', 3: '请输入IP地址' }
  return map[banDialog.form.banType] || '请输入目标'
})

const getList = async () => {
  const res = await getGameList({ page: 1, pageSize: 100 })
  if (res.code === 0) gameList.value = res.data.list || []
}

const connectSSE = () => {
  const params = searchInfo.appId ? `?app_ids=${searchInfo.appId}` : ''
  // SSE 连接到 api-server
  eventSource = new EventSource(`http://localhost:8080/internal/sse/stream${params}`)
  eventSource.addEventListener('message', (e) => {
    try {
      const event = JSON.parse(e.data)
      const raw = event.data
      // 统一字段名（兼容下划线和驼峰）
      const msg = {
        ID: raw.ID || raw.id,
        appId: raw.appId || raw.app_id,
        channel: raw.channel,
        senderUid: raw.senderUid || raw.sender_uid,
        senderName: raw.senderName || raw.sender_name,
        senderIp: raw.senderIp || raw.sender_ip,
        roleId: raw.roleId || raw.role_id,
        content: raw.content,
        isSensitive: raw.isSensitive || raw.is_sensitive,
        hitWords: raw.hitWords || raw.hit_words,
        riskLevel: raw.riskLevel || raw.risk_level,
        sentAt: raw.sentAt || raw.sent_at,
      }
      if (onlySensitive.value && !msg.isSensitive) return
      messages.value.push(msg)
      if (messages.value.length > 500) messages.value = messages.value.slice(-500)
    } catch (err) { /* ignore */ }
  })
  eventSource.onerror = () => {
    eventSource.close()
    setTimeout(connectSSE, 5000)
  }
}

const handleBan = (msg, banType) => {
  const uid = msg.senderUid || msg.sender_uid || ''
  const roleId = msg.roleId || msg.role_id || ''
  const ip = msg.senderIp || msg.sender_ip || ''
  const targetMap = { 1: uid, 2: roleId || uid, 3: ip }
  const hitWords = msg.hitWords || msg.hit_words || ''
  banDialog.form = {
    appId: msg.appId || msg.app_id, banType, target: targetMap[banType],
    duration: 3600, reason: hitWords ? `敏感词: ${hitWords}` : '',
  }
  banDialog.visible = true
}

const submitBan = async () => {
  if (!banDialog.form.target) {
    ElMessage.warning('请输入封禁目标')
    return
  }
  const res = await createBan(banDialog.form)
  if (res.code === 0) {
    ElMessage.success('封禁成功')
    banDialog.visible = false
  }
}

onMounted(() => { getList(); connectSSE() })
onUnmounted(() => { if (eventSource) eventSource.close() })
</script>

<style scoped>
.chat-stream { background: #1a1a2e; border-radius: 8px; padding: 16px; min-height: 500px; color: #eee; font-size: 13px; overflow-y: auto; max-height: 70vh; }
.chat-item { padding: 8px 12px; border-bottom: 1px solid #2a2a4a; display: flex; align-items: flex-start; gap: 12px; }
.chat-item.sensitive { background: rgba(245, 108, 108, 0.1); border-left: 3px solid #f56c6c; }
.chat-meta { display: flex; align-items: center; gap: 8px; min-width: 280px; }
.channel { color: #67c23a; }
.sender { color: #409eff; font-weight: bold; }
.time { color: #999; font-size: 12px; }
.ip { color: #666; font-size: 12px; }
.chat-content { flex: 1; word-break: break-all; }
.hit-words { color: #f56c6c; font-weight: bold; }
.chat-actions { display: flex; gap: 4px; }
.empty { text-align: center; color: #666; padding: 40px; }
</style>
