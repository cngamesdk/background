<template>
  <el-scrollbar>
    <div
      v-for="(item, index) in notices"
      :key="index"
      class="flex items-center mb-1.5 gap-3"
    >
      <el-tag :type="item.type" size="small">
        {{ item.typeTitle }}
      </el-tag>
      <el-tooltip effect="light" :content="item.title" placement="top">
        <div class="text-xs text-gray-700 dark:text-gray-300 line-clamp-1">
          {{ item.title }}
        </div>
      </el-tooltip>
    </div>
  </el-scrollbar>
</template>

<script setup>
  const notices = [
    {
      type: 'primary',
      typeTitle: '公告',
      title: '梅花互娱上线，欢迎新老用户试用'
    }
  ]
</script>

<style scoped lang="scss"></style>
