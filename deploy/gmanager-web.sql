/** 前后端分离菜单 **/
DELETE FROM `sys_menu`;
INSERT INTO `sys_menu` VALUES (1, 20, '系统首页', 'welcome', 'home', '/welcome', '', 1, 2, 10, 2, 1, '2019-12-06 10:37:44', 1, '2015-04-27 17:28:06', 1);
INSERT INTO `sys_menu` VALUES (2, 0, '系统管理', 'settings', 'system_root', NULL, NULL, 1, 1, 190, 1, 1, '2019-12-04 14:18:23', 1, '2015-04-27 17:28:06', 1);
INSERT INTO `sys_menu` VALUES (3, 2, '组织机构', 'depart', 'department', '/system/department/index', NULL, 1, 2, 191, 2, 1, '2019-12-04 14:17:51', 1, '2015-04-27 17:28:25', 1);
INSERT INTO `sys_menu` VALUES (4, 2, '用户管理', 'user_1', 'user', '/system/user/index', NULL, 1, 2, 192, 2, 1, '2019-12-04 14:14:06', 1, '2015-04-27 17:28:46', 1);
INSERT INTO `sys_menu` VALUES (5, 2, '角色管理', 'role', 'role', '/system/role/index', NULL, 1, 2, 194, 2, 1, '2019-12-04 14:14:12', 1, '2015-04-27 17:29:13', 1);
INSERT INTO `sys_menu` VALUES (6, 2, '菜单管理', 'menu', 'menu', '/system/menu/index', NULL, 1, 2, 196, 2, 1, '2019-12-04 14:14:33', 1, '2015-04-27 17:29:43', 1);
INSERT INTO `sys_menu` VALUES (20, 0, '业务处理', 'business', 'home', '', '', 1, 1, 10, 1, 1, '2019-12-04 14:14:55', 1, '2019-02-17 23:24:08', 1);
INSERT INTO `sys_menu` VALUES (37, 2, '参数配置', 'config', NULL, 'system/config/index', NULL, 1, 2, 198, 2, 1, '2019-12-12 15:31:40', 1, '2019-12-10 14:51:29', 1);
INSERT INTO `sys_menu` VALUES (38, 2, '日志管理', 'log', 'log', 'system/log/index', NULL, 1, 2, 199, 2, 1, '2019-12-12 15:31:24', 1, '2019-12-10 14:55:22', 1);

/**
DELETE FROM `sys_menu`;
INSERT INTO `sys_menu` VALUES (1, 20, '系统首页', 'fa fa-home', 'home', '/admin/welcome.html', '', 1, 2, 10, 2, 1, '2019-02-17 23:24:28', 1, '2015-04-27 17:28:06', 1);
INSERT INTO `sys_menu` VALUES (2, 0, '系统管理', 'fa fa-institution', 'system_root', NULL, NULL, 1, 1, 190, 1, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:28:06', 1);
INSERT INTO `sys_menu` VALUES (3, 2, '组织机构', 'fa fa-users', 'department', '/system/department/index', NULL, 1, 2, 191, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:28:25', 1);
INSERT INTO `sys_menu` VALUES (4, 2, '用户管理', 'fa fa-user-o', 'user', '/system/user/index', NULL, 1, 2, 192, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:28:46', 1);
INSERT INTO `sys_menu` VALUES (5, 2, '角色管理', 'fa fa-address-book-o', 'role', '/system/role/index', NULL, 1, 2, 194, 2, 1, '2015-04-27 17:28:06', 1, '2015-04-27 17:29:13', 1);
INSERT INTO `sys_menu` VALUES (6, 2, '菜单管理', 'fa fa-bars', 'menu', '/system/menu/index', NULL, 1, 2, 196, 2, 1, '2015-04-27 17:29:43', 1, '2015-04-27 17:29:43', 1);
INSERT INTO `sys_menu` VALUES (8, 2, '参数配置', 'fa fa-file-text-o', 'config', '/system/config/index', '', 1, 2, 198, 2, 1, '2017-09-15 14:53:36', 1, '2016-12-17 23:34:13', 1);
INSERT INTO `sys_menu` VALUES (9, 2, '日志管理', 'fa fa-line-chart', 'log', '/system/log/index', NULL, 1, 2, 199, 2, 1, '2015-04-27 17:28:06', 1, '2016-01-03 18:09:18', 1);
INSERT INTO `sys_menu` VALUES (20, 0, '业务处理', 'fa fa-home', 'home', '', '', 1, 1, 10, 1, 1, '2019-02-17 23:24:08', 1, '2019-02-17 23:24:08', 1);
 */