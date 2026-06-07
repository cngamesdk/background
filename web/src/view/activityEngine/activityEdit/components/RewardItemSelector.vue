<template>
  <div class="reward-item-selector" style="display:flex;gap:5px;">
    <el-select
      v-model="selectedCode"
      filterable
      remote
      reserve-keyword
      placeholder="搜索奖励道具"
      :remote-method="handleSearch"
      :loading="loading"
      size="small"
      style="width:200px"
      @change="handleSelect"
    >
      <el-option
        v-for="item in options"
        :key="item.item_code"
        :label="`${item.item_name} (${item.item_code})`"
        :value="item.item_code"
      />
    </el-select>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue'
import { rewardItemSearch } from '@/api/activityEngine'

const props = defineProps({
  itemCode: { type: String, default: '' },
  itemName: { type: String, default: '' },
})
const emit = defineEmits(['update:itemCode', 'update:itemName'])

const selectedCode = ref(props.itemCode)
const loading = ref(false)
const options = ref([])

watch(() => props.itemCode, (val) => { selectedCode.value = val })

const handleSearch = async (query) => {
  if (!query || query.length < 1) {
    options.value = []
    return
  }
  loading.value = true
  try {
    const res = await rewardItemSearch({ keyword: query, page: 1, pageSize: 20 })
    if (res.code === 0) {
      options.value = res.data.list || []
    }
  } finally {
    loading.value = false
  }
}

const handleSelect = (code) => {
  const item = options.value.find(i => i.item_code === code)
  emit('update:itemCode', code)
  emit('update:itemName', item ? item.item_name : '')
}
</script>
