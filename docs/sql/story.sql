/*
Navicat MySQL Data Transfer

Source Database       : story

Target Server Type    : MYSQL
Target Server Version : 50639
File Encoding         : 65001

Date: 2024-01-16 23:36:00
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for story
-- ----------------------------
DROP TABLE IF EXISTS `story`;
CREATE TABLE `story` (
  `id` int(10) unsigned NOT NULL,
  `name` varchar(100) DEFAULT '',
  `created_on` int(10) unsigned DEFAULT '0',
  `modified_on` int(10) unsigned DEFAULT '0',
  `story_score` BIGINT(15) unsigned DEFAULT '0',
  `confidence_level` int(10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`),
  CONSTRAINT chk_confidence_level CHECK (`confidence_level` >= 0 AND `confidence_level` <= 5)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
