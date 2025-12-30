<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openAddConfigDialog"
        >新增配置</el-button>
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
            min-width="150">
          <template #default="scope">
            {{ scope.row.platform_id }}({{ scope.row.platform_name }})
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="Java执行路径"
            min-width="150"
            prop="java_execution_path"
        />
        <el-table-column
            align="left"
            label="安卓打包工具"
            min-width="150"
            prop="game_packaging_tool_path"
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
        <el-form-item label="Java执行路径" prop="java_execution_path">
          <el-input
              v-model="configInfo.java_execution_path" placeholder="请输入java执行路径,如:/usr/bin/java"/>
        </el-form-item>
        <el-form-item label="安卓打包工具" prop="game_packaging_tool_path">
          <el-input
              autosize
              type="textarea"
              v-model="configInfo.game_packaging_tool_path"
              placeholder="请输入安卓打包工具或者上传" style="width:450px;margin: 0 .5rem;"/>
          <GlobalUploadCommon @on-success="handleSuccessCallback" :allowMimes="['application/java-archive']"></GlobalUploadCommon>
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

import GlobalUploadCommon from '../../../../components/upload/global.vue'

import { globalCommonConfigList, globalCommonConfigAdd, globalCommonConfigModify } from '@/api/operationManagement'
import { searchPlatform,  } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'
import { UploadFilled } from '@element-plus/icons-vue'

defineOptions({
  name: 'GlobalCommonConfigList'
})

const appStore = useAppStore()

const searchInfo = ref({
})

//游戏对话框
const configDialog = ref({
  show: false,
  add: true
})

const defaultConfigInfo = {
  id: 0,
  platform_id: 0,
  java_execution_path: '',
  game_packaging_tool_path: '',
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
  configInfo.value = defaultConfigInfo
  configDialog.value.show = false
}

const onSearchSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {
  }
  getTableData()
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const platforms = ref([])

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
  const table = await globalCommonConfigList({
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

const handleSuccessCallback = (data) => {
  configInfo.value.game_packaging_tool_path = data.file.url
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
  java_execution_path: [
    { required: true, message: '请输入java执行路径', trigger: 'blur' },
  ],
  game_packaging_tool_path: [
    { required: true, message: '请上传安卓打包工具', trigger: 'blur' },
  ]
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await globalCommonConfigAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await globalCommonConfigModify(configInfo.value)
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
