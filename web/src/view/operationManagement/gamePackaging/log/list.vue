<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="平台">
          <el-select
              filterable
              v-model="searchInfo.platform_id" placeholder="请选择所属平台">
            <el-option
                v-for="item in platforms"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="游戏">
          <el-select
              remote
              :remote-method="handleRemoteSearchGame"
              filterable
              v-model="searchInfo.game_id" placeholder="请选择游戏">
            <el-option
                v-for="item in games"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="广告位">
          <el-select
              remote
              :remote-method="handleRemoteSearchSiteId"
              filterable
              v-model="searchInfo.site_id" placeholder="请选择广告位">
            <el-option
                v-for="item in sites"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
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
        <el-button type="primary" icon="plus" @click="openAddConfigDialog"
        >新增打包</el-button>
      </div>
      <el-table :data="tableData" stripe row-key="id">
        <el-table-column
            align="left"
            label="ID"
            min-width="100"
            prop="id"
        />
        <el-table-column
            align="left"
            label="平台"
            min-width="100">
          <template #default="scope">
            {{ scope.row.platform_id }}({{ scope.row.platform_name }})
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="游戏"
            min-width="100">
          <template #default="scope">
            {{ scope.row.game_id }}({{ scope.row.game_name }})
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="渠道ID">
          <template #default="scope">
            {{ scope.row.agent_id }}({{ scope.row.agent_name }})
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="广告位ID">
          <template #default="scope">
            {{ scope.row.site_id }}({{ scope.row.site_name }})
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="状态"
            min-width="150"
            prop="status"
        />
        <el-table-column
            align="left"
            label="安装包地址"
            min-width="150"
            prop="game_package_path"
        />
        <el-table-column
            align="left"
            label="执行命令"
            min-width="150"
            prop="exec_cmd"
        />
        <el-table-column
            align="left"
            label="执行结果"
            min-width="150"
            prop="exec_cmd_result"
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
<!--        <el-table-column label="操作" fixed="right">-->
<!--          <template #default="scope">-->
<!--            <el-button-->
<!--                type="primary"-->
<!--                link-->
<!--                icon="edit"-->
<!--                @click="openConfigEditDialog(scope.row)"-->
<!--            >编辑</el-button>-->
<!--          </template>-->
<!--        </el-table-column>-->
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
        @close="closeConfigDialog"
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
        <el-form-item label="游戏" prop="game_id">
          <el-select
              filterable
              remote
              :remote-method="handleRemoteSearchGame"
              v-model="configInfo.game_id"
              placeholder="请选择游戏" style="width: 240px">
            <el-option
                v-for="item in games"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="广告位" prop="site_id">
          <el-select
              filterable
              remote
              :remote-method="handleRemoteSearchSiteId"
              v-model="configInfo.site_id"
              placeholder="请选择媒体" style="width: 240px">
            <el-option
                v-for="item in sites"
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

import { gamePackagingLogList, gamePackagingAdd } from '@/api/operationManagement'
import { searchPlatform, searchSubGame, searchSite } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'
import { UploadFilled } from '@element-plus/icons-vue'

defineOptions({
  name: 'GamePackagingLogList'
})

const appStore = useAppStore()

const defaultSearchInfo = {
  platform_id: 0,
  game_id: 0,
  site_id: 0,
}

const searchInfo = ref(defaultSearchInfo)

//游戏对话框
const configDialog = ref({
  show: false,
  add: true
})

const defaultConfigInfo = {
  id: 0,
  platform_id: 0,
  game_id: 0,
  site_id: 0,
}

//游戏信息
const configInfo = ref(defaultConfigInfo)

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
  configInfo.value = defaultConfigdefaultConfigInfoInfo
  configDialog.value.show = false
}

const onSearchSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = defaultSearchInfo
  getTableData()
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const platforms = ref([])
const games = ref([])
const sites = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleRemoteSearchGame = async (query) => {
  const result = await searchSubGame({keyword: query})
  if (result.code === 0) {
    games.value = result.data
  }
}

const handleRemoteSearchSiteId = async (query) => {
  const result = await searchSite({keyword: query})
  if (result.code === 0) {
    sites.value = result.data
  }
}

// 查询
const getTableData = async () => {
  const table = await gamePackagingLogList({
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
  const table = await searchPlatform()
  if (table.code === 0) {
    platforms.value = table.data
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
  game_id: [
    { required: true, message: '请选择游戏', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择游戏', trigger: 'blur' }
  ],
  site_id: [
    { required: true, message: '请选择广告位', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择广告位', trigger: 'blur' }
  ],
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await gamePackagingAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
      }
    }
  })
}

</script>