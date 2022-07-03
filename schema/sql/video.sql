CREATE TABLE `video`
(
    `id`             bigint unsigned NOT NULL AUTO_INCREMENT,
    `title`          varchar(255) NOT NULL,
    `file_url`       varchar(255) NOT NULL COMMENT 'file url in OSS',
    `cover_url`      varchar(255) NOT NULL DEFAULT '',
    `user_id`        bigint unsigned NOT NULL COMMENT 'author id',
    `create_time`    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `favorite_count` bigint NOT NULL DEFAULT 0,
    `comment_count`  bigint NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX idx_title (`title`),
    INDEX idx_user_id (`user_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
