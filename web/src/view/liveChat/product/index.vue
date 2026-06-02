<template>
  <div>
    <el-card>
      <el-form :inline="true" :model="searchForm" size="small">
        <el-form-item label="产品名称"><el-input v-model="searchForm.name" placeholder="产品名称" /></el-form-item>
        <el-form-item label="状态"><el-select v-model="searchForm.status" clearable><el-option label="正常" value="normal" /><el-option label="禁用" value="disabled" /></el-select></el-form-item>
        <el-form-item><el-button type="primary" @click="getList">查询</el-button><el-button @click="openDialog()">新增产品</el-button></el-form-item>
      </el-form>
    </el-card>
    <el-card style="margin-top:10px">
      <el-table :data="tableData" border stripe>
        <el-table-column prop="ID" label="ID" width="80" />
        <el-table-column prop="product_code" label="产品编码" width="150" />
        <el-table-column prop="name" label="产品名称" />
        <el-table-column prop="welcome_title" label="欢迎标题" />
        <el-table-column prop="status" label="状态" width="80"><template #default="{row}">{{ row.status === 'normal' ? '正常' : '禁用' }}</template></el-table-column>
        <el-table-column label="操作" width="200">
          <template #default="{row}">
            <el-button size="small" @click="openDialog(row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row.ID)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination v-model:current-page="page" :total="total" :page-size="pageSize" layout="total,prev,pager,next" @current-change="getList" style="margin-top:10px" />
    </el-card>

    <el-dialog :title="dialogTitle" v-model="dialogVisible" width="600px">
      <el-form :model="form" label-width="100px">
        <el-form-item label="产品编码"><el-input v-model="form.product_code" :disabled="!!form.ID" /></el-form-item>
        <el-form-item label="产品名称"><el-input v-model="form.name" /></el-form-item>
        <el-form-item label="Logo URL"><el-input v-model="form.logo" placeholder="http://..." /></el-form-item>
        <el-form-item label="欢迎标题"><el-input v-model="form.welcome_title" /></el-form-item>
        <el-form-item label="欢迎消息"><el-input v-model="form.welcome_message" type="textarea" :rows="3" /></el-form-item>
        <el-form-item label="状态"><el-select v-model="form.status"><el-option label="正常" value="normal" /><el-option label="禁用" value="disabled" /></el-select></el-form-item>
      </el-form>
      <template #footer><el-button @click="dialogVisible=false">取消</el-button><el-button type="primary" @click="handleSave">保存</el-button></template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { getProductList, createProduct, updateProduct, deleteProduct } from '@/api/liveChat'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({ name: 'LiveChatProduct' })

const searchForm = reactive({ name: '', status: '' })
const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const dialogVisible = ref(false)
const form = reactive({ ID: 0, product_code: '', name: '', logo: '', welcome_title: '', welcome_message: '', status: 'normal' })
const dialogTitle = computed(() => form.ID ? '编辑产品' : '新增产品')

const getList = async () => {
  const res = await getProductList({ ...searchForm, page: page.value, pageSize: pageSize.value })
  if (res.code === 0) { tableData.value = res.data?.list || []; total.value = res.data?.total || 0 }
}

const openDialog = (row) => {
  if (row) { Object.assign(form, row) } else { Object.assign(form, { ID: 0, product_code: '', name: '', logo: '', welcome_title: '', welcome_message: '', status: 'normal' }) }
  dialogVisible.value = true
}

const handleSave = async () => {
  const fn = form.ID ? updateProduct : createProduct
  const res = await fn(form)
  if (res.code === 0) { ElMessage.success('保存成功'); dialogVisible.value = false; getList() }
}

const handleDelete = (id) => {
  ElMessageBox.confirm('确定删除吗？', '提示', { type: 'warning' }).then(async () => {
    await deleteProduct({ ID: id }); ElMessage.success('删除成功'); getList()
  })
}

onMounted(getList)
</script>
