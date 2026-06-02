<template>
  <div>
    <el-card>
      <el-form :inline="true" :model="searchForm" size="small">
        <el-form-item label="产品">
          <el-select v-model="searchForm.product_id" filterable clearable placeholder="全部产品" style="width:200px">
            <el-option v-for="p in productList" :key="p.ID" :label="p.name" :value="p.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="开始日期"><el-date-picker v-model="searchForm.start_date" type="date" value-format="YYYY-MM-DD" /></el-form-item>
        <el-form-item label="结束日期"><el-date-picker v-model="searchForm.end_date" type="date" value-format="YYYY-MM-DD" /></el-form-item>
        <el-form-item><el-button type="primary" @click="fetchData">查询</el-button></el-form-item>
      </el-form>
    </el-card>

    <el-row :gutter="10" style="margin-top:10px">
      <el-col :span="6" v-for="card in overviewCards" :key="card.label">
        <el-card><el-statistic :title="card.label" :value="card.value"><template #suffix v-if="card.suffix">{{ card.suffix }}</template></el-statistic></el-card>
      </el-col>
    </el-row>

    <el-card header="会话趋势" style="margin-top:10px">
      <el-table :data="trendData" border stripe>
        <el-table-column prop="date" label="日期" />
        <el-table-column prop="total_sessions" label="会话数" />
      </el-table>
    </el-card>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { getReportOverview, getReportTrend } from '@/api/liveChat'
import { getProductList } from '@/api/liveChat'

defineOptions({ name: 'LiveChatReport' })

const searchForm = reactive({ product_id: '', start_date: new Date(Date.now()-7*86400000).toISOString().substring(0,10), end_date: new Date().toISOString().substring(0,10) })
const overview = ref({})
const trendData = ref([])
const productList = ref([])

const overviewCards = computed(() => [
  { label: '总会话数', value: overview.value.total_sessions || 0 },
  { label: 'FAQ解决', value: overview.value.faq_resolved || 0 },
  { label: '人工解决', value: overview.value.agent_resolved || 0 },
  { label: 'FAQ解决率', value: overview.value.faq_rate?.toFixed(1) || 0, suffix: '%' },
  { label: '在线客服', value: overview.value.online_agents || 0 },
  { label: '等待会话', value: overview.value.waiting_sessions || 0 }
])

const loadProducts = async () => {
  const res = await getProductList({ page: 1, pageSize: 100 })
  if (res.code === 0) productList.value = res.data?.list || []
}

const fetchData = async () => {
  const params = { ...searchForm }
  if (!searchForm.product_id) delete params.product_id
  const ov = await getReportOverview(params)
  if (ov.code === 0) overview.value = ov.data || {}
  const tr = await getReportTrend(params)
  if (tr.code === 0) trendData.value = tr.data || []
}

onMounted(() => { loadProducts(); fetchData() })
</script>
