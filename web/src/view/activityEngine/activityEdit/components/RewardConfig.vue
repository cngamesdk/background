<template>
  <div class="reward-config">
    <el-divider content-position="left">奖励策略</el-divider>
    <el-form label-width="100px">
      <el-form-item label="奖励模式">
        <el-select v-model="config.strategy">
          <el-option label="固定奖励" value="fixed" />
          <el-option label="阶梯奖励" value="tiered" />
          <el-option label="概率掉落" value="probability" />
        </el-select>
      </el-form-item>
      <el-form-item label="奖励档位">
        <el-table :data="config.tiers" border style="width:100%">
          <el-table-column label="目标值" width="120">
            <template #default="{ row }">
              <el-input-number v-model="row.threshold" :min="0" size="small" controls-position="right" />
            </template>
          </el-table-column>
          <el-table-column label="概率" width="100" v-if="config.strategy === 'probability'">
            <template #default="{ row }">
              <el-input-number v-model="row.probability" :min="0" :max="1" :step="0.1" size="small" />
            </template>
          </el-table-column>
          <el-table-column label="奖励物品" min-width="300">
            <template #default="{ row }">
              <div v-for="(item, i) in row.items" :key="i" style="display:flex;gap:5px;margin-bottom:8px;align-items:center;">
                <RewardItemSelector v-model:item-code="item.item_code" v-model:item-name="item.item_name" />
                <el-input-number v-model="item.quantity" :min="1" size="small" style="width:100px" placeholder="数量" />
                <el-button size="small" type="danger" :icon="Delete" circle @click="removeItem(row, i)" />
              </div>
              <el-button size="small" type="primary" plain @click="addItem(row)">添加物品</el-button>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="70">
            <template #default="{ $index }">
              <el-button size="small" type="danger" @click="removeTier($index)">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
        <el-button type="primary" plain style="margin-top:10px" @click="addTier">添加档位</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { Delete } from '@element-plus/icons-vue'
import RewardItemSelector from './RewardItemSelector.vue'

const props = defineProps({ modelValue: { type: Object, default: () => ({}) } })
const emit = defineEmits(['update:modelValue'])

const config = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val)
})

const addTier = () => {
  const tiers = [...(config.value.tiers || []), { threshold: 0, probability: 1, items: [] }]
  emit('update:modelValue', { ...config.value, tiers })
}

const removeTier = (idx) => {
  const tiers = [...config.value.tiers]
  tiers.splice(idx, 1)
  emit('update:modelValue', { ...config.value, tiers })
}

const addItem = (tier) => {
  if (!tier.items) tier.items = []
  tier.items.push({ item_code: '', item_name: '', quantity: 1 })
}

const removeItem = (tier, idx) => {
  tier.items.splice(idx, 1)
}
</script>
