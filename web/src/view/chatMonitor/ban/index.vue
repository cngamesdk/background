<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true">
        <el-form-item>
          <el-select v-model="searchInfo.appId" clearable placeholder="全部游戏" style="width: 150px" @change="getList">
            <el-option v-for="g in gameList" :key="g.appId" :label="g.name" :value="g.appId" />
          </el-select>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" border stripe>
        <el-table-column prop="appId" label="游戏" width="120" />
        <el-table-column prop="banType" label="类型" width="100">
          <template #default="{ row }">
            <el-tag :type="{ 1: 'danger', 2: 'warning', 3: 'info' }[row.banType]" size="small">{{ { 1: '封号', 2: '禁言', 3: '封IP' }[row.banType] }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="target" label="目标" width="160" />
        <el-table-column prop="reason" label="原因" show-overflow-tooltip />
        <el-table-column prop="duration" label="时长" width="100">
          <template #default="{ row }">{{ formatDuration(row.duration) }}</template>
        </el-table-column>
        <el-table-column prop="startAt" label="开始时间" width="170" />
        <el-table-column prop="expireAt" label="过期时间" width="170">
          <template #default="{ row }">{{ row.expireAt || '永久' }}</template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="{ 1: 'danger', 2: 'success', 3: 'info' }[row.status]" size="small">{{ { 1: '生效中', 2: '已解封', 3: '已过期' }[row.status] }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="100" fixed="right">
          <template #default="{ row }">
            <el-popconfirm v-if="row.status === 1" title="确认解封?" @confirm="handleRevoke(row.ID)">
              <template #reference><el-button size="small" type="success">解封</el-button></template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination v-model:current-page="searchInfo.page" :total="total" layout="total, prev, pager, next" style="margin-top: 16px" @current-change="getList" />
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getBanList, revokeBan, getGameList } from '@/api/chatMonitor'
import { ElMessage } from 'element-plus'

defineOptions({ name: 'ChatMonitorBan' })

const gameList = ref([])
const tableData = ref([])
const total = ref(0)
const searchInfo = reactive({ appId: '', page: 1, pageSize: 20 })

const formatDuration = (s) => {
  if (!s) return '永久'
  if (s < 3600) return `${Math.floor(s / 60)}分钟`
  if (s < 86400) return `${Math.floor(s / 3600)}小时`
  return `${Math.floor(s / 86400)}天`
}

const getList = async () => {
  const res = await getBanList(searchInfo)
  if (res.code === 0) { tableData.value = res.data.list || []; total.value = res.data.total }
}
const handleRevoke = async (id) => {
  await revokeBan({ id }); ElMessage.success('解封成功'); getList()
}

onMounted(async () => {
  const res = await getGameList({ page: 1, pageSize: 100 })
  if (res.code === 0) gameList.value = res.data.list || []
  getList()
})
</script>
