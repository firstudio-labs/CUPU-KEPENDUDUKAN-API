/*
 Navicat Premium Dump SQL

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80031 (8.0.31)
 Source Host           : localhost:3307
 Source Schema         : gmni_db

 Target Server Type    : MySQL
 Target Server Version : 80031 (8.0.31)
 File Encoding         : 65001

 Date: 31/12/2024 00:47:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for indonesia_provinces
-- ----------------------------
DROP TABLE IF EXISTS `indonesia_provinces`;
CREATE TABLE `indonesia_provinces`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `code` char(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `indonesia_provinces_code_unique`(`code` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of indonesia_provinces
-- ----------------------------
INSERT INTO `indonesia_provinces` VALUES (1, '11', 'ACEH', '{\"lat\":\"4.695135\",\"long\":\"96.7493993\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (2, '12', 'SUMATERA UTARA', '{\"lat\":\"2.1153547\",\"long\":\"99.5450974\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (3, '13', 'SUMATERA BARAT', '{\"lat\":\"-0.7399397\",\"long\":\"100.8000051\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (4, '14', 'RIAU', '{\"lat\":\"0.2933469\",\"long\":\"101.7068294\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (5, '15', 'JAMBI', '{\"lat\":\"-1.6101229\",\"long\":\"103.6131203\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (6, '16', 'SUMATERA SELATAN', '{\"lat\":\"-3.3194374\",\"long\":\"103.914399\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (7, '17', 'BENGKULU', '{\"lat\":\"-3.7928451\",\"long\":\"102.2607641\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (8, '18', 'LAMPUNG', '{\"lat\":\"-4.5585849\",\"long\":\"105.4068079\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (9, '19', 'KEPULAUAN BANGKA BELITUNG', '{\"lat\":\"-2.7410513\",\"long\":\"106.4405872\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (10, '21', 'KEPULAUAN RIAU', '{\"lat\":\"3.9456514\",\"long\":\"108.1428669\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (11, '31', 'DKI JAKARTA', '{\"lat\":\"-6.2087634\",\"long\":\"106.845599\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (12, '32', 'JAWA BARAT', '{\"lat\":\"-7.090911\",\"long\":\"107.668887\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (13, '33', 'JAWA TENGAH', '{\"lat\":\"-7.150975\",\"long\":\"110.1402594\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (14, '34', 'DAERAH ISTIMEWA YOGYAKARTA', '{\"lat\":\"-7.8753849\",\"long\":\"110.4262088\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (15, '35', 'JAWA TIMUR', '{\"lat\":\"-7.5360639\",\"long\":\"112.2384017\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (16, '36', 'BANTEN', '{\"lat\":\"-6.4058172\",\"long\":\"106.0640179\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (17, '51', 'BALI', '{\"lat\":\"-8.3405389\",\"long\":\"115.0919509\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (18, '52', 'NUSA TENGGARA BARAT', '{\"lat\":\"-8.6529334\",\"long\":\"117.3616476\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (19, '53', 'NUSA TENGGARA TIMUR', '{\"lat\":\"-8.6573819\",\"long\":\"121.0793705\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (20, '61', 'KALIMANTAN BARAT', '{\"lat\":\"-0.2787808\",\"long\":\"111.4752851\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (21, '62', 'KALIMANTAN TENGAH', '{\"lat\":\"-1.6814878\",\"long\":\"113.3823545\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (22, '63', 'KALIMANTAN SELATAN', '{\"lat\":\"-3.0926415\",\"long\":\"115.2837585\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (23, '64', 'KALIMANTAN TIMUR', '{\"lat\":\"0.5386586\",\"long\":\"116.419389\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (24, '65', 'KALIMANTAN UTARA', '{\"lat\":\"3.0730929\",\"long\":\"116.0413889\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (25, '71', 'SULAWESI UTARA', '{\"lat\":\"0.6246932\",\"long\":\"123.9750018\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (26, '72', 'SULAWESI TENGAH', '{\"lat\":\"-1.4300254\",\"long\":\"121.4456179\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (27, '73', 'SULAWESI SELATAN', '{\"lat\":\"-3.6687994\",\"long\":\"119.9740534\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (28, '74', 'SULAWESI TENGGARA', '{\"lat\":\"-4.14491\",\"long\":\"122.174605\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (29, '75', 'GORONTALO', '{\"lat\":\"0.5435442\",\"long\":\"123.0567693\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (30, '76', 'SULAWESI BARAT', '{\"lat\":\"-2.8441371\",\"long\":\"119.2320784\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (31, '81', 'MALUKU', '{\"lat\":\"-2.8646166\",\"long\":\"129.5765974\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (32, '82', 'MALUKU UTARA', '{\"lat\":\"1.5709993\",\"long\":\"127.8087693\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (33, '91', 'PAPUA', '{\"lat\":\"-4.269928\",\"long\":\"138.0803529\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');
INSERT INTO `indonesia_provinces` VALUES (34, '92', 'PAPUA BARAT', '{\"lat\":\"-1.3361154\",\"long\":\"133.1747162\"}', '2024-10-25 11:46:44', '2024-10-25 11:46:44');

SET FOREIGN_KEY_CHECKS = 1;
