
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(191) DEFAULT '' COMMENT '用户名',
  `nickname` varchar(191) DEFAULT '' COMMENT '昵称',
  `avatar` varchar(191) DEFAULT '' COMMENT '头像',
  `level` bigint DEFAULT '1' COMMENT '等级',
  `exp` bigint DEFAULT '0' COMMENT '经验值',
  `gold` bigint DEFAULT '0' COMMENT '金币',
  `diamond` bigint DEFAULT '0' COMMENT '钻石',
  `land_level` bigint DEFAULT '1' COMMENT '土地等级',
  `login_at` datetime(3) DEFAULT NULL COMMENT '登录时间',
  `extra` json NOT NULL COMMENT '额外信息',

  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_user_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
