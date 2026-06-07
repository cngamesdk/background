<template>
  <div class="activity-edit-container">
    <el-steps :active="activeStep" finish-status="success" align-center style="margin-bottom: 30px;">
      <el-step title="触发条件" />
      <el-step title="计算逻辑" />
      <el-step title="奖励策略" />
      <el-step title="约束规则" />
      <el-step title="预览确认" />
    </el-steps>

    <el-card>
      <el-form :model="form" label-width="120px" v-show="activeStep === 0">
        <el-form-item label="活动名称" required>
          <el-input v-model="form.activity_name" placeholder="请输入活动名称" />
        </el-form-item>
        <el-form-item label="活动类型" required>
          <el-select v-model="form.activity_type" placeholder="请选择活动类型">
            <el-option label="累计充值" value="recharge" />
            <el-option label="登录奖励" value="login" />
            <el-option label="签到" value="signin" />
            <el-option label="分享奖励" value="share" />
            <el-option label="自定义" value="custom" />
          </el-select>
        </el-form-item>
        <el-form-item label="生效时间" required>
          <el-date-picker v-model="timeRange" type="datetimerange" range-separator="至"
            start-placeholder="开始时间" end-placeholder="结束时间" value-format="YYYY-MM-DD HH:mm:ss" />
        </el-form-item>
        <el-form-item label="活动描述">
          <el-input v-model="form.description" type="textarea" :rows="3" placeholder="请输入活动描述" />
        </el-form-item>
        <TriggerConfig v-model="triggerConfig" :activity-type="form.activity_type" />
      </el-form>

      <CalculationConfig v-model="calculationConfig" v-show="activeStep === 1" />
      <RewardConfig v-model="rewardConfig" v-show="activeStep === 2" />
      <ConstraintConfig v-model="constraintConfig" v-show="activeStep === 3" />
      <ActivityPreview v-show="activeStep === 4" :form="form" :trigger-config="triggerConfig"
        :calculation-config="calculationConfig" :reward-config="rewardConfig" :constraint-config="constraintConfig" />
    </el-card>

    <div class="step-actions" style="margin-top: 20px; text-align: center;">
      <el-button @click="prevStep" v-if="activeStep > 0">上一步</el-button>
      <el-button type="primary" @click="nextStep" v-if="activeStep < 4">下一步</el-button>
      <el-button type="success" @click="handleSubmit" v-if="activeStep === 4">提交保存</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { activityAdd, activityModify, activityDetail } from '@/api/activityEngine'
import TriggerConfig from './components/TriggerConfig.vue'
import CalculationConfig from './components/CalculationConfig.vue'
import RewardConfig from './components/RewardConfig.vue'
import ConstraintConfig from './components/ConstraintConfig.vue'
import ActivityPreview from './components/ActivityPreview.vue'

const route = useRoute()
const router = useRouter()

const activeStep = ref(0)
const timeRange = ref([])
const form = ref({
  activity_name: '',
  activity_type: '',
  description: '',
  start_time: '',
  end_time: '',
  platform_id: 0,
  game_id: 0,
  priority: 0,
})
const triggerConfig = ref({ event_type: '', conditions: [] })
const calculationConfig = ref({ mode: 'accumulate', field: 'count', reset_cycle: 'never', dedup_key: '' })
const rewardConfig = ref({ strategy: 'tiered', tiers: [] })
const constraintConfig = ref({ user_segments: ['all'], daily_claim_max: 0, total_claim_max: 0, cooldown_sec: 0, time_windows: [] })

watch(timeRange, (val) => {
  if (val && val.length === 2) {
    form.value.start_time = val[0]
    form.value.end_time = val[1]
  }
})

watch(() => form.value.activity_type, (val) => {
  triggerConfig.value.event_type = val
})

const prevStep = () => { if (activeStep.value > 0) activeStep.value-- }
const nextStep = () => { if (activeStep.value < 4) activeStep.value++ }

const handleSubmit = async () => {
  const payload = {
    ...form.value,
    trigger_config: JSON.stringify(triggerConfig.value),
    calculation_config: JSON.stringify(calculationConfig.value),
    reward_config: JSON.stringify(rewardConfig.value),
    constraint_config: JSON.stringify(constraintConfig.value),
  }

  let res
  if (route.query.id) {
    payload.id = Number(route.query.id)
    res = await activityModify(payload)
  } else {
    res = await activityAdd(payload)
  }

  if (res.code === 0) {
    ElMessage.success('保存成功')
    router.push({ name: 'activityList' })
  }
}

onMounted(async () => {
  const id = route.query.id || route.query.clone
  if (id) {
    const res = await activityDetail({ id: Number(id) })
    if (res.code === 0) {
      const data = res.data
      form.value = { activity_name: data.activity_name, activity_type: data.activity_type, description: data.description,
        start_time: data.start_time, end_time: data.end_time, platform_id: data.platform_id, game_id: data.game_id, priority: data.priority }
      timeRange.value = [data.start_time, data.end_time]
      triggerConfig.value = JSON.parse(data.trigger_config || '{}')
      calculationConfig.value = JSON.parse(data.calculation_config || '{}')
      rewardConfig.value = JSON.parse(data.reward_config || '{}')
      constraintConfig.value = JSON.parse(data.constraint_config || '{}')
      if (route.query.clone) { form.value.activity_name += ' (副本)' }
    }
  }
})
</script>
