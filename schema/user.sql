CREATE TABLE `user`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT,
    `username`       varchar(32) NOT NULL DEFAULT '',
    `password`       varchar(32) NOT NULL DEFAULT '',
    `follow_count`   bigint NOT NULL DEFAULT 0,
    `follower_count` bigint NOT NULL DEFAULT 0,
    `create_time`    timestamp DEFAULT CURRENT_TIMESTAMP,
    `update_time`    timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE `uni_username` (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
