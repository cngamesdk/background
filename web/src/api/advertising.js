import service from '@/utils/request'

// @Router /advertising/media/list [post]
export const advertisingMediaList = (data) => {
    return service({
        url: '/advertising/media/list',
        method: 'post',
        data
    })
}

// @Router /advertising/media/add [post]
export const advertisingMediaAdd = (data) => {
    return service({
        url: '/advertising/media/add',
        method: 'post',
        data
    })
}

// @Router /advertising/media/modify [post]
export const advertisingMediaModify = (data) => {
    return service({
        url: '/advertising/media/modify',
        method: 'post',
        data
    })
}

// @Router /advertising/channel-group/list [post]
export const channelGroupList = (data) => {
    return service({
        url: '/advertising/channel-group/list',
        method: 'post',
        data
    })
}

// @Router /advertising/channel-group/add [post]
export const channelGroupAdd = (data) => {
    return service({
        url: '/advertising/channel-group/add',
        method: 'post',
        data
    })
}

// @Router /advertising/channel-group/modify [post]
export const channelGroupModify = (data) => {
    return service({
        url: '/advertising/channel-group/modify',
        method: 'post',
        data
    })
}

// @Router /advertising/agent/list [post]
export const agentList = (data) => {
    return service({
        url: '/advertising/agent/list',
        method: 'post',
        data
    })
}

// @Router /advertising/agent/add [post]
export const agentAdd = (data) => {
    return service({
        url: '/advertising/agent/add',
        method: 'post',
        data
    })
}

// @Router /advertising/agent/modify [post]
export const agentModify = (data) => {
    return service({
        url: '/advertising/agent/modify',
        method: 'post',
        data
    })
}

// @Router /advertising/site/list [post]
export const siteList = (data) => {
    return service({
        url: '/advertising/site/list',
        method: 'post',
        data
    })
}

// @Router /advertising/site/add [post]
export const siteAdd = (data) => {
    return service({
        url: '/advertising/site/add',
        method: 'post',
        data
    })
}

// @Router /advertising/site/modify [post]
export const siteModify = (data) => {
    return service({
        url: '/advertising/site/modify',
        method: 'post',
        data
    })
}