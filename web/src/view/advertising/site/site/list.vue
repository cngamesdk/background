<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="广告位ID/广告位名称">
          <el-input v-model="searchInfo.site_name" placeholder="广告位ID/广告位名称" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearchSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset"> 重置 </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openAddConfigDialog">新增广告位</el-button>
      </div>
      <el-table :data="tableData" stripe row-key="id">
        <el-table-column align="left" label="ID" min-width="50" prop="id" />
        <el-table-column
            align="left"
            label="平台"
            min-width="100">
          <template #default="scope">
            {{ scope.row.platform_id }}-{{ scope.row.platform_name }}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="渠道"
            min-width="100">
          <template #default="scope">
            {{ scope.row.agent_id }}-{{ scope.row.agent_name }}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="广告位名称"
            min-width="150"
            prop="site_name"
        />
        <el-table-column
            align="left"
            label="创建时间"
            min-width="180"
        >
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="更新时间"
            min-width="180"
        >
          <template #default="scope">
            {{ formatDate(scope.row.updated_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right">
        <template #default="scope">
          <el-button
              type="primary"
              link
              icon="edit"
              @click="openConfigEditDialog(scope.row)"
          >编辑</el-button>
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

    <!-- 游戏添加/编辑对话框 -->
    <el-dialog
        v-model="configDialog.show"
        title="配置"
        width="700px"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-form ref="configForm" :model="configInfo" :rules="rules" label-width="auto">
        <el-form-item label="平台" prop="platform_id">
          <el-select v-model="configInfo.platform_id" placeholder="请选择所属平台" style="width: 240px">
            <el-option
                v-for="item in platforms"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="渠道" prop="agent_id">
          <el-select v-model="configInfo.agent_id"
                     placeholder="请选择所属渠道"
                     filterable
                     remote
                     reserve-keyword
                     :remote-method="getAgents"
                     :loading="loading"
                     style="width: 240px">
            <el-option
                v-for="item in agents"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="广告位名称" prop="site_name">
          <el-input v-model="configInfo.site_name" placeholder="请输入广告位名称"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeConfigDialog">取 消</el-button>
          <el-button type="primary" @click="submitConfigInfo">提 交</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import { siteList, siteAdd, siteModify } from '@/api/advertising'
import { searchPlatform, searchAgent } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'SiteList'
})

const appStore = useAppStore()

const searchInfo = ref({
  site_name: '',
})

const configDialog = ref({
  show: false,
  add: true
})

const defaultData = {
  id: 0,
  platform_id: 0,
  agent_id: 0,
  site_name: '',
}

const configInfo = ref(defaultData)

const platforms = ref([])
const agents = ref([])

const openAddConfigDialog = () => {
  configDialog.value.show = true
  configDialog.value.add = true
}

const openConfigEditDialog = (row) => {
  configInfo.value = row
  configDialog.value.show = true
  configDialog.value.add = false
}

const closeConfigDialog = () => {
  configDialog.value.show = false
  configInfo.value = defaultData
  console.log(configInfo.value)
}

const onSearchSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {
    site_name: '',
  }
  getTableData()
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const getPlatforms = async () => {
  const table = await searchPlatform()
  if (table.code === 0) {
    platforms.value = table.data
  }
}

const getAgents = async (query) => {
  const table = await searchAgent({keyword: query})
  if (table.code === 0) {
    agents.value = table.data
  }
}

// 查询
const getTableData = async () => {
  const table = await siteList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const initPage = async () => {
  getTableData()
  getPlatforms()
}

initPage()

const configForm = ref(null)
const rules = ref({
  platform_id: [
    { required: true, message: '请选择平台', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择平台', trigger: 'blur' }
  ],
  agent_id: [
    { required: true, message: '请选择渠道', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择渠道', trigger: 'blur' }
  ],
  site_name: [
    { required: true, message: '请输入广告位名称', trigger: 'blur' },
  ]
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await siteAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await siteModify(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeConfigDialog()
        }
      }
    }
  })
}

</script>

<style lang="scss">
.header-img-box {
  @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
}
</style>
