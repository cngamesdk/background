<template>
  <div class="model-transfer-container">
    <el-card shadow="hover">
      <template #header>
        <div class="card-header">
          <span class="header-title">AI模型中转管理</span>
        </div>
      </template>

      <!-- 统计卡片 -->
      <el-row :gutter="20" style="margin-bottom: 30px">
        <el-col :span="8">
          <el-card shadow="hover" class="stat-card-wrapper">
            <div class="stat-card">
              <div class="stat-icon" style="background: #409eff">
                <el-icon size="30"><key /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ stats.totalTokens }}</div>
                <div class="stat-label">总Token数</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover" class="stat-card-wrapper">
            <div class="stat-card">
              <div class="stat-icon" style="background: #67c23a">
                <el-icon size="30"><data-analysis /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ stats.totalRequests }}</div>
                <div class="stat-label">今日请求</div>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card shadow="hover" class="stat-card-wrapper">
            <div class="stat-card">
              <div class="stat-icon" style="background: #e6a23c">
                <el-icon size="30"><coin /></el-icon>
              </div>
              <div class="stat-content">
                <div class="stat-value">{{ stats.totalUsage }}</div>
                <div class="stat-label">今日使用量</div>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <!-- 快捷入口 -->
      <div class="quick-links">
        <h3>快捷入口</h3>
        <el-row :gutter="20">
          <el-col :span="8">
            <el-card shadow="hover" class="link-card" @click="goTo('/model-transfer/token')">
              <div class="link-content">
                <el-icon size="40" color="#409eff"><key /></el-icon>
                <div class="link-text">
                  <div class="link-title">Token管理</div>
                  <div class="link-desc">创建、编辑、查看Token</div>
                </div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card shadow="hover" class="link-card" @click="goTo('/model-transfer/report')">
              <div class="link-content">
                <el-icon size="40" color="#67c23a"><data-analysis /></el-icon>
                <div class="link-text">
                  <div class="link-title">使用报表</div>
                  <div class="link-desc">查看使用统计和报表</div>
                </div>
              </div>
            </el-card>
          </el-col>
          <el-col :span="8">
            <el-card shadow="hover" class="link-card" @click="goTo('/model-transfer/daily')">
              <div class="link-content">
                <el-icon size="40" color="#e6a23c"><pie-chart /></el-icon>
                <div class="link-text">
                  <div class="link-title">日报统计</div>
                  <div class="link-desc">查看日报和趋势分析</div>
                </div>
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// 统计数据
const stats = ref({
  totalTokens: 5,
  totalRequests: 120,
  totalUsage: 15000
})

// 跳转
const goTo = (path) => {
  router.push(path)
}
</script>

<style scoped>
.model-transfer-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.header-title {
  font-size: 18px;
  font-weight: bold;
}

.stat-card-wrapper {
  cursor: pointer;
  transition: all 0.3s;
}

.stat-card-wrapper:hover {
  transform: translateY(-5px);
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 20px;
}

.stat-icon {
  width: 60px;
  height: 60px;
  border-radius: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}

.stat-content {
  flex: 1;
}

.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #303133;
}

.stat-label {
  font-size: 14px;
  color: #909399;
  margin-top: 5px;
}

.quick-links {
  margin-top: 30px;
}

.quick-links h3 {
  margin-bottom: 20px;
  color: #303133;
}

.link-card {
  cursor: pointer;
  transition: all 0.3s;
}

.link-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.link-content {
  display: flex;
  align-items: center;
  gap: 20px;
  padding: 10px;
}

.link-text {
  flex: 1;
}

.link-title {
  font-size: 16px;
  font-weight: bold;
  color: #303133;
  margin-bottom: 5px;
}

.link-desc {
  font-size: 13px;
  color: #909399;
}
</style>
