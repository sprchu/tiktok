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
