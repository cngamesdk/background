<template>
  <div>
    <el-tabs v-model="activeName" class="mix-tabs" @tab-click="handleClick" type="border-card">
      <el-tab-pane label="通用" name="common">
        <div>
        <AdvertisingMixCommonConfig></AdvertisingMixCommonConfig>
        </div>
      </el-tab-pane>
      <el-tab-pane label="账号" name="account" style="padding: 20px 55px;">账号</el-tab-pane>
      <el-tab-pane label="一级" name="ad1">一级</el-tab-pane>
      <el-tab-pane label="二级" name="ad2">二级</el-tab-pane>
      <el-tab-pane label="三级" name="ad3">三级</el-tab-pane>
      <el-tab-pane label="组合" name="mix">组合</el-tab-pane>
      <el-tab-pane label="确认" name="confirm">确认</el-tab-pane>
    </el-tabs>
  </div>
</template>

<script setup>

import { ref, defineModel } from 'vue'
import { getAggregationTime } from '@/api/systemManagement'
import AdvertisingMixCommonConfig from '../../../components/advertising/mix/commonConfig.vue'

defineOptions({
  name: 'AdvertisingMixConfig',
})

const activeName = ref('common')

</script>

<style>
.mix-tabs .el-tabs__item {
  text-align: center;
  font-weight: bold;
  display: flex;
  width: 150px;
}
.mix-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
.mix-tabs .custom-tabs-label .el-icon {
  vertical-align: middle;
}
.mix-tabs .custom-tabs-label span {
  vertical-align: middle;
  margin-left: 4px;
}
</style>