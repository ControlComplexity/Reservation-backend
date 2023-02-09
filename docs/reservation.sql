/*
Navicat MySQL Data Transfer

Source Server         : 101.43.39.188
Source Server Version : 80031
Source Host           : 101.43.39.188:3306
Source Database       : reservation

Target Server Type    : MYSQL
Target Server Version : 80031
File Encoding         : 65001

Date: 2023-02-09 20:34:01
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity` (
  `activity_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `name` varchar(255) NOT NULL,
  `label` varchar(255) NOT NULL,
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
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of activity
-- ----------------------------
INSERT INTO `activity` VALUES ('1', 'freetalk', '#联谊脱单,#国内名校&海归', '80', '2023-01-19 16:37:07', '上海市浦东新区金山饭店', 'https://5b0988e595225.cdn.sohucs.com/images/20190422/97d544867870416f92f39ec5fecd14da.jpeg', 'https://5b0988e595225.cdn.sohucs.com/images/20190422/97d544867870416f92f39ec5fecd14da.jpeg', '自由交流', '2023-01-19 16:37:07', '2023-01-19 16:37:07', '2022-12-24 19:27:38');
INSERT INTO `activity` VALUES ('2', 'mbti', '#联谊脱单,#985/211/海归为主', '70', '2023-01-19 16:37:11', '上海市浦东新区金山饭店', 'https://www.mianfeiwendang.com/pic/e9b4783148b5e1bcda623096/5-810-jpg_6-1080-0-0-1080.jpg', 'https://www.mianfeiwendang.com/pic/e9b4783148b5e1bcda623096/5-810-jpg_6-1080-0-0-1080.jpg', 'MBTI测试', '2023-01-19 16:37:11', '2023-01-19 16:37:11', '2022-12-24 19:27:38');
INSERT INTO `activity` VALUES ('3', '心理学', '#联谊脱单,#985/211/海归为主', '100', '2023-01-19 16:37:16', '上海市浦东新区金山饭店', 'https://txt22262.book118.com/2018/0117/book149273/149272065.jpg', 'https://txt22262.book118.com/2018/0117/book149273/149272065.jpg', '听心理分析师讲述个人故事', '2023-01-19 16:37:16', '2023-01-19 16:37:16', '2022-12-24 19:27:38');
INSERT INTO `activity` VALUES ('4', '羽毛球', '#运动', '120', '2023-01-18 22:44:09', '上海市浦东新区金山饭店', 'https://bpic.588ku.com/element_origin_min_pic/20/11/05/0cc8ad22082a8cbd4043b38309abad50.jpg', 'https://bpic.588ku.com/element_origin_min_pic/20/11/05/0cc8ad22082a8cbd4043b38309abad50.jpg', '打羽毛球锻炼身体', '2023-01-18 22:44:09', '2023-01-18 22:44:09', '2022-12-24 19:27:38');
INSERT INTO `activity` VALUES ('5', '乒乓球', '#运动', '120', '2023-01-18 22:44:30', '上海市浦东新区金山饭店', 'https://tyunfile.71360.com/UploadFile/bjside7/637643033252732091/2733169.png', 'https://tyunfile.71360.com/UploadFile/bjside7/637643033252732091/2733169.png', '打乒乓球锻炼身体', '2023-01-18 22:44:30', '2023-01-18 22:44:30', '2022-12-24 19:27:38');
INSERT INTO `activity` VALUES ('6', '狼人杀', '#游戏', '50', '2023-01-18 22:42:05', '上海市浦东新区金山饭店', 'https://img.haote.com/upload/news/20210325/161666440651353.jpeg', 'https://img.haote.com/upload/news/20210325/161666440651353.jpeg', '狼人杀一决高下', '2023-01-18 22:42:05', '2023-01-18 22:42:05', '2022-12-24 19:27:38');

-- ----------------------------
-- Table structure for enlist
-- ----------------------------
DROP TABLE IF EXISTS `enlist`;
CREATE TABLE `enlist` (
  `id` int NOT NULL,
  `user_1_id` int DEFAULT NULL,
  `user_2_id` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of enlist
-- ----------------------------
INSERT INTO `enlist` VALUES ('1', '2', '1');
INSERT INTO `enlist` VALUES ('2', '2', '3');
INSERT INTO `enlist` VALUES ('3', '4', '2');

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `order_id` int unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL,
  `activity_id` int DEFAULT NULL,
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES ('1', '1', '1');
INSERT INTO `orders` VALUES ('2', '1', '2');
INSERT INTO `orders` VALUES ('3', '2', '1');
INSERT INTO `orders` VALUES ('4', '2', '2');
INSERT INTO `orders` VALUES ('5', '2', '4');

-- ----------------------------
-- Table structure for swiper
-- ----------------------------
DROP TABLE IF EXISTS `swiper`;
CREATE TABLE `swiper` (
  `img` varchar(255) DEFAULT NULL,
  `url` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of swiper
-- ----------------------------
INSERT INTO `swiper` VALUES ('../../static/swiper1.jpg', 'www.baidu.com');
INSERT INTO `swiper` VALUES ('../../static/swiper1.jpg', 'www.baidu.com');
INSERT INTO `swiper` VALUES ('../../static/swiper1.jpg', 'www.baidu.com');
INSERT INTO `swiper` VALUES ('../../static/swiper1.jpg', 'www.baidu.com');
INSERT INTO `swiper` VALUES ('../../static/swiper1.jpg', 'www.baidu.com');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `head_image` varchar(255) DEFAULT NULL,
  `nick_name` varchar(255) DEFAULT NULL,
  `birth_date` varchar(20) DEFAULT NULL,
  `gender` enum('male','female') DEFAULT NULL,
  `height` int DEFAULT NULL,
  `weight` int DEFAULT NULL,
  `hometown` varchar(255) DEFAULT NULL,
  `emotional_status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `education` varchar(255) DEFAULT NULL,
  `university` varchar(255) DEFAULT NULL,
  `occupation` varchar(255) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `wechat_number` varchar(255) DEFAULT NULL,
  `phone_number` varchar(255) DEFAULT NULL,
  `open_id` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` timestamp(6) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6),
  `updated_at` timestamp(6) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', null, 'fang', '1999-12-12', 'male', '185', '71', '上海金山', 'SINGLE_WITHOUT_MARRIAGE_HISTORY', 'UNIFIED_BACHELOR', '清华大学', null, '新东方', '13167266118', '13167266118', null, null, '2023-02-09 20:33:26.398863', '2023-02-09 20:33:26.398863');
INSERT INTO `user` VALUES ('2', null, 'sdsdfds', '1989-02-12', 'female', '163', '0', null, null, null, null, null, null, null, '13166661111', 'oVGbv5EAoFhF5aHVr5tQaF5xrTec', null, '2023-02-09 20:33:46.013159', '2023-02-09 20:33:46.013159');
