<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item>

         <Filter label="维度筛选" v-model:init="searchInfo.dimension_filter" v-model:fields="displayFields"></Filter>

          <Dimensions v-model="searchInfo.dimensions" :dimensions="allDimensions"></Dimensions>

          <Indicators v-model="searchInfo.indicators" :indicators="allIndicators"></Indicators>

        </el-form-item>
        <el-form-item>

          <StatisticalCaliber v-model="searchInfo.statistical_caliber"></StatisticalCaliber>

        </el-form-item>
        <el-form-item>

          <AggregationTime v-model="searchInfo.aggregation_time"></AggregationTime>

        </el-form-item>
        <el-form-item>
          <DateRange v-model="dateRange"></DateRange>
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
      <el-table :data="tableData" stripe row-key="id">
        <el-table-column
            v-if="firstRowData.hasOwnProperty('stat_date')"
            align="left" label="日期" prop="stat_date"/>
        <el-table-column
            v-if="firstRowData.hasOwnProperty('platform_id')"
            align="left" label="平台">
          <template #default="scope">
            {{ scope.row.platform_id }}-{{ scope.row.platform_name }}
          </template>
        </el-table-column>
        <el-table-column
            v-if="firstRowData.hasOwnProperty('root_game_id')"
            align="left" label="根游戏">
          <template #default="scope">
            {{ scope.row.root_game_id }}-{{ scope.row.root_game_name }}
          </template>
        </el-table-column>
        <el-table-column
            v-if="firstRowData.hasOwnProperty('main_game_id')"
            align="left" label="主游戏">
          <template #default="scope">
            {{ scope.row.main_game_id }}-{{ scope.row.main_game_name }}
          </template>
        </el-table-column>
        <el-table-column
            v-if="firstRowData.hasOwnProperty('game_id')"
            align="left" label="子游戏">
          <template #default="scope">
            {{ scope.row.game_id }}-{{ scope.row.game_name }}
          </template>
        </el-table-column>
        <el-table-column
            v-if="firstRowData.hasOwnProperty('agent_id')"
            align="left" label="渠道ID">
          <template #default="scope">
            {{ scope.row.agent_id }}-{{ scope.row.agent_name }}
          </template>
        </el-table-column>
        <el-table-column
            v-if="firstRowData.hasOwnProperty('site_id')"
            align="left" label="广告位ID">
          <template #default="scope">
            {{ scope.row.site_id }}-{{ scope.row.site_name }}
          </template>
        </el-table-column>
        <el-table-column
            v-if="firstRowData.hasOwnProperty('ad3_id')"
            align="left" label="广告三级ID" prop="ad3_id"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('activation')"
            align="left" label="激活数" prop="activation"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('activation_device')"
            align="left" label="激活设备数" prop="activation_device"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('launch')"
            align="left" label="启动数" prop="launch"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('launch_device')"
            align="left" label="启动设备数" prop="launch_device"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('reg')"
            align="left" label="注册数" prop="reg"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('reg_device')"
            align="left" label="注册设备数" prop="reg_device"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('login')"
            align="left" label="登录数" prop="login"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('login_user')"
            align="left" label="登录用户数" prop="login_user"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('login_device')"
            align="left" label="登录设备数" prop="login_device"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('role')"
            align="left" label="创角数" prop="role"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('role_user')"
            align="left" label="创角用户数" prop="role_user"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('role_user')"
            align="left" label="创角设备数" prop="role_device"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('pay')"
            align="left" label="付费数" prop="pay"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('pay_user')"
            align="left" label="付费用户数" prop="pay_user"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('pay_device')"
            align="left" label="付费设备数" prop="pay_device"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('pay_money')"
            align="left" label="付费金额" prop="pay_money"
        />
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

import { dayOverviewList } from '@/api/dataReport'
import { dimensionFilter } from '@/utils/common'
import Filter from '../../../components/dataReport/filter.vue'
import Dimensions from '../../../components/dataReport/dimensions.vue'
import Indicators from '../../../components/dataReport/indicators.vue'
import StatisticalCaliber from '../../../components/dataReport/statisticalCaliber.vue'
import AggregationTime from '../../../components/dataReport/aggregationTime.vue'
import DateRange from '../../../components/dataReport/dateRange.vue'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatTimeToStr } from '@/utils/date'

defineOptions({
  name: 'DayOverviewList',
  components: {
    Filter,
    Dimensions,
    Indicators,
    StatisticalCaliber,
    AggregationTime,
    DateRange,
  },
})

const appStore = useAppStore()

const dateRange = ref([new Date(), new Date()])

//需要显示的维度筛选
const displayFields = ref([])

const searchInfo = ref({
  statistical_caliber : 'root-game-back-30',
  dimension_filter: [],
  dimensions: [],
  indicators: [],
  aggregation_time: 'day',
  start_time: '',
  end_time: ''
})

const allDimensions = [
  {value: '运营侧', childs: [
      {key: 'platform_id', value: '平台', childs: []},
      {key: 'root_game_id', value: '根游戏', childs: []},
      {key: 'main_game_id', value: '主游戏', childs: []},
      {key: 'game_id', value: '子游戏', childs: []},
    ]},
  {value: '市场侧', childs: [
      {key: 'agent_id', value: '渠道ID', childs: []},
      {key: 'site_id', value: '广告位', childs: []},
    ]},
    ]

const allIndicators = [
  {value: '整体情况', childs: [
      {key: 'reg_device', value: '注册设备数'},
      {key: 'reg', value: '注册数'},
      {key: 'login_user', value: '活跃用户数'},
      {key: 'role', value: '创角数'},
      {key: 'pay', value: '付费数'},
    ]}
]

const getDisplayFields = async () => {
  const result = await dimensionFilter()
  result.forEach(function (item) {
    displayFields.value.push(item.value)
  })
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
const firstRowData = ref([])

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
  const dataRange = dateRange.value
  searchInfo.value.start_time = formatTimeToStr(dataRange[0],'yyyy-MM-dd')
  searchInfo.value.end_time = formatTimeToStr(dataRange[1],'yyyy-MM-dd')
  const table = await dayOverviewList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    
    if (table.data.list.length > 0) {
      firstRowData.value = table.data.list[0]
    }
  }
}

const initData = () => {
  getDisplayFields()
}

initData()

</script>
