<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="配置ID/配置名称">
          <el-input v-model="searchInfo.remark" placeholder="配置ID/配置名称" />
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
        <el-table-column
            align="left"
            label="ID"
            min-width="50"
            prop="id"
        />
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
            label="配置名称"
            min-width="150"
            prop="remark"
        />
        <el-table-column
            align="left"
            label="游戏"
            min-width="100">
          <template #default="scope">
            {{ scope.row.game_id }}-{{ scope.row.game_name }}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="APP版本号"
            min-width="100"
            prop="app_version_code"
        />
        <el-table-column
            align="left"
            label="APP版本名"
            min-width="150"
            prop="app_version_name"
        />
        <el-table-column
            align="left"
            label="归属配置"
            min-width="100">
          <template #default="scope">
            {{ scope.row.product_config_id }}-{{ scope.row.config_name }}
          </template>
        </el-table-column>
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
        title="游戏版本配置"
        width="700px"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-form ref="configForm" :model="configInfo" :rules="rules" label-width="auto">
        <el-form-item label="配置名称" prop="remark">
          <el-input v-model="configInfo.remark" placeholder="请输入配置名称"/>
        </el-form-item>
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
        <el-form-item label="绑定游戏" prop="game_id">
          <el-select v-model="configInfo.game_id"
                     placeholder="请选择绑定游戏"
                     filterable
                     remote
                     reserve-keyword
                     :remote-method="getSubGames"
                     :loading="loading"
                     style="width: 240px">
            <el-option
                v-for="item in subGames"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="APP版本号" prop="app_version_code">
          <el-input-number v-model="configInfo.app_version_code" placeholder="APP版本号,如：100001"/>
        </el-form-item>
        <el-form-item label="APP版本名" prop="app_version_name">
          <el-input v-model="configInfo.app_version_name" placeholder="APP版本名，如1.2.3"/>
        </el-form-item>
        <el-form-item label="绑定配置" prop="product_config_id">
          <el-select v-model="configInfo.product_config_id"
                     placeholder="请选择配置"
                     filterable
                     remote
                     reserve-keyword
                     :remote-method="getProductCommonConfigs"
                     :loading="loading"
                     style="width: 240px">
            <el-option
                v-for="item in productCommonConfigs"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
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

import { gameAppVersionConfigurationList, gameAppVersionConfigurationAdd, gameAppVersionConfigurationModify } from '@/api/operationManagement'
import { searchPlatform, searchSubGame, searchProductCommonConfig } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'GameAppVersionConfigurationList'
})

const appStore = useAppStore()

const searchInfo = ref({
  remark_name: '',
})

const configDialog = ref({
  show: false,
  add: true
})

const configInfo = ref({
  id: 0,
  platform_id: 0,
  game_id: 0,
  app_version_code: 0,
  app_version_name: '',
  remark: '',
  product_config_id: 0,
})

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
    remark: '',
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

// 查询
const getTableData = async () => {
  const table = await gameAppVersionConfigurationList({
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

const getPlatforms = async () => {
  const result = await searchPlatform()
  if (result.code === 0) {
    platforms.value = result.data
  }
}

const getSubGames = async (query) => {
  const result = await searchSubGame({keyword: query})
  if (result.code === 0) {
    subGames.value = result.data
  }
}

const getProductCommonConfigs = async (query) => {
  const result = await searchPlatform({keyword: query})
  if (result.code === 0) {
    productCommonConfigs.value = result.data
  }
}

watch(
    () => tableData.value,
)

const initPage = async () => {
  getTableData()
  getPlatforms()
}

initPage()

const platforms = ref([])
const subGames = ref([])
const productCommonConfigs = ref([])
const configForm = ref(null)
const rules = ref({
  platform_id: [
    { required: true, message: '请选择平台ID', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择平台', trigger: 'blur' }
  ],
  game_id: [
    { required: true, message: '请选择游戏', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择游戏', trigger: 'blur' }
  ],
  app_version_code: [
    { required: true, message: '请输入APP版本号', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请输入APP版本号', trigger: 'blur' }
  ],
  app_version_name: [
    { required: true, message: '请输入APP版本名', trigger: 'blur' }
  ],
  remark: [
    { required: true, message: '请输入配置名称', trigger: 'blur' },
  ],
  product_config_id: [
    { required: true, message: '请选择绑定配置', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择绑定配置', trigger: 'blur' }
  ]
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await gameAppVersionConfigurationAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await gameAppVersionConfigurationModify(configInfo.value)
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