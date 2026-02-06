<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item>
          <el-input v-model="searchInfo.theme_name" placeholder="题材名称" />
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
        <el-button type="primary" icon="plus" @click="openAddConfigDialog">新增题材</el-button>
      </div>
      <el-table
          :data="tableData"
          stripe row-key="id">
        <el-table-column align="left" label="ID" min-width="50" prop="id" />
        <el-table-column
            align="left"
            label="平台"
            min-width="100">
          <template #default="scope">
            {{ scope.row.platform_id }}-{{scope.row.platform_name}}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="题材名称"
            min-width="150"
            prop="theme_name"
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
        <el-table-column label="操作" min-width="250">
        <template #default="scope">
          <el-button
              type="primary"
              link
              icon="edit"
              @click="openConfigEditDialog(scope.row)"
          >编辑</el-button>
          <el-button
              type="primary"
              link
              icon="view"
              @click="handleSubTheme(scope.row)"
          >子题材</el-button>
          <el-button
              type="primary"
              link
              icon="plus"
              @click="handleAddSubTheme(scope.row)"
          >添加子题材</el-button>
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

    <el-dialog
        v-model="subThemeDialog.show"
        title="子题材"
        width="700px"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-table :data="tableDataChildren">
        <el-table-column label="ID" prop="id" />
        <el-table-column label="题材名称" prop="theme_name" />
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
            :current-page="pageChildren"
            :page-size="pageSizeChildren"
            :page-sizes="[10, 30, 50, 100]"
            :total="totalChildren"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChangeChildren"
            @size-change="handleSizeChangeChildren"
        />
      </div>
    </el-dialog>

    <!-- 游戏添加/编辑对话框 -->
    <el-dialog
        v-model="configDialog.show"
        title="题材配置"
        width="700px"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-form ref="configForm" :model="configInfo" :rules="rules" label-width="80px">
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
        <el-form-item label="题材名称" prop="theme_name">
          <el-input v-model="configInfo.theme_name" placeholder="请输入题材名称"/>
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

import { materialThemeList, materialThemeAdd, materialThemeModify } from '@/api/material'
import { searchPlatform } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'MaterialThemeList'
})

const appStore = useAppStore()

const searchInfo = ref({
  theme_name: '',
})

//游戏对话框
const configDialog = ref({
  show: false,
  add: true
})

const subThemeDialog = ref({
  show: false,
  add: true,
  row: {}
})

const defaultConfigInfo = {
  id: 0,
  platform_id: 0,
  theme_name: '',
  parent_id: 0,
}

//游戏信息
const configInfo = ref(Object.assign({}, defaultConfigInfo))

const platforms = ref([])

const openAddConfigDialog = () => {
  configInfo.value = Object.assign({}, defaultConfigInfo)
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
}

const onSearchSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {
    theme_name: '',
  }
  getTableData()
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

const pageChildren = ref(1)
const totalChildren = ref(0)
const pageSizeChildren = ref(10)
const tableDataChildren = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleAddSubTheme = (row) => {
  openAddConfigDialog()
  configInfo.value.parent_id = row.id
}

// 子题材分页
const handleSizeChangeChildren = (val) => {
  pageSizeChildren.value = val
  getTableData()
}

const handleCurrentChangeChildren = (val) => {
  pageChildren.value = val
  getTableData()
}

// 查询游戏类型
const getPlatforms = async () => {
  const table = await searchPlatform()
  if (table.code === 0) {
    platforms.value = table.data
  }
}

const handleSubTheme = (row) => {
  subThemeDialog.value.show = true
  subThemeDialog.value.row = row
  getSubThemeList()
}

const getSubThemeList = async () => {
  const table = await materialThemeList({
    page: page.value,
    pageSize: pageSize.value,
    parent_id: subThemeDialog.value.row.id,
    ...searchInfo.value
  })
  if (table.code === 0) {
    tableDataChildren.value = table.data.list
    totalChildren.value = table.data.total
    pageChildren.value = table.data.page
    pageSizeChildren.value = table.data.pageSize
  }
}

// 查询
const getTableData = async () => {
  const table = await materialThemeList({
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
  theme_name: [
    { required: true, message: '请输入题材名称', trigger: 'blur' },
  ],
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await materialThemeAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await materialThemeModify(configInfo.value)
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
