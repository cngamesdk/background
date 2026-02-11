<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="配置名称">
          <el-input v-model="searchInfo.name" placeholder="配置名称" />
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
        <el-button type="primary" icon="plus" @click="openAddConfigDialog">新增配置</el-button>
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
            label="名称"
            prop="name"
        />
        <el-table-column
            align="left"
            label="媒体"
            min-width="150">
          <template #default="scope">
            {{ scope.row.code }}-{{ scope.row.media_name }}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="主体"
            min-width="150">
          <template #default="scope">
            {{ scope.row.company_id }}-{{ scope.row.company_name }}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="应用ID"
            prop="app_id"
        />
        <el-table-column
            align="left"
            label="密钥"
            prop="secret"
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
        <el-form-item label="配置名称" prop="name">
          <el-input v-model="configInfo.name" placeholder="请输入配置名称"/>
        </el-form-item>
        <el-form-item label="媒体" prop="code">
          <el-select v-model="configInfo.code" placeholder="请选择媒体" style="width: 240px">
            <el-option
                v-for="item in medias"
                :key="item.key_2"
                :label="item.value"
                :value="item.key_2"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="主体" prop="company_id">
          <el-select v-model="configInfo.company_id" placeholder="请选择主体" style="width: 240px">
            <el-option
                v-for="item in companys"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="应用ID(app_id)" prop="app_id">
          <el-input v-model="configInfo.app_id" placeholder="请输入应用ID"/>
        </el-form-item>
        <el-form-item label="配置密钥" prop="secret">
          <el-input v-model="configInfo.secret" placeholder="请输入密钥"/>
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

import { advertisingDeveloperConfigList, advertisingDeveloperConfigAdd, advertisingDeveloperConfigModify } from '@/api/advertising'
import { searchPlatform, searchMedia, searchCompany } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'AdvertisingDeveloperConfigList'
})

const appStore = useAppStore()

const searchInfo = ref({
  name: '',
})

//游戏对话框
const configDialog = ref({
  show: false,
  add: true
})
//游戏信息
const configInfo = ref({
  id: 0,
  platform_id: 0,
  name: '',
  code: '',
  company_id: 0,
  app_id: '',
  secret: '',
})

const platforms = ref([])
const medias = ref([])
const companys = ref([])

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
}

const onSearchSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {
    advertising_media_name: '',
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

const getMedias = async () => {
  const table = await searchMedia()
  if (table.code === 0) {
    medias.value = table.data
  }
}

const getCompanys = async () => {
  const table = await searchCompany()
  if (table.code === 0) {
    companys.value = table.data
  }
}

// 查询
const getTableData = async () => {
  const table = await advertisingDeveloperConfigList({
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
  getMedias()
  getCompanys()
}

initPage()

const configForm = ref(null)
const rules = ref({
  platform_id: [
    { required: true, message: '请选择平台', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择平台', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入配置名称', trigger: 'blur' },
  ],
  code: [
    { required: true, message: '请选择媒体', trigger: 'blur' },
  ],
  company_id: [
    { required: true, message: '请选择主体', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择主体', trigger: 'blur' }
  ],
  app_id: [
    { required: true, message: '请输入应用ID', trigger: 'blur' },
  ],
  secret: [
    { required: true, message: '请输入密钥', trigger: 'blur' },
  ],
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await advertisingDeveloperConfigAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await advertisingDeveloperConfigModify(configInfo.value)
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
