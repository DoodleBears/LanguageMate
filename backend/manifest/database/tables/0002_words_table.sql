CREATE TABLE IF NOT EXISTS `words`
(
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Word ID',
    `content`  varchar(45) NOT NULL COMMENT 'Word content',
    `created_at` datetime DEFAULT NULL COMMENT 'Created Time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_unicode_ci;