<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item>
          <Filter label="维度筛选" v-model:init="searchInfo.dimension_filter" v-mdoel:fields="displayDimensionFilter"></Filter>
          <Dimensions v-model="searchInfo.dimensions" :dimensions="allDimensions"></Dimensions>
          <Indicators v-model="searchInfo.indicators" :indicators="allIndicators"></Indicators>
          &nbsp;&nbsp;
          <StatisticalCaliber v-model="searchInfo.statistical_caliber"></StatisticalCaliber>
          &nbsp;&nbsp;
          <AggregationTime v-model="searchInfo.aggregation_time"></AggregationTime>
          &nbsp;&nbsp;
          <DateRange v-model="dateRange"></DateRange>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSearchSubmit">
            查询
          </el-button>
          <el-button icon="refresh" @click="onReset"> 重置 </el-button>
        </el-form-item>
        <el-form-item>
          <el-select style="width:8rem;" v-model="currPaymentOption" placeholder="请选择选项">
            <el-option
                v-for="item in paymentStatusOptions"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
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
            v-if="firstRowData.hasOwnProperty('reg')"
            align="left" label="注册数" prop="reg"
        />
        <el-table-column
            v-if="firstRowData.hasOwnProperty('cost')"
            align="left" label="消耗" prop="cost"
        />
        <el-table-column
            v-if="currPaymentOption === 'roi'"
            align="left" label="当前回本率" prop="roi_rate_str"
        />
        <el-table-column
            v-if="currPaymentOption === 'ltv'"
            align="left" label="当前LTV" prop="ltv"
        />
        <el-table-column
            v-if="currPaymentOption === 'cumulative-payment-amount'"
            align="left" label="当前累计付费" prop="cumulative_pay_amount"
        />
        <template v-for="(item, index) in nDayColumns" :key="index">
          <el-table-column v-if="item.show" align="left" :label="item.day + '日' ">
            <template #default="scope">
              <template v-if="currPaymentOption === 'roi'">{{ scope.row.n_day_container[index].roi_rate_str }}</template>
              <template v-if="currPaymentOption === 'ltv'">{{ scope.row.n_day_container[index].ltv }}</template>
              <template v-if="currPaymentOption === 'cumulative-payment-amount'">{{ scope.row.n_day_container[index].cumulative_payments }}</template>
            </template>
          </el-table-column>
        </template>
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

import { paymentStatusList } from '@/api/dataReport'
import { dimensionFilter } from '@/utils/common'
import Filter from '../../../components/dataReport/filter.vue'
import Dimensions from '../../../components/dataReport/dimensions.vue'
import Indicators from '../../../components/dataReport/indicators.vue'
import StatisticalCaliber from '../../../components/dataReport/statisticalCaliber.vue'
import AggregationTime from '../../../components/dataReport/aggregationTime.vue'
import DateRange from '../../../components/dataReport/dateRange.vue'

import { ref } from 'vue'
import { useAppStore } from "@/pinia";
import { formatTimeToStr } from '@/utils/date'

defineOptions({
  name: 'PaymentStatus',
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
const nDayColumns = ref([])
//需要显示的维度筛选
const displayDimensionFilter = ref([])

const searchInfo = ref({
  statistical_caliber : 'root-game-back-30',
  dimension_filter: [],
  dimensions: [],
  indicators: ['2'],
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
  {value: '留存', childs: [
      {key: '1', value: '1日'},
      {key: '2', value: '2日'},
      {key: '3', value: '3日'},
      {key: '4', value: '4日'},
      {key: '5', value: '5日'},
      {key: '6', value: '6日'},
      {key: '7', value: '7日'},
      {key: '8-14', value: '8-14日'},
      {key: '15-21', value: '15-21日'},
      {key: '22-28', value: '22-28日'},
    ]}
]

const currPaymentOption = ref('roi')
const paymentStatusOptions = ref([
  {key: 'roi', value: '回本率'},
  {key: 'ltv', value: 'LTV'},
  {key: 'cumulative-payment-amount', value: '累计付费'},
])

const getDisplayDimensionFilter = async () => {
  const result = await dimensionFilter()
  result.forEach(function (item) {
    displayDimensionFilter.value.push(item.value)
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
  const table = await paymentStatusList({
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
      nDayColumns.value = []
      firstRowData.value.n_day_container.forEach(function (item) {
        nDayColumns.value.push({day:item.n_day, show: item.show})
      })
    }
  }
}

const initData = () => {
  getDisplayDimensionFilter()
}

initData()

</script>
