-- 用户
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password_hash` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci COMMENT '邮箱',
    `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 问题
DROP TABLE IF EXISTS `question`;
CREATE TABLE `question` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `question_id` bigint(20) NOT NULL,
    `caption` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL,
    `author_id` bigint(20) NOT NULL,
    `category_id` bigint(20) NOT NULL,
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_question_id` (`question_id`),
    KEY `idx_category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 答案
DROP TABLE IF EXISTS `answer`;
CREATE TABLE `answer` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `answer_id` bigint(20) unsigned NOT NULL,
    `content` text COLLATE utf8mb4_general_ci NOT NULL,
    `author_id` bigint(20) NOT NULL,
    `question_id` bigint(20) NOT NULL,
    `vote_up_count` int(11) NOT NULL DEFAULT '0',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_answer_id` (`answer_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_question_id` (`question_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 分类
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
    `id` int(10) NOT NULL AUTO_INCREMENT,
    `category_id` int(10) unsigned NOT NULL,
    `category_name` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_category_id` (`category_id`),
    UNIQUE KEY `idx_category_name` (`category_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;