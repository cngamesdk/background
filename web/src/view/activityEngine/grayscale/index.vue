<template>
  <div class="grayscale-container">
    <el-card header="灰度发布管理">
      <el-form :model="form" label-width="100px">
        <el-form-item label="活动ID">
          <el-input-number v-model="form.id" :min="1" controls-position="right" />
        </el-form-item>
        <el-form-item label="灰度比例">
          <el-slider v-model="form.grayscale_ratio" :min="1" :max="100" show-input />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleUpdate" :loading="loading">更新灰度比例</el-button>
        </el-form-item>
      </el-form>

      <el-alert type="info" :closable="false" style="margin-top:20px">
        <template #title>
          灰度说明：设置为N%表示只有用户ID % 100 &lt; N 的用户能看到和参与该活动。设置为100%表示全量发布。
        </template>
      </el-alert>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { grayscaleUpdate } from '@/api/activityEngine'

const form = ref({ id: null, grayscale_ratio: 100 })
const loading = ref(false)

const handleUpdate = async () => {
  if (!form.value.id) {
    ElMessage.warning('请输入活动ID')
    return
  }
  await ElMessageBox.confirm(`确认将活动 #${form.value.id} 的灰度比例更新为 ${form.value.grayscale_ratio}%?`, '确认')
  loading.value = true
  try {
    const res = await grayscaleUpdate(form.value)
    if (res.code === 0) {
      ElMessage.success('更新成功')
    }
  } finally {
    loading.value = false
  }
}
</script>
