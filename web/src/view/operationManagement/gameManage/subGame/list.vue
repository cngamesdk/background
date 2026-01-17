<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="游戏ID/游戏名称">
          <el-input v-model="searchInfo.game_name" placeholder="游戏ID/游戏名称" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset"> 重置 </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openAddGameDialog"
        >新增游戏</el-button
        >
      </div>
      <el-table :data="tableData" stripe row-key="id">
        <el-table-column align="left" label="游戏ID" min-width="100" prop="id" />
        <el-table-column align="left" label="名称" min-width="100" prop="game_name" />
        <el-table-column align="left" label="所属平台" min-width="100">
          <template #default="scope">
            {{ scope.row.platform_id }}-{{ scope.row.platform_name }}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="包名"
            min-width="150"
            prop="package_name"
        />
        <el-table-column
            align="left"
            label="游戏类型"
            min-width="150"
            prop="game_type_str"
        />
        <el-table-column
            align="left"
            label="操作系统"
            min-width="180"
            prop="os_str"
        />
        <el-table-column
            align="left"
            label="所属主游戏"
            min-width="180">
          <template #default="scope">
            {{scope.row.main_game_id}}-{{scope.row.main_game_name}}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="兑换比例"
            min-width="180"
            prop="game_rate"
        />
        <el-table-column align="left" label="与研发对接游戏" min-width="150" prop="cp_game_id"></el-table-column>
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
        <el-table-column label="操作" :min-width="appStore.operateMinWith" fixed="right">
          <template #default="scope">
            <el-button
                type="primary"
                link
                icon="edit"
                @click="openEditDialog(scope.row)"
            >编辑</el-button>
            <el-button
                type="primary"
                link
                icon="view"
                @click="openConfigDialog(scope.row)"
            >查看配置</el-button>
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

    <!-- 获取游戏配置对话框 -->
    <el-dialog
        v-model="configData.show"
        title="游戏配置"
        width="700px"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-form :model="configData" label-width="0px">
        <el-form-item>
          <el-input
              v-model="configData.content"
              style="width: 650px"
              :rows="10"
              type="textarea"
              placeholder="Please input"
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeConfigDialog">取 消</el-button>
          <el-button type="primary" @click="copyConfig">复 制</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- 游戏添加/编辑对话框 -->
    <el-dialog
        v-model="gameInfoDialog.show"
        title="游戏配置"
        width="90%"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-form ref="configForm" :rules="rules" :model="gameInfo" label-width="auto">
        <el-form-item label="平台" prop="platform_id">
          <el-select v-model="gameInfo.platform_id" placeholder="请选择所属平台" style="width: 240px">
            <el-option
                v-for="item in platforms"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="游戏名称" prop="game_name">
          <el-input v-model="gameInfo.game_name" placeholder="请输入游戏名称"/>
        </el-form-item>
        <el-form-item label="包名" prop="package_name">
          <el-input v-model="gameInfo.package_name" placeholder="请输入包名，如：com.xxx.xxx"/>
        </el-form-item>
        <el-form-item label="应用ID" prop="app_id">
          <el-input v-model="gameInfo.app_id" placeholder="请输入应用ID，如：123456"/>
        </el-form-item>
        <el-form-item label="应用名称" prop="app_name">
          <el-input v-model="gameInfo.app_name" placeholder="请输入应用名称，如：显名"/>
        </el-form-item>
        <el-form-item label="游戏类型" prop="game_type">
          <el-select v-model="gameInfo.game_type" placeholder="请选择游戏类型" style="width: 240px">
            <el-option
                v-for="item in gameTypes"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="系统" prop="os">
          <el-select v-model="gameInfo.os" placeholder="请选择系统" style="width: 240px">
            <el-option
                v-for="item in gameOss"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="所属主游戏" prop="main_game_id">
          <el-select
              v-model="gameInfo.main_game_id"
              placeholder="请选择所属主游戏"
              filterable
              remote
              reserve-keyword
              :remote-method="searchMainGameMethod"
              :loading="loading"
              style="width: 240px">
            <el-option
                v-for="item in mainGames"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="游戏币名称" prop="game_coin_name">
          <el-input v-model="gameInfo.game_coin_name" placeholder="请输入游戏币名称，如：元宝"/>
        </el-form-item>
        <el-form-item label="游戏币兑换比例" prop="game_rate">
          <el-input-number :min="0" v-model="gameInfo.game_rate" placeholder="请输入游戏币兑换比例，如：10"/>
        </el-form-item>
        <el-form-item label="归属主体" prop="company_id">
          <el-select
              v-model="gameInfo.company_id"
              filterable
              remote
              reserve-keyword
              placeholder="请选择归属主体"
              :remote-method="searchCompanyMethod"
              :loading="loading"
              style="width: 240px"
          >
            <el-option
                v-for="item in companys"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="gameInfo.status" placeholder="请选择状态" style="width: 240px">
            <el-option
                v-for="item in gameStatuss"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="合作方式" prop="cooperation_model">
          <el-select v-model="gameInfo.cooperation_model" placeholder="请选择合作方式" style="width: 240px">
            <el-option
                v-for="item in gameCooperationModels"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="与研发对接游戏" prop="cp_game_id">
          <el-select v-model="gameInfo.cp_game_id" placeholder="与研发对接游戏" style="width: 240px"
                     filterable
                     remote
                     reserve-keyword
                     :remote-method="searchSubGameMethod"
                     :loading="loading"
          >
            <el-option :key="0" label="自身配置" :value="0"/>
            <el-option
                v-for="item in subGames"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeGameInfoDialog">取 消</el-button>
          <el-button type="primary" @click="submitGameInfo">提 交</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import { getSubGameList, getSubGameConfig, subGameAdd, subGameModify  } from '@/api/operationManagement'
import { searchPlatform, searchGameType, searchGameOs, searchGameStatus, searchGameCooperationModel, searchCompany, searchMainGame, searchSubGame } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'SubGameList'
})

const appStore = useAppStore()

const searchInfo = ref({
  game_name: '',
})

const configData = ref({
  show: false,
  content: '',
})

//游戏对话框
const gameInfoDialog = ref({
  show: false,
  add: true
})
//游戏信息
const gameInfo = ref({
  id: 0,
  game_name: '',
  package_name: '',
  app_id: '',
  app_name: '',
  game_type: '',
  os: '',
  cp_url: '',
  main_id: 0,
  game_coin_name: '',
  game_rate: 0,
  cp_game_id: 0,
  company_id: 0,
  status: '',
  cooperation_model: '',
})

const platforms = ref([])
const gameTypes = ref([])
const gameOss = ref([])
const gameStatuss = ref([])
const gameCooperationModels = ref([])
const companys = ref([])
const mainGames = ref([])
const subGames = ref([])

const openAddGameDialog = () => {
  gameInfoDialog.value.show = true
  gameInfoDialog.value.add = true
}

const closeGameInfoDialog = () => {
  gameInfoDialog.value.show = false
  resetGameInfo()
}

const resetGameInfo = () => {
  gameInfo.value.id = 0
}

const onSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {
    game_name: '',
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

const openConfigDialog = async (row) => {
  configData.value.show = true
  const result = await getSubGameConfig({
    id: row.id
  })
  if (result.code === 0) {
    configData.value.content = result.data.content
  }
}

const closeConfigDialog = () => {
  configData.value.show = false
}

const copyConfig = () => {
  navigator.clipboard.writeText(configData.value.content);
  ElMessage({ type: 'success', message: '复制成功' })
}

// 查询游戏类型
const getGameTypes = async () => {
  const table = await searchGameType()
  if (table.code === 0) {
    gameTypes.value = table.data
  }
}

const searchSubGameMethod = async (query) => {
  const table = await searchSubGame({keyword: query})
  if (table.code === 0) {
    subGames.value = table.data
  }
}

const searchMainGameMethod = async (query) => {
  const table = await searchMainGame({keyword: query})
  if (table.code === 0) {
    mainGames.value = table.data
  }
}

const searchCompanyMethod = async (query) => {
  const table = await searchCompany({keyword: query})
  if (table.code === 0) {
    companys.value = table.data
  }
}

const searchPlatformMethod = async () => {
  const table = await searchPlatform()
  if (table.code === 0) {
    platforms.value = table.data
  }
}

// 查询游戏系统
const getGameOss = async () => {
  const table = await searchGameOs()
  if (table.code === 0) {
    gameOss.value = table.data
  }
}

// 查询游戏状态
const getGameStatuss = async () => {
  const table = await searchGameStatus()
  if (table.code === 0) {
    gameStatuss.value = table.data
  }
}

// 查询游戏合作方式
const getGameCooperationModels = async () => {
  const table = await searchGameCooperationModel()
  if (table.code === 0) {
    gameCooperationModels.value = table.data
  }
}

// 查询
const getTableData = async () => {
  const table = await getSubGameList({
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
  searchPlatformMethod()
  getGameTypes()
  getGameOss()
  getGameStatuss()
  getGameCooperationModels()
}

initPage()

const rules = ref({
  platform_id: [
    { required: true, message: '请选择平台', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择平台', trigger: 'blur' }
  ],
  game_name: [
    { required: true, message: '请输入游戏名称', trigger: 'blur' }
  ],
  game_type: [
    { required: true, message: '请选择游戏类型', trigger: 'blur' }
  ],
  os: [
    { required: true, message: '请选择游戏操作系统', trigger: 'blur' }
  ],
  main_game_id: [
    { required: true, message: '请选择所属主游戏', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择所属主游戏', trigger: 'blur' }
  ],
  game_coin_name: [
    { required: true, message: '请选择游戏币名称', trigger: 'blur' }
  ],
  game_rate: [
    { required: true, message: '请输入游戏兑换比例', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请输入游戏兑换比例', trigger: 'blur' }
  ],
  company_id: [
    { required: true, message: '请选择所属主体', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择所属主体', trigger: 'blur' }
  ],
  status: [
    { required: true, message: '请选择游戏状态', trigger: 'blur' }
  ],
  cooperation_model: [
    { required: true, message: '请选择游戏合作方式', trigger: 'blur' }
  ]
})
const configForm = ref(null)
const submitGameInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (gameInfoDialog.value.add) {
        const res = await subGameAdd(gameInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeGameInfoDialog()
        }
      } else {
        const res = await subGameModify(gameInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '编辑成功' })
          await getTableData()
          closeGameInfoDialog()
        }
      }
    }
  })
}

const openEditDialog = (row) => {
  gameInfo.value = row
  console.log(row, gameInfo.value)
  gameInfoDialog.value.show = true
  gameInfoDialog.value.add = false
}

</script>

<style lang="scss">
.header-img-box {
  @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
}
</style>
