import service from '@/utils/request'

// ========== 游戏管理 ==========
export const getGameList = (data) => {
  return service({ url: '/chatMonitor/game/list', method: 'get', params: data })
}
export const createGame = (data) => {
  return service({ url: '/chatMonitor/game/create', method: 'post', data })
}
export const updateGame = (data) => {
  return service({ url: '/chatMonitor/game/update', method: 'put', data })
}
export const deleteGame = (data) => {
  return service({ url: '/chatMonitor/game/delete', method: 'delete', data })
}

// ========== 聊天记录 ==========
export const getChatHistory = (data) => {
  return service({ url: '/chatMonitor/chat/history', method: 'get', params: data })
}

// ========== 敏感词 ==========
export const getSensitiveList = (data) => {
  return service({ url: '/chatMonitor/sensitive/list', method: 'get', params: data })
}
export const createSensitive = (data) => {
  return service({ url: '/chatMonitor/sensitive/create', method: 'post', data })
}
export const importSensitive = (data) => {
  return service({ url: '/chatMonitor/sensitive/import', method: 'post', data })
}
export const updateSensitive = (data) => {
  return service({ url: '/chatMonitor/sensitive/update', method: 'put', data })
}
export const deleteSensitive = (data) => {
  return service({ url: '/chatMonitor/sensitive/delete', method: 'delete', data })
}

// ========== 白名单 ==========
export const getWhitelistList = (data) => {
  return service({ url: '/chatMonitor/whitelist/list', method: 'get', params: data })
}
export const createWhitelist = (data) => {
  return service({ url: '/chatMonitor/whitelist/create', method: 'post', data })
}

// ========== 封禁管理 ==========
export const createBan = (data) => {
  return service({ url: '/chatMonitor/ban/create', method: 'post', data })
}
export const revokeBan = (data) => {
  return service({ url: '/chatMonitor/ban/revoke', method: 'put', data })
}
export const getBanList = (data) => {
  return service({ url: '/chatMonitor/ban/list', method: 'get', params: data })
}

// ========== 数据报表 ==========
export const getStatsOverview = (data) => {
  return service({ url: '/chatMonitor/stats/overview', method: 'get', params: data })
}
export const getStatsTrend = (data) => {
  return service({ url: '/chatMonitor/stats/trend', method: 'get', params: data })
}
export const getViolators = (data) => {
  return service({ url: '/chatMonitor/stats/violators', method: 'get', params: data })
}
