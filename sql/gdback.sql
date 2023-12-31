CREATE DATABASE if NOT EXISTS gdback;
USE gdback;

CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account` varchar(128) NOT NULL DEFAULT '' COMMENT '账户',
  `password` varchar(128) NOT NULL DEFAULT '' COMMENT '密码',
  `createtime` timestamp NULL default CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `IDX_id` (`id`),
  KEY `IDX_account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS `register` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint NOT NULL COMMENT '用户id',
  `remote` varchar(40) NOT NULL,
  `ip` varchar(40) NOT NULL,
  `imei` varchar(128) NOT NULL,
  `os` varchar(20) NOT NULL,
  `model` varchar(20) NOT NULL,
  `app_id` varchar(32) NOT NULL,
  `channel_id` varchar(32) NOT NULL,
  `register_at` bigint NOT NULL,
  `register_type` tinyint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `IDX_register_register_at` (`register_at`),
  KEY `IDX_register_register_type` (`register_type`),
  KEY `IDX_register_uid` (`uid`),
  KEY `IDX_register_app_id` (`app_id`),
  KEY `IDX_register_channel_id` (`channel_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

