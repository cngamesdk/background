<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true">
        <el-form-item>
          <el-button type="primary" @click="showCreate">新增游戏</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-table :data="tableData" border stripe>
        <el-table-column prop="appId" label="AppID" width="180" />
        <el-table-column prop="name" label="游戏名称" width="200" />
        <el-table-column prop="callbackUrl" label="回调地址" show-overflow-tooltip />
        <el-table-column prop="status" label="状态" width="100">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'danger'">{{ row.status === 1 ? '启用' : '禁用' }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column prop="CreatedAt" label="创建时间" width="180" />
        <el-table-column label="操作" width="200" fixed="right">
          <template #default="{ row }">
            <el-button size="small" @click="showEdit(row)">编辑</el-button>
            <el-popconfirm title="确认删除?" @confirm="handleDelete(row.ID)">
              <template #reference><el-button size="small" type="danger">删除</el-button></template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination v-model:current-page="page" v-model:page-size="pageSize" :total="total"
        layout="total, prev, pager, next" style="margin-top: 16px" @current-change="getList" />
    </div>

    <el-dialog v-model="dialog.visible" :title="dialog.isEdit ? '编辑游戏' : '新增游戏'" width="500px">
      <el-form :model="dialog.form" label-width="100px">
        <el-form-item label="AppID" v-if="!dialog.isEdit">
          <el-input v-model="dialog.form.appId" placeholder="留空自动生成" />
        </el-form-item>
        <el-form-item label="AppSecret" v-if="!dialog.isEdit">
          <el-input v-model="dialog.form.appSecret" placeholder="签名密钥" />
        </el-form-item>
        <el-form-item label="游戏名称"><el-input v-model="dialog.form.name" /></el-form-item>
        <el-form-item label="回调地址"><el-input v-model="dialog.form.callbackUrl" /></el-form-item>
        <el-form-item label="状态" v-if="dialog.isEdit">
          <el-switch v-model="dialog.form.status" :active-value="1" :inactive-value="0" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialog.visible = false">取消</el-button>
        <el-button type="primary" @click="handleSubmit">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getGameList, createGame, updateGame, deleteGame } from '@/api/chatMonitor'
import { ElMessage } from 'element-plus'

defineOptions({ name: 'ChatMonitorGame' })

const tableData = ref([])
const page = ref(1)
const pageSize = ref(10)
const total = ref(0)
const dialog = reactive({ visible: false, isEdit: false, form: {} })

const getList = async () => {
  const res = await getGameList({ page: page.value, pageSize: pageSize.value })
  if (res.code === 0) { tableData.value = res.data.list || []; total.value = res.data.total }
}
const showCreate = () => { dialog.isEdit = false; dialog.form = { appId: '', appSecret: '', name: '', callbackUrl: '' }; dialog.visible = true }
const showEdit = (row) => { dialog.isEdit = true; dialog.form = { ...row }; dialog.visible = true }
const handleSubmit = async () => {
  dialog.isEdit ? await updateGame(dialog.form) : await createGame(dialog.form)
  ElMessage.success('操作成功'); dialog.visible = false; getList()
}
const handleDelete = async (id) => { await deleteGame({ id }); ElMessage.success('删除成功'); getList() }

onMounted(getList)
</script>
