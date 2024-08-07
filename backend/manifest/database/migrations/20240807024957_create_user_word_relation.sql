-- Create "user_word" table
CREATE TABLE `user_word` (`user_uid` varchar(10) NOT NULL COMMENT "User Unique ID", `word_id` int unsigned NOT NULL COMMENT "Word ID", `created_at` datetime NULL COMMENT "Created Time", `deleted_at` datetime NULL COMMENT "Soft deleted Time", PRIMARY KEY (`user_uid`, `word_id`), INDEX `word_id` (`word_id`), CONSTRAINT `user_word_ibfk_1` FOREIGN KEY (`user_uid`) REFERENCES `users` (`uid`) ON UPDATE RESTRICT ON DELETE CASCADE, CONSTRAINT `user_word_ibfk_2` FOREIGN KEY (`word_id`) REFERENCES `words` (`id`) ON UPDATE RESTRICT ON DELETE CASCADE) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;