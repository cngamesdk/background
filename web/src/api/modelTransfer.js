import service from '@/utils/request'

// ==================== Token管理 ====================

// 创建Token
export const createToken = (data) => {
  return service({
    url: '/model-transfer/token/create',
    method: 'post',
    data
  })
}

// 更新Token
export const updateToken = (data) => {
  return service({
    url: '/model-transfer/token/update',
    method: 'post',
    data
  })
}

// 删除Token
export const deleteToken = (data) => {
  return service({
    url: '/model-transfer/token/delete',
    method: 'post',
    data
  })
}

// Token列表
export const getTokenList = (data) => {
  return service({
    url: '/model-transfer/token/list',
    method: 'post',
    data
  })
}

// Token详情
export const getTokenDetail = (data) => {
  return service({
    url: '/model-transfer/token/detail',
    method: 'post',
    data
  })
}

// 重新生成Token
export const regenerateToken = (data) => {
  return service({
    url: '/model-transfer/token/regenerate',
    method: 'post',
    data
  })
}

// ==================== 报表查询 ====================

// 日报查询
export const getDailyReport = (data) => {
  return service({
    url: '/model-transfer/report/daily',
    method: 'post',
    data
  })
}

// Token使用详情
export const getTokenUsage = (data) => {
  return service({
    url: '/model-transfer/report/token-usage',
    method: 'post',
    data
  })
}

// 汇总报表
export const getSummaryReport = (data) => {
  return service({
    url: '/model-transfer/report/summary',
    method: 'post',
    data
  })
}
