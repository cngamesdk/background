// 加载网站配置文件夹
import { register } from './global'
import packageInfo from '../../package.json'

export default {
  install: (app) => {
    register(app)
    console.log(`
       欢迎使用 「${packageInfo.name}」
       当前版本:v${packageInfo.version}
       联系方式:微信：doudoualvin QQ群：827883750
       项目地址：https://github.com/cngamesdk
       默认自动化文档地址:http://127.0.0.1:${import.meta.env.VITE_SERVER_PORT}/swagger/index.html
       默认前端文件运行地址:http://127.0.0.1:${import.meta.env.VITE_CLI_PORT}
    `)
  }
}
