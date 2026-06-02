import service from '@/utils/request'

// Product
export const createProduct = (data) => service({ url: '/liveChat/product/create', method: 'post', data })
export const updateProduct = (data) => service({ url: '/liveChat/product/update', method: 'put', data })
export const deleteProduct = (data) => service({ url: '/liveChat/product/delete', method: 'delete', data })
export const getProductList = (params) => service({ url: '/liveChat/product/list', method: 'get', params })
export const getProductDetail = (id) => service({ url: `/liveChat/product/${id}`, method: 'get' })

// FAQ
export const createFaq = (data) => service({ url: '/liveChat/faq/create', method: 'post', data })
export const updateFaq = (data) => service({ url: '/liveChat/faq/update', method: 'put', data })
export const deleteFaq = (data) => service({ url: '/liveChat/faq/delete', method: 'delete', data })
export const getFaqList = (params) => service({ url: '/liveChat/faq/list', method: 'get', params })
export const importFaq = (data) => service({ url: '/liveChat/faq/import', method: 'post', data })
export const getFaqCategories = (params) => service({ url: '/liveChat/faq/categories', method: 'get', params })

// Agent
export const agentOnline = (data) => service({ url: '/liveChat/agent/online', method: 'post', data })
export const agentOffline = (data) => service({ url: '/liveChat/agent/offline', method: 'post', data })
export const updateAgent = (data) => service({ url: '/liveChat/agent/update', method: 'put', data })
export const getAgentList = (params) => service({ url: '/liveChat/agent/list', method: 'get', params })
export const getAgentStatus = (params) => service({ url: '/liveChat/agent/status', method: 'get', params })

// Chat
export const getChatSessions = (params) => service({ url: '/liveChat/chat/sessions', method: 'get', params })
export const getSessionDetail = (id) => service({ url: `/liveChat/chat/session/${id}`, method: 'get' })
export const assignSession = (data) => service({ url: '/liveChat/chat/assign', method: 'post', data })
export const agentReply = (data) => service({ url: '/liveChat/chat/reply', method: 'post', data })
export const closeSession = (data) => service({ url: '/liveChat/chat/close', method: 'post', data })

// Report
export const getReportOverview = (params) => service({ url: '/liveChat/report/overview', method: 'get', params })
export const getReportTrend = (params) => service({ url: '/liveChat/report/trend', method: 'get', params })
