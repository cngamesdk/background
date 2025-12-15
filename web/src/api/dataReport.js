import service from '@/utils/request'

// @Router /data-report/day-overview/list [post]
export const dayOverviewList = (data) => {
    return service({
        url: '/data-report/day-overview/list',
        method: 'post',
        data
    })
}