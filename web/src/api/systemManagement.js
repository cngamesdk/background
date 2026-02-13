import service from '@/utils/request'
const baseSearchUrl = '/system_management/search/search'
// @Router /system_management/search/search [post]
export const searchGameType = (data) => {
    data = {dim_type: 'game-type'}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchGameOs = (data) => {
    data = {dim_type: 'game-os'}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchGameStatus = (data) => {
    data = {dim_type: 'game-status'}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchGameCooperationModel = (data) => {
    data = {dim_type: 'game-cooperation-model'}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchCompany = (data) => {
    data = {dim_type: 'company', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchPlatform = (data) => {
    data = {dim_type: 'platform', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchRootGame = (data) => {
    data = {dim_type: 'root-game', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchMainGame = (data) => {
    data = {dim_type: 'main-game', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchSubGame = (data) => {
    data = {dim_type: 'sub-game', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchProductCommonConfig = (data) => {
    data = {dim_type: 'product-common-config', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchPayType = (data) => {
    data = {dim_type: 'pay-type', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchPayStatus = (data) => {
    data = {dim_type: 'pay-status', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchPayChannel = (data) => {
    data = {dim_type: 'pay-channel', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchPublishingChannel = (data) => {
    data = {dim_type: 'publishing-channel', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchAgent = (data) => {
    data = {dim_type: 'agent', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchSite = (data) => {
    data = {dim_type: 'site', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// @Router /system_management/search/search [post]
export const searchCommonMedia = (data) => {
    data = {dim_type: 'common-media', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

export const searchMedia = (data) => {
    data = {dim_type: 'media', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

export const searchChannelGroup = (data) => {
    data = {dim_type: 'channel-group', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

export const searchSettlementType = (data) => {
    data = {dim_type: 'settlement-type', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// getStatisticalCaliber 获取口径
export const getStatisticalCaliber = () => {
    return [
        {key:'root-game-back-30', value: '30天回流'}
    ]
}

// getStatisticalCaliber 获取聚合时间
export const getAggregationTime = () => {
    return [
        {key:'day', value: '按日'},
        {key:'month', value: '按月'},
        {key:'aggregation', value: '聚合'}
    ]
}

// getMaterialTheme 获取素材主题
export const getMaterialTheme = (data) => {
    data = {dim_type: 'material-theme', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

// getAdvertisingDeveloper 获取开发者
export const getAdvertisingDeveloper = (data) => {
    data = {dim_type: 'advertising-developer', ...data}
    return service({
        url: baseSearchUrl,
        method: 'post',
        data
    })
}

