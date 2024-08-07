CREATE TABLE IF NOT EXISTS `user_word`
(
    `user_uid` varchar(10) NOT NULL COMMENT 'User Unique ID',
    `word_id` int(10) unsigned NOT NULL COMMENT 'Word ID',
    `created_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `deleted_at` datetime DEFAULT NULL COMMENT 'Soft deleted Time',
    PRIMARY KEY (`user_uid`, `word_id`),
    FOREIGN KEY (`user_uid`) REFERENCES `users`(`uid`) ON DELETE CASCADE,
    FOREIGN KEY (`word_id`) REFERENCES `words`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;