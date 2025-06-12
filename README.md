# gmanager后端管理系统

- 基于GoFrame V2框架的后台管理系统。支持登录、认证、组织机构、用户帐号、角色权限、菜单、配置、操作日志等模块
- 前端基于 Vue3 + Vite + TypeScript + Element-Plus 的后台管理模板
- 登录组件使用gtoken完美支持集群部署

## 代码
* github地址： https://github.com/goflyfox/gmanager
* gitee地址： https://gitee.com/goflyfox/gmanager

## 功能模块

1. 部门管理：配置系统组织机构信息
2. 用户管理：处理用户添加、用户配置、权限分配
3. 角色管理：角色菜单权限分配，支持菜单和按钮权限设置
4. 菜单管理：配置系统菜单、按钮权限等
5. 配置管理：支持对系统参数动态配置及数据字典配置
6. 日志管理：支持登录、登出、业务增删改操作记录
7. 其他：登录、认证、登出、访问统计

> gmanager开源以来得到了大家的很多支持，本项目初衷只为互相学习交流，没有任何盈利性目的！欢迎为gmanager贡献代码或提供建议！

## 部署说明

以下为部署简要说明，详情参考[部署文档](docs/installation.md) 

### 后端

1. 从git下载项目： git clone https://github.com/goflyfox/gmanager
2. 安装mysql数据库运行resource/sql/gmanager.sql脚本
3. 复制`server/manifest/config/config.example.yaml`配置文件，改名为`config.yaml`,修改数据库配置

```toml
# 数据库配置
database:
  default:
    link: "mysql:root:123456@tcp(127.0.0.1:3306)/gmanager"
```

4. 启动项目

```bash
cd server
go mod tidy
go run main.go
```

5. 访问 http://localhost:8000/ping 后端接口，返回`pong`验证部署成功

### 前端部署


```bash
cd web
pnpm install
pnpm run dev
```

浏览器访问 [http://localhost:3000](http://localhost:3000) 即可看到登录页面，默认账号/密码：admin/123456

## 使用文档

[说明文档](docs/README.md) · [部署文档](docs/installation.md) · [更新说明](docs/ChangeLog.md)

## 效果截图
登录：
![image](docs/images/pic_login.png)

组织机构：
![image](docs/images/pic_dept.png)

用户管理：
![image](docs/images/pic_user.png)

日志管理：
![image](docs/images/pic_log.png)

##  感谢

- gf框架 [https://github.com/gogf/gf](https://gitee.com/link?target=https%3A%2F%2Fgithub.com%2Fgogf%2Fgf)

## 项目支持

- 项目的发展，离不开大家得支持~！~
- 可以请作者喝一杯咖啡:)

![jflyfox](https://raw.githubusercontent.com/jflyfox/jfinal_cms/master/doc/pay01.jpg "Open source support")

[捐赠列表](docs/Donate.md)
