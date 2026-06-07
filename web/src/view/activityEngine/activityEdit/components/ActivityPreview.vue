<template>
  <div class="activity-preview">
    <el-divider content-position="left">配置预览</el-divider>
    <el-descriptions :column="2" border>
      <el-descriptions-item label="活动名称">{{ form.activity_name }}</el-descriptions-item>
      <el-descriptions-item label="活动类型">{{ form.activity_type }}</el-descriptions-item>
      <el-descriptions-item label="开始时间">{{ form.start_time }}</el-descriptions-item>
      <el-descriptions-item label="结束时间">{{ form.end_time }}</el-descriptions-item>
      <el-descriptions-item label="描述" :span="2">{{ form.description || '-' }}</el-descriptions-item>
    </el-descriptions>

    <el-divider content-position="left">触发条件 JSON</el-divider>
    <el-input type="textarea" :rows="4" :model-value="JSON.stringify(triggerConfig, null, 2)" readonly />

    <el-divider content-position="left">计算逻辑 JSON</el-divider>
    <el-input type="textarea" :rows="4" :model-value="JSON.stringify(calculationConfig, null, 2)" readonly />

    <el-divider content-position="left">奖励策略 JSON</el-divider>
    <el-input type="textarea" :rows="6" :model-value="JSON.stringify(rewardConfig, null, 2)" readonly />

    <el-divider content-position="left">约束规则 JSON</el-divider>
    <el-input type="textarea" :rows="4" :model-value="JSON.stringify(constraintConfig, null, 2)" readonly />
  </div>
</template>

<script setup>
defineProps({
  form: { type: Object, default: () => ({}) },
  triggerConfig: { type: Object, default: () => ({}) },
  calculationConfig: { type: Object, default: () => ({}) },
  rewardConfig: { type: Object, default: () => ({}) },
  constraintConfig: { type: Object, default: () => ({}) },
})
</script>
