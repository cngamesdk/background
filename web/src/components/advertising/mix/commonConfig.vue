<template>
  <el-space wrap style="align-items: normal;">
    <el-card class="box-card nav-left">
      <div>
        <el-form ref="searchForm" :model="searchInfo">
          <el-form-item label="平台">
            <el-select
                @change="handleChangePlatform"
                filterable
                v-model="configInfo.platform_id" placeholder="请选择所属平台">
              <el-option
                  v-for="item in platforms"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="媒体">
            <el-select
                filterable
                v-model="configInfo.code" placeholder="请选择媒体">
              <el-option
                  v-for="item in medias"
                  :key="item.key_2"
                  :label="item.value"
                  :value="item.key_2"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="主体">
            <el-select
                remote
                filterable
                :remote-method="handleSearchCompany"
                v-model="configInfo.company_id" placeholder="请选择主体">
              <el-option
                  v-for="item in companys"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
            </el-select>
          </el-form-item>

        </el-form>
      </div>
    </el-card>
    <el-card class="box-card nav-right">
      <div>
        <el-form ref="searchForm" :model="configInfo">
          <el-form-item label="广告位名称">
            <el-input
                v-model="configInfo.site_name"
                style="width: 340px"
                placeholder="请输入广告位名称"
                clearable
            />
          </el-form-item>
          <el-form-item label="游戏">
            <el-select
                remote
                filterable
                :remote-method="handleSearchGameId"
                v-model="configInfo.game_id"
                placeholder="请选择投放游戏"
                style="width:150px;"
            >
              <el-option
                  v-for="item in games"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
            </el-select>
            <el-select
                v-model="configInfo.game_packageing"
                placeholder="请选择是否打包"
                style="width:150px;border-left: none;"
            >
              <el-option
                  v-for="item in gamePackageings"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="转化来源">
            <el-select
                v-model="configInfo.convert_source"
                placeholder="请选择转化来源"
                style="width:200px;"
            >
              <el-option
                  v-for="item in convertSources"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="转化类型">
            <el-select
                v-model="configInfo.convert_type"
                placeholder="请选择转化类型"
                style="width:200px;"
            >
              <el-option
                  v-for="item in convertTypes"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
            </el-select>
          </el-form-item>

          <el-form-item label="深度转化类型">
            <el-select
                v-model="configInfo.deep_convert_type"
                placeholder="请选择深度转化类型"
                style="width:200px;"
            >
              <el-option
                  v-for="item in deepConvertTypes"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
            </el-select>
          </el-form-item>

        </el-form>
      </div>
    </el-card>
  </el-space>
</template>
<script setup>

import { ref } from 'vue'
import { searchPlatform, searchMedia, searchCompany, searchSubGame } from '@/api/systemManagement'

const defaultConfigInfo = {
  platform_id: 0,
  code: '',
  company_id: 0,
  site_name: '',
  game_id: 0,
  game_packageing: 'none',
  convert_source: '',
  convert_type: '',
  deep_convert_type: '',
}

const platforms = ref([])
const medias = ref([])
const companys = ref([])
const games = ref([])
// 游戏打包配置
const gamePackageings = ref([
  {label: '不打包', value: 'none'},
  {label: '平台打包', value: 'plat-pack'},
  {label: '媒体分包', value: 'media-pack'},
])
//转化来源配置
const convertSources = ref([
  {key: 'api', value: 'API'},
  {key: 'sdk', value: 'SDK'},
])
//转化类型配置
const convertTypes = ref([
  {key: 'activate', value: '激活'},
  {key: 'registration', value: '注册'},
  {key: 'key-behaviors', value: '关键行为'},
  {key: 'pay', value: '付费'},
])

//深度转化类型配置
const deepConvertTypes = ref([
  {key: 'registration', value: '注册'},
  {key: 'pay', value: '付费'},
  {key: 'day1-retention', value: '次留'},
  {key: 'pay-roi', value: '付费ROI'},
])

const configInfo = ref(Object.assign({}, defaultConfigInfo))

defineOptions({
  name: 'AdvertisingMixCommonConfig',
})

const getPlatforms = async () => {
  const result = await searchPlatform()
  if (result.code === 0) {
    platforms.value = result.data
  }
}

const handleChangePlatform = async () => {
  const result = await searchMedia({platform_id: configInfo.value.platform_id})
  if (result.code === 0) {
    medias.value = result.data
  }
  getCompanys({})
  getGames({})
}

const handleSearchCompany = (query) => {
  getCompanys({keyword: query})
}

const getCompanys = async (data) => {
  const result = await searchCompany({platform_id: configInfo.value.platform_id, ...data})
  if (result.code === 0) {
    companys.value = result.data
  }
}

const handleSearchGameId = (query) => {
  getGames({keyword: query})
}

const getGames = async (data) => {
  const result = await searchSubGame({platform_id: configInfo.value.platform_id, ...data})
  if (result.code === 0) {
    games.value = result.data
  }
}

const initData = () => {
  getPlatforms()
}

initData()

</script>

<style>
.nav-left{
  width: 25rem;
}
.nav-right{
  width: 53rem;
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