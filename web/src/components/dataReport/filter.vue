<template>
  <el-popover v-mdoel="" placement="bottom-start" :width="500" trigger="click" @hide="hideComponent">
    <template #reference>
      <el-button>{{ label }}({{ selectedNum }})</el-button>
    </template>
    <el-form label-position="right" label-width="100px">

      <el-form-item v-if="displayFilters[platformIdFilterKey]" label="平台">
        <el-select
            :teleported="false"
            v-model="displayFilters[platformIdFilterKey].operator"
            placeholder="请选择操作符"
            class="middle">
          <el-option
              v-for="item in operators"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
        <el-select
            :teleported="false"
            placement="right-start"
            class="dimension-select"
            v-model="displayFilters[platformIdFilterKey].value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectPlatform"
            placeholder="请选择平台">
          <el-option
              v-for="item in platforms"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>

      <el-form-item v-if="displayFilters[rootGameIdFilterKey]" label="根游戏">
        <el-select
            :teleported="false"
            v-model="displayFilters[rootGameIdFilterKey].operator"
            placeholder="请选择操作符" class="middle">
          <el-option
              v-for="item in operators"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
        <el-select
            :teleported="false"
            placement="right-start"
            class="dimension-select"
            v-model="displayFilters[rootGameIdFilterKey].value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectRootGameId"
            placeholder="请选择根游戏">
          <el-option
              v-for="item in rootGameIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>

      <el-form-item v-if="displayFilters[mainGameIdFilterKey]" label="主游戏">
        <el-select
            :teleported="false"
            v-model="displayFilters[mainGameIdFilterKey].operator"
            placeholder="请选择操作符" class="middle">
          <el-option
              v-for="item in operators"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
        <el-select
            :teleported="false"
            placement="right-start"
            class="dimension-select"
            v-model="displayFilters[mainGameIdFilterKey].value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectMainGameId"
            placeholder="请选择主游戏">
          <el-option
              v-for="item in mainGameIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>

      <el-form-item v-if="displayFilters[gameIdFilterKey]" label="子游戏">
        <el-select
            :teleported="false"
            v-model="displayFilters[gameIdFilterKey].operator"
            placeholder="请选择操作符" class="middle">
          <el-option
              v-for="item in operators"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
        <el-select
            :teleported="false"
            placement="right-start"
            class="dimension-select"
            v-model="displayFilters[gameIdFilterKey].value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectSubGameId"
            placeholder="请选择子游戏">
          <el-option
              v-for="item in subGameIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>

      <el-form-item v-if="displayFilters[agentIdFilterKey]" label="渠道ID">
        <el-select
            :teleported="false"
            v-model="displayFilters[agentIdFilterKey].operator"
            placeholder="请选择操作符" class="middle">
          <el-option
              v-for="item in operators"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
        <el-select
            :teleported="false"
            placement="right-start"
            class="dimension-select"
            v-model="displayFilters[agentIdFilterKey].value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectAgentId"
            placeholder="请选择渠道ID">
          <el-option
              v-for="item in agentIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>

      <el-form-item v-if="displayFilters[siteIdFilterKey]" label="广告位ID">
        <el-select
            :teleported="false"
            v-model="displayFilters[siteIdFilterKey].operator"
            placeholder="请选择操作符" class="middle">
          <el-option
              v-for="item in operators"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
        <el-select
            :teleported="false"
            placement="right-start"
            class="dimension-select"
            v-model="displayFilters[siteIdFilterKey].value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectSiteId"
            placeholder="请选择广告位ID">
          <el-option
              v-for="item in siteIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>

      <el-form-item v-if="displayFilters[userIdFilterKey]" label="用户ID">
        <el-select
            :teleported="false"
            v-model="displayFilters[userIdFilterKey].operator"
            placeholder="请选择操作符"
            class="middle">
          <el-option
              v-for="item in operators"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
        <el-input-tag v-model="displayFilters[userIdFilterKey].value" placeholder="请输入用户ID，空格添加" class="dimension-select" clearable/>
      </el-form-item>

    </el-form>
  </el-popover>
</template>

<script setup>

import { watch } from 'vue'
import { ref, nextTick, defineModel, onMounted } from 'vue'
import { searchPlatform, searchRootGame, searchMainGame, searchSubGame, searchAgent, searchSite } from '@/api/systemManagement'
import {
  platformIdFilterKey,
  rootGameIdFilterKey,
  mainGameIdFilterKey,
    gameIdFilterKey,
    agentIdFilterKey,
    siteIdFilterKey,
    userIdFilterKey,
} from '@/utils/common'

//初始化数据模型
const initModel = defineModel('init',{ type: Array, default: () => [] })
const fieldsModel = defineModel('fields',{ type: Array, default: () => [] })

defineOptions({
  name: 'Filter',
})

onMounted(() => {
  initModelData()
  initFieldModelData()
})

// 监听模型的变化
watch(initModel, () => {
      initModelData()
},
    { deep: true })

// 监听模型的变化
watch(fieldsModel, () => {
      initFieldModelData()
},
    { deep: true })

//显示的筛选器
const displayFilters = ref({})

watch(displayFilters, (newVal) => {
  changeSelect()
},{ deep: true })

defineProps({
  label: {
    type: String,
    default() {
      return '维度筛选'
    }
  }
})

const operators = ref([
    {key: 'in', value: '包含'},
    {key: 'not-in', value: '不包含'}
])

const selectedNum = ref(0)

const initFieldModelData = () => {
  if (fieldsModel.value.length > 0) {
    fieldsModel.value.forEach(function (item) {
      if (!displayFilters.value[item]) {
        displayFilters.value[item] = {key: item, operator: 'in', value: []}
      }
    })
  }
}

const initModelData = () => {
  const initValues = initModel.value
  console.log('initValues', initValues)
  if ( initValues.length > 0 ) {
    initValues.forEach((item) => {
      if(item.value.length > 0) {
        displayFilters.value[item.key] = item
      }
    });
  }
  console.log('displayFilters', displayFilters.value)
}

const platforms = ref([])
const remoteSelectPlatform = async (keyword) => {
  const table = await searchPlatform({
    keyword: keyword
  })
  if (table.code === 0) {
    platforms.value = table.data
  }
}

const rootGameIds = ref([])
const remoteSelectRootGameId = async (keyword) => {
  const table = await searchRootGame({
    keyword: keyword
  })
  if (table.code === 0) {
    rootGameIds.value = table.data
  }
}

const mainGameIds = ref([])
const remoteSelectMainGameId = async (keyword) => {
  const table = await searchMainGame({
    keyword: keyword
  })
  if (table.code === 0) {
    mainGameIds.value = table.data
  }
}

const subGameIds = ref([])
const remoteSelectSubGameId = async (keyword) => {
  const table = await searchSubGame({
    keyword: keyword
  })
  if (table.code === 0) {
    subGameIds.value = table.data
  }
}

const agentIds = ref([])
const remoteSelectAgentId = async (keyword) => {
  const table = await searchAgent({
    keyword: keyword
  })
  if (table.code === 0) {
    agentIds.value = table.data
  }
}

const siteIds = ref([])
const remoteSelectSiteId = async (keyword) => {
  const table = await searchSite({
    keyword: keyword
  })
  if (table.code === 0) {
    siteIds.value = table.data
  }
}

const hideComponent = () => {
  initModel.value = calcSelected()
}

const calcSelected = () => {
  let filterValues = []
  const filters = displayFilters.value
  for (let key in filters) {
    if (filters[key].value.length > 0) {
      filterValues.push(filters[key])
    }
  }
  return filterValues
}

const changeSelect = (currentValue, oldValue) => {
  let num = 0
  const tmpFilters = displayFilters.value
  for(let key in tmpFilters) {
    if (tmpFilters[key].value.length > 0) {
      num++
    }
  }
  selectedNum.value = num
}

</script>

<style lang="scss">
.middle {
  width:7rem;margin:auto 1rem;
}
.dimension-select{
  width:15rem;
}
</style>