<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="标题">
          <el-input v-model="searchInfo.title" placeholder="标题" />
        </el-form-item>
        <el-form-item label="短码">
          <el-input v-model="searchInfo.short_code" placeholder="短码" />
        </el-form-item>
        <el-form-item label="状态">
          <el-select v-model="searchInfo.status" placeholder="全部" clearable>
            <el-option label="启用" :value="1" />
            <el-option label="禁用" :value="0" />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearch">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openCreateDialog">创建短链接</el-button>
      </div>
      <el-table :data="tableData" stripe row-key="id">
        <el-table-column align="left" label="ID" width="80" prop="id" />
        <el-table-column align="left" label="标题" min-width="150" prop="title" />
        <el-table-column align="left" label="短码" width="120" prop="short_code" />
        <el-table-column align="left" label="短链接" min-width="200">
          <template #default="scope">
            <el-link type="primary" :href="scope.row.domain + '/' + scope.row.short_code" target="_blank">
              {{ scope.row.domain }}/{{ scope.row.short_code }}
            </el-link>
          </template>
        </el-table-column>
        <el-table-column align="left" label="原始链接" min-width="250" show-overflow-tooltip>
          <template #default="scope">
            <el-link type="info" :href="scope.row.original_url" target="_blank">
              {{ scope.row.original_url }}
            </el-link>
          </template>
        </el-table-column>
        <el-table-column align="center" label="状态" width="80">
          <template #default="scope">
            <el-switch
              v-model="scope.row.status"
              :active-value="1"
              :inactive-value="0"
              @change="handleStatusChange(scope.row)"
            />
          </template>
        </el-table-column>
        <el-table-column align="left" label="点击量" width="80" prop="total_clicks" />
        <el-table-column align="left" label="过期时间" width="170">
          <template #default="scope">
            <span v-if="scope.row.expire_at">{{ formatDate(scope.row.expire_at) }}</span>
            <el-tag v-else type="success" size="small">永久</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="创建时间" width="170">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width="200">
          <template #default="scope">
            <el-button type="primary" link icon="view" @click="openClickLog(scope.row)">统计</el-button>
            <el-button type="primary" link icon="edit" @click="openEditDialog(scope.row)">编辑</el-button>
            <el-button type="danger" link icon="delete" @click="handleDelete(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <!-- 创建对话框 -->
    <el-dialog v-model="createDialog.show" title="创建短链接" width="550px" :close-on-click-modal="false">
      <el-form ref="createFormRef" :model="createForm" :rules="createRules" label-width="100px">
        <el-form-item label="原始链接" prop="original_url">
          <el-input v-model="createForm.original_url" placeholder="请输入需要缩短的链接" />
        </el-form-item>
        <el-form-item label="标题" prop="title">
          <el-input v-model="createForm.title" placeholder="请输入备注标题（可选）" />
        </el-form-item>
        <el-form-item label="过期天数">
          <el-input-number v-model="createForm.expire_days" :min="0" :max="3650" />
          <span style="margin-left: 8px; color: #909399;">0 表示永久有效</span>
        </el-form-item>
      </el-form>
      <div v-if="createResult.short_url" style="margin-top: 16px; padding: 12px; background: #f0f9eb; border-radius: 4px;">
        <p style="margin: 0 0 8px; font-weight: bold; color: #67c23a;">创建成功</p>
        <p style="margin: 0;">
          短链接：
          <el-link type="primary" :href="createResult.short_url" target="_blank">{{ createResult.short_url }}</el-link>
          <el-button type="primary" link size="small" @click="copyUrl(createResult.short_url)" style="margin-left: 8px;">复制</el-button>
        </p>
      </div>
      <template #footer>
        <el-button @click="closeCreateDialog">关闭</el-button>
        <el-button type="primary" @click="submitCreate" :loading="createLoading">创建</el-button>
      </template>
    </el-dialog>

    <!-- 编辑对话框 -->
    <el-dialog v-model="editDialog.show" title="编辑短链接" width="500px">
      <el-form :model="editForm" label-width="80px">
        <el-form-item label="标题">
          <el-input v-model="editForm.title" placeholder="请输入标题" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="editDialog.show = false">取消</el-button>
        <el-button type="primary" @click="submitEdit">保存</el-button>
      </template>
    </el-dialog>

    <!-- 点击统计对话框 -->
    <el-dialog v-model="clickLogDialog.show" title="点击统计" width="80%">
      <el-table :data="clickLogData" stripe>
        <el-table-column label="访问IP" prop="ip" width="150" />
        <el-table-column label="User-Agent" prop="user_agent" min-width="300" show-overflow-tooltip />
        <el-table-column label="来源" prop="referer" min-width="200" show-overflow-tooltip />
        <el-table-column label="访问时间" width="170">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="clickLogPage"
          :page-size="clickLogPageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="clickLogTotal"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleClickLogPageChange"
          @size-change="handleClickLogSizeChange"
        />
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { shortLinkList, shortLinkCreate, shortLinkUpdate, shortLinkDelete, clickLogList } from '@/api/shortLink'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'ShortLinkList'
})

const searchInfo = ref({ title: '', short_code: '', status: null })
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const tableData = ref([])

const getTableData = async () => {
  const params = {
    page: page.value,
    page_size: pageSize.value,
    title: searchInfo.value.title,
    short_code: searchInfo.value.short_code,
  }
  if (searchInfo.value.status !== null && searchInfo.value.status !== '') {
    params.status = searchInfo.value.status
  }
  const res = await shortLinkList(params)
  if (res.code === 0) {
    tableData.value = res.data.list || []
    total.value = res.data.total
  }
}

const onSearch = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = { title: '', short_code: '', status: null }
  getTableData()
}

const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 创建短链接
const createDialog = ref({ show: false })
const createFormRef = ref(null)
const createForm = ref({ original_url: '', title: '', expire_days: 0 })
const createResult = ref({ short_url: '', short_code: '' })
const createLoading = ref(false)
const createRules = {
  original_url: [{ required: true, message: '请输入原始链接', trigger: 'blur' }],
}

const openCreateDialog = () => {
  createForm.value = { original_url: '', title: '', expire_days: 0 }
  createResult.value = { short_url: '', short_code: '' }
  createDialog.value.show = true
}

const closeCreateDialog = () => {
  createDialog.value.show = false
  if (createResult.value.short_url) {
    getTableData()
  }
}

const submitCreate = async () => {
  if (!createFormRef.value) return
  await createFormRef.value.validate(async (valid) => {
    if (!valid) return
    createLoading.value = true
    try {
      const res = await shortLinkCreate(createForm.value)
      if (res.code === 0) {
        createResult.value = res.data
        ElMessage.success('创建成功')
      }
    } finally {
      createLoading.value = false
    }
  })
}

const copyUrl = (url) => {
  navigator.clipboard.writeText(url).then(() => {
    ElMessage.success('已复制到剪贴板')
  })
}

// 状态切换
const handleStatusChange = async (row) => {
  const res = await shortLinkUpdate({ id: row.id, status: row.status })
  if (res.code === 0) {
    ElMessage.success('状态更新成功')
  } else {
    row.status = row.status === 1 ? 0 : 1
  }
}

// 编辑
const editDialog = ref({ show: false })
const editForm = ref({ id: 0, title: '' })

const openEditDialog = (row) => {
  editForm.value = { id: row.id, title: row.title }
  editDialog.value.show = true
}

const submitEdit = async () => {
  const res = await shortLinkUpdate(editForm.value)
  if (res.code === 0) {
    ElMessage.success('保存成功')
    editDialog.value.show = false
    getTableData()
  }
}

// 删除
const handleDelete = (row) => {
  ElMessageBox.confirm('确认删除该短链接？删除后无法恢复', '提示', {
    confirmButtonText: '确认',
    cancelButtonText: '取消',
    type: 'warning',
  }).then(async () => {
    const res = await shortLinkDelete({ id: row.id })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      getTableData()
    }
  })
}

// 点击统计
const clickLogDialog = ref({ show: false })
const clickLogData = ref([])
const clickLogPage = ref(1)
const clickLogPageSize = ref(10)
const clickLogTotal = ref(0)
const currentShortCode = ref('')

const openClickLog = (row) => {
  currentShortCode.value = row.short_code
  clickLogPage.value = 1
  clickLogDialog.value.show = true
  getClickLogData()
}

const getClickLogData = async () => {
  const res = await clickLogList({
    short_code: currentShortCode.value,
    page: clickLogPage.value,
    pageSize: clickLogPageSize.value,
  })
  if (res.code === 0) {
    clickLogData.value = res.data.list || []
    clickLogTotal.value = res.data.total
  }
}

const handleClickLogPageChange = (val) => {
  clickLogPage.value = val
  getClickLogData()
}

const handleClickLogSizeChange = (val) => {
  clickLogPageSize.value = val
  getClickLogData()
}

// 初始化
getTableData()
</script>
