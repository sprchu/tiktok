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

CREATE TABLE `relation`
(
    `id`        bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`   bigint unsigned NOT NULL,
    `follow_id` bigint unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`user_id`, `follow_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user`(`id`),
    FOREIGN KEY (`follow_id`) REFERENCES `user`(`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `favorite`
(
    `id`       bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`  bigint unsigned NOT NULL,
    `video_id` bigint unsigned NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY (`user_id`, `video_id`),
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
    FOREIGN KEY (`video_id`) REFERENCES `video` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `comment`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     bigint unsigned NOT NULL,
    `video_id`    bigint unsigned NOT NULL,
    `content`     varchar(255) NOT NULL DEFAULT "",
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
    FOREIGN KEY (`video_id`) REFERENCES `video` (`id`),
    INDEX idx_user_id (`user_id`),
    INDEX idx_video_id (`video_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
