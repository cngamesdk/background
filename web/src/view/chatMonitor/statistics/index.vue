<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true">
        <el-form-item>
          <el-select v-model="appId" clearable placeholder="全部游戏" style="width: 150px" @change="refresh">
            <el-option v-for="g in gameList" :key="g.appId" :label="g.name" :value="g.appId" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select v-model="days" style="width: 120px" @change="refresh">
            <el-option label="近7天" :value="7" /><el-option label="近14天" :value="14" /><el-option label="近30天" :value="30" />
          </el-select>
        </el-form-item>
      </el-form>
    </div>

    <el-row :gutter="16" style="margin-bottom: 16px">
      <el-col :span="6"><el-card shadow="hover"><div class="card-title">今日消息</div><div class="card-value">{{ overview.todayMessages }}</div></el-card></el-col>
      <el-col :span="6"><el-card shadow="hover"><div class="card-title">敏感消息</div><div class="card-value danger">{{ overview.todaySensitive }}</div></el-card></el-col>
      <el-col :span="6"><el-card shadow="hover"><div class="card-title">活跃用户</div><div class="card-value">{{ overview.todayActive }}</div></el-card></el-col>
      <el-col :span="6"><el-card shadow="hover"><div class="card-title">今日封禁</div><div class="card-value warning">{{ overview.todayBans }}</div></el-card></el-col>
    </el-row>

    <el-card style="margin-bottom: 16px"><template #header>消息趋势</template><div ref="chartRef" style="height: 350px"></div></el-card>

    <el-card>
      <template #header>违规用户 TOP20</template>
      <el-table :data="violators" border stripe>
        <el-table-column type="index" label="排名" width="60" />
        <el-table-column prop="senderUid" label="账号" />
        <el-table-column prop="senderName" label="昵称" />
        <el-table-column prop="appId" label="游戏" width="120" />
        <el-table-column prop="count" label="违规次数" width="100"><template #default="{ row }"><el-tag type="danger">{{ row.count }}</el-tag></template></el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'
import * as echarts from 'echarts'
import { getStatsOverview, getStatsTrend, getViolators, getGameList } from '@/api/chatMonitor'

defineOptions({ name: 'ChatMonitorStatistics' })

const gameList = ref([])
const appId = ref('')
const days = ref(7)
const overview = ref({ todayMessages: 0, todaySensitive: 0, todayActive: 0, todayBans: 0 })
const violators = ref([])
const chartRef = ref(null)
let chart = null

const refresh = async () => {
  const ov = await getStatsOverview({ appId: appId.value })
  if (ov.code === 0) overview.value = ov.data

  const tr = await getStatsTrend({ appId: appId.value, days: days.value })
  if (tr.code === 0) renderChart(tr.data || [])

  const vl = await getViolators({ appId: appId.value, limit: 20 })
  if (vl.code === 0) violators.value = vl.data || []
}

const renderChart = (data) => {
  if (!chart) chart = echarts.init(chartRef.value)
  chart.setOption({
    tooltip: { trigger: 'axis' },
    legend: { data: ['总消息', '敏感消息', '活跃用户'] },
    xAxis: { type: 'category', data: data.map(d => d.date) },
    yAxis: [{ type: 'value', name: '消息数' }, { type: 'value', name: '用户数' }],
    series: [
      { name: '总消息', type: 'bar', data: data.map(d => d.totalMessages), itemStyle: { color: '#409eff' } },
      { name: '敏感消息', type: 'bar', data: data.map(d => d.sensitiveCount), itemStyle: { color: '#f56c6c' } },
      { name: '活跃用户', type: 'line', yAxisIndex: 1, data: data.map(d => d.activeUsers), itemStyle: { color: '#67c23a' } },
    ],
  })
}

onMounted(async () => {
  const res = await getGameList({ page: 1, pageSize: 100 })
  if (res.code === 0) gameList.value = res.data.list || []
  await nextTick()
  refresh()
})
</script>

<style scoped>
.card-title { font-size: 14px; color: #999; margin-bottom: 8px; text-align: center; }
.card-value { font-size: 28px; font-weight: bold; color: #333; text-align: center; }
.card-value.danger { color: #f56c6c; }
.card-value.warning { color: #e6a23c; }
</style>
