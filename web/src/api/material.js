import service from '@/utils/request'

// @Router /material/theme/list [post]
export const materialThemeList = (data) => {
    return service({
        url: '/material/theme/list',
        method: 'post',
        data
    })
}

// @Router /material/theme/add [post]
export const materialThemeAdd = (data) => {
    return service({
        url: '/material/theme/add',
        method: 'post',
        data
    })
}

// @Router /material/theme/modify [post]
export const materialThemeModify = (data) => {
    return service({
        url: '/material/theme/modify',
        method: 'post',
        data
    })
}