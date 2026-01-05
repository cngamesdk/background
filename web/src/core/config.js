/**
 * 网站配置文件
 */
import packageInfo from '../../package.json'

const greenText = (text) => `\x1b[32m${text}\x1b[0m`

export const config = {
  appName: '梅花互娱',
  showViteLogo: true,
  KeepAliveTabs: true,
  logs: []
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    console.log(
      greenText(
        `> 欢迎使用「`+ config.appName +`」，开源地址：https://github.com/cngamesdk`
      )
    )
  }
}

export default config
