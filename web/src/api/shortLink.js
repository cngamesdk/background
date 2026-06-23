import service from '@/utils/request'

// @Router /short-link/create [post]
export const shortLinkCreate = (data) => {
    return service({
        url: '/short-link/create',
        method: 'post',
        data
    })
}

// @Router /short-link/list [post]
export const shortLinkList = (data) => {
    return service({
        url: '/short-link/list',
        method: 'post',
        data
    })
}

// @Router /short-link/detail [post]
export const shortLinkDetail = (data) => {
    return service({
        url: '/short-link/detail',
        method: 'post',
        data
    })
}

// @Router /short-link/update [post]
export const shortLinkUpdate = (data) => {
    return service({
        url: '/short-link/update',
        method: 'post',
        data
    })
}

// @Router /short-link/delete [post]
export const shortLinkDelete = (data) => {
    return service({
        url: '/short-link/delete',
        method: 'post',
        data
    })
}

// @Router /short-link/click-log/list [post]
export const clickLogList = (data) => {
    return service({
        url: '/short-link/click-log/list',
        method: 'post',
        data
    })
}
