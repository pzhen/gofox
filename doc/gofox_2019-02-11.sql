

DROP TABLE IF EXISTS `sys_log`;

CREATE TABLE `sys_log` (
  `log_id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(512) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '操作地址',
  `urlfor` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `user_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '用户名称',
  `form_data` text COLLATE utf8_unicode_ci COMMENT '操作数据',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '操作时间',
  PRIMARY KEY (`log_id`),
  KEY `user_id` (`user_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='操作日志表';

LOCK TABLES `sys_log` WRITE;
/*!40000 ALTER TABLE `sys_log` DISABLE KEYS */;

INSERT INTO `sys_log` (`log_id`, `url`, `urlfor`, `user_id`, `user_name`, `form_data`, `create_time`)
VALUES
	(1,'/sys_menu/form_modify_sysmenu?menu_id=37',NULL,1,'abc123','{\"menu_id\":[\"37\"]}',0),
	(2,'/sys_menu/save_sysmenu',NULL,1,'abc123','{\"Id\":[\"37\"],\"MenuFuncs\":[\"\"],\"MenuIcon\":[\"\\u0026#xe6ce;\"],\"MenuLevel\":[\"1\"],\"MenuName\":[\"统计管理\"],\"MenuRootid\":[\"0\"],\"MenuStatus\":[\"1\"],\"MenuUrl\":[\"\"]}',0),
	(3,'/sys_menu/list_sysmenu',NULL,1,'abc123','{}',0),
	(4,'/sys_menu/form_modify_sysmenu?menu_id=3','SysMenuController.FormModifySysMenu',1,'abc123','{\"menu_id\":[\"3\"]}',0),
	(5,'/sys_menu/list_sysmenu','SysMenuController.ListSysMenu',1,'abc123','{}',0),
	(6,'/sys_menu/list_sysmenu','SysMenuController.ListSysMenu',1,'abc123','{}',1531476255),
	(7,'/sys_role/list_sysrole','SysRoleController.GetSysRoleListByPage',1,'abc123','{}',1531476266),
	(8,'/sys_role/form_modify_sysrole?role_id=121','SysRoleController.FormModifySysRole',1,'abc123','{\"role_id\":[\"121\"]}',1531476275),
	(9,'/sys_role/save_sysrole','SysRoleController.SaveSysRole',1,'abc123','{\"Intro\":[\"qwwww\"],\"MenuMap\":[\"1,2,2-1,2-4,3,3-2,3-3,21,22,1,\"],\"RoleId\":[\"121\"],\"RoleName\":[\"测试123\"],\"RoleStatus\":[\"1\"]}',1531476282),
	(10,'/sys_role/list_sysrole','SysRoleController.GetSysRoleListByPage',1,'abc123','{}',1531476284);

/*!40000 ALTER TABLE `sys_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `menu_id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_rootid` int(11) DEFAULT '0' COMMENT '上级id',
  `menu_name` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '菜单名称',
  `menu_funcs` varchar(1024) COLLATE utf8_unicode_ci DEFAULT '',
  `menu_url` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '所属类',
  `menu_icon` varchar(50) COLLATE utf8_unicode_ci NOT NULL COMMENT '图标',
  `menu_lock` tinyint(4) NOT NULL COMMENT '锁定',
  `menu_status` tinyint(4) NOT NULL COMMENT '状态',
  `menu_level` tinyint(4) unsigned DEFAULT '0',
  `menu_path` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`menu_id`),
  KEY `menu_status` (`menu_status`),
  KEY `menu_rootid` (`menu_rootid`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='菜单表';

LOCK TABLES `sys_menu` WRITE;
/*!40000 ALTER TABLE `sys_menu` DISABLE KEYS */;

INSERT INTO `sys_menu` (`menu_id`, `menu_rootid`, `menu_name`, `menu_funcs`, `menu_url`, `menu_icon`, `menu_lock`, `menu_status`, `menu_level`, `menu_path`)
VALUES
	(1,0,'系统管理','','','&#xe6ae;',0,1,1,'-1-'),
	(2,1,'菜单列表','[\r\n  {\r\n    \"func_id\": 1,\r\n    \"func_name\": \"GetSysMenuList\",\r\n    \"func_desc\": \"菜单列表\"\r\n  },\r\n  {\r\n    \"func_id\": 2,\r\n    \"func_name\": \"FormSysMenu\",\r\n    \"func_desc\": \"菜单表单\"\r\n  },\r\n  {\r\n    \"func_id\": 3,\r\n    \"func_name\": \"SaveSysMenu\",\r\n    \"func_desc\": \"菜单保存\"\r\n  },\r\n  {\r\n    \"func_id\": 4,\r\n    \"func_name\": \"ModifySysMenuStatus\",\r\n    \"func_desc\": \"菜单修改状态\"\r\n  },\r\n  {\r\n    \"func_id\": 5,\r\n    \"func_name\": \"DeleteSysMenu\",\r\n    \"func_desc\": \"菜单删除\"\r\n  }\r\n]','SysMenuController.GetSysMenuList','',0,1,2,'-1-2-'),
	(3,1,'角色列表','[\r\n  {\r\n    \"func_id\": 1,\r\n    \"func_name\": \"GetSysRoleListByPage\",\r\n    \"func_desc\": \"角色列表\"\r\n  },\r\n  {\r\n    \"func_id\": 2,\r\n    \"func_name\": \"FormSysRole\",\r\n    \"func_desc\": \"角色表单\"\r\n  },\r\n  {\r\n    \"func_id\": 3,\r\n    \"func_name\": \"SaveSysRole\",\r\n    \"func_desc\": \"角色保存\"\r\n  },\r\n  {\r\n    \"func_id\": 4,\r\n    \"func_name\": \"ModifySysRoleStatus\",\r\n    \"func_desc\": \"角色修改状态\"\r\n  },\r\n  {\r\n    \"func_id\": 5,\r\n    \"func_name\": \"DeleteSysRole\",\r\n    \"func_desc\": \"角色删除\"\r\n  }\r\n]','SysRoleController.GetSysRoleListByPage','',0,1,2,'-1-3-'),
	(37,0,'统计管理','','','&#xe6ce;',0,0,1,'-37-'),
	(36,35,'商品列表','','','',0,0,2,'-35-36-'),
	(32,1,'操作日志','[\r\n  {\r\n    \"func_id\": 1,\r\n    \"func_name\": \"GetSysLogListByPage\",\r\n    \"func_desc\": \"日志列表\"\r\n  },\r\n  {\r\n    \"func_id\": 2,\r\n    \"func_name\": \"DeleteSysLog\",\r\n    \"func_desc\": \"日志删除\"\r\n  }\r\n]','SysLogController.GetSysLogListByPage','',0,1,2,'-1-32-'),
	(33,0,'广告管理','','','&#xe6fc;',0,0,1,'-33-'),
	(34,33,'广告列表','','','',0,0,2,'-33-34-'),
	(35,0,'商品管理','','','&#xe728;',0,0,1,'-35-'),
	(21,0,'会员管理','','','&#xe6b8;',0,0,1,'-21-'),
	(22,21,'会员列表','','','',0,0,2,'-21-22-'),
	(23,0,'订单管理','','','&#xe723;',0,0,1,'-23-'),
	(26,1,'用户列表','[\r\n  {\r\n    \"func_id\": 1,\r\n    \"func_name\": \"GetSysUserListByPage\",\r\n    \"func_desc\": \"用户列表\"\r\n  },\r\n  {\r\n    \"func_id\": 2,\r\n    \"func_name\": \"FormSysUser\",\r\n    \"func_desc\": \"用户表单\"\r\n  },\r\n  {\r\n    \"func_id\": 3,\r\n    \"func_name\": \"SaveSysUser\",\r\n    \"func_desc\": \"用户保存\"\r\n  },\r\n  {\r\n    \"func_id\": 4,\r\n    \"func_name\": \"ModifySysUserStatus\",\r\n    \"func_desc\": \"用户修改状态\"\r\n  },\r\n  {\r\n    \"func_id\": 5,\r\n    \"func_name\": \"DeleteSysUser\",\r\n    \"func_desc\": \"用户删除\"\r\n  }\r\n]','SysUserController.GetSysUserListByPage','',0,1,2,'-1-26-'),
	(27,0,'文章管理','','','&#xe705;',0,0,1,'-27-'),
	(28,27,'文章列表','','','',0,0,2,'-27-28-'),
	(29,21,'评论列表','','','',0,0,2,'-21-29-'),
	(30,0,'分类管理','','','&#xe699;',0,0,1,'-30-'),
	(31,30,'分类列表','','','',0,0,2,'-30-31-');

/*!40000 ALTER TABLE `sys_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `role_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(60) COLLATE utf8_unicode_ci NOT NULL COMMENT '角色名称',
  `intro` text COLLATE utf8_unicode_ci COMMENT '角色介绍',
  `role_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`role_id`),
  KEY `role_status` (`role_status`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='角色表';

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;

INSERT INTO `sys_role` (`role_id`, `role_name`, `intro`, `role_status`, `create_time`, `update_time`)
VALUES
	(152,'角色1','',1,1548406252,1548406325),
	(153,'角色2','',1,1548406471,1548406471);

/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role_menu_map
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role_menu_map`;

CREATE TABLE `sys_role_menu_map` (
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `menu_id` int(11) NOT NULL COMMENT '菜单id',
  `action_id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '操作权限',
  KEY `role_id` (`role_id`),
  KEY `menu_id` (`menu_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='权限表';

LOCK TABLES `sys_role_menu_map` WRITE;
/*!40000 ALTER TABLE `sys_role_menu_map` DISABLE KEYS */;

INSERT INTO `sys_role_menu_map` (`role_id`, `menu_id`, `action_id`)
VALUES
	(153,32,'2'),
	(153,32,'1'),
	(153,32,'0'),
	(153,2,'4'),
	(153,2,'1'),
	(153,2,'0'),
	(153,1,'0'),
	(152,2,'1'),
	(152,2,'0'),
	(152,1,'0'),
	(146,35,'0'),
	(146,32,'2'),
	(146,32,'1'),
	(146,32,'0'),
	(146,3,'5'),
	(146,3,'4'),
	(146,3,'3'),
	(146,3,'2'),
	(146,3,'1'),
	(146,3,'0'),
	(146,26,'0'),
	(146,2,'5'),
	(146,2,'4'),
	(146,2,'3'),
	(146,2,'2'),
	(146,2,'1'),
	(146,2,'0'),
	(146,1,'0'),
	(147,2,'5'),
	(147,2,'4'),
	(147,2,'3'),
	(147,2,'2'),
	(147,2,'1'),
	(147,2,'0'),
	(147,1,'0'),
	(146,36,'0'),
	(146,37,'0'),
	(146,1,'0');

/*!40000 ALTER TABLE `sys_role_menu_map` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `user_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `user_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT '登录名',
  `nick_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT '昵称',
  `role_id` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT '角色id',
  `photo` char(128) COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '头像',
  `password` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '密码',
  `salt` char(6) COLLATE utf8_unicode_ci NOT NULL COMMENT '密码盐值',
  `email` char(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `mobile` char(32) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0',
  `update_time` int(10) unsigned zerofill DEFAULT NULL,
  `last_time` int(10) unsigned NOT NULL DEFAULT '0',
  `last_ip` char(15) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `login_count` int(11) unsigned NOT NULL DEFAULT '0',
  `user_type` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `user_status` tinyint(4) unsigned NOT NULL DEFAULT '1',
  PRIMARY KEY (`user_id`),
  KEY `status` (`user_status`),
  KEY `email` (`email`),
  KEY `mobile` (`mobile`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户表';

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user` DISABLE KEYS */;

INSERT INTO `sys_user` (`user_id`, `user_name`, `nick_name`, `role_id`, `photo`, `password`, `salt`, `email`, `mobile`, `create_time`, `update_time`, `last_time`, `last_ip`, `login_count`, `user_type`, `user_status`)
VALUES
	(11,'user1','','153,152','','96e79218965eb72c92a549dd5a330112','','','',1548406275,1548406651,0,'',0,0,1),
	(1,'abc123','超级管理员','','','e99a18c428cb38d5f260853678922e03','','','',0,NULL,0,'',0,1,1);
