CREATE TABLE `user`
(
    `id`          bigint unsigned    NOT NULL AUTO_INCREMENT,
    `created_at`  timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp          NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  timestamp          NULL     DEFAULT NULL,
    `name`        varchar(255)       NOT NULL DEFAULT '' COMMENT '姓名',
    `sex`         enum ('F','M','O') NOT NULL DEFAULT 'O' COMMENT '性别 1男 2女 3未知',
    `birth_date`  varchar(64)        NOT NULL DEFAULT '' COMMENT '生日',
    `id_card`     varchar(64)        NOT NULL DEFAULT '' COMMENT '身份证',
    `mobile`      varchar(32)        NOT NULL DEFAULT '' COMMENT '手机号',
    `avatar`      varchar(255)       NOT NULL DEFAULT '照片',
    `description` varchar(255)       NOT NULL DEFAULT '' COMMENT '简介',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_id_card` (`id_card`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 0
  DEFAULT CHARSET = utf8 COMMENT ='用户信息表';