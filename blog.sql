/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : blog

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 29/03/2022 22:35:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `article_id` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '文章id',
  `article_title` varchar(50) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '文章标题',
  `article_content` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '文章内容',
  `article_summary` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '文章摘要',
  `is_original` int NOT NULL DEFAULT '1' COMMENT '是否原创',
  `view_count` int unsigned NOT NULL DEFAULT '0' COMMENT '浏览量',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL COMMENT '最后修改时间',
  PRIMARY KEY (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of articles
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `category_id` char(36) COLLATE utf8_unicode_ci NOT NULL COMMENT '分类id',
  `category_name` varchar(20) COLLATE utf8_unicode_ci NOT NULL COMMENT '分类名称',
  `article_id` char(36) COLLATE utf8_unicode_ci NOT NULL COMMENT '文章id',
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for comment_table
-- ----------------------------
DROP TABLE IF EXISTS `comment_table`;
CREATE TABLE `comment_table` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `article_id` int NOT NULL COMMENT '文章id',
  `comment_type` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '评论类型',
  `comment_content` varchar(1000) COLLATE utf8_unicode_ci NOT NULL COMMENT '评论内容',
  `uid` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of comment_table
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for replay_table
-- ----------------------------
DROP TABLE IF EXISTS `replay_table`;
CREATE TABLE `replay_table` (
  `reply_id` int NOT NULL AUTO_INCREMENT COMMENT '回复id',
  `comment_id` int NOT NULL COMMENT '评论id',
  `content` varchar(1000) COLLATE utf8_unicode_ci NOT NULL COMMENT '回复内容',
  `type` varchar(20) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '回复类型',
  `replay_time` datetime NOT NULL COMMENT '回复时间',
  PRIMARY KEY (`reply_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of replay_table
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `tag_id` char(36) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '标签id',
  `tag_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL COMMENT '标签名称',
  `atricle_id` char(36) COLLATE utf8_unicode_ci NOT NULL COMMENT '文章id',
  PRIMARY KEY (`tag_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of tag
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `user_id` char(36) COLLATE utf8_unicode_ci NOT NULL COMMENT '用户ID',
  `user_name` varchar(128) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '用户名',
  `user_pwd` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '用户密码',
  `user_email` varchar(64) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '用户邮箱',
  `user_avatar` varchar(128) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '用户头像',
  `user_role` int unsigned DEFAULT '0' COMMENT '用户角色（0普通用户，1管理员）',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '注册时间',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='用户 ';

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES ('939beb07-30c7-4f61-9e02-700fc734e4d9', 'Kepler', '123456', 'anyi123520@gmail.com', 'https://www.baidu.com', 1, '2022-03-28 21:43:08');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
