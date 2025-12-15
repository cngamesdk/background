<template>
  <el-select v-model="modelValue" placeholder="请选择聚合时间" style="width:6rem">
    <el-option
        v-for="item in localData"
        :key="item.key"
        :label="item.value"
        :value="item.key"
    />
  </el-select>
</template>

<script setup>

import { ref, defineModel } from 'vue'
import { getAggregationTime } from '@/api/systemManagement'

const modelValue = defineModel({ type: String, default: 'day' })

defineOptions({
  name: 'AggregationTime',
})

const localData = ref([])

const { aggregationTimes } = defineProps({
  aggregationTimes: {
    type: Array,
    default() {
      return []
    }
  }
})

if (aggregationTimes.length > 0 ) {
  aggregationTimes.forEach(function (item) {
    localData.value.push(item)
  })
} else {
  localData.value = getAggregationTime()
}

</script>