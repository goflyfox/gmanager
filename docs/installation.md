# 项目部署

## 一、代码地址

* github地址：https://github.com/goflyfox/gmanager
* gitee地址：https://gitee.com/goflyfox/gmanager

## 二、环境准备（已具备开发环境，可忽略）

### 2.1 环境要求

- node版本 >= v20.0.0
- golang版本 >= v1.23
- goframe版本 >=v2.9.0
- mysql版本 >=8.0

###  2.2 前端环境

1. 前往：https://nodejs.org/zh-cn/download
2. 命令行运行 `node -v` 若控制台输出版本号则node安装成功
3. node 版本需大于等于 `20.0`
4. 安装pnpm：`npm install -g pnpm`
5. 命令行运行 `pnpm -v` 若控制台输出版本号则前端环境搭建成功

### 2.3 后端环境

1. 下载golang安装 版本号需>=1.23
2. 国际: [https://golang.org/dl/](https://golang.org/dl/)  国内: [https://golang.google.cn/dl/](https://golang.google.cn/dl/)
4. 命令行运行 go 若控制台输出各类提示命令 则安装成功 输入 `go version` 确认版本大于1.23
5. 开发工具推荐 [Goland](https://www.jetbrains.com/go/)

## 三、后端部署

### 3.1 部署
1. 从git下载项目： git clone https://github.com/goflyfox/gmanager
2. 安装mysql数据库，创建db，运行resource/sql/gmanager.sql脚本
3. 复制`server/manifest/config/config.example.yaml`配置文件，改名为`config.yaml`,修改数据库配置
```toml
# 数据库配置
database:
  default:
    link: "mysql:root:123456@tcp(127.0.0.1:3306)/gmanager"
```
4. 启动后端服务

```bash
# 进入服务端目录
cd server
# 编译
go mod tidy
# 启动
go run main.go
```

5. 访问 http://localhost:8000/ping 后端接口，返回`pong`验证部署成功

### 3.2 服务端打包
1. 打包可以使用gf提供的gf-cli进行打包
2. 也可通过原始交叉编译命令

打linux环境包
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

打本地环境包
```
go build main.go
```

## 四、前端部署

### 4.1 部署


```bash
# 进入项目目录
cd web

# 安装依赖
pnpm install

# 启动服务
pnpm run dev
```

浏览器访问 [http://localhost:3000](http://localhost:3000) 即可看到登录页面，默认账号/密码：admin/123456

### 4.2 前端打包发布

```bash
# 构建生产环境
pnpm run build
```

项目会在`dist`目录生成打包好的文件，可以通过Nginx等静态服务器进行部署。







