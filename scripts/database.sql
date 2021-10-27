-- CREATE TABLE 'user'(
-- 	`id` int unsigned NOT NULL AUTO_INCREMENT,
-- 	`name` varchar(100) DEFAULT '' COMMENT '用户姓名',
-- 	`password` varchar(100) DEFAULT '' COMMENT '用户密码',
-- 	`privilege` int DEFAULT 0 COMMENT '特权级，1管理员，0普通用户',
-- 	PRIMARY KEY (`name`)
-- ) ENGINE=InnoDB DEFALUT CHARSET=utf8mb4 COMMENT='用户管理'
-- `

CREATE TABLE `user` (
  `id` int(10) unsigned  PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '用户姓名',
  `password` varchar(100) DEFAULT '' COMMENT '用户密码',
  `privilege` int DEFAULT 0 COMMENT '特权级，1管理员，0普通用户',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户管理';
