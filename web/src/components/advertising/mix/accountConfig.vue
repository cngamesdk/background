<template>
  <div>
    <el-form ref="searchForm" :inline="true" :model="searchInfo">
      <el-form-item>
        <el-input
            v-model="searchInfo.account_id"
            autosize
            type="textarea"
            placeholder="请输入帐户ID"
        />
      </el-form-item>
      <el-form-item>
        <el-input
            v-model="searchInfo.account_name"
            autosize
            type="textarea"
            placeholder="请输入帐户名"
        />
      </el-form-item>
      <el-form-item>
        <el-input
            v-model="searchInfo.user_name"
            autosize
            type="textarea"
            placeholder="请输入用户名"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="search" @click="onSearchSubmit">
          查询
        </el-button>
        <el-button icon="refresh" @click="onReset"> 重置 </el-button>
      </el-form-item>
    </el-form>
  </div>

  <div>
    <el-space wrap style="align-items: normal;">
      <el-card class="box-card nav-left">
        <div class="gva-table-box">
          <el-table :data="accountListData" stripe row-key="id">
            <el-table-column
                align="left"
                label="主体"
                min-width="100">
              <template #default="scope">
                {{ scope.row.account_name }}
              </template>
            </el-table-column>
            <el-table-column
                align="left"
                label="帐户ID"
                min-width="100">
              <template #default="scope">
                {{ scope.row.account_id }}
              </template>
            </el-table-column>
            <el-table-column
                align="left"
                min-width="150"
                label="帐户名称">
              <template #default="scope">
                {{ scope.row.account_name }}
              </template>
            </el-table-column>
            <el-table-column
                align="left"
                label="用户名"
                min-width="150"
            >
              <template #default="scope">
                {{ scope.row.user_name }}
              </template>
            </el-table-column>
            <el-table-column
                align="left"
                label="操作"
                min-width="180"
            >
              <template #default="scope">
                <el-button type="primary">使用</el-button>
              </template>
            </el-table-column>
          </el-table>
          <div class="gva-pagination">
            <el-pagination
                :current-page="page"
                :page-size="pageSize"
                :page-sizes="[10, 30, 50, 100]"
                :total="total"
                layout="total, sizes, prev, pager, next, jumper"
                @current-change="handleCurrentChange"
                @size-change="handleSizeChange"
            />
          </div>
        </div>
      </el-card>
      <el-card class="box-card nav-right">
        <div class="gva-table-box">
          <el-table :data="tableData" stripe row-key="id">
            <el-table-column
                align="left"
                label="媒体帐户"
                min-width="100">
              <template #default="scope">
                {{ scope.row.account_id }}-{{ scope.row.account_name }}
              </template>
            </el-table-column>
            <el-table-column
                align="left"
                label="操作"
                min-width="100">
              <template #default="scope">
                <el-button type="primary" :icon="Delete">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
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

const searchInfo = ref(Object.assign({}, defaultSearchInfo))

const defaultConfigInfo = {
  list: [],
}

const configInfo = ref(Object.assign({}, defaultConfigInfo))

defineOptions({
  name: 'AdvertisingMixAccountConfig',
})

</script>

<style scoped>
.nav-left{
  width: 50rem;
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