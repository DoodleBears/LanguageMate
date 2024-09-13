-- Modify "sentences" table
ALTER TABLE `sentences` ADD COLUMN `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT "Sentence ID", DROP PRIMARY KEY, ADD PRIMARY KEY (`id`), ADD INDEX `owner_idx` (`user_uid`, `word_id`);
