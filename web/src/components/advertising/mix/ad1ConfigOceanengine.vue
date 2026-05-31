<template>
  <div>
    <el-space wrap style="align-items: normal;" direction="vertical">
      <el-card class="box-card">
       <el-form :model="configInfo">
         <el-form-item>
           <el-segmented v-model="configInfo.ad_type" :options="adTypes" @change="handleChangeAdType"/>
         </el-form-item>
         <el-form-item>
           <el-text class="mx-1" size="large">营销产品与目标</el-text>
         </el-form-item>
         <el-form-item label="营销目的">
           <el-segmented v-model="configInfo.landing_type" :options="landingTypes" />
         </el-form-item>
         <el-form-item>
           <el-segmented v-if="configInfo.landing_type === 'APP'" v-model="configInfo.app_promotion_type" :options="appPromotionTypes" />
         </el-form-item>
         <el-form-item label="营销场景">
           <el-segmented v-model="configInfo.marketing_goal" :options="marketingGoals" />
         </el-form-item>
       </el-form>
      </el-card>

      <el-card class="box-card">
        <el-form :model="configInfo">
          <el-form-item label="投放模式">
            <el-segmented v-model="configInfo.delivery_mode" :options="deliveryModes" />
          </el-form-item>
          <el-form-item label="投放类型">
            <el-segmented v-model="configInfo.delivery_type" :options="deliveryTypes" />
          </el-form-item>
        </el-form>
      </el-card>

      <el-card class="box-card">
        <el-form :model="configInfo">
          <el-form-item label="AIGC动态创意">
            <el-switch
                v-model="configInfo.aigc_dynamic_creative_switch"
                active-value="ON"
                inactive-value="OFF"
            />
          </el-form-item>
        </el-form>
      </el-card>

      <el-card class="box-card">
        <el-form :model="configInfo">
          <el-form-item label="状态">
            <el-segmented v-model="configInfo.operation" :options="operations" />
          </el-form-item>
          <el-form-item label="项目名称">
            <el-input v-model="configInfo.name" placeholder="请输入项目名称"></el-input>
          </el-form-item>
        </el-form>
      </el-card>

    </el-space>
  </div>

</template>
<script setup>

import { ref } from 'vue'
import { searchPlatform, searchMedia, searchCompany, searchSubGame } from '@/api/systemManagement'

const defaultSearchInfo = {
  account_id: 0,
  account_name: '',
  user_name: '',
}

const adTypes = ref([
  {label: '通投',value: 'ALL'},
  {label: '搜索',value: 'SEARCH'},
])

// 营销目的
const landingTypes = ref([
  {label: '应用',value: 'APP'},
  {label: '销售线索投放',value: 'LINK'},
  {label: '小程序',value: 'MICRO_GAME'},
  {label: '电商',value: 'SHOP'},
  {label: '快应用',value: 'QUICK_APP'},
  {label: '原生互动',value: 'NATIVE_ACTION'},
  {label: '商品目录',value: 'DPA'},
])

const appPromotionTypes = ref([
  {label: '应用下载',value: 'DOWNLOAD'},
  {label: '应用调起',value: 'LAUNCH'},
  {label: '预约下载',value: 'RESERVE'},
])

const marketingGoals = ref([
  {label: '短视频/图片',value: 'VIDEO_AND_IMAGE'},
  {label: '直播',value: 'LIVE'},
])

// 项目状态
const operations = ref([
  {label: '开启',value: 'ENABLE'},
  {label: '关闭',value: 'DISABLE'},
])

// 投放模式
const deliveryModes = ref([
  {label: '手动投放',value: 'MANUAL'},
  {label: '自动投放',value: 'PROCEDURAL'},
])

// 投放类型
const deliveryTypes = ref([
  {label: '常规投放',value: 'NORMAL'},
  {label: '智能托管',value: 'UBX_INTELLIGENT'},
])

const searchInfo = ref(Object.assign({}, defaultSearchInfo))

const defaultConfigInfo = {
  name: '',
  delivery_mode: 'MANUAL',
  delivery_type: 'NORMAL',
  operation: 'ENABLE',
  ad_type: 'ALL',
  landing_type: 'APP',
  app_promotion_type: 'DOWNLOAD',
  marketing_goal: 'VIDEO_AND_IMAGE',
  aigc_dynamic_creative_switch: 'OFF',
}

const configInfo = ref(Object.assign({}, defaultConfigInfo))

defineOptions({
  name: 'AdvertisingMixAd1ConfigOceanengine',
})

const handleChangeAdType = (val) => {
  let tempDeliveryTypes = []
  deliveryTypes.value.forEach(function (item) {
    if (item.value !== 'DURATION' ) {
      tempDeliveryTypes.push(item)
    }
  })
  if (val === 'SEARCH') {
    tempDeliveryTypes.push({label: '极速智投',value: 'DURATION'})
  }
  deliveryTypes.value = tempDeliveryTypes
}

</script>

<style scoped>
.nav-left{
  width: 65rem;
}
.nav-right{
  width: 25rem;
}
/* 响应式设计 */
@media (max-width: 1100px) {
  .nav-left{
    width: 25rem;
  }
  .nav-right{
    width: 53rem;
  }
}

@media (max-width: 768px) {
  .nav-left{
    width: 25rem;
  }
  .nav-right{
    width: 53rem;
  }
}

@media (max-width: 576px) {
  .nav-left{
    width: 25rem;
  }
  .nav-right{
    width: 53rem;
  }
}
</style>