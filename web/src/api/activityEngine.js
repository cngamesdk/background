import service from '@/utils/request'

// 活动管理
export const activityList = (data) => {
  return service({ url: '/activity_engine/activity/list', method: 'post', data })
}

export const activityAdd = (data) => {
  return service({ url: '/activity_engine/activity/add', method: 'post', data })
}

export const activityModify = (data) => {
  return service({ url: '/activity_engine/activity/modify', method: 'post', data })
}

export const activityDetail = (data) => {
  return service({ url: '/activity_engine/activity/detail', method: 'post', data })
}

export const activityPublish = (data) => {
  return service({ url: '/activity_engine/activity/publish', method: 'post', data })
}

export const activityOffline = (data) => {
  return service({ url: '/activity_engine/activity/offline', method: 'post', data })
}

// 模板管理
export const templateList = (data) => {
  return service({ url: '/activity_engine/template/list', method: 'post', data })
}

export const templateAdd = (data) => {
  return service({ url: '/activity_engine/template/add', method: 'post', data })
}

export const templateClone = (data) => {
  return service({ url: '/activity_engine/template/clone', method: 'post', data })
}

// 奖励道具
export const rewardItemSearch = (data) => {
  return service({ url: '/activity_engine/reward_item/search', method: 'post', data })
}

// 灰度管理
export const grayscaleUpdate = (data) => {
  return service({ url: '/activity_engine/grayscale/update', method: 'post', data })
}

// 沙箱测试
export const sandboxSimulate = (data) => {
  return service({ url: '/activity_engine/sandbox/simulate', method: 'post', data })
}
