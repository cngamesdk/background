<template>
  <div class="activity-list-container">
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="活动名称">
          <el-input v-model="searchForm.activity_name" placeholder="请输入活动名称" clearable />
        </el-form-item>
        <el-form-item label="活动类型">
          <el-select v-model="searchForm.activity_type" placeholder="请选择" clearable>
            <el-option label="累计充值" value="recharge" />
            <el-option label="登录奖励" value="login" />
            <el-option label="签到" value="signin" />
            <el-option label="分享" value="share" />
            <el-option label="自定义" value="custom" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchForm.status" placeholder="请选择" clearable>
            <el-option label="未开始" value="not-started" />
            <el-option label="进行中" value="normal" />
            <el-option label="已下线" value="remove" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
          <el-button @click="handleReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" @click="handleAdd">新增活动</el-button>
        <el-button @click="handleFromTemplate">从模板创建</el-button>
      </div>

      <el-table :data="tableData" border stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="activity_name" label="活动名称" min-width="150" />
        <el-table-column prop="activity_type" label="类型" width="100">
          <template #default="{ row }">
            {{ typeMap[row.activity_type] || row.activity_type }}
          </template>
        </el-table-column>
        <el-table-column prop="start_time" label="开始时间" width="170" />
        <el-table-column prop="end_time" label="结束时间" width="170" />
        <el-table-column prop="status" label="状态" width="90">
          <template #default="{ row }">
            <el-tag :type="statusTagType(row.status)">{{ statusMap[row.status] }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="grayscale_ratio" label="灰度" width="80">
          <template #default="{ row }">{{ row.grayscale_ratio }}%</template>
        </el-table-column>
        <el-table-column label="操作" width="280" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="success" v-if="row.status === 'not-started'" @click="handlePublish(row)">发布</el-button>
            <el-button size="small" type="danger" v-if="row.status === 'normal'" @click="handleOffline(row)">下线</el-button>
            <el-button size="small" @click="handleClone(row)">克隆</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :total="total"
          layout="total, sizes, prev, pager, next"
          @current-change="handlePageChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { activityList, activityPublish, activityOffline } from '@/api/activityEngine'

const router = useRouter()

const typeMap = { recharge: '累计充值', login: '登录奖励', signin: '签到', share: '分享', custom: '自定义' }
const statusMap = { 'not-started': '未开始', normal: '进行中', remove: '已下线' }

const searchForm = ref({ activity_name: '', activity_type: '', status: '' })
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)

const statusTagType = (status) => {
  const map = { 'not-started': 'info', normal: 'success', remove: 'danger' }
  return map[status] || 'info'
}

const fetchData = async () => {
  const res = await activityList({ ...searchForm.value, page: page.value, pageSize: pageSize.value })
  if (res.code === 0) {
    tableData.value = res.data.list || []
    total.value = res.data.total || 0
  }
}

const handleSearch = () => { page.value = 1; fetchData() }
const handleReset = () => { searchForm.value = { activity_name: '', activity_type: '', status: '' }; handleSearch() }
const handlePageChange = (val) => { page.value = val; fetchData() }
const handleSizeChange = (val) => { pageSize.value = val; fetchData() }

const handleAdd = () => { router.push({ name: 'activityEdit' }) }
const handleEdit = (row) => { router.push({ name: 'activityEdit', query: { id: row.id } }) }
const handleFromTemplate = () => { router.push({ name: 'activityTemplate' }) }
const handleClone = (row) => { router.push({ name: 'activityEdit', query: { clone: row.id } }) }

const handlePublish = async (row) => {
  await ElMessageBox.confirm('确认发布该活动?', '提示')
  const res = await activityPublish({ id: row.id })
  if (res.code === 0) { ElMessage.success('发布成功'); fetchData() }
}

const handleOffline = async (row) => {
  await ElMessageBox.confirm('确认下线该活动?', '提示')
  const res = await activityOffline({ id: row.id })
  if (res.code === 0) { ElMessage.success('下线成功'); fetchData() }
}

onMounted(() => { fetchData() })
</script>
