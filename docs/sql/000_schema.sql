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
