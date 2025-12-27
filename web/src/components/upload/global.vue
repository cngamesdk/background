<template>
  <div>
    <el-upload
        :action="`${getBaseUrl()}/fileUploadAndDownload/upload`"
        :before-upload="checkFile"
        :on-error="uploadError"
        :on-success="uploadSuccess"
        :show-file-list="false"
        :data="{'classId': props.classId}"
        :headers="{'x-token': token}"
        multiple
        class="upload-btn"
    >
      <el-button type="primary" :icon="Upload">点击上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { isJarMime, } from '@/utils/zap'
import { getBaseUrl } from '@/utils/format'
import { Upload } from "@element-plus/icons-vue";
import { useUserStore } from "@/pinia";

defineOptions({
  name: 'GlobalUploadCommon'
})

const userStore = useUserStore()

const token = userStore.token

const props = defineProps({
  classId: {
    type: Number,
    default: 0
  },
  size: {
    type: Number,
    default: 0
  },
  allowMimes: {
    type: Array,
    default: []
  }
})

const emit = defineEmits(['on-success'])

const fullscreenLoading = ref(false)

const checkFile = (file) => {
  fullscreenLoading.value = true
  if (props.size > 0) {
    if (file.size > props.size) {
      ElMessage.error(
          '文件大小超过限制'
      )
      fullscreenLoading.value = false
      return false
    }
  }
  if (props.allowMimes.length > 0) {
    if (props.allowMimes.indexOf(file.type) < 0) {
       ElMessage.error(
          '文件格式非法'
      )
      fullscreenLoading.value = false
      return false
    }
  }

  return true
}

const uploadSuccess = (res) => {
  const { data } = res
  if (data) {
    emit('on-success', data)
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: '上传失败'
  })
  fullscreenLoading.value = false
}
</script>
