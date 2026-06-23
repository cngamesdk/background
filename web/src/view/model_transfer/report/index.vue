<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="Token名称">
          <el-input v-model="searchInfo.tokenName" placeholder="请输入" clearable />
        </el-form-item>
        <el-form-item label="提供商">
          <el-select v-model="searchInfo.provider" placeholder="请选择" clearable>
            <el-option label="全部" value="" />
            <el-option label="OpenAI" value="openai" />
            <el-option label="Anthropic" value="anthropic" />
            <el-option label="Google" value="google" />
          </el-select>
        </el-form-item>
        <el-form-item label="模型">
          <el-input v-model="searchInfo.model" placeholder="请输入" clearable />
        </el-form-item>
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
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <!-- 统计卡片 -->
    <el-row :gutter="20" style="margin-bottom: 20px">
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <el-icon size="40" color="#409eff"><data-analysis /></el-icon>
            <div class="stat-content">
              <div class="stat-value">{{ summary.totalRequests }}</div>
              <div class="stat-label">总请求数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <el-icon size="40" color="#67c23a"><success-filled /></el-icon>
            <div class="stat-content">
              <div class="stat-value">{{ summary.successRate }}%</div>
              <div class="stat-label">成功率</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <el-icon size="40" color="#e6a23c"><coin /></el-icon>
            <div class="stat-content">
              <div class="stat-value">{{ summary.totalTokens }}</div>
              <div class="stat-label">总Token数</div>
            </div>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6">
        <el-card shadow="hover">
          <div class="stat-card">
            <el-icon size="40" color="#f56c6c"><timer /></el-icon>
            <div class="stat-content">
              <div class="stat-value">{{ summary.avgDurationMs }}ms</div>
              <div class="stat-label">平均耗时</div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 日报表格 -->
    <div class="gva-table-box">
      <el-table v-loading="loading" :data="tableData" style="width: 100%" border>
        <el-table-column label="日期" prop="statDate" width="120" />
        <el-table-column label="Token名称" prop="tokenName" width="150" />
        <el-table-column label="提供商" prop="provider" width="120">
          <template #default="scope">
            <el-tag>{{ scope.row.provider }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column label="模型" prop="model" width="150" />
        <el-table-column label="请求数" prop="requestCount" width="100" />
        <el-table-column label="成功数" prop="successCount" width="100" />
        <el-table-column label="失败数" prop="errorCount" width="100" />
        <el-table-column label="Token数" prop="totalTokens" width="120" />
        <el-table-column label="平均耗时" prop="avgDurationMs" width="100">
          <template #default="scope">
            {{ scope.row.avgDurationMs }}ms
          </template>
        </el-table-column>
        <el-table-column label="成功率" prop="successRate" width="100">
          <template #default="scope">
            <el-tag :type="getSuccessRateType(scope.row.successRate)">
              {{ scope.row.successRate }}%
            </el-tag>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { getDailyReport, getSummaryReport } from '@/api/modelTransfer'

// 搜索条件
const searchInfo = ref({
  tokenName: '',
  provider: '',
  model: ''
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
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const loading = ref(false)

// 汇总数据
const summary = ref({
  totalRequests: 0,
  successRate: 0,
  totalTokens: 0,
  avgDurationMs: 0
})

// 获取日报数据
const getTableData = async() => {
  loading.value = true
  try {
    const params = {
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    }

    if (dateRange.value && dateRange.value.length === 2) {
      params.startDate = dateRange.value[0]
      params.endDate = dateRange.value[1]
    }

    const res = await getDailyReport(params)
    if (res.code === 0) {
      tableData.value = res.data.list || []
      total.value = res.data.total || 0
    }
  } catch (error) {
    ElMessage.error('获取日报数据失败: ' + error.message)
  } finally {
    loading.value = false
  }
}

// 获取汇总数据
const getSummary = async() => {
  try {
    const params = {}

    if (dateRange.value && dateRange.value.length === 2) {
      params.startDate = dateRange.value[0]
      params.endDate = dateRange.value[1]
    }

    const res = await getSummaryReport(params)
    if (res.code === 0) {
      summary.value = res.data
    }
  } catch (error) {
    ElMessage.error('获取汇总数据失败: ' + error.message)
  }
}

// 查询
const onSubmit = () => {
  page.value = 1
  getTableData()
  getSummary()
}

// 重置
const onReset = () => {
  searchInfo.value = {
    tokenName: '',
    provider: '',
    model: ''
  }
  dateRange.value = []
  onSubmit()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 成功率标签类型
const getSuccessRateType = (rate) => {
  if (rate >= 95) return 'success'
  if (rate >= 80) return 'warning'
  return 'danger'
}

// 初始化
getTableData()
getSummary()
</script>

<style scoped>
.stat-card {
  display: flex;
  align-items: center;
  gap: 15px;
}

.stat-content {
  flex: 1;
}

.stat-value {
  font-size: 24px;
  font-weight: bold;
  color: #303133;
}

.stat-label {
  font-size: 14px;
  color: #909399;
  margin-top: 5px;
}
</style>
