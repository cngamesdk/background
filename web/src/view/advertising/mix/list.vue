<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item>

          <el-popover placement="bottom-start" trigger="click" :width="600">
            <template #reference>
              <el-button>筛选</el-button>
            </template>
            <el-form ref="searchForm" :model="searchInfo" label-width="7rem" label-position="left">
              <el-form-item label="平台">
                <el-select
                    filterable
                    v-model="searchInfo.platform_id"
                    placeholder="请选择平台"
                    style="width:150px;"
                >
                  <el-option
                      v-for="item in platforms"
                      :key="item.key"
                      :label="item.value"
                      :value="item.key"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="组合ID">
                <el-input-number v-model="searchInfo.id" />
              </el-form-item>
              <el-form-item label="组合名称">
                <el-input v-model="searchInfo.name" placeholder="请输入组合名称" style="width: 250px;"/>
              </el-form-item>
              <el-form-item label="媒体">
                <el-select
                    filterable
                    v-model="searchInfo.code"
                    placeholder="请选择媒体"
                    style="width:150px;"
                >
                  <el-option
                      v-for="item in medias"
                      :key="item.key_2"
                      :label="item.value"
                      :value="item.key_2"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="创建时间">
                <el-date-picker
                    v-model="searchInfo.created_at_datetime_range"
                    type="datetimerange"
                    :shortcuts="shortcuts"
                    range-separator="To"
                    start-placeholder="Start date"
                    end-placeholder="End date"
                />
              </el-form-item>
              <el-form-item label="更新时间">
                <el-date-picker
                    v-model="searchInfo.updated_at_datetime_range"
                    type="datetimerange"
                    :shortcuts="shortcuts"
                    range-separator="To"
                    start-placeholder="Start date"
                    end-placeholder="End date"
                />
              </el-form-item>
            </el-form>
          </el-popover>

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearchSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset"> 重置 </el-button>
          <el-button type="primary" icon="plus" @click="handleAddMix">新增组合</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" stripe row-key="id">
        <el-table-column align="left" label="组合ID" min-width="60" prop="id" />
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
            label="用户"
            min-width="100"
            prop="user_name"
        />
        <el-table-column
            align="left"
            label="媒体"
            min-width="180">
          <template #default="scope">
            {{ scope.row.code }}-{{ scope.row.media_name }}
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="通用配置"
            min-width="100">
          <template #default="scope">
            <el-button type="primary" link>查看</el-button>
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="帐户配置"
            min-width="100">
          <template #default="scope">
            <el-button type="primary" link>查看</el-button>
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="一级配置"
            min-width="100">
          <template #default="scope">
            <el-button type="primary" link>查看</el-button>
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="二级配置"
            min-width="100">
          <template #default="scope">
            <el-button type="primary" link>查看</el-button>
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="三级配置"
            min-width="100">
          <template #default="scope">
            <el-button type="primary" link>查看</el-button>
          </template>
        </el-table-column>
        <el-table-column
            align="left"
            label="组合配置"
            min-width="100">
          <template #default="scope">
            <el-button type="primary" link>查看</el-button>
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

  </div>
</template>

<script setup>

import { advertisingMixList } from '@/api/advertising'
import { searchPlatform, searchMedia } from '@/api/systemManagement'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'
import { useRouter } from 'vue-router'

const router = useRouter()

defineOptions({
  name: 'AdvertisingMixList'
})

const appStore = useAppStore()

const searchInfo = ref({
  platform_id: 0,
  id: 0,
  name: '',
  code: '',
  created_at_datetime_range: '',
  updated_at_datetime_range: '',
})

const shortcuts = [
  {
    text: 'Last week',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setDate(start.getDate() - 7)
      return [start, end]
    },
  },
  {
    text: 'Last month',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setMonth(start.getMonth() - 1)
      return [start, end]
    },
  },
  {
    text: 'Last 3 months',
    value: () => {
      const end = new Date()
      const start = new Date()
      start.setMonth(start.getMonth() - 3)
      return [start, end]
    },
  },
]

const platforms = ref([])
const medias = ref([])

const handleAddMix = () => {
  router.push({ name: 'advertisingMixConfig', params: { id: 123 } })
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
  const table = await searchMedia({platform_id: searchInfo.value.platform_id})
  if (table.code === 0) {
    medias.value = table.data
  }
}

// 查询
const getTableData = async () => {
  const table = await advertisingMixList({
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
}

initPage()

</script>
