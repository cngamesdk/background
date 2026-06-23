<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="日期范围">
          <el-date-picker
            v-model="dateRange"
            type="daterange"
            range-separator="至"
            start-placeholder="开始日期"
            end-placeholder="结束日期"
            format="YYYY-MM-DD"
            value-format="YYYY-MM-DD"
          />
        </el-form-item>
        <el-form-item label="提供商">
          <el-select v-model="searchInfo.provider" placeholder="请选择" clearable>
            <el-option label="全部" value="" />
            <el-option label="OpenAI" value="openai" />
            <el-option label="Anthropic" value="anthropic" />
            <el-option label="Google" value="google" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button type="success" icon="download" @click="exportData">导出</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 趋势图表 -->
    <el-card shadow="hover" style="margin-bottom: 20px">
      <template #header>
        <div class="card-header">
          <span>请求趋势</span>
        </div>
      </template>
      <div ref="chartRef" style="height: 400px">
        <el-empty v-if="!tableData.length" description="暂无数据" />
      </div>
    </el-card>

    <!-- 提供商统计 -->
    <el-row :gutter="20" style="margin-bottom: 20px">
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <span>提供商请求分布</span>
            </div>
          </template>
          <div ref="pieChartRef" style="height: 300px">
            <el-empty v-if="!tableData.length" description="暂无数据" />
          </div>
        </el-card>
      </el-col>
      <el-col :span="12">
        <el-card shadow="hover">
          <template #header>
            <div class="card-header">
              <span>提供商Token分布</span>
            </div>
          </template>
          <div ref="tokenPieChartRef" style="height: 300px">
            <el-empty v-if="!tableData.length" description="暂无数据" />
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 数据表格 -->
    <div class="gva-table-box">
      <el-table v-loading="loading" :data="tableData" style="width: 100%" border>
        <el-table-column label="日期" prop="statDate" width="120" />
        <el-table-column label="提供商" prop="provider" width="120">
          <template #default="scope">
            <el-tag>{{ scope.row.provider }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="模型" prop="model" width="150" />
        <el-table-column label="请求数" prop="requestCount" width="100" sortable />
        <el-table-column label="Token数" prop="totalTokens" width="120" sortable />
        <el-table-column label="成功率" prop="successRate" width="100" sortable>
          <template #default="scope">
            <el-progress
              :percentage="scope.row.successRate"
              :color="getProgressColor(scope.row.successRate)"
            />
          </template>
        </el-table-column>
        <el-table-column label="平均耗时" prop="avgDurationMs" width="100">
          <template #default="scope">
            {{ scope.row.avgDurationMs }}ms
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import * as echarts from 'echarts'
import { ElMessage } from 'element-plus'
import { getDailyReport } from '@/api/modelTransfer'

// 搜索条件
const searchInfo = ref({
  provider: ''
})

// 获取今天的日期字符串 YYYY-MM-DD
const getTodayDate = () => {
  const today = new Date()
  const year = today.getFullYear()
  const month = String(today.getMonth() + 1).padStart(2, '0')
  const day = String(today.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
}

const dateRange = ref([getTodayDate(), getTodayDate()])

// 表格数据
const tableData = ref([])
const loading = ref(false)

// 图表引用
const chartRef = ref(null)
const pieChartRef = ref(null)
const tokenPieChartRef = ref(null)

// 初始化趋势图
const initChart = () => {
  if (!chartRef.value) return

  const myChart = echarts.init(chartRef.value)

  // 从tableData提取数据
  const dates = tableData.value.map(item => item.statDate)
  const requests = tableData.value.map(item => item.requestCount)
  const tokens = tableData.value.map(item => item.totalTokens)

  const option = {
    title: { text: '请求趋势' },
    tooltip: { trigger: 'axis' },
    legend: { data: ['请求数', 'Token数'] },
    xAxis: {
      type: 'category',
      data: dates
    },
    yAxis: [
      {
        type: 'value',
        name: '请求数'
      },
      {
        type: 'value',
        name: 'Token数'
      }
    ],
    series: [
      {
        name: '请求数',
        type: 'line',
        data: requests,
        yAxisIndex: 0
      },
      {
        name: 'Token数',
        type: 'line',
        data: tokens,
        yAxisIndex: 1
      }
    ]
  }
  myChart.setOption(option)
}

// 初始化饼图 - 提供商请求分布
const initPieChart = () => {
  if (!pieChartRef.value) return

  const myChart = echarts.init(pieChartRef.value)

  // 按提供商聚合数据
  const providerData = {}
  tableData.value.forEach(item => {
    if (!providerData[item.provider]) {
      providerData[item.provider] = 0
    }
    providerData[item.provider] += item.requestCount
  })

  const data = Object.keys(providerData).map(key => ({
    name: key,
    value: providerData[key]
  }))

  const option = {
    title: { text: '提供商请求分布', left: 'center' },
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)'
    },
    series: [
      {
        name: '请求分布',
        type: 'pie',
        radius: '60%',
        data: data,
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0, 0, 0, 0.5)'
          }
        }
      }
    ]
  }
  myChart.setOption(option)
}

// 初始化Token饼图 - 提供商Token分布
const initTokenPieChart = () => {
  if (!tokenPieChartRef.value) return

  const myChart = echarts.init(tokenPieChartRef.value)

  // 按提供商聚合Token数据
  const providerTokens = {}
  tableData.value.forEach(item => {
    if (!providerTokens[item.provider]) {
      providerTokens[item.provider] = 0
    }
    providerTokens[item.provider] += item.totalTokens
  })

  const data = Object.keys(providerTokens).map(key => ({
    name: key,
    value: providerTokens[key]
  }))

  const option = {
    title: { text: '提供商Token分布', left: 'center' },
    tooltip: {
      trigger: 'item',
      formatter: '{a} <br/>{b}: {c} ({d}%)'
    },
    series: [
      {
        name: 'Token分布',
        type: 'pie',
        radius: '60%',
        data: data,
        emphasis: {
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0, 0, 0, 0.5)'
          }
        }
      }
    ]
  }
  myChart.setOption(option)
}

// 获取数据
const getTableData = async() => {
  loading.value = true
  try {
    const params = {
      ...searchInfo.value
    }

    if (dateRange.value && dateRange.value.length === 2) {
      params.startDate = dateRange.value[0]
      params.endDate = dateRange.value[1]
    }

    const res = await getDailyReport(params)
    if (res.code === 0) {
      tableData.value = res.data.list || []

      // 等待DOM更新后初始化图表
      if (tableData.value.length > 0) {
        setTimeout(() => {
          initChart()
          initPieChart()
          initTokenPieChart()
        }, 100)
      }
    }
  } catch (error) {
    ElMessage.error('获取数据失败: ' + error.message)
  } finally {
    loading.value = false
  }
}

// 查询
const onSubmit = () => {
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = { provider: '' }
  dateRange.value = []
  onSubmit()
}

// 导出
const exportData = () => {
  ElMessage.info('导出功能开发中')
}

// 进度条颜色
const getProgressColor = (percent) => {
  if (percent >= 95) return '#67c23a'
  if (percent >= 80) return '#e6a23c'
  return '#f56c6c'
}

// 初始化
onMounted(() => {
  getTableData()
})
</script>

<style scoped>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}
</style>
