# gmanager使用说明

## 项目目录结构
```bash
gmanager:.
│  ChangeLog.md                 更新说明
│  go.mod                       go mod 
│  LICENSE                      开源协议
│  main.go                      go入口文件
│  README.md                    项目介绍
├─app                           业务代码目录
│  ├─api                        业务接口
│  ├─component
│  │  ├─middle
│  │  └─     MiddlewareLog.go  访问日志中间件
│  ├─constants                  常量
│  ├─model                     实体对象   
│  └─service                业务层代码
│      └─user
│              userSvc.go
├─boot                      启动类
│      boot.go
├─config                    配置文件
│      config.toml
├─deploy
│  │  Donate.md             贡献者记录
│  │  gmanager-web.sql      初始化脚本（前后端分离）
│  │  gmanager.sql          初始化脚本
│  │  ReadMe.md             使用说明
├─library                   公共方法  
├─logs                      日志目录  
├─public                    静态文件
│  │  favicon.ico
│  └─component              第三方组件                          
├─router                    项目路由
├─template                  模板页面
└─test                      测试代码
```

## model生成
使用gf自动生成model

```bash
gf gen model -l "mysql:root:123123@tcp(10.130.44.180:3306)/db" -p sys_
gf gen model -l "mysql:root:A123456@tcp(127.0.0.1:33306)/gmanager" -p sys_
```