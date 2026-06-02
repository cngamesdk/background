<template>
  <div>
    <el-card>
      <el-form :inline="true" :model="searchForm" size="small">
        <el-form-item label="产品">
          <el-select v-model="searchForm.product_id" filterable clearable placeholder="选择产品" style="width:200px">
            <el-option v-for="p in productList" :key="p.ID" :label="p.name" :value="p.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="分类">
          <el-select v-model="searchForm.category" clearable filterable placeholder="选择分类">
            <el-option v-for="c in categories" :key="c" :label="c" :value="c" />
          </el-select>
        </el-form-item>
        <el-form-item label="关键词"><el-input v-model="searchForm.keyword" /></el-form-item>
        <el-form-item><el-button type="primary" @click="getList">查询</el-button><el-button @click="openDialog()">新增FAQ</el-button></el-form-item>
      </el-form>
    </el-card>
    <el-card style="margin-top:10px">
      <el-table :data="tableData" border stripe>
        <el-table-column prop="ID" label="ID" width="60" />
        <el-table-column label="产品名称" width="120">
          <template #default="{row}">{{ getProductName(row.product_id) }}</template>
        </el-table-column>
        <el-table-column prop="category" label="分类" width="100" />
        <el-table-column prop="question" label="问题" min-width="200" show-overflow-tooltip />
        <el-table-column prop="answer" label="答案" min-width="200" show-overflow-tooltip />
        <el-table-column prop="match_count" label="匹配次数" width="100" />
        <el-table-column prop="priority" label="优先级" width="80" />
        <el-table-column label="操作" width="160">
          <template #default="{row}"><el-button size="small" @click="openDialog(row)">编辑</el-button><el-button size="small" type="danger" @click="handleDelete(row.ID)">删除</el-button></template>
        </el-table-column>
      </el-table>
      <el-pagination v-model:current-page="page" :total="total" :page-size="pageSize" layout="total,prev,pager,next" @current-change="getList" style="margin-top:10px" />
    </el-card>

    <el-dialog :title="dialogTitle" v-model="dialogVisible" width="700px">
      <el-form :model="form" label-width="80px">
        <el-form-item label="产品">
          <el-select v-model="form.product_id" filterable placeholder="选择产品" style="width:100%">
            <el-option v-for="p in productList" :key="p.ID" :label="p.name" :value="p.ID" />
          </el-select>
        </el-form-item>
        <el-form-item label="分类"><el-input v-model="form.category" /></el-form-item>
        <el-form-item label="问题"><el-input v-model="form.question" /></el-form-item>
        <el-form-item label="答案"><el-input v-model="form.answer" type="textarea" :rows="4" /></el-form-item>
        <el-form-item label="关键词"><el-input v-model="form.keywords" placeholder="逗号分隔" /></el-form-item>
        <el-form-item label="优先级"><el-input-number v-model="form.priority" :min="0" /></el-form-item>
        <el-form-item label="状态"><el-select v-model="form.status"><el-option label="正常" value="normal" /><el-option label="禁用" value="disabled" /></el-select></el-form-item>
      </el-form>
      <template #footer><el-button @click="dialogVisible=false">取消</el-button><el-button type="primary" @click="handleSave">保存</el-button></template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { getFaqList, createFaq, updateFaq, deleteFaq, getFaqCategories } from '@/api/liveChat'
import { getProductList } from '@/api/liveChat'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({ name: 'LiveChatFaq' })

const searchForm = reactive({ product_id: '', category: '', keyword: '', status: '' })
const tableData = ref([])
const productList = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const dialogVisible = ref(false)
const categories = ref([])
const form = reactive({ ID: 0, product_id: '', category: '', question: '', answer: '', keywords: '', priority: 0, status: 'normal' })
const dialogTitle = computed(() => form.ID ? '编辑FAQ' : '新增FAQ')

const loadProducts = async () => {
  const res = await getProductList({ page: 1, pageSize: 100 })
  if (res.code === 0) productList.value = res.data?.list || []
}

const getProductName = (pid) => {
  const p = productList.value.find(i => i.ID === pid)
  return p ? p.name : pid
}

const getList = async () => {
  const params = { ...searchForm, page: page.value, pageSize: pageSize.value }
  if (searchForm.product_id) params.product_id = searchForm.product_id
  const res = await getFaqList(params)
  if (res.code === 0) { tableData.value = res.data?.list || []; total.value = res.data?.total || 0 }
}

const getCategories = async () => {
  const res = await getFaqCategories({ product_id: searchForm.product_id || '' })
  if (res.code === 0) categories.value = res.data || []
}

const openDialog = (row) => {
  if (row) { Object.assign(form, row) } else { Object.assign(form, { ID: 0, product_id: '', category: '', question: '', answer: '', keywords: '', priority: 0, status: 'normal' }) }
  dialogVisible.value = true
}

const handleSave = async () => {
  const fn = form.ID ? updateFaq : createFaq
  const res = await fn(form)
  if (res.code === 0) { ElMessage.success('保存成功'); dialogVisible.value = false; getList() }
}

const handleDelete = (id) => {
  ElMessageBox.confirm('确定删除吗？', '提示', { type: 'warning' }).then(async () => {
    await deleteFaq({ ID: id }); ElMessage.success('删除成功'); getList()
  })
}

onMounted(() => { loadProducts(); getList(); getCategories() })
</script>
