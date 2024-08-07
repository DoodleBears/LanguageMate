-- Create "words" table
CREATE TABLE `words` (`id` int unsigned NOT NULL AUTO_INCREMENT COMMENT "Word ID", `content` varchar(45) NOT NULL COMMENT "Word content", `created_at` datetime NULL COMMENT "Created Time", PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
