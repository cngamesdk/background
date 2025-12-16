<template>
  <el-date-picker style="width:16rem"
                  v-model="dateRange"
                  type="daterange"
                  :shortcuts="shortcuts"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
  />
</template>
<script setup>
import { ref, defineModel } from 'vue'
const dateRange = defineModel({ type: Array, default: [new Date(), new Date()] })

defineOptions({
  name: 'DateRange',
})

const shortcuts = [
  {
    text: '今天',
    value: [new Date(), new Date()],
  },
  {
    text: '昨天',
    value: () => {
      const date = new Date()
      date.setTime(date.getTime() - 3600 * 1000 * 24)
      return [date, date]
    },
  },
  {
    text: '最近一周',
    value: () => {
      const date = new Date()
      date.setTime(date.getTime() - 3600 * 1000 * 24 * 7)
      return [date, new Date()]
    },
  },
]

</script>