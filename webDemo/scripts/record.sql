CREATE TABLE `record` (
  `record_id` int(10) unsigned PRIMARY KEY AUTO_INCREMENT,
  `user_id` int(10) unsigned,
  `program_id` int(10) unsigned,
  `status` varchar(50) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='提交记录'; 