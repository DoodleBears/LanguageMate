CREATE TABLE IF NOT EXISTS `users`
(
    `uid` varchar(10) NOT NULL UNIQUE COMMENT 'User Unique ID',
    `avatar_url`  varchar(45) NOT NULL COMMENT 'User Avatar',
    `username`  varchar(25) NOT NULL UNIQUE COMMENT 'User Name',
    `email`  varchar(45) NOT NULL UNIQUE COMMENT 'User Email',
    `email_verified`  bit(1) DEFAULT 0  COMMENT 'User Email Verification Flag',
    `password`  varchar(45) NOT NULL COMMENT 'User Password',
    `nickname`  varchar(25) NOT NULL COMMENT 'User Nickname',
    `created_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `updated_at` datetime DEFAULT NULL COMMENT 'Updated Time',
    `deleted_at` datetime DEFAULT NULL COMMENT 'Soft deleted Time',
    `last_login` datetime DEFAULT NULL COMMENT 'Last Login Time',
    `login_attempts`  int(2) DEFAULT 0 COMMENT 'User Login Attempts',
    `lock`  bit(1) DEFAULT 0 COMMENT 'User Login Lock',
    `banned`  bit(1) DEFAULT 0 COMMENT 'User Banned control',
    PRIMARY KEY (`uid`),
    INDEX `idx_username` (`username`),
    INDEX `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;