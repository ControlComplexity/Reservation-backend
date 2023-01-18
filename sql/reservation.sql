/*
Navicat MySQL Data Transfer

Source Server         : 101.43.39.188
Source Server Version : 80031
Source Host           : 101.43.39.188:3306
Source Database       : reservation

Target Server Type    : MYSQL
Target Server Version : 80031
File Encoding         : 65001

Date: 2023-01-18 20:59:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity` (
  `activity_id` int NOT NULL COMMENT 'id',
  `name` varchar(255) NOT NULL,
  `price` float(5,0) NOT NULL,
  `time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `location` varchar(255) NOT NULL,
  `small_img` varchar(255) NOT NULL,
  `big_img` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`activity_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of activity
-- ----------------------------
INSERT INTO `activity` VALUES ('1', 'freetalk', '80', '2022-12-24 00:00:00', '上海市浦东新区金山饭店', 'a.png', 'b.png', '自由交流', '2022-12-24 19:27:38', '2022-12-24 19:28:12', '2022-12-24 19:27:38');
INSERT INTO `activity` VALUES ('2', 'mbti', '70', '2022-12-25 00:00:00', '上海市浦东新区金山饭店', 'a.png', 'b.png', 'MBTI测试', '2022-12-24 19:27:38', '2022-12-24 19:28:13', '2022-12-24 19:27:38');

-- ----------------------------
-- Table structure for order
-- ----------------------------
DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `order_id` int NOT NULL,
  `user_id` int DEFAULT NULL,
  `activity_id` int DEFAULT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of order
-- ----------------------------
INSERT INTO `order` VALUES ('1', '1', '1');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL,
  `open_id` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `phone_number` bigint DEFAULT NULL,
  `head_image` varchar(255) DEFAULT NULL,
  `created_at` datetime(6) DEFAULT NULL,
  `updated_at` datetime(6) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'jinzhu', 'sdsdfds', 'oVGbv5EAoFhF5aHVr5tQaF5xrTec', null, '13166661111', null, '2023-01-18 20:56:24.622467', '2023-01-18 20:56:24.622467');
