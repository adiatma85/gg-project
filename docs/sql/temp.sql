DROP DATABASE IF EXISTS gg_project;\n
CREATE DATABASE IF NOT EXISTS gg_project;\n
USE gg_project;\n
SET foreign_key_checks = 0;\n
-- Create Table Users
DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `email` VARCHAR(255) NOT NULL DEFAULT '',
  `username` VARCHAR(255) NOT NULL DEFAULT '',
  `password` VARCHAR(255) NOT NULL DEFAULT '',
  `display_name` VARCHAR(255) NOT NULL DEFAULT '',

  -- Utility columns
  `status` SMALLINT NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` VARCHAR(255),
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` VARCHAR(255),
  `deleted_at`TIMESTAMP,
  `deleted_by` VARCHAR(255),
  PRIMARY KEY (`id`),
  UNIQUE (`username`)
) ENGINE = INNODB COMMENT='User table';

-- Create Task Table
DROP TABLE IF EXISTS `task`;
CREATE TABLE IF NOT EXISTS `task` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `fk_user_id` INT COMMENT 'Foreign Key To User Id',
  `title` VARCHAR(255) NOT NULL DEFAULT '',
  `priority` INT NOT NULL DEFAULT 1,
  `task_status` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'todo, ongoing, done',
  `periodic` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'none, daily, weekly, monthly, yearly',
  `due_time` TIMESTAMP,

  -- Utility columns
  `status` SMALLINT NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` VARCHAR(255),
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` VARCHAR(255),
  `deleted_at`TIMESTAMP,
  `deleted_by` VARCHAR(255),
  PRIMARY KEY (`id`)
) ENGINE = INNODB COMMENT='Task Table';

-- Create Category Table
DROP TABLE IF EXISTS `category`;
CREATE TABLE IF NOT EXISTS `category` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL DEFAULT '',

  -- Utility columns
  `status` SMALLINT NOT NULL DEFAULT '1',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_by` VARCHAR(255),
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` VARCHAR(255),
  `deleted_at`TIMESTAMP,
  `deleted_by` VARCHAR(255),
  PRIMARY KEY (`id`)
) ENGINE = INNODB COMMENT='Category Table';
-- [DML] Insert Dummy Data for User and for Category Tables
INSERT INTO `user` (`email`, `username`, `password`, `display_name`)
VALUES 
('adiatma85@gmail.com', 'adiatma85', '$2a$10$hEU84tig1.W0TcoSe5.zwushMbDarsTnaXadC5/Y/difKiatAHuGO', 'Luki');

INSERT INTO `category` (`name`)
VALUES 
('Work'),
('Hobby'),
('School');
-- [DDL] Create new column to accomodate task
ALTER TABLE `task` ADD `fk_category_id` INT COMMENT 'Foreign Key To User Id' AFTER `fk_user_id`;
-- [DDL] Create new table for Role
DROP TABLE IF EXISTS `role`;
CREATE TABLE IF NOT EXISTS `role` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT '',
    `type` VARCHAR(255) NOT NULL DEFAULT '',
    `rank` INT NOT NULL DEFAULT 0,

    -- Utility columns
    `status` SMALLINT NOT NULL DEFAULT '1',
    `flag` INT NOT NULL DEFAULT '0',
    `meta` VARCHAR(255),
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `created_by` VARCHAR(255),
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `updated_by` VARCHAR(255),
    `deleted_at`TIMESTAMP,
    `deleted_by` VARCHAR(255),
    PRIMARY KEY (`id`)
);

-- [DDL] Add new column `fk_role_id` in user table
ALTER TABLE `user` ADD `fk_role_id` INT COMMENT 'Foreign Key To role Id' AFTER `id`;


-- [DML] Populate role admin and user
INSERT INTO `role` (`name`, `type`, `rank`) VALUES
('Super Admin', 'admin', 1),
('User', 'user', 2)
;
