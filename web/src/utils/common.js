const platformIdFilterKey = 'platform_id'
const rootGameIdFilterKey = 'root_game_id'
const mainGameIdFilterKey = 'main_game_id'
const gameIdFilterKey = 'game_id'
const agentIdFilterKey = 'agent_id'
const siteIdFilterKey = 'site_id'
const ipFilterKey = 'ip'
const oaidFilterKey = 'oaid'
const imeiFilterKey = 'imei'
const idfvFilterKey = 'idfv'
const userIdFilterKey = 'user_id'
const roleIdFilterKey = 'role_id'
const regTimeFilterKey = 'reg_time'
const totalPayCountFilterKey = 'total_pay_count'
const totalPayAmountFilterKey = 'total_pay_amount'

export {
    platformIdFilterKey,
    rootGameIdFilterKey,
    mainGameIdFilterKey,
    gameIdFilterKey,
    agentIdFilterKey,
    siteIdFilterKey,
    ipFilterKey,
    oaidFilterKey,
    imeiFilterKey,
    idfvFilterKey,
    userIdFilterKey,
    roleIdFilterKey,
    regTimeFilterKey,
    totalPayCountFilterKey,
    totalPayAmountFilterKey,
}

/**
 * 获取所有维度筛选
 * @returns {({name: string, value: string}|{name: string, value: string}|{name: string, value: string}|{name: string, value: string}|{name: string, value: string})[]}
 */
export const dimensionFilter = () => {
    return [
        {name: '平台', value: platformIdFilterKey},
        {name: '根游戏', value: rootGameIdFilterKey},
        {name: '主游戏', value: mainGameIdFilterKey},
        {name: '子游戏', value: gameIdFilterKey},
        {name: '渠道ID', value: agentIdFilterKey},
        {name: '广告位ID', value: siteIdFilterKey},
    ]
}

/**
 * 位置筛选
 * @returns {[{name: string, value: string}]}
 */
export const locationFilter = () => {
    return [
        {name: 'Ip', value: ipFilterKey},
    ]
}

/**
 * 设备筛选
 */
export const deviceFilter = () => {
    return [
        {name: 'Oaid', value: oaidFilterKey},
        {name: 'Imei/Idfa', value: imeiFilterKey},
        {name: 'Idfv', value: idfvFilterKey},
    ]
}

/**
 * 用户筛选
 */
export const userFilter = () => {
    return [
        {name: '用户ID', value: userIdFilterKey},
        {name: '角色ID', value: roleIdFilterKey},
        {name: '注册时间', value: regTimeFilterKey},
    ]
}

/**
 * 指标筛选
 * @returns {[{name: string, value: string}, {name: string, value: string}]}
 */
export const indicatorFilter = () => {
    return [
        {name: '累充次数', value: totalPayCountFilterKey},
        {name: '累充金额', value: totalPayAmountFilterKey},
    ]
}

/**
 * 所有筛选
 * @returns {({name: string, value: string}|{name: string, value: string}|{name: string, value: string}|{name: string, value: string}|{name: string, value: string})[]}
 */
export const allFilter = () => {
    return [...dimensionFilter(), ...locationFilter(), ...deviceFilter(), ...userFilter(), ...indicatorFilter()]
}