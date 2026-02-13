<template>
  <div>
    <el-form ref="configForm" :model="configInfo" :rules="rules" label-width="auto"  style="max-width: 600px;margin:2rem;">
      <el-form-item label="平台" prop="platform_id">
        <el-select v-model="configInfo.platform_id" placeholder="请选择所属平台" style="width: 240px" @change="handlePlatform">
          <el-option
              v-for="item in platforms"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="媒体" prop="code">
        <el-select v-model="configInfo.code" placeholder="请选择媒体" style="width: 240px" @change="handleMedia">
          <el-option
              v-for="item in medias"
              :key="item.key_2"
              :label="item.value"
              :value="item.key_2"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="开发者" prop="id">
        <el-select v-model="configInfo.id" placeholder="请选择开发者" style="width: 240px">
          <el-option
              v-for="item in developers"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="search" @click="submitConfigInfo">
          去授权
        </el-button>
        <el-button icon="refresh" @click="onReset"> 重置 </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>

import { advertisingAuthRedirect } from '@/api/advertising'
import { searchPlatform, searchMedia, getAdvertisingDeveloper } from '@/api/systemManagement'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'AdvertisingAuthRedirect'
})

const appStore = useAppStore()

const searchInfo = ref({
  name: '',
})

const defaultConfigInfo = {
  id: 0,
  platform_id: 0,
  code: '',
}

//游戏信息
const configInfo = ref(Object.assign({}, defaultConfigInfo))

const platforms = ref([])
const medias = ref([])
const developers = ref([])

const onReset = () => {
  configInfo.value = Object.assign({}, defaultConfigInfo)
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

const handlePlatform = (val) => {
  configInfo.value.platform_id = val
  getMedias()
}

const handleMedia = (val) => {
  configInfo.value.code = val
  getDevelopers()
}

const getMedias = async () => {
  const table = await searchMedia({platform_id: configInfo.value.platform_id})
  if (table.code === 0) {
    medias.value = table.data
  }
}

const getDevelopers = async () => {
  const table = await getAdvertisingDeveloper({platform_id: configInfo.value.platform_id, keyword: configInfo.value.code})
  if (table.code === 0) {
    developers.value = table.data
  }
}

const initPage = async () => {
  getPlatforms()
}

initPage()

const configForm = ref(null)
const rules = ref({
  platform_id: [
    { required: true, message: '请选择平台', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择平台', trigger: 'blur' }
  ],
  code: [
    { required: true, message: '请选择媒体', trigger: 'blur' },
  ],
  id: [
    { required: true, message: '请选择开发者', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择开发者', trigger: 'blur' }
  ],
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      const res = await advertisingAuthRedirect(configInfo.value)
      if (res.code === 0) {
        window.open(res.data.url)
      }
    }
  })
}

</script>
