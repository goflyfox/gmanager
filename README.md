# gmanager

#### 介绍
基于gf框架的管理平台，支撑登录、认证、组织机构、用户、角色、菜单、日志

注：此项目基本功能已具备，暂时仅供学习go或者gf框架的小伙伴参考；但还在孵化中，革命仍需努力！！！

#### 安装教程

1. 从git下载项目： git clone https://github.com/goflyfox/gmanager
2. 安装mysql数据库，创建db，运行deploy下gmanager.sql脚本
3. 修改config下config.toml配置文件
```toml
# 数据库配置
[database]
    [[database.default]]
        type     = "mysql"
        linkinfo = "root:123456@tcp(127.0.0.1:33306)/gmanager"
```
4. go run main.go
5. 访问http://localhost即可看到登录页面，账号/密码：admin/123456

#### 功能模块

1. 登录、认证、登出
2. 组织机构管理
3. 用户管理
4. 角色管理
5. 菜单管理
6. 日志管理

#### 平台截图


#### 感谢

1. gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf) 
