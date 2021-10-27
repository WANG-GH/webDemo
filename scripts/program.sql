CREATE TABLE `program` (
  `program_id` int(10) unsigned  PRIMARY KEY AUTO_INCREMENT,
  `program_name` varchar(100) DEFAULT '' COMMENT '题目名称',
  `content` varchar(500) DEFAULT '' COMMENT '内容',
  `ptype` varchar(50) DEFAULT '' COMMENT '题型',
  `answer` varchar(100) DEFAULT '' COMMENT '答案',
  `difficulty` varchar(50) DEFAULT '' COMMENT '难度',
  PRIMARY KEY (`program_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='题目管理';