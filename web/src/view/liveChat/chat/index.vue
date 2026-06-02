<template>
  <div class="chat-admin">
    <el-row :gutter="10" style="height:calc(100vh - 120px)">
      <el-col :span="8" style="height:100%">
        <el-card header="会话列表" style="height:100%; display: flex; flex-direction: column;">
          <el-form :inline="true" size="small">
            <el-form-item>
              <el-input v-model="searchForm.user_id" placeholder="用户ID" clearable @clear="getSessions" />
            </el-form-item>
            <el-form-item>
              <el-select v-model="searchForm.status" placeholder="状态" clearable @change="getSessions">
                <el-option label="全部" value="" />
                <el-option label="等待中" value="waiting" />
                <el-option label="进行中" value="active" />
                <el-option label="已关闭" value="closed" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" size="small" @click="getSessions">搜索</el-button>
            </el-form-item>
          </el-form>
          <div class="session-list">
            <div v-for="s in sessions" :key="s.ID" class="session-item" :class="{ active: currentSession?.ID === s.ID }" @click="selectSession(s)">
              <div class="session-header">
                <div class="session-name">{{ s.user_name || s.user_id }}</div>
                <el-tag size="small" :type="s.status==='waiting'?'warning':s.status==='active'?'success':'info'">
                  {{ statusText(s.status) }}
                </el-tag>
              </div>
              <div class="session-meta">
                <span>{{ formatTime(s.created_at) }}</span>
                <el-badge v-if="s.unread_count > 0" :value="s.unread_count" class="unread-badge" />
              </div>
            </div>
            <el-empty v-if="sessions.length === 0" description="暂无会话" />
          </div>
          <el-pagination
            v-model:current-page="page"
            :total="total"
            :page-size="10"
            layout="prev, pager, next"
            @current-change="getSessions"
            small
            style="margin-top: 10px; text-align: center;"
          />
        </el-card>
      </el-col>
      <el-col :span="16" style="height:100%">
        <el-card style="height:100%; display: flex; flex-direction: column;" v-if="currentSession">
          <template #header>
            <div style="display: flex; justify-content: space-between; align-items: center;">
              <span>{{ currentSession.user_name || currentSession.user_id }} ({{ statusText(currentSession.status) }})</span>
              <div>
                <el-button v-if="currentSession.status==='waiting'" size="small" type="success" @click="handleAssign">接单</el-button>
                <el-button v-if="currentSession.status!=='closed'" size="small" type="danger" @click="handleClose">关闭</el-button>
              </div>
            </div>
          </template>
          <div class="chat-messages" ref="msgContainer">
            <div v-for="msg in messages" :key="msg.id" class="chat-msg" :class="getMsgClass(msg)">
              <div class="msg-sender">{{ msg.sender_name }}</div>
              <div class="msg-content">
                <span>{{ msg.content }}</span>
                <el-tag v-if="msg.is_faq_reply" size="small" type="warning" style="margin-left: 8px;">自动回复</el-tag>
              </div>
              <div v-if="msg.attachment_url" class="msg-attachment">
                <img v-if="msg.msg_type==='image'" :src="msg.attachment_url" style="max-width:200px; border-radius: 4px;" />
                <video v-else-if="msg.msg_type==='video'" :src="msg.attachment_url" controls style="max-width:250px; border-radius: 4px;" />
              </div>
              <div class="msg-time">{{ formatTime(msg.created_at) }}</div>
            </div>
          </div>
          <div class="chat-reply" v-if="currentSession.status!=='closed'">
            <el-input
              v-model="replyText"
              type="textarea"
              :rows="3"
              placeholder="输入回复内容..."
              @keydown.ctrl.enter="handleReply"
            />
            <div style="margin-top: 8px; text-align: right;">
              <el-button size="small" @click="replyText = ''">清空</el-button>
              <el-button size="small" type="primary" @click="handleReply">发送 (Ctrl+Enter)</el-button>
            </div>
          </div>
        </el-card>
        <el-empty v-else description="选择一个会话开始对话" style="height:100%; display: flex; align-items: center; justify-content: center;" />
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, reactive, nextTick, onMounted, onUnmounted, watch } from 'vue'
import { getChatSessions, getSessionDetail, assignSession, agentReply, closeSession } from '@/api/liveChat'
import { ElMessage } from 'element-plus'

defineOptions({ name: 'LiveChatChat' })

const searchForm = reactive({ user_id: '', status: 'waiting' })
const sessions = ref([])
const page = ref(1)
const total = ref(0)
const currentSession = ref(null)
const messages = ref([])
const replyText = ref('')
const msgContainer = ref(null)
let pollingTimer = null

const statusText = (s) => ({ waiting: '等待中', active: '进行中', closed: '已关闭' }[s] || s)

const formatTime = (time) => {
  if (!time) return ''
  const date = new Date(time)
  const now = new Date()
  const diff = now - date

  // 今天
  if (diff < 86400000 && date.getDate() === now.getDate()) {
    return date.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' })
  }
  // 昨天
  if (diff < 172800000 && date.getDate() === now.getDate() - 1) {
    return '昨天 ' + date.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit' })
  }
  // 其他
  return date.toLocaleString('zh-CN', {
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getMsgClass = (msg) => {
  return {
    fromUser: msg.sender_type === 'user',
    fromAgent: msg.sender_type === 'agent',
    fromSystem: msg.sender_type === 'system'
  }
}

const scrollToBottom = () => {
  nextTick(() => {
    if (msgContainer.value) {
      msgContainer.value.scrollTop = msgContainer.value.scrollHeight
    }
  })
}

const getSessions = async () => {
  try {
    const res = await getChatSessions({ ...searchForm, page: page.value, pageSize: 10 })
    if (res.code === 0) {
      sessions.value = res.data?.list || []
      total.value = res.data?.total || 0
    }
  } catch (error) {
    console.error('获取会话列表失败:', error)
  }
}

const selectSession = async (s) => {
  currentSession.value = s
  try {
    const res = await getSessionDetail(s.ID)
    if (res.code === 0) {
      messages.value = res.data?.messages || []
      scrollToBottom()
    }
  } catch (error) {
    console.error('获取会话详情失败:', error)
  }
}

const handleAssign = async () => {
  try {
    await assignSession({ session_id: currentSession.value.ID })
    ElMessage.success('已接单')
    getSessions()
    selectSession(currentSession.value)
  } catch (error) {
    ElMessage.error('接单失败')
  }
}

const handleReply = async () => {
  if (!replyText.value.trim()) {
    ElMessage.warning('请输入回复内容')
    return
  }
  try {
    await agentReply({
      session_id: currentSession.value.ID,
      content: replyText.value,
      msg_type: 'text'
    })
    replyText.value = ''
    // 立即刷新消息
    selectSession(currentSession.value)
  } catch (error) {
    ElMessage.error('发送失败')
  }
}

const handleClose = async () => {
  try {
    await closeSession({ ID: currentSession.value.ID })
    ElMessage.success('已关闭')
    currentSession.value = null
    getSessions()
  } catch (error) {
    ElMessage.error('关闭失败')
  }
}

// 实时更新当前会话的消息
const refreshCurrentSession = async () => {
  if (currentSession.value) {
    try {
      const res = await getSessionDetail(currentSession.value.ID)
      if (res.code === 0) {
        const newMessages = res.data?.messages || []
        // 只有消息数量变化时才更新并滚动
        if (newMessages.length !== messages.value.length) {
          messages.value = newMessages
          scrollToBottom()
        }
      }
    } catch (error) {
      console.error('刷新消息失败:', error)
    }
  }
}

// 监听当前会话变化，自动滚动
watch(() => messages.value.length, () => {
  scrollToBottom()
})

onMounted(() => {
  getSessions()
  // 每5秒刷新会话列表
  pollingTimer = setInterval(() => {
    getSessions()
    refreshCurrentSession()
  }, 5000)
})

onUnmounted(() => {
  if (pollingTimer) {
    clearInterval(pollingTimer)
  }
})
</script>

<style scoped>
.chat-admin {
  padding: 0;
  height: 100%;
}

.session-list {
  flex: 1;
  overflow-y: auto;
  margin-bottom: 10px;
}

.session-item {
  padding: 12px;
  border-bottom: 1px solid #eee;
  cursor: pointer;
  transition: background 0.2s;
}

.session-item.active {
  background: #ecf5ff;
  border-left: 3px solid #409eff;
}

.session-item:hover {
  background: #f5f7fa;
}

.session-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 6px;
}

.session-name {
  font-weight: 600;
  font-size: 14px;
  color: #303133;
}

.session-meta {
  font-size: 12px;
  color: #909399;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.unread-badge {
  margin-left: 8px;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  background: #f5f7fa;
  border-radius: 4px;
  margin-bottom: 16px;
  max-height: calc(100vh - 400px);
  min-height: 400px;
}

.chat-msg {
  margin-bottom: 16px;
  padding: 10px 14px;
  border-radius: 8px;
  max-width: 70%;
  word-wrap: break-word;
  animation: fadeIn 0.3s;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.chat-msg.fromUser {
  background: #e6f7ff;
  margin-left: auto;
  border: 1px solid #91d5ff;
}

.chat-msg.fromAgent {
  background: #f6ffed;
  margin-right: auto;
  border: 1px solid #b7eb8f;
}

.chat-msg.fromSystem {
  text-align: center;
  color: #909399;
  font-size: 12px;
  max-width: 100%;
  background: #f0f0f0;
  margin: 8px auto;
  padding: 6px 12px;
}

.msg-sender {
  font-size: 12px;
  color: #909399;
  margin-bottom: 6px;
  font-weight: 500;
}

.msg-content {
  line-height: 1.6;
  color: #303133;
  font-size: 14px;
}

.msg-attachment {
  margin-top: 8px;
}

.msg-time {
  font-size: 11px;
  color: #c0c4cc;
  margin-top: 4px;
  text-align: right;
}

.chat-reply {
  padding: 16px;
  background: #fff;
  border-top: 1px solid #ebeef5;
  border-radius: 4px;
}

:deep(.el-card__body) {
  display: flex;
  flex-direction: column;
  height: calc(100% - 60px);
  padding: 16px;
}

:deep(.el-card__header) {
  padding: 16px;
  border-bottom: 1px solid #ebeef5;
}
</style>
