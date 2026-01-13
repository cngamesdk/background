<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="searchForm" :inline="true" :model="searchInfo">
        <el-form-item>
          <el-select v-model="searchInfo.pay_type" placeholder="请选择支付网关" style="width: 240px">
            <el-option
                v-for="item in payTypes"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-select
              placeholder="请选择支付渠道"
              filterable
              remote
              reserve-keyword
              :remote-method="handleGetPayChannels"
              :loading="loading"
              style="width: 240px"
              v-model="searchInfo.pay_channel_id">
            <el-option
                v-for="item in payChannels"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
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
        <el-button type="primary" icon="plus" @click="openAddConfigDialog">新增切换规则</el-button>
      </div>
      <el-table :data="tableData" stripe row-key="id">
        <el-table-column align="left" label="ID" prop="id" />
        <el-table-column align="left" label="平台" prop="platform_id"/>
        <el-table-column align="left" label="支付网关" prop="pay_type"/>
        <el-table-column align="left" label="规则">
          <template #default="scope">
            {{ JSON.stringify(scope.row.rules) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="支付渠道">
          <template #default="scope">
            {{ JSON.stringify(scope.row.pay_channels) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="排序" prop="sort"/>
        <el-table-column align="left" label="状态" prop="status"/>
        <el-table-column align="left" label="创建时间">
          <template #default="scope">
            {{ formatDate(scope.row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="更新时间">
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
        destroy-on-close
        center
        v-model="configDialog.show"
        title="配置"
        width="700px"
        @close="closeConfigDialog"
        :close-on-click-modal="false"
        :close-on-press-escape="false"
    >
      <el-form ref="configForm" :model="configInfo" :rules="rules" label-width="auto">
        <el-form-item label="平台" prop="platform_id">
          <el-select
              filterable
              v-model="configInfo.platform_id" placeholder="请选择所属平台" style="width: 240px">
            <el-option
                v-for="item in platforms"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="支付网关" prop="pay_type">
          <el-select v-model="configInfo.pay_type"
                     placeholder="请选择支付网关"
                     filterable
                     style="width: 240px">
            <el-option
                v-for="item in payTypes"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="切换规则" prop="rules">
          <Filter label="规则" v-model:init="configInfo.rules" v-model:fields="displayFields"></Filter>
        </el-form-item>

        <el-form-item v-if="configInfo.pay_type !== ''" label="支付渠道" prop="pay_channels">

          <div class="pay-channel-container">
            <div class="pay-channel-container-item" v-for="(itemPayChannel, indexPayChannel) in configInfo.pay_channels">
              <el-select class="pay-channel-container-item-child" v-model="configInfo.pay_channels[indexPayChannel].pay_channel_id"
                         placeholder="请选择支付渠道"
                         filterable
                         remote
                         reserve-keyword
                         :remote-method="handleGetPayChannels"
                         :loading="loading"
                         style="width: 140px">
                <el-option
                    v-for="item in payChannels"
                    :key="item.key"
                    :label="item.value"
                    :value="item.key"
                />
              </el-select>
              <el-input-number class="pay-channel-container-item-child" v-model="configInfo.pay_channels[indexPayChannel].weight" :min="1">
                <template #prefix>
                  <span>权重</span>
                </template>
              </el-input-number>
              <el-button class="pay-channel-container-item-child" type="primary" size="small" :icon="Plus" @click="handleInsertPayChannel(indexPayChannel)"/>
              <el-button class="pay-channel-container-item-child" type="primary" size="small" :icon="Minus" @click="handleDelPayChannel(indexPayChannel)"/>
            </div>

            <div>
              <el-button type="primary" size="small" :icon="Plus" @click="handleAppendPayChannel"/>
            </div>
          </div>

        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="configInfo.status" placeholder="请选择状态" style="width: 240px">
            <el-option
                v-for="item in statuss"
                :key="item.key"
                :label="item.value"
                :value="item.key"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="排序">
          <el-input-number v-model="configInfo.sort"></el-input-number>
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

import Filter from '../../../../components/dataReport/filter.vue'
import { payChannelSwitchList, payChannelSwitchAdd, payChannelSwitchModify } from '@/api/operationManagement'
import { searchPlatform, searchPayChannel, searchPayType } from '@/api/systemManagement'
import { allFilter } from '@/utils/common'
import { Plus, Minus } from '@element-plus/icons-vue'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { useAppStore } from "@/pinia";
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'PayChannelSwitchList',
  components: {
    Filter,
  },
})

const appStore = useAppStore()

const defaultSearchInfo = {
  pay_type: '',
  pay_channel_id: 0,
}

const searchInfo = ref(defaultSearchInfo)

//游戏对话框
const configDialog = ref({
  show: false,
  add: true
})

const defaultConfigInfo = {
  id: 0,
  platform_id: 0,
  pay_type: '',
  sort: 0,
  rules: [],
  pay_channels: [],
  status: '',
}

//游戏信息
const configInfo = ref({...defaultConfigInfo})

const defaultPayChannelItem = {pay_channel_id: 0, weight: 1}

const platforms = ref([])
const payTypes = ref([])
const payChannels = ref([])
const statuss = ref([
  {key: 'normal', value: '正常'},
  {key: 'remove', value: '下架'},
  {key: 'delete', value: '删除'},
])
// 需要显示的筛选
const displayFields = ref([])

const handleAppendPayChannel = () => {
  configInfo.value.pay_channels.push({...defaultPayChannelItem})
}

const handleInsertPayChannel = (index) => {console.log('pay_channels', index, configInfo.value.pay_channels)
  configInfo.value.pay_channels.splice(index + 1, 0, {...defaultPayChannelItem})
}

const handleDelPayChannel = (index) => {
  if (configInfo.value.pay_channels.length <= 1) {
    ElMessage({ type: 'error', message: '必须配置一个支付渠道' })
    return
  }
  configInfo.value.pay_channels.splice(index, 1)
}

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
  configInfo.value = {...defaultConfigInfo}
}

const onSearchSubmit = () => {
  page.value = 1
  getTableData()
}

const onReset = () => {
  searchInfo.value = {...defaultSearchInfo}
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

const getPayTypes = async () => {
  const table = await searchPayType()
  if (table.code === 0) {
    payTypes.value = table.data
  }
}

const handleGetPayChannels = async (query) => {
  const table = await searchPayChannel({
    status: 'normal',
    keyword: query,
    pay_type: configInfo.value.pay_type})
  if (table.code === 0) {
    payChannels.value = table.data
  }
}

// 查询
const getTableData = async () => {
  const table = await payChannelSwitchList({
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

const getDisplayFields = async () => {
  const result = await allFilter()
  result.forEach(function (item) {
    displayFields.value.push(item.value)
  })
}

const initPage = async () => {
  getTableData()
  getPlatforms()
  getPayTypes()
  getDisplayFields()
}

initPage()

const validateRules = (rule, value, callback) => {
  if (value.length <= 0) {
    callback(new Error('请配置切换规则'))
  }else {
    callback()
  }
}

const validatePayChannels = (rule, value, callback) => {
  if (value.length <= 0) {
    callback(new Error('请配置切换充值渠道'))
  }else {
    callback()
  }
}

const configForm = ref(null)
const rules = ref({
  platform_id: [
    { required: true, message: '请选择平台', trigger: 'blur' },
    { pattern: /^[1-9]\d*$/, message: '请选择平台', trigger: 'blur' }
  ],
  pay_type: [
    { required: true, message: '请选择支付网关', trigger: 'blur' },
  ],
  status: [
    { required: true, message: '请选择状态', trigger: 'blur' },
  ],
  rules: [
    { required: true, message: '请配置支付规则', trigger: 'change' },
    { validator: validateRules, trigger: 'blur' }
  ],
  pay_channels: [
    { required: true, message: '请配置支付渠道', trigger: 'change' },
    { validator: validatePayChannels, trigger: 'blur' }
  ],
})

const submitConfigInfo = async () => {
  configForm.value.validate(async (valid) => {
    if (valid) {
      if (configDialog.value.add) {
        const res = await payChannelSwitchAdd(configInfo.value)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: '创建成功' })
          await getTableData()
          closeConfigDialog()
        }
      } else {
        const res = await payChannelSwitchModify(configInfo.value)
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
.pay-channel-container {
  display: block;
}
.pay-channel-container-item {
  display: flex;
  justify-content: center;    /* 水平居中 */
  align-items: center;        /* 垂直居中 */
  margin: .5rem 0;
}

.pay-channel-container-item-child {
  margin: 0 .5rem;
}

</style>
