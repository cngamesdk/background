<template>
  <div class="calculation-config">
    <el-divider content-position="left">计算逻辑</el-divider>
    <el-form label-width="100px">
      <el-form-item label="计算模式">
        <el-select v-model="config.mode">
          <el-option label="累加" value="accumulate" />
          <el-option label="每日重置" value="daily_reset" />
          <el-option label="去重计数" value="dedup" />
          <el-option label="限时累加" value="time_limited" />
        </el-select>
      </el-form-item>
      <el-form-item label="累加字段">
        <el-input v-model="config.field" placeholder="如: amount, count" />
        <div class="el-form-item__tip">count表示计次，amount表示累加金额</div>
      </el-form-item>
      <el-form-item label="重置周期" v-if="config.mode === 'daily_reset'">
        <el-select v-model="config.reset_cycle">
          <el-option label="每日" value="daily" />
          <el-option label="每周" value="weekly" />
          <el-option label="不重置" value="never" />
        </el-select>
      </el-form-item>
      <el-form-item label="去重字段" v-if="config.mode === 'dedup'">
        <el-input v-model="config.dedup_key" placeholder="如: order_id" />
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({ modelValue: { type: Object, default: () => ({}) } })
const emit = defineEmits(['update:modelValue'])

const config = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})
</script>
