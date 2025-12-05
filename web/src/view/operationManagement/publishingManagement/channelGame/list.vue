<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openAddConfigDialog"
        >新增发行游戏配置</el-button
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
            label="子游戏"
            min-width="100"
            prop="game_id"
        />
        <el-table-column
            align="left"
            label="发行渠道"
            min-width="150"
            prop="channel_id"
        />
        <el-table-column
            align="left"
            label="渠道ID"
            min-width="150"
            prop="agent_id"
        />
        <el-table-column
            align="left"
            label="广告位ID"
            min-width="150"
            prop="site_id"
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
        title="主游戏配置"
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
        <el-form-item label="发行游戏" prop="game_id">
          <el-select v-model="configInfo.game_id"
                     placeholder="请选择绑定子游戏"
                     filterable
                     remote
                     reserve-keyword
                     :remote-method="getGames"
                     :loading="loading"
                     style="width: 240px">
            <el-option
                v-for="item in games"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="所属渠道" prop="channel_id">
          <el-select v-model="configInfo.channel_id"
                     placeholder="请选择所属渠道"
                     filterable
                     remote
                     reserve-keyword
                     :remote-method="getChannels"
                     :loading="loading"
                     style="width: 240px">
            <el-option
                v-for="item in channels"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="渠道ID" prop="agent_id">
          <el-select v-model="configInfo.agent_id"
                     placeholder="请选择渠道ID"
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
        <el-form-item label="广告位ID" prop="site_id">
          <el-select v-model="configInfo.site_id"
                     placeholder="请选择广告位ID"
                     filterable
                     remote
                     reserve-keyword
                     :remote-method="getSites"
                     :loading="loading"
                     style="width: 240px">
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

import { publishingChannelGameList, publishingChannelGameAdd, publishingChannelGameModify } from '@/api/operationManagement'
import { searchPlatform, searchSubGame, searchPublishingChannel, searchAgent, searchSite } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'ChannelGameList'
})

const appStore = useAppStore()

const searchInfo = ref({})

//游戏对话框
const configDialog = ref({
  show: false,
  add: true
})
//游戏信息
const configInfo = ref({
  id: 0,
  platform_id: 0,
  game_id: 0,
  channel_id: 0,
  agent_id: 0,
  site_id: 0,
})

const platforms = ref([])
const games = ref([])
const channels = ref([])
const agents = ref([])
const sites = ref([])

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

const getGames = async (query) => {
  const table = await searchSubGame({keyword: query})
  if (table.code === 0) {
    games.value = table.data
  }
}

const getChannels = async (query) => {
  const table = await searchPublishingChannel({keyword: query})
  if (table.code === 0) {
    channels.value = table.data
  }
}

const getAgents = async (query) => {
  const table = await searchAgent({keyword: query})
  if (table.code === 0) {
    agents.value = table.data
  }
}

const getSites = async (query) => {
  const table = await searchSite({keyword: query})
  if (table.code === 0) {
    sites.value = table.data
  }
}

// 查询
const getTableData = async () => {
  const table = await publishingChannelGameList({
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
}

initPage()

const configForm = ref(null)
const rules = ref({
  platform_id: [
    { required: true, message: '请选择平台', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择平台', trigger: 'blur' }
  ],
  game_id: [
    { required: true, message: '请选择发行游戏', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择发行游戏', trigger: 'blur' }
  ],
  channel_id: [
    { required: true, message: '请选择发行渠道', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择发行渠道', trigger: 'blur' }
  ],
  agent_id: [
    { required: true, message: '请选择渠道ID', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择渠道ID', trigger: 'blur' }
  ],
  site_id: [
    { required: true, message: '请选择广告位ID', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择广告位ID', trigger: 'blur' }
  ],
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await publishingChannelGameAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await publishingChannelGameModify(configInfo.value)
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
