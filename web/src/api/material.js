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

// @Router /material/material/list [post]
export const materialList = (data) => {
    return service({
        url: '/material/material/list',
        method: 'post',
        data
    })
}

// @Router /material/material/add [post]
export const materialAdd= (data) => {
    return service({
        url: '/material/material/add',
        method: 'post',
        data
    })
}

// @Router /material/material/modify [post]
export const materialModify= (data) => {
    return service({
        url: '/material/material/modify',
        method: 'post',
        data
    })
}

// @Router /material/material-file/list [post]
export const materialFileList = (data) => {
    return service({
        url: '/material/material-file/list',
        method: 'post',
        data
    })
}

// @Router /material/material-file/add [post]
export const materialFileAdd = (data) => {
    return service({
        url: '/material/material-file/add',
        method: 'post',
        data
    })
}

// @Router /material/material-file/modify [post]
export const materialFileModify = (data) => {
    return service({
        url: '/material/material-file/modify',
        method: 'post',
        data
    })
}