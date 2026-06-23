<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
        <el-form-item label="Token名称">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" clearable />
        </el-form-item>
        <el-form-item label="类型">
          <el-select v-model="searchInfo.type" placeholder="请选择" clearable>
            <el-option label="全部" :value="0" />
            <el-option label="企业" :value="1" />
            <el-option label="个人" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
            <el-option label="全部" :value="0" />
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="2" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增Token</el-button>
      </div>
      <el-table
        ref="multipleTable"
        v-loading="loading"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
      >
        <el-table-column align="left" label="ID" prop="id" width="80" />
        <el-table-column align="left" label="名称" prop="name" width="150" />
        <el-table-column align="left" label="Token" prop="token" min-width="300">
          <template #default="scope">
            <el-input v-model="scope.row.token" readonly>
              <template #append>
                <el-button icon="document-copy" @click="copyToken(scope.row.token)" />
              </template>
            </el-input>
          </template>
        </el-table-column>
        <el-table-column align="left" label="类型" prop="type" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.type === 1 ? 'primary' : 'success'">
              {{ scope.row.type === 1 ? '企业' : '个人' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="使用量/限额" prop="usedTokens" width="150">
          <template #default="scope">
            {{ scope.row.usedTokens }} / {{ scope.row.tokenLimit || '无限制' }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="限流(次/分)" prop="requestLimit" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="100">
          <template #default="scope">
            <el-tag :type="scope.row.status === 1 ? 'success' : 'danger'">
              {{ scope.row.status === 1 ? '启用' : '禁用' }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="过期时间" prop="expireAt" width="180" />
        <el-table-column align="left" label="操作" fixed="right" width="280">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="viewDetail(scope.row)">详情</el-button>
            <el-button type="primary" link icon="edit" @click="updateTokenRow(scope.row)">编辑</el-button>
            <el-button type="warning" link icon="refresh" @click="regenerate(scope.row)">重新生成</el-button>
            <el-button type="danger" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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

    <!-- 表单对话框 -->
    <el-dialog v-model="dialogFormVisible" :title="type === 'create' ? '新增Token' : '编辑Token'" destroy-on-close>
      <el-form :model="formData" :rules="rules" label-position="right" label-width="100px">
        <el-form-item label="名称" prop="name" required>
          <el-input v-model="formData.name" clearable placeholder="请输入Token名称" />
        </el-form-item>
        <el-form-item label="类型" prop="type" required>
          <el-radio-group v-model="formData.type">
            <el-radio :label="1">企业</el-radio>
            <el-radio :label="2">个人</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="Token限额" prop="tokenLimit">
          <el-input v-model.number="formData.tokenLimit" type="number" placeholder="0表示无限制">
            <template #append>tokens</template>
          </el-input>
        </el-form-item>
        <el-form-item label="请求限流" prop="requestLimit">
          <el-input v-model.number="formData.requestLimit" type="number" placeholder="0表示无限制">
            <template #append>次/分钟</template>
          </el-input>
        </el-form-item>
        <el-form-item label="过期时间" prop="expireAt">
          <el-date-picker
            v-model="formData.expireAt"
            type="datetime"
            placeholder="选择日期时间"
            format="YYYY-MM-DD HH:mm:ss"
            value-format="YYYY-MM-DD HH:mm:ss"
            style="width: 100%"
          />
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-radio-group v-model="formData.status">
            <el-radio :label="1">启用</el-radio>
            <el-radio :label="2">禁用</el-radio>
          </el-radio-group>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取消</el-button>
          <el-button type="primary" @click="enterDialog">确定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 详情对话框 -->
    <el-dialog v-model="detailVisible" title="Token详情" width="700px">
      <el-descriptions :column="2" border>
        <el-descriptions-item label="ID">{{ detailData.id }}</el-descriptions-item>
        <el-descriptions-item label="名称">{{ detailData.name }}</el-descriptions-item>
        <el-descriptions-item label="Token" :span="2">
          <el-input v-model="detailData.token" readonly>
            <template #append>
              <el-button icon="document-copy" @click="copyToken(detailData.token)" />
            </template>
          </el-input>
        </el-descriptions-item>
        <el-descriptions-item label="类型">
          <el-tag :type="detailData.type === 1 ? 'primary' : 'success'">
            {{ detailData.type === 1 ? '企业' : '个人' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="状态">
          <el-tag :type="detailData.status === 1 ? 'success' : 'danger'">
            {{ detailData.status === 1 ? '启用' : '禁用' }}
          </el-tag>
        </el-descriptions-item>
        <el-descriptions-item label="已使用">{{ detailData.usedTokens }}</el-descriptions-item>
        <el-descriptions-item label="Token限额">{{ detailData.tokenLimit || '无限制' }}</el-descriptions-item>
        <el-descriptions-item label="请求限流">{{ detailData.requestLimit || '无限制' }} 次/分钟</el-descriptions-item>
        <el-descriptions-item label="今日请求">{{ detailData.todayRequests }}</el-descriptions-item>
        <el-descriptions-item label="今日Token">{{ detailData.todayTokens }}</el-descriptions-item>
        <el-descriptions-item label="使用率">
          <el-progress :percentage="detailData.usagePercent || 0" :color="getProgressColor(detailData.usagePercent)" />
        </el-descriptions-item>
        <el-descriptions-item label="创建时间">{{ detailData.createdAt }}</el-descriptions-item>
        <el-descriptions-item label="过期时间">{{ detailData.expireAt || '永不过期' }}</el-descriptions-item>
      </el-descriptions>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { getTokenList, createToken, updateToken, deleteToken, getTokenDetail, regenerateToken } from '@/api/modelTransfer'

// 搜索条件
const searchInfo = ref({
  name: '',
  type: 0,
  status: 0
})

// 表格数据
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const loading = ref(false)

// 对话框
const dialogFormVisible = ref(false)
const type = ref('')
const formData = ref({
  name: '',
  type: 1,
  tokenLimit: 0,
  requestLimit: 0,
  expireAt: '',
  status: 1
})

// 表单验证规则
const rules = ref({
  name: [{ required: true, message: '请输入Token名称', trigger: 'blur' }],
  type: [{ required: true, message: '请选择类型', trigger: 'change' }]
})

// 详情对话框
const detailVisible = ref(false)
const detailData = ref({})

// 获取列表
const getTableData = async() => {
  loading.value = true
  try {
    const res = await getTokenList({
      page: page.value,
      pageSize: pageSize.value,
      ...searchInfo.value
    })
    if (res.code === 0) {
      tableData.value = res.data.list || []
      total.value = res.data.total || 0
      page.value = res.data.page || page.value
      pageSize.value = res.data.pageSize || pageSize.value
    }
  } catch (error) {
    ElMessage.error('获取列表失败: ' + error.message)
  } finally {
    loading.value = false
  }
}

// 查询
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {
    name: '',
    type: 0,
    status: 0
  }
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

// 打开对话框
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭对话框
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    type: 1,
    tokenLimit: 0,
    requestLimit: 0,
    expireAt: '',
    status: 1
  }
}

// 确定
const enterDialog = async() => {
  try {
    if (type.value === 'create') {
      const res = await createToken(formData.value)
      if (res.code === 0) {
        ElMessage.success('创建成功')
        closeDialog()
        getTableData()
      }
    } else {
      const res = await updateToken(formData.value)
      if (res.code === 0) {
        ElMessage.success('更新成功')
        closeDialog()
        getTableData()
      }
    }
  } catch (error) {
    ElMessage.error('操作失败: ' + error.message)
  }
}

// 编辑
const updateTokenRow = (row) => {
  formData.value = { ...row }
  type.value = 'update'
  dialogFormVisible.value = true
}

// 删除
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    try {
      const res = await deleteToken({ id: row.id })
      if (res.code === 0) {
        ElMessage.success('删除成功')
        getTableData()
      }
    } catch (error) {
      ElMessage.error('删除失败: ' + error.message)
    }
  })
}

// 查看详情
const viewDetail = async(row) => {
  try {
    const res = await getTokenDetail({ id: row.id })
    if (res.code === 0) {
      detailData.value = res.data
      detailVisible.value = true
    }
  } catch (error) {
    ElMessage.error('获取详情失败: ' + error.message)
  }
}

// 重新生成
const regenerate = (row) => {
  ElMessageBox.confirm('重新生成Token后，旧Token将失效，确定要继续吗?', '警告', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
    try {
      const res = await regenerateToken({ id: row.id })
      if (res.code === 0) {
        ElMessage.success('重新生成成功，新Token：' + res.data.token)
        getTableData()
      }
    } catch (error) {
      ElMessage.error('重新生成失败: ' + error.message)
    }
  })
}

// 复制Token
const copyToken = (token) => {
  navigator.clipboard.writeText(token).then(() => {
    ElMessage.success('复制成功')
  }).catch(() => {
    ElMessage.error('复制失败')
  })
}

// 进度条颜色
const getProgressColor = (percent) => {
  if (percent < 50) return '#67c23a'
  if (percent < 80) return '#e6a23c'
  return '#f56c6c'
}

// 初始化
getTableData()
</script>

<style scoped>
.demo-form-inline {
  display: flex;
  flex-wrap: wrap;
}
</style>
