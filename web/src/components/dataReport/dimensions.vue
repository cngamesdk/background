<template>
  <el-popover placement="bottom-start" :width="500" trigger="click" @hide="hideComponent">
    <template #reference>
      <el-button>维度选择({{ selectedNum }})</el-button>
    </template>

    <div class="item-container" v-for="(group, index) in localDimensions" :key="index">
      <div class="title">
        <el-checkbox
            :indeterminate="group.indeterminate"
            v-model="group.allChecked"
            @change="(val) => handleGroupAllChange(val, index)"
        >
          {{ group.value }}
        </el-checkbox>
      </div>
      <el-checkbox-group
          v-model="group.selected"
          @change="(val) => handleGroupChange(val, index)"
      >
        <el-checkbox
            v-for="(item, indexChild) in group.childs" :key="indexChild" :label="item.value" :value="item.key"
            v-model="item.key"
        >
          {{ item.value }}
        </el-checkbox>
      </el-checkbox-group>
    </div>
  </el-popover>
</template>

<script setup>

import { ref, defineModel, reactive } from 'vue'

const modelValue = defineModel({ type: Array, default: [] })

defineOptions({
  name: 'Dimensions',
})

const { dimensions } = defineProps({
  dimensions: {
    type: Array,
    default() {
      return []
    }
  }
})

const localDimensions = reactive([])

// 为每组添加计算属性
dimensions.forEach((group, index) => {
  group.selected = []
  modelValue.value.forEach(function (item) {
    group.childs.forEach(function (childItem) {
      if (childItem.key === item) {
        group.selected.push(item)
      }
    })
  })
  group.allChecked = group.selected.length === group.childs.length
  group.indeterminate = (group.selected.length > 0 && group.selected.length < group.childs.length)
  localDimensions.push(group)
})

const selectedNum = ref(0)

const hideComponent = () => {
  let selectedValues = []
  localDimensions.forEach(function (item) {
    selectedValues.push(...item.selected)
  })
  modelValue.value = selectedValues
}

const handleGroupAllChange = (val, index) => {
  const indexValue = localDimensions[index]
  localDimensions[index].selected = val ? indexValue.childs.map(opt => opt.key) : []
  localDimensions[index].indeterminate = (indexValue.selected.length > 0 && indexValue.selected.length < indexValue.childs.length)

  calcSelectedNum()
}

const handleGroupChange = (val, index) => {
  const indexValue = localDimensions[index]
  localDimensions[index].selected = val ? val : []
  localDimensions[index].indeterminate = (indexValue.selected.length > 0 && indexValue.selected.length < indexValue.childs.length)
  localDimensions[index].allChecked = (indexValue.selected.length > 0 && indexValue.selected.length >= indexValue.childs.length)

  calcSelectedNum()
}

const calcSelectedNum = () => {
  let num = 0
  localDimensions.forEach(function (item) {
    num = num + item.selected.length
  })
  selectedNum.value = num
}

</script>

<style lang="scss">
.item-container{
  margin: 0 0 1rem 0;
}
.title {
  padding: 0 0 .5rem 0;
  margin: 0 0 .5rem 0;
  border-bottom: 1px solid #d8d4d4;
}
</style>