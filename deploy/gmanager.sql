SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `key` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '键',
  `value` varchar(4000) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '值',
  `code` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '编码',
  `data_type` int(2) NULL DEFAULT NULL COMMENT '数据类型//radio/1,KV,2,字典,3,数组',
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '类型',
  `parent_key` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` int(11) NOT NULL DEFAULT 10 COMMENT '排序号',
  `project_id` bigint(20) NULL DEFAULT 1 COMMENT '项目ID',
  `copy_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '拷贝状态 1 拷贝  2  不拷贝',
  `change_status` tinyint(1) NOT NULL DEFAULT 2 COMMENT '1 不可更改 2 可以更改',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 59 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '系统配置表' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (24, '系统参数', 'system', '', '', NULL, 0, NULL, 15, 1, 1, 2, 1, '2017-09-15 17:02:30', 4, '2017-09-15 16:54:52', 4);
INSERT INTO `sys_config` VALUES (46, '日志控制配置', 'system.debug', 'false', '', NULL, 24, 'system', 15, 1, 1, 1, 1, '2019-02-24 00:00:08', 0, '2017-09-15 17:06:21', 4);
INSERT INTO `sys_config` VALUES (47, '短信配置', 'sms', '', '', NULL, 0, '', 15, 1, 1, 2, 1, '2019-02-20 22:45:41', 1, '2017-09-15 17:06:56', 4);
INSERT INTO `sys_config` VALUES (50, '短信账号', 'sms.username', 'test', '', NULL, 47, 'sms', 10, 1, 1, 2, 1, '2019-02-20 22:26:29', 1, '2019-02-18 01:07:47', 1);
INSERT INTO `sys_config` VALUES (51, '短信密码', 'sms.passwd', '111111', '', NULL, 47, 'sms', 10, 1, 1, 2, 1, '2019-02-18 01:08:16', 1, '2019-02-18 01:08:16', 1);
INSERT INTO `sys_config` VALUES (52, '短信类型', 'sms.type', '阿里云', '', NULL, 47, 'sms', 10, 1, 1, 2, 1, '2019-02-20 22:26:21', 1, '2019-02-20 22:26:21', 1);
INSERT INTO `sys_config` VALUES (53, '性别', 'sex', '', '', NULL, 0, NULL, 90, 1, 1, 2, 1, '2019-02-20 23:35:18', 1, '2019-02-20 23:35:18', 1);
INSERT INTO `sys_config` VALUES (54, '性别男', 'sex.male', '男', '1', NULL, 53, 'sex', 91, 1, 1, 2, 1, '2019-02-20 23:40:19', 1, '2019-02-20 23:35:45', 1);
INSERT INTO `sys_config` VALUES (55, '性别女', 'sex.female', '女', '2', NULL, 53, 'sex', 92, 1, 1, 2, 1, '2019-02-20 23:40:24', 1, '2019-02-20 23:36:12', 1);
INSERT INTO `sys_config` VALUES (56, '性别未知', 'sex.unknown', '未知', '3', NULL, 53, 'sex', 93, 1, 1, 2, 1, '2019-02-20 23:40:29', 1, '2019-02-20 23:36:46', 1);

-- ----------------------------
-- Table structure for sys_department
-- ----------------------------
DROP TABLE IF EXISTS `sys_department`;
CREATE TABLE `sys_department`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(11) NULL DEFAULT 0 COMMENT '上级机构',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '部门/11111',
  `code` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '机构编码',
  `sort` int(11) NULL DEFAULT 0 COMMENT '序号',
  `linkman` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系人',
  `linkman_no` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '联系人电话',
  `remark` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '机构描述',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_depart_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10015 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '组织机构' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_department
-- ----------------------------
INSERT INTO `sys_department` VALUES (10001, 0, 'FLY的狐狸', 'ABC000', 100, '', '', '', 1, '2017-04-28 01:16:43', 1, '2016-07-31 18:12:30', 1);
INSERT INTO `sys_department` VALUES (10002, 10001, '开发组', 'ABC001', 101, NULL, NULL, NULL, 1, '2016-07-31 18:15:29', 1, '2016-07-31 18:15:29', 1);
INSERT INTO `sys_department` VALUES (10003, 10001, '产品组', 'ABC003', 103, '', '', '', 1, '2017-04-28 00:58:41', 1, '2016-07-31 18:16:06', 1);
INSERT INTO `sys_department` VALUES (10004, 10001, '运营组', 'ABC004', 104, NULL, NULL, NULL, 1, '2016-07-31 18:16:30', 1, '2016-07-31 18:16:30', 1);
INSERT INTO `sys_department` VALUES (10005, 10001, '测试组', '12323', 10, '', '', '', 0, '2019-06-30 22:33:44', 1, '2017-10-18 18:13:09', 1);

-- ----------------------------
-- Table structure for sys_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_log`;
CREATE TABLE `sys_log`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `log_type` int(11) NOT NULL COMMENT '类型',
  `oper_object` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作对象',
  `oper_table` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '操作表',
  `oper_id` int(11) NULL DEFAULT 0 COMMENT '操作主键',
  `oper_type` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作类型',
  `oper_remark` varchar(2000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '操作备注',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11813 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '日志' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parentid` int(11) NOT NULL DEFAULT 0 COMMENT '父id',
  `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称/11111',
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `urlkey` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单key',
  `url` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '链接地址',
  `perms` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '授权(多个用逗号分隔，如：user:list,user:create)',
  `status` int(11) NULL DEFAULT 1 COMMENT '状态//radio/2,隐藏,1,显示',
  `type` int(11) NULL DEFAULT 1 COMMENT '类型//select/1,目录,2,菜单,3,按钮',
  `sort` int(11) NULL DEFAULT 1 COMMENT '排序',
  `level` int(11) NULL DEFAULT 1 COMMENT '级别',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '菜单' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, 20, '系统首页', 'fa fa-home', 'home', '/admin/welcome.html', '', 1, 2, 10, 2, 1, '2019-02-17 23:24:28', 1, '2015-04-27 17:28:06', 1);
INSERT INTO `sys_menu` VALUES (2, 0, '系统管理', 'fa fa-institution', 'system_root', NULL, NULL, 1, 1, 190, 1, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:28:06', 1);
INSERT INTO `sys_menu` VALUES (3, 2, '组织机构', 'fa fa-users', 'department', '/system/department/index', NULL, 1, 2, 191, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:28:25', 1);
INSERT INTO `sys_menu` VALUES (4, 2, '用户管理', 'fa fa-user-o', 'user', '/system/user/index', NULL, 1, 2, 192, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:28:46', 1);
INSERT INTO `sys_menu` VALUES (5, 2, '角色管理', 'fa fa-address-book-o', 'role', '/system/role/index', NULL, 1, 2, 194, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:29:13', 1);
INSERT INTO `sys_menu` VALUES (6, 2, '菜单管理', 'fa fa-bars', 'menu', '/system/menu/index', NULL, 1, 2, 196, 2015, 1, '1', 2, '2015-04-27 17:29:43', 1);
INSERT INTO `sys_menu` VALUES (8, 2, '参数配置', 'fa fa-file-text-o', 'config', '/system/config/index', '', 1, 2, 198, 2, 1, '2017-09-15 14:53:36', 1, '2016-12-17 23:34:13', 1);
INSERT INTO `sys_menu` VALUES (9, 2, '日志管理', 'fa fa-line-chart', 'log', '/system/log/index', NULL, 1, 2, 199, 2, 1, '2015-04-27 17:28:06', 1, '2016-01-03 18:09:18', 1);
INSERT INTO `sys_menu` VALUES (20, 0, '业务处理', 'fa fa-home', 'home', '', '', 1, 1, 10, 1, 1, '2019-02-17 23:24:08', 1, '2019-02-17 23:24:08', 1);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称/11111/',
  `status` int(11) NULL DEFAULT 1 COMMENT '状态//radio/2,隐藏,1,显示',
  `sort` int(11) NULL DEFAULT 1 COMMENT '排序',
  `remark` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '说明//textarea',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '测试角色', 1, 10, '', 1, '2019-07-03 00:55:45', 1, '2017-09-15 14:54:26', 1);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  `menu_id` bigint(20) NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 50 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色和菜单关联' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
INSERT INTO `sys_role_menu` VALUES (48, 1, 20);
INSERT INTO `sys_role_menu` VALUES (49, 1, 1);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'UUID',
  `username` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '登录名/11111',
  `password` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `salt` varchar(16) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '1111' COMMENT '密码盐',
  `real_name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '真实姓名',
  `depart_id` int(11) NULL DEFAULT 0 COMMENT '部门/11111/dict',
  `user_type` int(11) NULL DEFAULT 2 COMMENT '类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户',
  `status` int(11) NULL DEFAULT 10 COMMENT '状态',
  `thirdid` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '第三方ID',
  `endtime` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '结束时间',
  `email` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT 'email',
  `tel` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '手机号',
  `address` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '地址',
  `title_url` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像地址',
  `remark` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '说明',
  `theme` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'default' COMMENT '主题',
  `back_site_id` int(11) NULL DEFAULT 0 COMMENT '后台选择站点ID',
  `create_site_id` int(11) NULL DEFAULT 1 COMMENT '创建站点ID',
  `project_id` bigint(20) NULL DEFAULT 0 COMMENT '项目ID',
  `project_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '项目名称',
  `enable` tinyint(1) NULL DEFAULT 1 COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) NULL DEFAULT 0 COMMENT '更新人',
  `create_time` varchar(24) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) NULL DEFAULT 0 COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_user_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, '94091b1fa6ac4a27a06c0b92155aea6a', 'admin', '9fb3dc842c899aa63d6944a55080b795', '1111', '系统管理员', 10001, 1, 10, '', '', 'zcool321@sina.com', '123', '', '', '时间是最好的老师，但遗憾的是&mdash;&mdash;最后他把所有的学生都弄死了', 'flat', 5, 1, 1, 'test', 1, '2019-07-08 18:12:28', 1, '2017-03-19 20:41:25', 1);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `role_id` bigint(20) NOT NULL COMMENT '角色id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户和角色关联' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, 1, 1);

SET FOREIGN_KEY_CHECKS = 1;
