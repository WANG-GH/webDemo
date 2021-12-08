CREATE TABLE `email` (
  `record_id` int(10) unsigned PRIMARY KEY AUTO_INCREMENT,
  `email_num` varchar(50) DEFAULT '',
  `verify_code` int(10) unsigned
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='邮箱'; 