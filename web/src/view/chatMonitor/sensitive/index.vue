<template>
  <div>
    <el-tabs v-model="activeTab">
      <el-tab-pane label="敏感词库" name="words">
        <div class="gva-search-box">
          <el-form :inline="true">
            <el-form-item><el-button type="primary" @click="showAdd">添加</el-button></el-form-item>
            <el-form-item><el-button @click="showImport">批量导入</el-button></el-form-item>
            <el-form-item>
              <el-select v-model="filter.appId" clearable placeholder="全部游戏" style="width: 150px" @change="fetchWords">
                <el-option label="全局" value="" />
                <el-option v-for="g in gameList" :key="g.appId" :label="g.name" :value="g.appId" />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-select v-model="filter.category" clearable placeholder="全部分类" style="width: 120px" @change="fetchWords">
                <el-option label="政治" value="politics" /><el-option label="色情" value="porn" />
                <el-option label="广告" value="ad" /><el-option label="自定义" value="default" />
              </el-select>
            </el-form-item>
          </el-form>
        </div>
        <div class="gva-table-box">
          <el-table :data="wordList" border stripe>
            <el-table-column prop="word" label="敏感词" />
            <el-table-column prop="category" label="分类" width="100" />
            <el-table-column prop="level" label="等级" width="80">
              <template #default="{ row }">
                <el-tag :type="{ 1: 'info', 2: 'warning', 3: 'danger' }[row.level]" size="small">{{ { 1: '低', 2: '中', 3: '高' }[row.level] }}</el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="appId" label="范围" width="100"><template #default="{ row }">{{ row.appId || '全局' }}</template></el-table-column>
            <el-table-column prop="isRegex" label="正则" width="60"><template #default="{ row }">{{ row.isRegex ? '是' : '否' }}</template></el-table-column>
            <el-table-column label="操作" width="150">
              <template #default="{ row }">
                <el-button size="small" @click="handleEdit(row)">编辑</el-button>
                <el-popconfirm title="确认删除?" @confirm="handleDelete(row.ID)">
                  <template #reference><el-button size="small" type="danger">删除</el-button></template>
                </el-popconfirm>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination v-model:current-page="filter.page" :total="wordTotal" layout="total, prev, pager, next" style="margin-top: 16px" @current-change="fetchWords" />
        </div>
      </el-tab-pane>

      <el-tab-pane label="白名单" name="whitelist">
        <div class="gva-search-box">
          <el-form :inline="true">
            <el-form-item><el-input v-model="newWhite" placeholder="白名单词" style="width: 200px" /></el-form-item>
            <el-form-item><el-button type="primary" @click="addWhite">添加</el-button></el-form-item>
          </el-form>
        </div>
        <div class="gva-table-box">
          <el-table :data="whiteList" border stripe>
            <el-table-column prop="word" label="白名单词" />
            <el-table-column prop="appId" label="范围" width="100"><template #default="{ row }">{{ row.appId || '全局' }}</template></el-table-column>
            <el-table-column prop="CreatedAt" label="创建时间" width="180" />
          </el-table>
        </div>
      </el-tab-pane>
    </el-tabs>

    <el-dialog v-model="dialog.visible" :title="dialog.isEdit ? '编辑' : '添加'" width="450px">
      <el-form :model="dialog.form" label-width="80px">
        <el-form-item label="敏感词"><el-input v-model="dialog.form.word" /></el-form-item>
        <el-form-item label="分类">
          <el-select v-model="dialog.form.category" style="width: 100%">
            <el-option label="政治" value="politics" /><el-option label="色情" value="porn" />
            <el-option label="广告" value="ad" /><el-option label="自定义" value="default" />
          </el-select>
        </el-form-item>
        <el-form-item label="等级">
          <el-radio-group v-model="dialog.form.level"><el-radio :value="1">低</el-radio><el-radio :value="2">中</el-radio><el-radio :value="3">高</el-radio></el-radio-group>
        </el-form-item>
        <el-form-item label="正则"><el-switch v-model="dialog.form.isRegex" :active-value="1" :inactive-value="0" /></el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialog.visible = false">取消</el-button>
        <el-button type="primary" @click="submitWord">确认</el-button>
      </template>
    </el-dialog>

    <el-dialog v-model="importVisible" title="批量导入" width="450px">
      <el-input v-model="importText" type="textarea" :rows="8" placeholder="每行一个敏感词" />
      <template #footer>
        <el-button @click="importVisible = false">取消</el-button>
        <el-button type="primary" @click="submitImport">导入</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { getSensitiveList, createSensitive, updateSensitive, deleteSensitive, importSensitive, getWhitelistList, createWhitelist, getGameList } from '@/api/chatMonitor'
import { ElMessage } from 'element-plus'

defineOptions({ name: 'ChatMonitorSensitive' })

const activeTab = ref('words')
const gameList = ref([])
const wordList = ref([])
const whiteList = ref([])
const wordTotal = ref(0)
const newWhite = ref('')
const importVisible = ref(false)
const importText = ref('')
const filter = reactive({ appId: '', category: '', page: 1, pageSize: 20 })
const dialog = reactive({ visible: false, isEdit: false, form: {} })

const fetchWords = async () => {
  const res = await getSensitiveList(filter)
  if (res.code === 0) { wordList.value = res.data.list || []; wordTotal.value = res.data.total }
}
const fetchWhitelist = async () => {
  const res = await getWhitelistList({ page: 1, pageSize: 100 })
  if (res.code === 0) whiteList.value = res.data.list || []
}
const showAdd = () => { dialog.isEdit = false; dialog.form = { word: '', category: 'default', level: 2, isRegex: 0 }; dialog.visible = true }
const showImport = () => { importVisible.value = true }
const handleEdit = (row) => { dialog.isEdit = true; dialog.form = { ...row }; dialog.visible = true }
const submitWord = async () => {
  dialog.isEdit ? await updateSensitive(dialog.form) : await createSensitive(dialog.form)
  ElMessage.success('操作成功'); dialog.visible = false; fetchWords()
}
const handleDelete = async (id) => { await deleteSensitive({ id }); ElMessage.success('删除成功'); fetchWords() }
const submitImport = async () => {
  const words = importText.value.split('\n').map(w => w.trim()).filter(Boolean)
  if (!words.length) return
  await importSensitive({ words, category: 'default', level: 2 })
  ElMessage.success(`导入 ${words.length} 个词`); importVisible.value = false; importText.value = ''; fetchWords()
}
const addWhite = async () => {
  if (!newWhite.value) return
  await createWhitelist({ word: newWhite.value }); ElMessage.success('添加成功'); newWhite.value = ''; fetchWhitelist()
}

onMounted(async () => {
  const res = await getGameList({ page: 1, pageSize: 100 })
  if (res.code === 0) gameList.value = res.data.list || []
  fetchWords(); fetchWhitelist()
})
</script>
