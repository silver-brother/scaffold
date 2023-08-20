CREATE TABLE `user_access_log`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` timestamp       NULL     DEFAULT NULL,
    `user_id`    bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `ip`         varchar(255)    NOT NULL DEFAULT '' COMMENT 'ip地址',
    `path`       varchar(255)    NOT NULL DEFAULT '' COMMENT '访问路径',
    `method`     varchar(255)    NOT NULL DEFAULT '' COMMENT '请求方式',
    `params`     varchar(32)     NOT NULL DEFAULT '' COMMENT '请求参数',
    `body`       text not null default '' COMMENT '请求体',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 0
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci COMMENT ='用户访问日志表';