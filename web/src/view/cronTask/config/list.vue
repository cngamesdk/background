<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item label="配置ID/配置名称">
          <el-input v-model="searchInfo.name" placeholder="配置ID/配置名称" />
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
            label="名称"
            min-width="150"
            prop="name"
        />
        <el-table-column
            align="left"
            label="计划"
            min-width="100"
            prop="spec"
        />
        <el-table-column
            align="left"
            label="状态"
            min-width="150"
            prop="status"
        />
        <el-table-column
            align="left"
            label="执行模式"
            min-width="150"
            prop="execution_mode"
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
            width="180"
        >
          <template #default="scope">
            {{ formatDate(scope.row.updated_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" fixed="right" width="180">
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
              @click="openLogDialog(scope.row)"
          >日志</el-button>
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

    <!-- 任务日志 -->
    <el-dialog
        v-model="logDialog.show"
        title="任务日志"
        width="90%"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <div class="gva-table-box">
        <el-table :data="logData" stripe row-key="id">
          <el-table-column align="left" label="ID" min-width="50" prop="id" />
          <el-table-column
              align="left"
              label="开始时间"
              min-width="100">
            <template #default="scope">
              {{ formatDate(scope.row.start_time) }}
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              label="结束时间"
              min-width="100">
            <template #default="scope">
              {{ formatDate(scope.row.end_time) }}
            </template>
          </el-table-column>
          <el-table-column
              align="left"
              label="耗时"
              min-width="100"
              prop="latency"
          />
          <el-table-column
              align="left"
              label="结果"
              min-width="150"
              prop="status"
          />
          <el-table-column
              align="left"
              label="影响行数"
              min-width="150"
              prop="rows_affected"
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
                  @click="showLogResultDialog(scope.row)"
              >查看脚本</el-button>
            </template>
          </el-table-column>
        </el-table>
        <div class="gva-pagination">
          <el-pagination
              :current-page="pageLog"
              :page-size="pageSizeLog"
              :page-sizes="[10, 30, 50, 100]"
              :total="totalLog"
              layout="total, sizes, prev, pager, next, jumper"
              @current-change="handleCurrentChangeLog"
              @size-change="handleSizeChangeLog"
          />
        </div>
      </div>
    </el-dialog>

    <!-- 日志结果展示 -->
    <el-dialog
        v-model="logResultDialog.show"
        title="日志执行结果"
        width="700px"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-form ref="configForm" :model="logResultDialog.data" label-width="auto">
        <el-form-item>
          <el-input :rows="10" type="textarea" v-model="logResultDialog.data.result"/>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!-- 任务操作 -->
    <el-dialog
        v-model="configDialog.show"
        title="配置"
        width="90%"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-form ref="configForm" :model="configInfo" :rules="rules" label-width="auto">
        <el-form-item label="名称" prop="name">
          <el-input v-model="configInfo.name" placeholder="请输入任务名称"/>
        </el-form-item>
        <el-form-item label="计划" prop="spec">
          <el-input v-model="configInfo.spec" placeholder="请输入执行计划，如: * */5 * * * *"/>
        </el-form-item>
      </el-form>
      <el-form-item label="备注" prop="remark">
        <el-input autosize
                  type="textarea"
                  v-model="configInfo.remark"
                  placeholder="请输入备注"/>
      </el-form-item>
      <el-form-item label="状态" prop="status">
        <el-select v-model="configInfo.status" placeholder="请选择状态">
          <el-option
              v-for="item in statuss"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="计划内容" prop="content">
        <el-input :autosize="{ minRows: 2, maxRows: 10 }"
                  type="textarea"
                  v-model="configInfo.content"
                  placeholder="请输入计划内容"/>
      </el-form-item>
      <el-form-item label="排序" prop="sort">
        <el-input-number v-model="configInfo.sort"/>
      </el-form-item>
      <el-form-item label="类型" prop="task_type">
        <el-select v-model="configInfo.task_type" placeholder="请选择任务类型">
          <el-option
              v-for="item in taskTypes"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="模式" prop="execution_mode">
        <el-select v-model="configInfo.execution_mode" placeholder="请选择执行模式">
          <el-option
              v-for="item in executionModes"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
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

import { configList, configAdd, configModify, logList } from '@/api/cronTask'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'CronTaskConfigList'
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

const logDialog = ref({
  show : false,
  logData: {},
})

const logResultDialog = ref({
  show : false,
  data: {},
})

//游戏信息
const configInfo = ref({
  id: 0,
  name: '',
  spec: '',
  remark: '',
  status: '',
  content: '',
  sort: 0,
  parent_id: 0,
  task_type: '',
  execution_mode: '',
})

const statuss = ref([
  {key: 'normal', value: '正常'},
  {key: 'delete', value: '下架'},
])
const taskTypes = ref([
  {key: 'sql-cleaning', value: 'SQL清洗'},
  {key: 'application-packaging', value: '应用打包'},
])
const executionModes = ref([
  {key: 'sync', value: '同步'},
  {key: 'async', value: '异步'},
])

const openAddConfigDialog = () => {
  configDialog.value.show = true
  configDialog.value.add = true
}

const openConfigEditDialog = (row) => {
  configInfo.value = Object.assign({}, row)
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
    name: '',
  }
  getTableData()
}

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

const pageLog = ref(1)
const totalLog = ref(0)
const pageSizeLog = ref(10)
const logData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 分页
const handleSizeChangeLog = (val) => {
  pageSizeLog.value = val
  getLogData()
}

const handleCurrentChangeLog = (val) => {
  pageLog.value = val
  getLogData()
}

const showLogResultDialog = (row) => {
  logResultDialog.value.show = true
  logResultDialog.value.data = row
}

const openLogDialog = (row) => {
  logDialog.value.show = true
  logDialog.value.logData = row
  getLogData()
}

const getLogData = async () => {
  const table = await logList({
    page: pageLog.value,
    pageSize: pageSizeLog.value,
    config_id: logDialog.value.logData.id,
  })
  if (table.code === 0) {
    logData.value = table.data.list
    totalLog.value = table.data.total
    pageLog.value = table.data.page
    pageSizeLog.value = table.data.pageSize
  }
}

// 查询
const getTableData = async () => {
  const table = await configList({
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
}

initPage()

const configForm = ref(null)
const rules = ref({
  name: [
    { required: true, message: '请输入任务名称', trigger: 'blur' },
  ],
  spec: [
    { required: true, message: '请输入执行计划', trigger: 'blur' },
  ],
  status: [
    { required: true, message: '请选择状态', trigger: 'change' },
  ],
  content: [
    { required: true, message: '请输入计划内容', trigger: 'blur' },
  ],
  task_type: [
    { required: true, message: '请选择任务类型', trigger: 'change' },
  ],
  execution_mode: [
    { required: true, message: '请选择执行模式', trigger: 'blur' },
  ],
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await configAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await configModify(configInfo.value)
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
