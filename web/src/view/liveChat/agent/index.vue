<template>
  <div>
    <el-card>
      <el-form :inline="true" :model="searchForm" size="small">
        <el-form-item label="产品">
          <el-select v-model="searchForm.product_id" filterable clearable placeholder="选择产品" style="width:200px">
            <el-option v-for="p in productList" :key="p.ID" :label="p.name" :value="p.ID" />
          </el-select>
        </el-form-item>
        <el-form-item><el-button type="primary" @click="getList">查询</el-button></el-form-item>
      </el-form>
    </el-card>
    <el-card style="margin-top:10px">
      <el-table :data="tableData" border stripe>
        <el-table-column prop="ID" label="ID" width="60" />
        <el-table-column prop="user_id" label="用户ID" />
        <el-table-column label="产品名称" width="120">
          <template #default="{row}">{{ getProductName(row.product_id) }}</template>
        </el-table-column>
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{row}"><el-tag :type="row.status==='online'?'success':'info'">{{ row.status === 'online' ? '在线' : row.status === 'offline' ? '离线' : '忙碌' }}</el-tag></template>
        </el-table-column>
        <el-table-column prop="current_sessions" label="当前会话" width="100" />
        <el-table-column prop="total_served" label="总服务量" width="100" />
        <el-table-column prop="max_concurrent" label="最大并发" width="100" />
        <el-table-column label="操作" width="200">
          <template #default="{row}">
            <el-button size="small" :type="row.status==='online'?'warning':'success'" @click="toggleStatus(row)">{{ row.status==='online'?'下线':'上线' }}</el-button>
            <el-button size="small" @click="openConfig(row)">配置</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination v-model:current-page="page" :total="total" :page-size="pageSize" layout="total,prev,pager,next" @current-change="getList" style="margin-top:10px" />
    </el-card>

    <el-dialog title="客服配置" v-model="configVisible" width="400px">
      <el-form :model="configForm" label-width="100px">
        <el-form-item label="最大并发数"><el-input-number v-model="configForm.max_concurrent" :min="1" :max="50" /></el-form-item>
      </el-form>
      <template #footer><el-button @click="configVisible=false">取消</el-button><el-button type="primary" @click="saveConfig">保存</el-button></template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getAgentList, agentOnline, agentOffline, updateAgent } from '@/api/liveChat'
import { getProductList } from '@/api/liveChat'
import { ElMessage } from 'element-plus'

defineOptions({ name: 'LiveChatAgent' })

const searchForm = reactive({ product_id: '' })
const tableData = ref([])
const productList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const configVisible = ref(false)
const configForm = reactive({ id: 0, max_concurrent: 5 })

const loadProducts = async () => {
  const res = await getProductList({ page: 1, pageSize: 100 })
  if (res.code === 0) productList.value = res.data?.list || []
}

const getProductName = (pid) => {
  const p = productList.value.find(i => i.ID === pid)
  return p ? p.name : pid
}

const getList = async () => {
  const params = { ...searchForm, page: page.value, pageSize: pageSize.value }
  if (searchForm.product_id) params.product_id = searchForm.product_id
  const res = await getAgentList(params)
  if (res.code === 0) { tableData.value = res.data?.list || []; total.value = res.data?.total || 0 }
}

const toggleStatus = async (row) => {
  const fn = row.status === 'online' ? agentOffline : agentOnline
  const res = await fn({ product_id: row.product_id })
  if (res.code === 0) { ElMessage.success('操作成功'); getList() }
}

const openConfig = (row) => { configForm.id = row.ID; configForm.max_concurrent = row.max_concurrent; configVisible.value = true }
const saveConfig = async () => {
  await updateAgent(configForm); ElMessage.success('保存成功'); configVisible.value = false; getList()
}

onMounted(() => { loadProducts(); getList() })
</script>
