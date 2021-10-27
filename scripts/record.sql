CREATE TABLE `record` (
  `record_id` int(10) unsigned DEFAULT 0,
  `id` int(10) unsigned DEFAULT 0,
  `program_id` int(10) unsigned DEFAULT 0,
  `status` varchar(50) DEFAULT '' COMMENT '提交状态',
  PRIMARY KEY (`record_id`),
  foreign key(id) references database(id),
  foreign key(program_id) references program(program_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='题目提交状态';