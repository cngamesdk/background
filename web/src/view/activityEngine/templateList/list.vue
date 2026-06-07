<template>
  <div class="template-list-container">
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchForm">
        <el-form-item label="模板名称">
          <el-input v-model="searchForm.template_name" placeholder="请输入模板名称" clearable />
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
        <el-form-item>
          <el-button type="primary" @click="handleSearch">查询</el-button>
        </el-form-item>
      </el-form>
    </div>

    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" @click="showAddDialog = true">新增模板</el-button>
      </div>

      <el-table :data="tableData" border stripe>
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="template_name" label="模板名称" min-width="150" />
        <el-table-column prop="activity_type" label="活动类型" width="120">
          <template #default="{ row }">{{ typeMap[row.activity_type] || row.activity_type }}</template>
        </el-table-column>
        <el-table-column prop="description" label="说明" min-width="200" />
        <el-table-column prop="created_at" label="创建时间" width="170" />
        <el-table-column label="操作" width="150">
          <template #default="{ row }">
            <el-button size="small" type="primary" @click="handleClone(row)">一键创建活动</el-button>
          </template>
        </el-table-column>
      </el-table>

      <div class="gva-pagination">
        <el-pagination :current-page="page" :page-size="pageSize" :total="total"
          layout="total, prev, pager, next" @current-change="handlePageChange" />
      </div>
    </div>

    <!-- 新增模板弹窗 -->
    <el-dialog v-model="showAddDialog" title="新增模板" width="600px" destroy-on-close>
      <el-form :model="addForm" label-width="100px">
        <el-form-item label="模板名称" required>
          <el-input v-model="addForm.template_name" placeholder="请输入模板名称" />
        </el-form-item>
        <el-form-item label="活动类型" required>
          <el-select v-model="addForm.activity_type" placeholder="请选择活动类型">
            <el-option label="累计充值" value="recharge" />
            <el-option label="登录奖励" value="login" />
            <el-option label="签到" value="signin" />
            <el-option label="分享奖励" value="share" />
            <el-option label="自定义" value="custom" />
          </el-select>
        </el-form-item>
        <el-form-item label="模板说明">
          <el-input v-model="addForm.description" type="textarea" :rows="3" placeholder="请输入模板说明" />
        </el-form-item>
        <el-form-item label="触发条件">
          <el-input v-model="addForm.trigger_config" type="textarea" :rows="3" placeholder='{"event_type":"recharge","conditions":[]}' />
        </el-form-item>
        <el-form-item label="计算逻辑">
          <el-input v-model="addForm.calculation_config" type="textarea" :rows="3" placeholder='{"mode":"accumulate","field":"amount","reset_cycle":"never"}' />
        </el-form-item>
        <el-form-item label="奖励策略">
          <el-input v-model="addForm.reward_config" type="textarea" :rows="3" placeholder='{"strategy":"tiered","tiers":[]}' />
        </el-form-item>
        <el-form-item label="约束规则">
          <el-input v-model="addForm.constraint_config" type="textarea" :rows="3" placeholder='{"user_segments":["all"],"daily_claim_max":0,"total_claim_max":0}' />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="showAddDialog = false">取消</el-button>
        <el-button type="primary" @click="handleAdd" :loading="addLoading">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { templateList, templateAdd } from '@/api/activityEngine'

const router = useRouter()
const typeMap = { recharge: '累计充值', login: '登录奖励', signin: '签到', share: '分享', custom: '自定义' }

const searchForm = ref({ template_name: '', activity_type: '' })
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const showAddDialog = ref(false)
const addLoading = ref(false)
const addForm = ref({
  template_name: '',
  activity_type: '',
  description: '',
  trigger_config: '{"event_type":"","conditions":[]}',
  calculation_config: '{"mode":"accumulate","field":"count","reset_cycle":"never","dedup_key":""}',
  reward_config: '{"strategy":"tiered","tiers":[]}',
  constraint_config: '{"user_segments":["all"],"daily_claim_max":0,"total_claim_max":0,"cooldown_sec":0,"time_windows":[]}',
})

const fetchData = async () => {
  const res = await templateList({ ...searchForm.value, page: page.value, pageSize: pageSize.value })
  if (res.code === 0) {
    tableData.value = res.data.list || []
    total.value = res.data.total || 0
  }
}

const handleSearch = () => { page.value = 1; fetchData() }
const handlePageChange = (val) => { page.value = val; fetchData() }
const handleClone = (row) => { router.push({ name: 'activityEdit', query: { clone_template: row.id } }) }

const handleAdd = async () => {
  if (!addForm.value.template_name || !addForm.value.activity_type) {
    ElMessage.warning('请填写模板名称和活动类型')
    return
  }
  addLoading.value = true
  try {
    const res = await templateAdd(addForm.value)
    if (res.code === 0) {
      ElMessage.success('创建成功')
      showAddDialog.value = false
      addForm.value = {
        template_name: '', activity_type: '', description: '',
        trigger_config: '{"event_type":"","conditions":[]}',
        calculation_config: '{"mode":"accumulate","field":"count","reset_cycle":"never","dedup_key":""}',
        reward_config: '{"strategy":"tiered","tiers":[]}',
        constraint_config: '{"user_segments":["all"],"daily_claim_max":0,"total_claim_max":0,"cooldown_sec":0,"time_windows":[]}',
      }
      fetchData()
    }
  } finally {
    addLoading.value = false
  }
}

onMounted(() => { fetchData() })
</script>
