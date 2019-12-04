# gmanager

#### 介绍
基于gf框架的管理平台，支持登录、认证、组织机构、用户、角色、菜单、日志

* github地址：https://github.com/goflyfox/gmanager
* gitee地址：https://gitee.com/goflyfox/gmanager

#### 安装教程

1. 从git下载项目： git clone https://github.com/goflyfox/gmanager
2. 安装mysql数据库，创建db，运行deploy下gmanager.sql脚本
3. 修改config下config.toml配置文件
```toml
# 数据库配置
[database]
    link = "root:123456@tcp(127.0.0.1:3306)/gmanager"
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
7. 支持登录、登出、业务增删改操作记录
8. 支持接口调用、返回参数打印，便于问题排查

#### 代码生成
如需代码生成，请参考java项目：https://gitee.com/jflyfox/AutoCreate

#### 平台截图

登录：
![image](https://raw.githubusercontent.com/goflyfox/gmanager/master/deploy/image/1.png)

组织机构：
![image](https://raw.githubusercontent.com/goflyfox/gmanager/master/deploy/image/2.png)

用户管理：
![image](https://raw.githubusercontent.com/goflyfox/gmanager/master/deploy/image/3.png)

日志管理：
![image](https://raw.githubusercontent.com/goflyfox/gmanager/master/deploy/image/4.png)

#### 感谢

1. gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf) 

#### 项目支持

- 项目的发展，离不开大家得支持~！~

- [【双12】主会场 低至1折；请点击这里](https://www.aliyun.com/1212/2019/home?userCode=c4hsn0gc)
- [阿里云：ECS云服务器2折起；请点击这里](https://www.aliyun.com/acts/limit-buy?spm=5176.11544616.khv0c5cu5.1.1d8e23e8XHvEIq&userCode=c4hsn0gc)
- [阿里云：ECS云服务器新人优惠券；请点击这里](https://promotion.aliyun.com/ntms/yunparter/invite.html?userCode=c4hsn0gc)

- 也可以请作者喝一杯咖啡:)

![jflyfox](https://raw.githubusercontent.com/jflyfox/jfinal_cms/master/doc/pay01.jpg "Open source support")
