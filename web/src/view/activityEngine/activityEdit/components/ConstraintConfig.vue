<template>
  <div class="constraint-config">
    <el-divider content-position="left">约束规则</el-divider>
    <el-form label-width="120px">
      <el-form-item label="用户分群">
        <el-select v-model="config.user_segments" multiple placeholder="选择用户分群">
          <el-option label="全部用户" value="all" />
          <el-option label="新用户" value="new_user" />
          <el-option label="VIP用户" value="vip" />
          <el-option label="付费用户" value="paid" />
          <el-option label="活跃用户" value="active" />
        </el-select>
      </el-form-item>
      <el-form-item label="每日领取上限">
        <el-input-number v-model="config.daily_claim_max" :min="0" controls-position="right" />
        <span style="margin-left:10px;color:#909399">0表示不限制</span>
      </el-form-item>
      <el-form-item label="总领取次数">
        <el-input-number v-model="config.total_claim_max" :min="0" controls-position="right" />
        <span style="margin-left:10px;color:#909399">0表示不限制</span>
      </el-form-item>
      <el-form-item label="冷却时间(秒)">
        <el-input-number v-model="config.cooldown_sec" :min="0" controls-position="right" />
      </el-form-item>
      <el-form-item label="时间窗口">
        <div v-for="(win, idx) in config.time_windows" :key="idx" style="display:flex;gap:10px;margin-bottom:10px;">
          <el-input-number v-model="win.start_hour" :min="0" :max="23" placeholder="开始时" size="small" />
          <span>时 至</span>
          <el-input-number v-model="win.end_hour" :min="0" :max="23" placeholder="结束时" size="small" />
          <span>时</span>
          <el-button type="danger" size="small" @click="removeWindow(idx)">删除</el-button>
        </div>
        <el-button type="primary" plain size="small" @click="addWindow">添加时间窗口</el-button>
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

const addWindow = () => {
  const windows = [...(config.value.time_windows || []), { start_hour: 0, end_hour: 24 }]
  emit('update:modelValue', { ...config.value, time_windows: windows })
}

const removeWindow = (idx) => {
  const windows = [...config.value.time_windows]
  windows.splice(idx, 1)
  emit('update:modelValue', { ...config.value, time_windows: windows })
}
</script>
