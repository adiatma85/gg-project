-- [DDL] Create new column to accomodate task
ALTER TABLE `task` ADD `fk_category_id` INT COMMENT 'Foreign Key To User Id' AFTER `fk_user_id`;
