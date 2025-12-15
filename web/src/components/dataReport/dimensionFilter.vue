<template>
  <el-popover v-mdoel="" placement="bottom-start" :width="500" trigger="click" @hide="hideComponent">
    <template #reference>
      <el-button>维度筛选({{ selectedDimensionNum }})</el-button>
    </template>
    <el-form label-position="right" label-width="100px">
      <el-form-item v-if="dimensions.indexOf(dimensionFilters.platform_id.key) > -1" label="平台">
        <el-select
            :teleported="false"
            v-model="dimensionFilters.platform_id.operator"
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
            v-model="dimensionFilters.platform_id.value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectPlatform"
            @change="changeSelect"
            placeholder="请选择平台">
          <el-option
              v-for="item in platforms"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item v-if="dimensions.indexOf(dimensionFilters.root_game_id.key) > -1" label="根游戏">
        <el-select
            :teleported="false"
            v-model="dimensionFilters.root_game_id.operator"
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
            v-model="dimensionFilters.root_game_id.value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectRootGameId"
            @change="changeSelect"
            placeholder="请选择根游戏">
          <el-option
              v-for="item in rootGameIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item v-if="dimensions.indexOf(dimensionFilters.main_game_id.key) > -1" label="主游戏">
        <el-select
            :teleported="false"
            v-model="dimensionFilters.main_game_id.operator"
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
            v-model="dimensionFilters.main_game_id.value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectMainGameId"
            @change="changeSelect"
            placeholder="请选择主游戏">
          <el-option
              v-for="item in mainGameIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item v-if="dimensions.indexOf(dimensionFilters.game_id.key) > -1" label="子游戏">
        <el-select
            :teleported="false"
            v-model="dimensionFilters.game_id.operator"
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
            v-model="dimensionFilters.game_id.value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectSubGameId"
            @change="changeSelect"
            placeholder="请选择子游戏">
          <el-option
              v-for="item in subGameIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item v-if="dimensions.indexOf(dimensionFilters.agent_id.key) > -1" label="渠道ID">
        <el-select
            :teleported="false"
            v-model="dimensionFilters.agent_id.operator"
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
            v-model="dimensionFilters.agent_id.value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectAgentId"
            @change="changeSelect"
            placeholder="请选择渠道ID">
          <el-option
              v-for="item in agentIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
      <el-form-item v-if="dimensions.indexOf(dimensionFilters.site_id.key) > -1" label="广告位ID">
        <el-select
            :teleported="false"
            v-model="dimensionFilters.site_id.operator"
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
            v-model="dimensionFilters.site_id.value"
            multiple
            filterable
            remote
            :remote-method="remoteSelectSiteId"
            @change="changeSelect"
            placeholder="请选择广告位ID">
          <el-option
              v-for="item in siteIds"
              :key="item.key"
              :label="item.value"
              :value="item.key"
          />
        </el-select>
      </el-form-item>
    </el-form>
  </el-popover>
</template>

<script setup>

import { ref, nextTick, defineModel, onMounted } from 'vue'
import { searchPlatform, searchRootGame, searchMainGame, searchSubGame, searchAgent, searchSite } from '@/api/systemManagement'

const modelValue = defineModel({ type: Array, default: [] })

defineOptions({
  name: 'DimensionFilter',
})

onMounted(() => {
  const initValues = modelValue.value
  const allDimensionFilters = dimensionFilters.value
  initValues.forEach((item) => {
    if(allDimensionFilters[item.key] && item.value.length > 0) {
      allDimensionFilters[item.key] = item
    }
  });
  changeSelect()
})

const dimensionFilters = ref({
  platform_id: {key: 'platform_id', operator: 'in', value: []},
  root_game_id: {key: 'root_game_id', operator: 'in', value: []},
  main_game_id: {key: 'main_game_id', operator: 'in', value: []},
  game_id: {key: 'game_id', operator: 'in', value: []},
  agent_id: {key: 'agent_id', operator: 'in', value: []},
  site_id: {key: 'site_id', operator: 'in', value: []},
})

defineProps({
  dimensions: {
    type: Array,
    default() {
      return []
    }
  }
})

const operators = ref([
    {key: 'in', value: '包含'},
    {key: 'not-in', value: '不包含'}
    ])

const selectedDimensionNum = ref(0)

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
  modelValue.value = calcSelectedDimensions()
}

const calcSelectedDimensions = () => {
  let filterValues = []
  const filters = dimensionFilters.value
  for (let key in filters) {
    if (filters[key].value.length > 0) {
      filterValues.push(filters[key])
    }
  }
  return filterValues
}

const changeSelect = (currentValue, oldValue) => {
  let num = 0
  const allDimensionFilters = dimensionFilters.value
  for(let key in allDimensionFilters) {
    if (allDimensionFilters[key].value.length > 0) {
      num++
    }
  }
  selectedDimensionNum.value = num
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