<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item>
          <el-input v-model="searchInfo.material_name" placeholder="素材名称" />
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
        <el-button type="primary" icon="plus" @click="openAddConfigDialog">新增素材</el-button>
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
            label="素材名称"
            min-width="150"
            prop="material_name"
        />
        <el-table-column
            align="left"
            label="类型"
            prop="material_type_name"
        />
        <el-table-column
            align="left"
            label="题材"
            prop="theme_name"
            min-width="150"
        />
        <el-table-column
            align="left"
            label="作者"
            prop="author"
        />
        <el-table-column
            align="left"
            label="来源"
            prop="source_name"
        />
        <el-table-column
            align="left"
            label="状态"
            prop="status_name"
        />
        <el-table-column
            align="left"
            label="可见性"
            prop="visibility_name"
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
        <el-table-column label="操作" min-width="200">
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
              @click="handleViewMaterialFile(scope.row)"
          >素材文件</el-button>
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
        v-model="materialFileDialog.show"
        title="素材文件"
        width="90%"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <div class="gva-btn-list">
        <GlobalUploadCommon label="上传文件" :data="uploadFileConfig.data" :accept="uploadFileConfig.accept" :size="uploadFileConfig.size" :allowMimes="uploadFileConfig.allowMimes" @on-success="handleSuccessUploadFile"></GlobalUploadCommon>
      </div>
      <el-table :data="tableDataMaterialFile">
        <el-table-column label="ID" prop="id" />
        <el-table-column label="名称" prop="file_name" />
        <el-table-column label="状态" prop="status_name" />
        <el-table-column label="可见性" prop="visibility_name" />
        <el-table-column label="签名" prop="signature" />
        <el-table-column label="宽高">
          <template #default="scope">
           {{ scope.row.width }}*{{ scope.row.height }}
          </template>
        </el-table-column>duration
        <el-table-column label="时长(秒)">
          <template #default="scope">
            {{ scope.row.duration / 1000 }}
          </template>
        </el-table-column>
        <el-table-column label="码率kbps">
          <template #default="scope">
            {{ scope.row.bitrate }}
          </template>
        </el-table-column>
        <el-table-column label="文件大小(bit)">
          <template #default="scope">
            {{ scope.row.size }}
          </template>
        </el-table-column>
        <el-table-column label="类型" prop="file_type_name" />
        <el-table-column label="操作" fixed="right">
          <template #default="scope">
            <el-button
                type="primary"
                link
                icon="edit"
                @click="openConfigEditDialog(scope.row)"
            >扩展</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
            :current-page="pageMaterialFile"
            :page-size="pageSizeMaterialFile"
            :page-sizes="[10, 30, 50, 100]"
            :total="totalMaterialFile"
            layout="total, sizes, prev, pager, next, jumper"
            @current-change="handleCurrentChangeCMaterialFile"
            @size-change="handleSizeChangeMaterialFile"
        />
      </div>
    </el-dialog>

    <!-- 游戏添加/编辑对话框 -->
    <el-dialog
        v-model="configDialog.show"
        title="素材配置"
        width="700px"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-form ref="configForm" :model="configInfo" :rules="rules" label-width="80px">
        <el-form-item label="平台" prop="platform_id">
          <el-select @change="handleChangePlatform" v-model="configInfo.platform_id" placeholder="请选择所属平台" style="width: 240px">
            <el-option
                v-for="item in platforms"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="素材名称" prop="material_name">
          <el-input v-model="configInfo.material_name" placeholder="请输入素材名称"/>
        </el-form-item>
        <el-form-item label="题材" prop="theme_id">
          <el-cascader v-model="configInfo.theme_id_cascader" style="width: 240px" :options="themes" @change="handleChangeTheme" />
        </el-form-item>
        <el-form-item label="来源" prop="source">
          <el-select v-model="configInfo.source" placeholder="请选择来源" style="width: 240px">
            <el-option
                key="original"
                label="原创"
                value="original"
            />
            <el-option
                key="editing"
                label="改动"
                value="editing"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="configInfo.status" placeholder="请选择状态" style="width: 240px">
            <el-option
                key="normal"
                label="正常"
                value="normal"
            />
            <el-option
                key="delete"
                label="删除"
                value="delete"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="可见性" prop="visibility">
          <el-select v-model="configInfo.visibility" placeholder="请选择可见性" style="width: 240px">
            <el-option
                key="public"
                label="公开"
                value="public"
            />
            <el-option
                key="private"
                label="私密"
                value="private"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="类型" prop="material_type">
          <el-select v-model="configInfo.material_type" placeholder="请选择类型" style="width: 240px">
            <el-option
                key="image"
                label="图片"
                value="image"
            />
            <el-option
                key="video"
                label="视频"
                value="video"
            />
            <el-option
                key="audio"
                label="音频"
                value="audio"
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

import { materialList, materialAdd, materialModify, materialFileList, materialFileAdd, materialFileModify } from '@/api/material'
import { searchPlatform, getMaterialTheme } from '@/api/systemManagement'
import GlobalUploadCommon from '@/components/upload/global.vue'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'MaterialList'
})

const appStore = useAppStore()

const searchInfo = ref({
  material_name: '',
})

//游戏对话框
const configDialog = ref({
  show: false,
  add: true
})

const materialFileDialog = ref({
  show: false,
  add: true,
})

const defaultConfigInfo = {
  id: 0,
  platform_id: 0,
  material_name: '',
  theme_id: 0,
  theme_id_cascader: [],
  source: '',
  status: '',
  author: '',
  visibility: '',
  material_type: '',
}

const uploadFileLimit = {
  'image': {size: 5 * 1024 * 1024, accept:'image/*', allowMimes:['image/jpeg', 'image/jpg', 'image/png']},
  'video': {size: 5 * 1024 * 1024, accept:'video/*', allowMimes:['video/mp4', 'video/wmv']},
  'audio': {size: 5 * 1024 * 1024, accept:'audio/*', allowMimes:['audio/mp3']},
}


const uploadFileConfig = ref({
  size: 0,
  allowMimes: [],
  accept: '',
  data: {}
})

const materialFileInfo = ref({
  platform_id: 0,
  material_id: 0,
  source: '',
  file_name: '',
  url: '',
  status: '',
  visibility: '',
  signature: '',
  width: 0,
  height: 0,
  file_type: '',
  duration: 0,
  bitrate: 0,
  size: 0,
})

//游戏信息
const configInfo = ref(Object.assign({}, defaultConfigInfo))

const platforms = ref([])
const themes = ref([])

const openAddConfigDialog = () => {
  configInfo.value = Object.assign({}, defaultConfigInfo)
  configDialog.value.show = true
  configDialog.value.add = true
}

const openConfigEditDialog = (row) => {
  configInfo.value = row
  configDialog.value.show = true
  configDialog.value.add = false
  getTheme()
}

const handleViewMaterialFile = (row) => {
  if (uploadFileLimit[row.material_type]) {
    uploadFileConfig.value = uploadFileLimit[row.material_type]
    uploadFileConfig.value.data = {platform_id: row.platform_id, biz: 'material'}
  } else {
    ElMessage({ type: 'error', message: '不能识别的素材类型' })
    return
  }
  materialFileDialog.value.show = true
  materialFileInfo.value.material_id = row.id
  materialFileInfo.value.platform_id = row.platform_id

  getMaterialFileList()
}

const handleSuccessUploadFile = (val) => {
  const data = Object.assign({}, materialFileInfo.value)
  const fileInfo = val.file
  data.file_name = fileInfo.name
  data.url = fileInfo.url
  data.signature = fileInfo.hash
  data.width = val.width
  data.height = val.height
  data.file_type = fileInfo.tag
  data.duration = val.duration
  data.bitrate = val.bitrate
  data.size = val.size
  data.fps = val.fps
  uploadMaterialFile(data)
}

const uploadMaterialFile = async (data) => {
  const result = await materialFileAdd(data)
  if (result.code === 0) {
    ElMessage({ type: 'success', message: '文件上传成功' })
    getMaterialFileList()
  } else {
    ElMessage({ type: 'error', message: '文件上传成功但是新增记录失败' })
  }
}

const handleConfigEditThemeCascader = () => {
  themes.value.forEach(function (item) {
    if(item.children && item.children.length > 0) {
      item.children.forEach(function (childItem) {
        if(childItem.value === configInfo.value.theme_id) {
          configInfo.value.theme_id_cascader = [item.value, childItem.value]
        }
      })
    }
  })
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

const pageMaterialFile = ref(1)
const totalMaterialFile = ref(0)
const pageSizeMaterialFile = ref(10)
const tableDataMaterialFile = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

const handleChangePlatform = (val) => {
  getTheme()
}

const handleAddSubTheme = (row) => {
  openAddConfigDialog()
}

const handleChangeTheme = (val) => {
  configInfo.value.theme_id = val[val.length - 1]
}

// 子题材分页
const handleSizeChangeMaterialFile = (val) => {
  pageSizeMaterialFile.value = val
  getMaterialFileList()
}

const handleCurrentChangeMaterialFile = (val) => {
  pageMaterialFile.value = val
  getMaterialFileList()
}

// 查询平台
const getPlatforms = async () => {
  const table = await searchPlatform()
  if (table.code === 0) {
    platforms.value = table.data
  }
}

// 获取题材
const getTheme = async () => {
  const table = await getMaterialTheme({platform_id: configInfo.value.platform_id})
  if (table.code === 0) {
    themes.value = table.data
    handleConfigEditThemeCascader()
  }
}

// 查询
const getTableData = async () => {
  const table = await materialList({
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

const getMaterialFileList = async () => {
  const table = await materialFileList({
    page: pageMaterialFile.value,
    pageSize: pageSizeMaterialFile.value,
    material_id: materialFileInfo.value.material_id,
    ...searchInfo.value
  })
  if (table.code === 0) {
    tableDataMaterialFile.value = table.data.list
    totalMaterialFile.value = table.data.total
    pageMaterialFile.value = table.data.page
    pageSizeMaterialFile.value = table.data.pageSize
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
  material_name: [
    { required: true, message: '请输入素材名称', trigger: 'blur' },
  ],
  theme_id: [
    { required: true, message: '请选择细分题材', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择细分题材', trigger: 'blur' }
  ],
  source: [
    { required: true, message: '请选择素材来源', trigger: 'blur' },
  ],
  status: [
    { required: true, message: '请选择素材状态', trigger: 'blur' },
  ],
  visibility: [
    { required: true, message: '请选择素材可见性', trigger: 'blur' },
  ],
  material_type: [
    { required: true, message: '请选择素材类型', trigger: 'blur' },
  ],
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await materialAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await materialModify(configInfo.value)
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
