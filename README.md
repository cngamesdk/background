## background
后台项目

## 介绍

该项目为发行后台，涉及两个目录：
- server：后端接口目录；
- web：前端接口目录；

<div align=center>
<img src="https://img.shields.io/badge/golang-1.22-blue"/>
<img src="https://img.shields.io/badge/gin-1.9.1-lightBlue"/>
<img src="https://img.shields.io/badge/vue-3.3.4-brightgreen"/>
<img src="https://img.shields.io/badge/element--plus-2.3.8-green"/>
<img src="https://img.shields.io/badge/gorm-1.25.2-red"/>
</div>

## ✨一分钟生成前后端基础代码

<table>
  <tr>
    <td width="250">
	  <p>⭐️ <a href="https://www.bilibili.com/video/BV1B3htzqEf1/?spm_id_from=333.1387.homepage.video_card.click" target="__blank"> 高度适配AI编辑器的MCP </a></p>
      <p>📄 创建基础模板</p>
      <p>🤖 AI生成结构</p>
      <p>⏰ 生成代码</p>
      <p>🏷️ 分配权限</p>
      <p>🎉 基础CURD生成完成</p>   
    </td>
    <td>
      <video src="https://private-user-images.githubusercontent.com/165128580/384700666-4d039215-af29-4f86-bb4f-60dbab38f58e.mp4?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MzEyNTIxNDYsIm5iZiI6MTczMTI1MTg0NiwicGF0aCI6Ii8xNjUxMjg1ODAvMzg0NzAwNjY2LTRkMDM5MjE1LWFmMjktNGY4Ni1iYjRmLTYwZGJhYjM4ZjU4ZS5tcDQ_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjQxMTEwJTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI0MTExMFQxNTE3MjZaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT00NjJkMDcwZjJkMjAyMmU1N2I2MzQxY2RhODFlNzgzNGRiMDFhMmY2NTYyM2ZmODdhNDVmMWE1NzlhMDdlOTI5JlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCJ9.ZJbswpLzF2RHjemcGirKOP0L1fvpl3FUqIiQ_-yjeUo" data-canonical-src="https://private-user-images.githubusercontent.com/165128580/384700666-4d039215-af29-4f86-bb4f-60dbab38f58e.mp4?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MzEyNTIxNDYsIm5iZiI6MTczMTI1MTg0NiwicGF0aCI6Ii8xNjUxMjg1ODAvMzg0NzAwNjY2LTRkMDM5MjE1LWFmMjktNGY4Ni1iYjRmLTYwZGJhYjM4ZjU4ZS5tcDQ_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjQxMTEwJTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI0MTExMFQxNTE3MjZaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT00NjJkMDcwZjJkMjAyMmU1N2I2MzQxY2RhODFlNzgzNGRiMDFhMmY2NTYyM2ZmODdhNDVmMWE1NzlhMDdlOTI5JlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCJ9.ZJbswpLzF2RHjemcGirKOP0L1fvpl3FUqIiQ_-yjeUo" controls="controls" muted="muted" class="d-block rounded-bottom-2 border-top width-fit" style="max-height:640px; min-height: 200px">
</video>
    </td>
  </tr>
</table>

## 1. 基本介绍

### 1.1 项目介绍

> 本项目是一个基于 [vue](https://vuejs.org) 和 [gin](https://gin-gonic.com) 开发的全栈前后端分离的开发基础平台，集成jwt鉴权，动态路由，动态菜单，casbin鉴权，表单生成器，代码生成器等功能，提供多种示例文件，让您把更多时间专注在业务开发上。

## 2. 使用说明

```
- node版本 > v18.16.0
- golang版本 >= v1.22
- IDE推荐：Goland
```

```bash

# 克隆项目
git clone https://github.com/cngamesdk/background.git

```

### 2.1 server项目

使用 `Goland` 等编辑工具，打开server目录，**不可以**打开 backgroud 根目录


```bash

# 进入server文件夹
cd server

# 使用 go mod 并安装go依赖包
go generate

# 运行
go run . 

```

### 2.2 web项目

```bash
# 进入web文件夹
cd web

# 安装依赖
npm install

# 启动web项目
npm run serve
```

### 2.3 swagger自动化API文档

#### 2.3.1 安装 swagger

``` shell
go install github.com/swaggo/swag/cmd/swag@latest
```

#### 2.3.2 生成API文档

```` shell
cd server
swag init
````

> 执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后, 在浏览器输入 [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) 即可查看swagger文档

### 2.4 VSCode工作区

#### 2.4.1 开发

使用`VSCode`打开根目录下的工作区文件`gin-vue-admin.code-workspace`，在边栏可以看到三个虚拟目录：`backend`、`frontend`、`root`。

#### 2.4.2 运行/调试

在运行和调试中也可以看到三个task：`Backend`、`Frontend`、`Both (Backend & Frontend)`。运行`Both (Backend & Frontend)`可以同时启动前后端项目。

#### 2.4.3 settings

在工作区配置文件中有`go.toolsEnvVars`字段，是用于`VSCode`自身的go工具环境变量。此外在多go版本的系统中，可以通过`gopath`、`go.goroot`指定运行版本。

```json
    "go.gopath": null,
    "go.goroot": null,
```

## 3. 技术选型

- 前端：用基于 [Vue](https://vuejs.org) 的 [Element](https://github.com/ElemeFE/element) 构建基础页面。
- 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
- 数据库：采用`MySql` > (5.7) 版本 数据库引擎 InnoDB，使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用 [fsnotify](https://github.com/fsnotify/fsnotify) 和 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。

## 4. 项目架构

### 4.1 系统架构图

![系统架构图](http://qmplusimg.henrongyi.top/gva/gin-vue-admin.png)

### 4.2 前端详细设计图 

![前端详细设计图](http://qmplusimg.henrongyi.top/naotu.png)

### 4.3 目录结构

```
    ├── server
        ├── api             (api层)
        │   └── v1          (v1版本接口)
        ├── config          (配置包)
        ├── core            (核心文件)
        ├── docs            (swagger文档目录)
        ├── global          (全局对象)                    
        ├── initialize      (初始化)                        
        │   └── internal    (初始化内部函数)                            
        ├── middleware      (中间件层)                        
        ├── model           (模型层)                    
        │   ├── request     (入参结构体)                        
        │   └── response    (出参结构体)                            
        ├── packfile        (静态文件打包)                        
        ├── resource        (静态资源文件夹)                        
        │   ├── excel       (excel导入导出默认路径)                        
        │   ├── page        (表单生成器)                        
        │   └── template    (模板)                            
        ├── router          (路由层)                    
        ├── service         (service层)                    
        ├── source          (source层)                    
        └── utils           (工具包)                    
            ├── timer       (定时器接口封装)                        
            └── upload      (oss接口封装)                        
    
            web
        ├── babel.config.js
        ├── Dockerfile
        ├── favicon.ico
        ├── index.html                 -- 主页面
        ├── limit.js                   -- 助手代码
        ├── package.json               -- 包管理器代码
        ├── src                        -- 源代码
        │   ├── api                    -- api 组
        │   ├── App.vue                -- 主页面
        │   ├── assets                 -- 静态资源
        │   ├── components             -- 全局组件
        │   ├── core                   -- gva 组件包
        │   │   ├── config.js          -- gva网站配置文件
        │   │   ├── gin-vue-admin.js   -- 注册欢迎文件
        │   │   └── global.js          -- 统一导入文件
        │   ├── directive              -- v-auth 注册文件
        │   ├── main.js                -- 主文件
        │   ├── permission.js          -- 路由中间件
        │   ├── pinia                  -- pinia 状态管理器，取代vuex
        │   │   ├── index.js           -- 入口文件
        │   │   └── modules            -- modules
        │   │       ├── dictionary.js
        │   │       ├── router.js
        │   │       └── user.js
        │   ├── router                 -- 路由声明文件
        │   │   └── index.js
        │   ├── style                  -- 全局样式
        │   │   ├── base.scss
        │   │   ├── basics.scss
        │   │   ├── element_visiable.scss  -- 此处可以全局覆盖 element-plus 样式
        │   │   ├── iconfont.css           -- 顶部几个icon的样式文件
        │   │   ├── main.scss
        │   │   ├── mobile.scss
        │   │   └── newLogin.scss
        │   ├── utils                  -- 方法包库
        │   │   ├── asyncRouter.js     -- 动态路由相关
        │   │   ├── btnAuth.js         -- 动态权限按钮相关
        │   │   ├── bus.js             -- 全局mitt声明文件
        │   │   ├── date.js            -- 日期相关
        │   │   ├── dictionary.js      -- 获取字典方法 
        │   │   ├── downloadImg.js     -- 下载图片方法
        │   │   ├── format.js          -- 格式整理相关
        │   │   ├── image.js           -- 图片相关方法
        │   │   ├── page.js            -- 设置页面标题
        │   │   ├── request.js         -- 请求
        │   │   └── stringFun.js       -- 字符串文件
        |   ├── view -- 主要view代码
        |   |   ├── about -- 关于我们
        |   |   ├── dashboard -- 面板
        |   |   ├── error -- 错误
        |   |   ├── example --上传案例
        |   |   ├── iconList -- icon列表
        |   |   ├── init -- 初始化数据  
        |   |   |   ├── index -- 新版本
        |   |   |   ├── init -- 旧版本
        |   |   ├── layout  --  layout约束页面 
        |   |   |   ├── aside 
        |   |   |   ├── bottomInfo     -- bottomInfo
        |   |   |   ├── screenfull     -- 全屏设置
        |   |   |   ├── setting        -- 系统设置
        |   |   |   └── index.vue      -- base 约束
        |   |   ├── login              --登录 
        |   |   ├── person             --个人中心 
        |   |   ├── superAdmin         -- 超级管理员操作
        |   |   ├── system             -- 系统检测页面
        |   |   ├── systemTools        -- 系统配置相关页面
        |   |   └── routerHolder.vue   -- page 入口页面 
        ├── vite.config.js             -- vite 配置文件
        └── yarn.lock

```

## 5. 主要功能

- 权限管理：基于`jwt`和`casbin`实现的权限管理。
- 文件上传下载：实现基于`七牛云`, `阿里云`, `腾讯云` 的文件上传操作(请开发自己去各个平台的申请对应 `token` 或者对应`key`)。
- 分页封装：前端使用 `mixins` 封装分页，分页方法调用 `mixins` 即可。
- 用户管理：系统管理员分配用户角色和角色权限。
- 角色管理：创建权限控制的主要对象，可以给角色分配不同api权限和菜单权限。
- 菜单管理：实现用户动态菜单配置，实现不同角色不同菜单。
- api管理：不同用户可调用的api接口的权限不同。
- 配置管理：配置文件可前台修改(在线体验站点不开放此功能)。
- 条件搜索：增加条件搜索示例。
- restful示例：可以参考用户管理模块中的示例API。
	- 前端文件参考: [web/src/view/superAdmin/api/api.vue](https://github.com/cngamesdk/background/blob/master/web/src/view/superAdmin/api/api.vue)
    - 后台文件参考: [server/router/sys_api.go](https://github.com/cngamesdk/background/blob/master/server/router/system/sys_api.go)
- 多点登录限制：需要在`config.yaml`中把`system`中的`use-multipoint`修改为true(需要自行配置Redis和Config中的Redis参数，测试阶段，有bug请及时反馈)。
- 分片上传：提供文件分片上传和大文件分片上传功能示例。
- 表单生成器：表单生成器借助 [@Variant Form](https://github.com/vform666/variant-form) 。
- 代码生成器：后台基础逻辑以及简单curd的代码生成器。
