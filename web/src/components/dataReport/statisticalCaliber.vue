<template>
  <el-select v-model="modelValue" placeholder="请选择口径" style="width:8rem">
    <el-option
        v-for="item in localStatisticalCalibers"
        :key="item.key"
        :label="item.value"
        :value="item.key"
    />
  </el-select>
</template>

<script setup>

import { ref, defineModel, reactive } from 'vue'
import { getStatisticalCaliber, } from '@/api/systemManagement'

const modelValue = defineModel({ type: String, default: 'root-game-back-30' })

defineOptions({
  name: 'StatisticalCaliber',
})

const localStatisticalCalibers = ref([])

const { statisticalCalibers } = defineProps({
  statisticalCalibers: {
    type: Array,
    default() {
      return []
    }
  }
})

if (statisticalCalibers.length > 0 ) {
  statisticalCalibers.forEach(function (item) {
    localStatisticalCalibers.value.push(item)
  })
} else {
  localStatisticalCalibers.value = getStatisticalCaliber()
}

</script>