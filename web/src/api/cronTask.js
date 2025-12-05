import service from '@/utils/request'

// @Router /cron-task/config/list [post]
export const configList = (data) => {
    return service({
        url: '/cron-task/config/list',
        method: 'post',
        data
    })
}

// @Router /cron-task/config/add [post]
export const configAdd = (data) => {
    return service({
        url: '/cron-task/config/add',
        method: 'post',
        data
    })
}

// @Router /cron-task/config/modify [post]
export const configModify = (data) => {
    return service({
        url: '/cron-task/config/modify',
        method: 'post',
        data
    })
}

// @Router /cron-task/log/list [post]
export const logList = (data) => {
    return service({
        url: '/cron-task/log/list',
        method: 'post',
        data
    })
}