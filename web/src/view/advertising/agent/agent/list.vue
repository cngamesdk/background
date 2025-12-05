<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="渠道ID/渠道名称">
          <el-input v-model="searchInfo.agent_name" placeholder="渠道ID/渠道名" />
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
        >新增渠道</el-button
        >
      </div>
      <el-table :data="tableData" stripe row-key="id">
        <el-table-column align="left" label="ID" min-width="50" prop="id" />
        <el-table-column
            align="left"
            label="平台"
            min-width="100"
            prop="platform_id"
        />
        <el-table-column
            align="left"
            label="渠道组"
            min-width="100"
            prop="channel_group_id"
        />
        <el-table-column
            align="left"
            label="渠道名称"
            min-width="100"
            prop="agent_name"
        />
        <el-table-column
            align="left"
            label="结算类型"
            min-width="150"
            prop="settlement_type"
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
        <el-form-item label="渠道组" prop="channel_group_id">
          <el-select v-model="configInfo.channel_group_id"
                     placeholder="请选择所属渠道组"
                     filterable
                     remote
                     reserve-keyword
                     :remote-method="getChannelGroups"
                     :loading="loading"
                     style="width: 240px">
            <el-option
                v-for="item in channelGroups"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="渠道名称" prop="agent_name">
          <el-input v-model="configInfo.agent_name" placeholder="请输入渠道名称"/>
        </el-form-item>
        <el-form-item label="结算类型" prop="settlement_type">
          <el-select v-model="configInfo.settlement_type" placeholder="请选择结算类型" style="width: 240px">
            <el-option
                v-for="item in settlementTypes"
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

import { agentList, agentAdd, agentModify } from '@/api/advertising'
import { searchPlatform, searchChannelGroup, searchSettlementType } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'AgentList'
})

const appStore = useAppStore()

const searchInfo = ref({
  agent_name: '',
})

//游戏对话框
const configDialog = ref({
  show: false,
  add: true
})
const defaultData = {
  id: 0,
  platform_id: 0,
  channel_group_id: 0,
  agent_name: '',
  settlement_type: '',
}
//游戏信息
const configInfo = ref(defaultData)

const platforms = ref([])
const channelGroups = ref([])
const settlementTypes = ref([])

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
}

const onSearchSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {
    channel_group_name: '',
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

const getChannelGroups = async (query) => {
  const table = await searchChannelGroup({keyword: query})
  if (table.code === 0) {
    channelGroups.value = table.data
  }
}

const getSettlementTypes = async () => {
  const table = await searchSettlementType()
  if (table.code === 0) {
    settlementTypes.value = table.data
  }
}

// 查询
const getTableData = async () => {
  const table = await agentList({
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

watch(
    () => tableData.value,
)

const initPage = async () => {
  getTableData()
  getPlatforms()
  getSettlementTypes()
}

initPage()

const configForm = ref(null)
const rules = ref({
  platform_id: [
    { required: true, message: '请选择平台', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择平台', trigger: 'blur' }
  ],
  channel_group_id: [
    { required: true, message: '请选择渠道组', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择渠道组', trigger: 'blur' }
  ],
  agent_name: [
    { required: true, message: '请输入渠道组名称', trigger: 'blur' },
  ],
  settlement_type: [
    { required: true, message: '请选择结算类型', trigger: 'blur' },
  ]
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await agentAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await agentModify(configInfo.value)
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
