CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_name`  varchar(128) NOT NULL DEFAULT '' COMMENT 'UserName',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `follow_count` bigint unsigned DEFAULT 0 COMMENT 'follow_count',
    `follower_count` bigint unsigned DEFAULT 0 COMMENT 'follower_count',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_name` (`user_name`) COMMENT 'UserName index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

CREATE TABLE `video`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id`  bigint unsigned NOT NULL COMMENT 'user_id',
    `title`   varchar(128) NOT NULL DEFAULT '' COMMENT 'title',
    `play_url`   varchar(1000) NOT NULL DEFAULT '' COMMENT 'play_url',
    `cover_url`   varchar(1000) NOT NULL DEFAULT '' COMMENT 'cover_url',
    `favorite_count`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'favorite_count',
    `comment_count`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'comment_count',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_title` (`title`) COMMENT 'title index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='video table';

CREATE TABLE `favorite`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id` bigint unsigned DEFAULT 1 NOT NULL,
    `video_id` bigint unsigned DEFAULT 1 NOT NULL,
    `status` int unsigned default 0 NOT NULL ,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Favorite record create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Favorite record update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Favorite record delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`user_id`) COMMENT 'user_id index'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='favorite table';

CREATE TABLE `follow`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id` bigint unsigned DEFAULT 1 NOT NULL,
    `follow_user_id` bigint unsigned DEFAULT 1 NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Follow record create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Follow record update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Follow record delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`user_id`) COMMENT 'user_id index'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='follow table';

CREATE TABLE `comment`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id` bigint unsigned DEFAULT 1 NOT NULL,
    `video_id` bigint unsigned DEFAULT 1 NOT NULL,
    `content` text DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'comment record create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'comment record update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'comment record delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`user_id`) COMMENT 'user_id index',
    KEY          `idx_video_id` (`video_id`) COMMENT 'video_id index'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='comment table';

CREATE TABLE `message`
(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `user_id` bigint unsigned DEFAULT 1 NOT NULL,
    `to_user_id` bigint unsigned DEFAULT 1 NOT NULL,
    `content` text DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'message record create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'message record update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'message record delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`user_id`) COMMENT 'user_id index',
    KEY          `idx_to_user_id` (`to_user_id`) COMMENT 'to_user_id index'
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='chat table';