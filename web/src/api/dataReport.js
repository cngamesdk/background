import service from '@/utils/request'

// @Router /data-report/day-overview/list [post]
export const dayOverviewList = (data) => {
    return service({
        url: '/data-report/day-overview/list',
        method: 'post',
        data
    })
}

// @Router /data-report/retention-status/list [post]
export const retentionStatusList = (data) => {
    return service({
        url: '/data-report/retention-status/list',
        method: 'post',
        data
    })
}

// @Router /data-report/payment-status/list [post]
export const paymentStatusList = (data) => {
    return service({
        url: '/data-report/payment-status/list',
        method: 'post',
        data
    })
}