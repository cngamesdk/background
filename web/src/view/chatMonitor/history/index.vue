<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo">
        <el-form-item label="游戏">
          <el-select v-model="searchInfo.appId" clearable placeholder="全部" style="width: 150px">
            <el-option v-for="g in gameList" :key="g.appId" :label="g.name" :value="g.appId" />
          </el-select>
        </el-form-item>
        <el-form-item label="账号"><el-input v-model="searchInfo.senderUid" clearable style="width: 140px" /></el-form-item>
        <el-form-item label="关键词"><el-input v-model="searchInfo.keyword" clearable style="width: 140px" /></el-form-item>
        <el-form-item label="时间">
          <el-date-picker v-model="dateRange" type="datetimerange" range-separator="至"
            start-placeholder="开始" end-placeholder="结束" value-format="YYYY-MM-DD HH:mm:ss" />
        </el-form-item>
        <el-form-item label="仅敏感"><el-switch v-model="searchInfo.onlySensitive" /></el-form-item>
        <el-form-item>
          <el-button type="primary" @click="getList">查询</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" border stripe>
        <el-table-column prop="sentAt" label="时间" width="170" />
        <el-table-column prop="appId" label="游戏" width="120" />
        <el-table-column prop="channel" label="频道" width="100" />
        <el-table-column prop="senderName" label="昵称" width="120" />
        <el-table-column prop="senderUid" label="账号" width="140" />
        <el-table-column prop="senderIp" label="IP" width="130" />
        <el-table-column prop="content" label="内容" show-overflow-tooltip />
        <el-table-column prop="hitWords" label="敏感词" width="150">
          <template #default="{ row }"><el-tag v-if="row.hitWords" type="danger" size="small">{{ row.hitWords }}</el-tag></template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" type="danger" @click="handleBan(row, 1)">封号</el-button>
            <el-button size="small" type="warning" @click="handleBan(row, 2)">禁言</el-button>
            <el-button size="small" type="info" @click="handleBan(row, 3)">封IP</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination v-model:current-page="searchInfo.page" v-model:page-size="searchInfo.pageSize" :total="total"
        layout="total, sizes, prev, pager, next" :page-sizes="[20, 50, 100]" style="margin-top: 16px" @current-change="getList" @size-change="getList" />
    </div>

    <!-- 封禁弹窗 -->
    <el-dialog v-model="banDialog.visible" title="执行封禁" width="400px">
      <el-form :model="banDialog.form" label-width="80px">
        <el-form-item label="类型">
          <el-tag>{{ { 1: '账号封禁', 2: '角色禁言', 3: 'IP封禁' }[banDialog.form.banType] }}</el-tag>
        </el-form-item>
        <el-form-item label="目标">
          <el-input v-model="banDialog.form.target" disabled />
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
import { ref, reactive, onMounted } from 'vue'
import { getChatHistory, getGameList, createBan } from '@/api/chatMonitor'
import { ElMessage } from 'element-plus'

defineOptions({ name: 'ChatMonitorHistory' })

const gameList = ref([])
const tableData = ref([])
const total = ref(0)
const dateRange = ref([])
const searchInfo = reactive({ appId: '', senderUid: '', keyword: '', onlySensitive: false, page: 1, pageSize: 20 })

const getList = async () => {
  const params = { ...searchInfo }
  if (dateRange.value?.length === 2) { params.startTime = dateRange.value[0]; params.endTime = dateRange.value[1] }
  const res = await getChatHistory(params)
  if (res.code === 0) { tableData.value = res.data.list || []; total.value = res.data.total }
}

onMounted(async () => {
  const res = await getGameList({ page: 1, pageSize: 100 })
  if (res.code === 0) gameList.value = res.data.list || []
  getList()
})

const banDialog = reactive({
  visible: false,
  form: { appId: '', banType: 1, target: '', duration: 3600, reason: '' },
})

const handleBan = (row, banType) => {
  const uid = row.senderUid || row.sender_uid || ''
  const roleId = row.roleId || row.role_id || ''
  const ip = row.senderIp || row.sender_ip || ''
  const targetMap = { 1: uid, 2: roleId || uid, 3: ip }
  const hitWords = row.hitWords || row.hit_words || ''
  banDialog.form = {
    appId: row.appId || row.app_id, banType, target: targetMap[banType],
    duration: 3600, reason: hitWords ? `敏感词: ${hitWords}` : '',
  }
  banDialog.visible = true
}

const submitBan = async () => {
  const res = await createBan(banDialog.form)
  if (res.code === 0) {
    ElMessage.success('封禁成功')
    banDialog.visible = false
  }
}
</script>
