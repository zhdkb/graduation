DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE
            CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        UNIQUE INDEX `idx_username` (`username`) USING BTREE,
        UNIQUE INDEX `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- DROP TABLE IF EXISTS `emotional`;
-- CREATE TABLE `emotional` (
--     `id` bigint(20) NOT NULL AUTO_INCREMENT,
--     `user_id` bigint(20) NOT NULL COMMENT '用户id',
--     `text` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '情绪文本',
--     `msg` varchar(64) COLLATE utf8mb4_general_ci NOT NULL  COMMENT '情绪描述',
--     `data` varchar(255) COLLATE utf8mb4_general_ci NOT NULL  COMMENT '情绪具体说明',
--     `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
--     `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE
--             CURRENT_TIMESTAMP   COMMENT '更新时间',
--         PRIMARY KEY (`id`),
--         UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `emotional_infos`;
CREATE TABLE emotional_infos (
    id BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '自增主键',
    user_id BIGINT NOT NULL COMMENT '用户id',
    emotional_text VARCHAR(255) NOT NULL COMMENT "情绪文本",
    emotional_type INT NOT NULL COMMENT '情绪类型',
    emotional_msg VARCHAR(255) COMMENT '情绪描述',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    modify_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_user_id (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


-- 创建打卡时间关系表
DROP TABLE IF EXISTS `check_in_records`;
CREATE TABLE check_in_records (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    check_in_time DATETIME NOT NULL COMMENT '打卡时间',
    streak_days INT DEFAULT 1 COMMENT '连续打卡天数',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    UNIQUE INDEX idx_user_time (user_id, check_in_time)  -- 联合索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 创建每个用户的打卡统计表
DROP TABLE IF EXISTS `user_check_in_counts`;
CREATE TABLE user_check_in_counts (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '自增主键',
    user_id BIGINT NOT NULL COMMENT '用户ID',
    check_in_count INT DEFAULT 0 COMMENT '打卡次数',
    max_streak_days INT DEFAULT 0 COMMENT '最大连续打卡天数',
    create_time DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    INDEX idx_user_id (user_id)  -- 为 user_id 字段创建索引
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;