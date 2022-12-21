# gin-demo

## 一、项目内容

基于gin框架的demo项目,本项目基于B站视频Go语言 Gin+Vue 前后端分离实战 - OceanLearn
[https://www.bilibili.com/video/BV1CE411H7bQ?t=7](https://www.bilibili.com/video/BV1CE411H7bQ?t=7)

**master**分支为后端go代码,**vue**分支为前端vue代码

## 二、怎样运行该项目

### 2.1 运行后端程序

> 先确保你电脑上正确安装了 golang 环境

从master分支拉取后端golang代码

```bash
# 拉取代码
git clone -b main https://github.com/kotlindev/gin-demo.git backend
# 进入项目目录
cd  backend
# 安装项目依赖
go get
```

如果没有mysql数据库，可以使用docker启动一个mysql数据库，项目中提供脚本快速启动，脚本在`db`目录下，相关操作如下

```bash
# 进入db目录
cd db
# 启动mysql数据库
./startup.sh
# 销毁数据库
./shutdown.sh
```

启动数据库后，在浏览器中访问`http://localhost:8080`，进入数据库管理页面。默认两个账号如下

| 账号       | 密码       | 级别    |
|----------|----------|-------|
| root     | root     | 超级管理员 |
| gin_demo | gin_demo | 普通用户  |

打开 `config/application.yaml` 文件，修改数据库链接配置，修改项目运行端口，确保端口不被占用，参考如下

```yaml
server:
  port: 8081
datasource:
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: gin_demo
  username: gin_demo
  password: gin_demo
  charset: utf8mb4
  loc: Asia/Shanghai
```

启动项目

```bash
go run routes.go main.go
```

如果看到命令行终端输出以下路由信息，代表项目运行正常。如果不正常，检查一下数据库地址还有账号密码是否正确，同时确保运行的端口没有被占用

```bash
[GIN-debug] POST   /api/auth/register        --> jkdev.cn/api/controller.Register (5 handlers)
[GIN-debug] POST   /api/auth/login           --> jkdev.cn/api/controller.Login (5 handlers)
[GIN-debug] GET    /api/auth/info            --> jkdev.cn/api/controller.Info (6 handlers)
[GIN-debug] POST   /categories               --> jkdev.cn/api/controller.ICategoryController.Create-fm (5 handlers)
[GIN-debug] PUT    /categories/:id           --> jkdev.cn/api/controller.ICategoryController.Update-fm (5 handlers)
[GIN-debug] GET    /categories/:id           --> jkdev.cn/api/controller.ICategoryController.Show-fm (5 handlers)
[GIN-debug] DELETE /categories/:id           --> jkdev.cn/api/controller.ICategoryController.Delete-fm (5 handlers)
[GIN-debug] POST   /posts                    --> jkdev.cn/api/controller.IPostController.Create-fm (6 handlers)
[GIN-debug] PUT    /posts/:id                --> jkdev.cn/api/controller.IPostController.Update-fm (6 handlers)
[GIN-debug] GET    /posts/:id                --> jkdev.cn/api/controller.IPostController.Show-fm (6 handlers)
[GIN-debug] DELETE /posts/:id                --> jkdev.cn/api/controller.IPostController.Delete-fm (6 handlers)
[GIN-debug] POST   /posts/page/list          --> jkdev.cn/api/controller.IPostController.PageList-fm (6 handlers)
[GIN-debug] Listening and serving HTTP on :8081
```

### 2.2 运行前端程序

> 先确保你电脑上正确安装了 npm 环境，并安装了 vue、yarn

从vue分支拉取前端vue代码

```bash
# 拉取代码
git clone -b vue https://github.com/kotlindev/gin-demo.git vue
# 进入项目目录
cd  vue
# 安装项目依赖
yarn install
```

根据1中的 后端代码的运行端口，修改 `.env.development.local` 和 `.env.development` 两个配置文件，修改配置如下为

```bash
VUE_APP_BASE_URL = http://localhost:8081/api/
```

在运行项目

```bash
yarn serve
```
