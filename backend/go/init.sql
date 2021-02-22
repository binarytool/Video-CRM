-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema video-crm
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema video-crm
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `video-crm` DEFAULT CHARACTER SET utf8 ;
USE `video-crm` ;

-- -----------------------------------------------------
-- Table `video-crm`.`device`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `video-crm`.`device` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `hardware` VARCHAR(512) NULL COMMENT '512 symbol string to describe hw details such as screen resolution, cpu, platform...',
  `owner` VARCHAR(256) NULL COMMENT '256 symbol string has onwer id',
  `status` INT NULL COMMENT 'status is: active or not, and so on. Just int could be a bit mask\n\n',
  `create_date` DATETIME NULL COMMENT 'when this device was registered in system\n',
  `uptime` INT NULL COMMENT 'uptime in hours\n',
  `update_time` DATETIME NULL COMMENT 'when last message were received\n',
  `info` VARCHAR(128) NULL COMMENT 'serial numbers and other details\n',
  `token` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'Device description';


-- -----------------------------------------------------
-- Table `video-crm`.`timestamps`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `video-crm`.`timestamps` (
  `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` TIMESTAMP NULL);


-- -----------------------------------------------------
-- Table `video-crm`.`content`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `video-crm`.`content` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `path` VARCHAR(45) NULL COMMENT 'path where file is located\n',
  `caption` VARCHAR(45) NULL,
  `size` INT NULL COMMENT 'size in bytes',
  `info` VARCHAR(45) NULL COMMENT 'general info about content',
  `mediainfo` VARCHAR(45) NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'Content description table';


-- -----------------------------------------------------
-- Table `video-crm`.`playlist`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `video-crm`.`playlist` (
  `device_id` INT NOT NULL,
  `content_id` INT NOT NULL,
  PRIMARY KEY (`device_id`, `content_id`),
  INDEX `fk_playlist_content1_idx` (`content_id` ASC),
  CONSTRAINT `fk_playlist_device`
    FOREIGN KEY (`device_id`)
    REFERENCES `video-crm`.`device` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_playlist_content1`
    FOREIGN KEY (`content_id`)
    REFERENCES `video-crm`.`content` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'relation between device and content which should be displayed on it.';


-- -----------------------------------------------------
-- Table `video-crm`.`stat`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `video-crm`.`stat` (
  `device_id` INT NOT NULL,
  `date` DATETIME NOT NULL,
  `event` VARCHAR(45) NULL,
  `value` VARCHAR(255) NULL,
  PRIMARY KEY (`device_id`),
  CONSTRAINT `fk_stat_device1`
    FOREIGN KEY (`device_id`)
    REFERENCES `video-crm`.`device` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'stat table which handles device activity';


-- -----------------------------------------------------
-- Table `video-crm`.`strategy`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `video-crm`.`strategy` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `caption` VARCHAR(45) NULL,
  `content_id` INT NOT NULL,
  `status` VARCHAR(45) NULL,
  `show` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_strategy_content1_idx` (`content_id` ASC),
  CONSTRAINT `fk_strategy_content1`
    FOREIGN KEY (`content_id`)
    REFERENCES `video-crm`.`content` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB
COMMENT = 'Description for advertisement ';


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
