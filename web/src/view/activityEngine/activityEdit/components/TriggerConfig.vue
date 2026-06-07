<template>
  <div class="trigger-config">
    <el-divider content-position="left">触发条件</el-divider>
    <el-form label-width="100px">
      <el-form-item label="事件类型">
        <el-select v-model="config.event_type" disabled>
          <el-option label="充值" value="recharge" />
          <el-option label="登录" value="login" />
          <el-option label="签到" value="signin" />
          <el-option label="分享" value="share" />
          <el-option label="自定义" value="custom" />
        </el-select>
      </el-form-item>
      <el-form-item label="附加条件">
        <div v-for="(cond, idx) in config.conditions" :key="idx" style="display:flex;gap:10px;margin-bottom:10px;">
          <el-input v-model="cond.field" placeholder="字段名" style="width:120px" />
          <el-select v-model="cond.operator" style="width:100px">
            <el-option label="等于" value="eq" />
            <el-option label="大于" value="gt" />
            <el-option label="大于等于" value="gte" />
            <el-option label="小于" value="lt" />
            <el-option label="小于等于" value="lte" />
            <el-option label="包含" value="in" />
          </el-select>
          <el-input v-model="cond.value" placeholder="值" style="width:150px" />
          <el-button type="danger" :icon="Delete" circle @click="removeCond(idx)" />
        </div>
        <el-button type="primary" plain @click="addCond">添加条件</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Delete } from '@element-plus/icons-vue'

const props = defineProps({ modelValue: { type: Object, default: () => ({}) }, activityType: String })
const emit = defineEmits(['update:modelValue'])

const config = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const addCond = () => {
  const newConfig = { ...config.value, conditions: [...(config.value.conditions || []), { field: '', operator: 'eq', value: '' }] }
  emit('update:modelValue', newConfig)
}

const removeCond = (idx) => {
  const conditions = [...config.value.conditions]
  conditions.splice(idx, 1)
  emit('update:modelValue', { ...config.value, conditions })
}
</script>
