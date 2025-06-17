-- ----------------------------
-- 创建数据库
-- ----------------------------
-- create schema gmanager collate utf8mb4_bin;
CREATE DATABASE IF NOT EXISTS gmanager CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_bin;


-- ----------------------------
-- 创建表 && 数据初始化
-- ----------------------------
use gmanager;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Drop Table
-- ----------------------------

drop table if exists sys_config;
drop table if exists sys_dept;
drop table if exists sys_log;
drop table if exists sys_menu;
drop table if exists sys_role;
drop table if exists sys_role_menu;
drop table if exists sys_user;
drop table if exists sys_user_role;


-- ----------------------------
-- Create Table
-- ----------------------------
create table sys_config
(
    id            bigint auto_increment comment '主键'
        primary key,
    name          varchar(128)          not null comment '名称',
    `key`         varchar(64)           not null comment '键',
    value         varchar(4000)         not null comment '值',
    code          varchar(256)          null comment '编码',
    data_type     int                   null comment '数据类型//radio/1,KV配置,2,字典,3,字典数据',
    parent_id     bigint     default 0  not null comment '类型',
    parent_key    varchar(64)           null,
    remark        varchar(4000)         null comment '备注',
    sort          int        default 10 not null comment '排序号',
    copy_status   tinyint(1) default 1  not null comment '拷贝状态 1 拷贝  2  不拷贝',
    change_status tinyint(1) default 1  not null comment '1 可以更改 2 不可更改 ',
    enable        tinyint(1) default 1  null comment '是否启用//radio/1,启用,2,禁用',
    update_at     datetime              null comment '更新时间',
    update_id     bigint     default 0  null comment '更新人',
    create_at     datetime              null comment '创建时间',
    create_id     bigint     default 0  null comment '创建者',
    constraint uni_config_key
        unique (`key`) comment '配置key唯一键'
)
    comment '系统配置表' charset = utf8mb4
                         row_format = DYNAMIC;

create table sys_dept
(
    id         bigint auto_increment comment '主键'
        primary key,
    parent_id  bigint     default 0         null comment '上级机构',
    name       varchar(32) charset utf8mb4  not null comment '部门/11111',
    code       varchar(128) charset utf8mb4 null comment '机构编码',
    sort       int        default 0         null comment '序号',
    linkman    varchar(64) charset utf8mb4  null comment '联系人',
    linkman_no varchar(32) charset utf8mb4  null comment '联系人电话',
    remark     varchar(128) charset utf8mb4 null comment '机构描述',
    enable     tinyint(1) default 1         null comment '是否启用//radio/1,启用,2,禁用',
    update_at  datetime                     null comment '更新时间',
    update_id  bigint     default 0         null comment '更新人',
    create_at  datetime                     null comment '创建时间',
    create_id  bigint     default 0         null comment '创建者',
    constraint uni_depart_name
        unique (name)
)
    comment '组织机构' row_format = DYNAMIC;

create table sys_log
(
    id             bigint auto_increment comment '主键'
        primary key,
    log_type       int                  not null comment '类型',
    oper_object    varchar(2000)        null comment '操作对象',
    oper_table     varchar(64)          not null comment '操作表',
    oper_id        bigint     default 0 null comment '操作主键',
    oper_type      varchar(64)          null comment '操作类型',
    oper_remark    varchar(2000)        null comment '操作备注',
    url            varchar(1000)        null comment '提交url',
    method         varchar(128)         null comment '请求方式',
    ip             varchar(128)         null comment 'IP地址',
    user_agent     varchar(512)         null comment 'UA信息',
    execution_time bigint     default 0 null comment '响应时间',
    operator       varchar(200)         null comment '操作人',
    enable         tinyint(1) default 1 null comment '是否启用//radio/1,启用,2,禁用',
    update_at      datetime             null comment '更新时间',
    update_id      bigint     default 0 null comment '更新人',
    create_at      datetime             null comment '创建时间',
    create_id      bigint     default 0 null comment '创建者'
)
    comment '日志' charset = utf8mb4
                   row_format = DYNAMIC;

create table sys_menu
(
    id          bigint auto_increment comment 'ID'
        primary key,
    parent_id   bigint               not null comment '父菜单ID',
    name        varchar(64)          not null comment '菜单名称',
    type        tinyint              not null comment '菜单类型（1-菜单 2-目录 3-外链 4-按钮）',
    route_name  varchar(255)         null comment '路由名称（Vue Router 中用于命名路由）',
    route_path  varchar(128)         null comment '路由路径（Vue Router 中定义的 URL 路径）',
    component   varchar(128)         null comment '组件路径（组件页面完整路径，相对于 src/views/，缺省后缀 .vue）',
    perm        varchar(128)         null comment '【按钮】权限标识',
    always_show tinyint    default 0 null comment '【目录】只有一个子路由是否始终显示（1-是 0-否）',
    keep_alive  tinyint    default 0 null comment '【菜单】是否开启页面缓存（1-是 0-否）',
    sort        int        default 0 null comment '排序',
    icon        varchar(64)          null comment '菜单图标',
    redirect    varchar(128)         null comment '跳转路径',
    params      varchar(255)         null comment '路由参数',
    enable      tinyint(1) default 1 null comment '是否启用//radio/1,启用,2,禁用',
    update_at   datetime             null comment '更新时间',
    update_id   bigint               null comment '更新人',
    create_at   datetime             null comment '创建时间',
    create_id   bigint               null comment '创建者'
)
    comment '菜单管理' charset = utf8mb4;

create table sys_role
(
    id         bigint auto_increment comment '主键'
        primary key,
    name       varchar(200) default '' not null comment '名称/11111/',
    code       varchar(200)            null comment '角色编码',
    data_scope tinyint                 null comment '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
    sort       int          default 1  null comment '排序',
    remark     text                    null comment '说明//textarea',
    enable     tinyint(1)   default 1  null comment '是否启用//radio/1,启用,2,禁用',
    update_at  datetime                null comment '更新时间',
    update_id  bigint       default 0  null comment '更新人',
    create_at  datetime                null comment '创建时间',
    create_id  bigint       default 0  null comment '创建者'
)
    comment '角色' charset = utf8mb4
                   row_format = DYNAMIC;

create table sys_role_menu
(
    role_id bigint not null comment '角色id',
    menu_id bigint not null comment '菜单id'
)
    comment '角色和菜单关联' charset = utf8mb4
                             row_format = DYNAMIC;

create table sys_user
(
    id        bigint auto_increment comment '主键'
        primary key,
    uuid      varchar(32)                null comment 'UUID',
    user_name varchar(32)                not null comment '登录名/11111',
    mobile    varchar(32)                null comment '手机号',
    email     varchar(64)                null comment 'email',
    password  varchar(32)                not null comment '密码',
    salt      varchar(16) default '1111' not null comment '密码盐',
    dept_id   bigint      default 0      null comment '部门/11111/dict',
    user_type int         default 2      null comment '类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户',
    status    int         default 10     null comment '状态',
    thirdid   varchar(200)               null comment '第三方ID',
    endtime   varchar(32)                null comment '结束时间',
    nick_name varchar(32)                null comment '昵称',
    gender    tinyint                    null comment '性别;0:保密,1:男,2:女',
    address   varchar(32)                null comment '地址',
    avatar    varchar(200)               null comment '头像地址',
    birthday  int                        null comment '生日',
    remark    varchar(1000)              null comment '说明',
    enable    tinyint(1)  default 1      null comment '是否启用//radio/1,启用,2,禁用',
    update_at datetime                   null comment '更新时间',
    update_id bigint      default 0      null comment '更新人',
    create_at datetime                   null comment '创建时间',
    create_id bigint      default 0      null comment '创建者',
    constraint uni_user_username
        unique (user_name)
)
    comment '用户' charset = utf8mb4
                   row_format = DYNAMIC;

grant select on table sys_user to uroot@localhost;

create table sys_user_role
(
    user_id bigint not null comment '用户id',
    role_id bigint not null comment '角色id'
)
    comment '用户和角色关联' charset = utf8mb4
                             row_format = DYNAMIC;



-- ----------------------------
-- Records
-- ----------------------------
INSERT INTO sys_dept (id, parent_id, name, code, sort, linkman, linkman_no, remark, enable, update_at, update_id, create_at, create_id) VALUES (10001, 0, 'FLY的狐狸', 'ABC000', 100, '', '', '', 1, '2025-06-03 15:05:18', 1, '2025-06-03 15:05:18', 1);
INSERT INTO sys_dept (id, parent_id, name, code, sort, linkman, linkman_no, remark, enable, update_at, update_id, create_at, create_id) VALUES (10002, 10001, '开发组', 'ABC001', 101, null, null, null, 1, '2016-07-31 18:15:29', 1, '2016-07-31 18:15:29', 1);
INSERT INTO sys_dept (id, parent_id, name, code, sort, linkman, linkman_no, remark, enable, update_at, update_id, create_at, create_id) VALUES (10003, 10001, '产品组', 'ABC003', 103, '', '', '', 1, '2017-04-28 00:58:41', 1, '2016-07-31 18:16:06', 1);
INSERT INTO sys_dept (id, parent_id, name, code, sort, linkman, linkman_no, remark, enable, update_at, update_id, create_at, create_id) VALUES (10004, 10001, '运营组', 'ABC004', 104, null, null, null, 1, '2016-07-31 18:16:30', 1, '2016-07-31 18:16:30', 1);
INSERT INTO sys_dept (id, parent_id, name, code, sort, linkman, linkman_no, remark, enable, update_at, update_id, create_at, create_id) VALUES (10005, 10001, '测试组', 'ABC005', 105, '', '', '', 0, '2025-06-11 15:09:46', 1, '2025-06-11 15:09:46', 1);


INSERT INTO sys_user (id, uuid, user_name, mobile, email, password, salt, dept_id, user_type, status, thirdid, endtime, nick_name, gender, address, avatar, birthday, remark, enable, update_at, update_id, create_at, create_id) VALUES (1, '94091b1fa6ac4a27a06c0b92155aea6a', 'admin', '15812345678', '330627517@qq.com', '9fb3dc842c899aa63d6944a55080b795', '1111', 10001, 1, 1, null, '', '管理员', 1, '', 'https://www.jflyfox.com/static/img/logo.png', 0, '时间是最好的老师，但遗憾的是&mdash;&mdash;最后他把所有的学生都弄死了', 1, '2025-06-17 14:19:10', 1, '2025-05-24 00:23:23', 1);
INSERT INTO sys_user (id, uuid, user_name, mobile, email, password, salt, dept_id, user_type, status, thirdid, endtime, nick_name, gender, address, avatar, birthday, remark, enable, update_at, update_id, create_at, create_id) VALUES (6, '1ltefoo1ed6dah7rjog2tyw1001dzhd1', 'test', '15812345671', '1581@158.com', '8ae97053f637c9b0dc826dfe40976fd1', 'WpBnHl', 10003, 2, 0, null, '', '测试', 2, '', 'https://www.jflyfox.com/static/img/logo.png', 0, '', 1, '2025-06-17 14:19:05', 1, '2025-06-08 14:16:00', 1);
INSERT INTO sys_user (id, uuid, user_name, mobile, email, password, salt, dept_id, user_type, status, thirdid, endtime, nick_name, gender, address, avatar, birthday, remark, enable, update_at, update_id, create_at, create_id) VALUES (7, '1fp076ol9d0dakd766thvzs100zhed31', 'test2', '15612345679', '15612345679@158.com', '30adf34d6e5d6060b955dae6137e8ace', 'EKQhLT', 10003, 2, 0, null, '', '测试员2', 1, '', '', 0, '', 2, '2025-06-17 14:18:58', 1, '2025-06-12 07:09:30', 1);

INSERT INTO sys_config (id, name, `key`, value, code, data_type, parent_id, parent_key, remark, sort, copy_status, change_status, enable, update_at, update_id, create_at, create_id) VALUES (24, '系统参数', 'system', '', '', 2, 0, null, null, 15, 1, 2, 1, '2017-09-15 17:02:30', 4, '2017-09-15 16:54:52', 4);
INSERT INTO sys_config (id, name, `key`, value, code, data_type, parent_id, parent_key, remark, sort, copy_status, change_status, enable, update_at, update_id, create_at, create_id) VALUES (46, '日志控制配置', 'system.debug', 'false', '', 3, 24, 'system', null, 15, 1, 1, 1, '2019-02-24 00:00:08', 0, '2017-09-15 17:06:21', 4);
INSERT INTO sys_config (id, name, `key`, value, code, data_type, parent_id, parent_key, remark, sort, copy_status, change_status, enable, update_at, update_id, create_at, create_id) VALUES (51, '性别', 'gender', '', '', 2, 0, null, '', 11, 0, 0, 1, '2025-06-09 15:37:28', 1, '2025-06-09 15:37:28', 1);
INSERT INTO sys_config (id, name, `key`, value, code, data_type, parent_id, parent_key, remark, sort, copy_status, change_status, enable, update_at, update_id, create_at, create_id) VALUES (52, '男', 'gender.male', '1', '', 3, 51, 'gender', '', 12, 0, 0, 1, '2025-06-09 16:16:42', 1, '2025-06-09 16:16:42', 1);
INSERT INTO sys_config (id, name, `key`, value, code, data_type, parent_id, parent_key, remark, sort, copy_status, change_status, enable, update_at, update_id, create_at, create_id) VALUES (53, '女', 'gender.female', '2', '', 3, 51, 'gender', '', 12, 0, 0, 1, '2025-06-09 16:16:48', 1, '2025-06-09 16:16:48', 1);
INSERT INTO sys_config (id, name, `key`, value, code, data_type, parent_id, parent_key, remark, sort, copy_status, change_status, enable, update_at, update_id, create_at, create_id) VALUES (54, '未知', 'gender.unkown', '3', '', 3, 51, 'gender', '', 12, 0, 0, 1, '2025-06-09 16:16:53', 1, '2025-06-09 16:16:53', 1);

INSERT INTO sys_role (id, name, code, data_scope, sort, remark, enable, update_at, update_id, create_at, create_id) VALUES (1, '管理员', 'ADMIN', 1, 8, '', 1, '2025-06-12 07:27:39', 1, '2025-06-12 07:27:39', 1);
INSERT INTO sys_role (id, name, code, data_scope, sort, remark, enable, update_at, update_id, create_at, create_id) VALUES (2, '默认角色', 'DEFAULT', 2, 9, '', 1, '2025-06-12 07:27:47', 1, '2025-06-12 07:27:47', 1);
INSERT INTO sys_role (id, name, code, data_scope, sort, remark, enable, update_at, update_id, create_at, create_id) VALUES (4, '测试', 'TEST', 2, 10, '', 1, '2025-06-12 07:27:50', 1, '2025-06-12 07:27:50', 1);

INSERT INTO sys_user_role (user_id, role_id) VALUES (7, 2);
INSERT INTO sys_user_role (user_id, role_id) VALUES (6, 4);
INSERT INTO sys_user_role (user_id, role_id) VALUES (1, 1);

INSERT INTO sys_role_menu (role_id, menu_id) VALUES (1, 20);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (1, 1);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 1);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 2);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 105);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 31);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 32);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 33);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 88);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 106);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 107);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 3);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 139);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 71);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 72);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 4);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 140);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 75);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 74);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 5);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 141);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 77);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 78);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 117);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 120);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 121);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 123);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 124);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 125);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (2, 40);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 1);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 2);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 105);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 31);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 32);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 33);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 88);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 106);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 107);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 3);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 139);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 71);
INSERT INTO sys_role_menu (role_id, menu_id) VALUES (4, 72);

INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (1, 0, '系统管理', 2, '', '/system', 'Layout', null, null, null, 1, 'system', '/system/user', null, 1, '2025-06-02 23:41:33', null, '2025-06-02 23:41:33', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (2, 1, '用户管理', 1, 'User', 'user', 'system/user/index', null, null, 1, 1, 'el-icon-User', null, null, 1, '2025-06-02 23:41:33', null, '2025-06-02 23:41:33', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (3, 1, '角色管理', 1, 'Role', 'role', 'system/role/index', null, null, 1, 2, 'role', null, null, 1, '2025-06-02 23:41:33', null, '2025-06-02 23:41:33', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (4, 1, '菜单管理', 1, 'SysMenu', 'menu', 'system/menu/index', null, null, 1, 3, 'menu', null, null, 1, '2025-06-02 23:41:33', null, '2025-06-02 23:41:33', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (5, 1, '部门管理', 1, 'Dept', 'dept', 'system/dept/index', null, null, 1, 4, 'tree', null, null, 1, '2025-06-02 23:41:34', null, '2025-06-02 23:41:34', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (31, 2, '用户新增', 4, null, '', null, 'admin:user:save', null, null, 1, '', '', null, 1, '2025-06-02 23:41:35', null, '2025-06-02 23:41:35', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (32, 2, '用户编辑', 4, null, '', null, 'admin:user:save', null, null, 2, '', '', null, 1, '2025-06-02 23:41:35', null, '2025-06-02 23:41:35', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (33, 2, '用户删除', 4, null, '', null, 'admin:user:delete', null, null, 3, '', '', null, 1, '2025-06-02 23:41:35', null, '2025-06-02 23:41:35', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (40, 1, '接口文档', 3, 'Api', 'http://127.0.0.1:8000/swagger/', '', '', 1, 0, 20, 'api', '', '[]', 1, '2025-06-11 13:27:02', 1, '2025-06-11 13:27:02', 1);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (71, 3, '角色编辑', 4, null, '', null, 'admin:role:save', null, null, 3, '', null, null, 1, '2025-06-02 23:41:36', null, '2025-06-02 23:41:36', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (72, 3, '角色删除', 4, null, '', null, 'admin:role:delete', null, null, 4, '', null, null, 1, '2025-06-02 23:41:36', null, '2025-06-02 23:41:36', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (74, 4, '菜单编辑', 4, null, '', null, 'admin:menu:save', null, null, 3, '', null, null, 1, '2025-06-02 23:41:37', null, '2025-06-02 23:41:37', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (75, 4, '菜单删除', 4, null, '', null, 'admin:menu:delete', null, null, 3, '', null, null, 1, '2025-06-02 23:41:37', null, '2025-06-02 23:41:37', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (77, 5, '部门编辑', 4, null, '', null, 'admin:dept:save', null, null, 2, '', null, null, 1, '2025-06-02 23:41:37', null, '2025-06-02 23:41:37', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (78, 5, '部门删除', 4, null, '', null, 'admin:dept:delete', null, null, 3, '', null, null, 1, '2025-06-02 23:41:37', null, '2025-06-02 23:41:37', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (88, 2, '重置密码', 4, null, '', null, 'admin:user:reset-password', null, null, 4, '', null, null, 1, '2025-06-02 23:41:38', null, '2025-06-02 23:41:38', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (105, 2, '用户查询', 4, null, '', null, 'admin:user:query', 0, 0, 0, '', null, null, 1, '2025-06-02 23:41:39', null, '2025-06-02 23:41:39', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (106, 2, '用户导入', 4, null, '', null, 'admin:user:import', null, null, 5, '', null, null, 1, '2025-06-02 23:41:39', null, '2025-06-02 23:41:39', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (107, 2, '用户导出', 4, null, '', null, 'admin:user:export', null, null, 6, '', null, null, 1, '2025-06-02 23:41:39', null, '2025-06-02 23:41:39', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (108, 36, '增删改查', 1, null, 'curd', 'demo/curd/index', null, null, 1, 0, '', '', null, 1, null, null, null, null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (117, 1, '系统日志', 1, 'Log', 'log', 'system/log/index', null, 0, 1, 6, 'document', null, null, 1, '2025-06-02 23:41:40', null, '2025-06-02 23:41:40', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (120, 1, '系统配置', 1, 'Config', 'config', 'system/config/index', null, 0, 1, 7, 'setting', null, null, 1, '2025-06-02 23:41:40', null, '2025-06-02 23:41:40', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (121, 120, '系统配置查询', 4, null, '', null, 'admin:config:query', 0, 1, 1, '', null, null, 1, '2025-06-02 23:41:40', null, '2025-06-02 23:41:40', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (123, 120, '系统配置修改', 4, null, '', null, 'admin:config:save', 0, 1, 3, '', null, null, 1, '2025-06-02 23:41:40', null, '2025-06-02 23:41:40', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (124, 120, '系统配置删除', 4, null, '', null, 'admin:config:delete', 0, 1, 4, '', null, null, 1, '2025-06-02 23:41:41', null, '2025-06-02 23:41:41', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (125, 120, '系统配置刷新', 4, null, '', null, 'admin:config:refresh', 0, 1, 5, '', null, null, 1, '2025-06-02 23:41:41', null, '2025-06-02 23:41:41', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (139, 3, '角色查询', 4, null, '', null, 'admin:role:query', null, null, 1, '', null, null, 1, '2025-06-02 23:41:42', null, '2025-06-02 23:41:42', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (140, 4, '菜单查询', 4, null, '', null, 'admin:menu:query', null, null, 1, '', null, null, 1, '2025-06-02 23:41:43', null, '2025-06-02 23:41:43', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (141, 5, '部门查询', 4, null, '', null, 'admin:dept:query', null, null, 1, '', null, null, 1, '2025-06-02 23:41:43', null, '2025-06-02 23:41:43', null);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (149, 0, '客户管理', 2, '', '/customer', '', '', 2, 1, 20, 'bell', '', '[]', 2, '2025-06-12 07:28:30', 1, '2025-06-12 07:28:30', 1);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (150, 149, '客户信息', 1, 'Custumer', 'custumer', 'customer/index', '', 2, 2, 1, 'fullscreen', '', '[]', 2, '2025-06-08 16:02:53', 1, '2025-06-08 16:02:53', 1);
INSERT INTO sys_menu (id, parent_id, name, type, route_name, route_path, component, perm, always_show, keep_alive, sort, icon, redirect, params, enable, update_at, update_id, create_at, create_id) VALUES (151, 149, '客户资料', 1, 'Info', 'ifno', 'customer/info', '', 2, 1, 1, '', '', '[]', 2, '2025-06-13 06:06:05', 1, '2025-06-13 06:06:05', 1);

SET FOREIGN_KEY_CHECKS = 1;
